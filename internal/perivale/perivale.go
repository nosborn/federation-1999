package perivale

import (
	"container/list"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"time"

	"github.com/nosborn/federation-1999/internal/ioctl"
	"github.com/nosborn/federation-1999/internal/link"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/nosborn/federation-1999/pkg/ibgames/rules"
	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

type InState int

const (
	inReady InState = iota
	inAckWait
	inThrottle
)

var (
	winch int32 = 1 // Fake it the first time through

	rows int //nolint:unused
	cols int

	inState      = inReady
	unthrottleAt time.Time

	iq = list.New()
	oq = list.New()

	inputLine    = make([]byte, 256)
	inputLineLen = 0
)

func Perivale(uid ibgames.AccountID, remoteAddr string, robobod int) {
	// //
	// term := os.Getenv("TERM")
	// if term != "" {
	// 	//
	// }

	// Check for lock out!
	if rules.IsLockedOut(uid) {
		lockout()
	}

	var remoteHostname string
	if names, err := net.LookupAddr(remoteAddr); err != nil {
		remoteHostname = names[0]
	} else {
		remoteHostname = remoteAddr // Use the dotted quad instead
	}

	login := fmt.Sprintf("%d %s %s %d\n", uid, os.Args[2], remoteHostname, robobod)

	unixAddr, err := net.ResolveUnixAddr("unix", link.SocketPath)
	if err != nil {
		log.Fatalf("socket: %v", err)
	}
	unixConn, err := net.DialUnix("unix", nil, unixAddr)
	if err != nil {
		unavailable()
	}
	defer unixConn.Close()

	_, err = unixConn.Write([]byte(login))
	if err != nil {
		unavailable()
	}

	//
	stdinFd := int(os.Stdin.Fd())
	if err = unix.SetNonblock(stdinFd, true); err != nil {
		log.Fatalf("unix.SetNonblock: %v", err)
	}
	stdoutFd := int(os.Stdout.Fd())
	if err = unix.SetNonblock(stdoutFd, true); err != nil {
		log.Fatalf("unix.SetNonblock: %v", err)
	}

	f, _ := unixConn.File()
	connFd := int(f.Fd())

	oldTermios, err := unix.IoctlGetTermios(int(os.Stdin.Fd()), ioctl.GetTerminalAttrs)
	if err != nil {
		log.Printf("IoctlGetTermios: %v", err)
	} else {
		defer func() {
			unix.IoctlSetTermios(int(os.Stdin.Fd()), ioctl.SetTerminalAttrsNow, oldTermios)
		}()

		newTermios := *oldTermios
		newTermios.Iflag |= unix.ICRNL | unix.IGNBRK | unix.ISTRIP
		newTermios.Lflag |= unix.ICANON
		newTermios.Lflag &^= unix.ECHO | unix.ISIG
		newTermios.Cc[unix.VINTR] = 0
		newTermios.Cc[unix.VQUIT] = 0
		newTermios.Cc[unix.VKILL] = 0

		if err := unix.IoctlSetTermios(int(os.Stdin.Fd()), ioctl.SetTerminalAttrsNow, &newTermios); err != nil {
			log.Printf("IoctlSetTermios: %v", err)
		}
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, unix.SIGWINCH)
	go func() {
		for range sigCh {
			sigWinch()
		}
	}()

	selector, err := NewSelector()
	if err != nil {
		log.Fatalf("NewSelector: %v", err)
	}
	defer selector.Close()

	selector.AddRead(stdinFd)
	selector.AddRead(connFd)

	for {
		selector.UpdateWrite(stdoutFd, oq.Len() > 0)

		timeout := -1 * time.Second // Negative timeout blocks indefinitely
		if inState == inThrottle {
			now := time.Now()
			if now.After(unthrottleAt) {
				inState = inReady
				timeout = 0 // Poll, don't block
			} else {
				timeout = unthrottleAt.Sub(now)
			}
		}

		events, err := selector.Wait(timeout)
		if err != nil {
			log.Fatalf("wait: %v", err)
		}

		if atomic.CompareAndSwapInt32(&winch, 1, 0) {
			var err error
			cols, rows, err = term.GetSize(int(os.Stdin.Fd()))
			if err != nil {
				log.Fatal("TIOCGWINSZ")
			}
		}

		if len(events) == 0 && inState == inThrottle {
			inState = inReady
		}

		for _, event := range events {
			switch event.Fd {
			case connFd:
				if event.IsEOF {
					flushAllOutput()
					time.Sleep(1 * time.Second)
					os.Exit(0)
				}
				buf := make([]byte, 4096)
				n, err := unixConn.Read(buf)
				if err != nil {
					if opErr, ok := err.(*net.OpError); ok && opErr.Err == unix.EAGAIN {
						break // Nothing to read
					}
					log.Printf("read: %v", err)
					flushAllOutput()
					os.Exit(1)
				}

				if n == 0 { // EOF from game server
					flushAllOutput()
					os.Exit(0)
				}

				if n > 0 {
					output(buf[:n])
				}
			case stdoutFd:
				flushOutput()
			case stdinFd:
				readInput()
			}
		}

		if iq.Len() > 0 && inState == inReady {
			element := iq.Front()
			line := element.Value.([]byte)
			_, err := unixConn.Write(line)
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && opErr.Err == unix.EAGAIN {
					return
				}
				log.Fatalf("write: %v", err)
			}
			iq.Remove(element)
			inState = inAckWait
		}
	}
}

func flushAllOutput() {
	unix.SetNonblock(int(os.Stdout.Fd()), false)
	for oq.Len() > 0 {
		flushOutput()
	}
}

func flushOutput() {
	for oq.Len() > 0 {
		element := oq.Front()
		data := element.Value.([]byte)

		n, err := os.Stdout.Write(data)
		if err != nil {
			if errors.Is(err, unix.EAGAIN) {
				if n > 0 {
					element.Value = data[n:]
				}
				return
			}
			log.Fatalf("write to stdout failed: %v", err)
		}

		if n < len(data) {
			element.Value = data[n:]
			return
		}
		oq.Remove(element)
	}
}

func lockout() {
	fmt.Println("")
	fmt.Println("Your account has been locked out of the Federation game. If you")
	fmt.Println("are unsure about why you have been locked out of the game, please")
	fmt.Println("read the rules at <URL:https://federation-1999.fly.dev/ibinfo/t&c.html>")
	fmt.Printf("or send e-mail to %s.\n\n", rules.INFO_EMAIL)
	time.Sleep(5 * time.Second)
	os.Exit(0)
}

type State int

const (
	tsData State = iota
	tsSpyDepth
	tsTrace
	tsEscape
)

var (
	state        State  = tsData
	trace        bool   = false
	blankPending bool   = true
	column       int    = 1
	allWS        bool   = false
	eatBlanks    bool   = false
	lastBlank    bool   = false
	lastSpyDepth int    = 0
	prefix       string = ""
	spyDepth     int    = 0
	continuation bool   = false
)

func output(buf []byte) {
	var output, line strings.Builder

	if trace {
		fmt.Fprintf(&output, "[LEN=%d]", len(buf))
	}

	i := 0
	for i < len(buf) {
		ch := buf[i]
		i++

		switch state {
		case tsData:
			switch ch {
			case model.DLE:
				state = tsEscape
			case '\a':
				line.WriteByte(ch)
			case '\n':
				if allWS {
					if !lastBlank {
						lastBlank = true
					}
					line.Reset()
					blankPending = true
				} else {
					lastBlank = false
				}

				if line.Len() > 0 || continuation {
					trimmedLine := strings.TrimRight(line.String(), " ")
					if blankPending {
						if trace {
							output.WriteString("[PEND]")
						}
						if spyDepth == lastSpyDepth {
							output.WriteString(prefix)
						}
						output.WriteByte('\n')
						blankPending = false
					}
					if spyDepth > 0 {
						output.WriteString(prefix)
						lastSpyDepth = spyDepth
					}
					output.WriteString(trimmedLine)
					if trace {
						output.WriteString("[PARA]")
					}
					output.WriteByte('\n')
					line.Reset()
					continuation = false
				}
				column = 1
				allWS = true
				eatBlanks = false
			case '\r':
				continue
			case ' ':
				if eatBlanks {
					continue
				}
				fallthrough
			default:
				allWS = allWS && (ch == ' ')
				if eatBlanks {
					if ch == '/' || ch == '>' {
						continue
					}
					eatBlanks = false
				}

				if ch < 32 || ch > 126 {
					line.WriteByte('^')
					line.WriteByte(ch ^ 0x40)
				} else {
					line.WriteByte(ch)
				}

				column++
				if cols > 0 && column >= cols-len(prefix) {
					lineStr := line.String()
					lastSpace := strings.LastIndex(lineStr, " ")
					if lastSpace != -1 {
						chop := len(lineStr) - lastSpace - 1
						if chop > 0 {
							i -= chop
							line.Reset()
							line.WriteString(lineStr[:lastSpace+1])
						}
					}
					trimmedLine := strings.TrimRight(line.String(), " ")
					line.Reset()
					line.WriteString(trimmedLine)

					if blankPending {
						if trace {
							output.WriteString("[PEND]")
						}
						if spyDepth == lastSpyDepth {
							output.WriteString(prefix)
						}
						output.WriteByte('\n')
						blankPending = false
					}
					if spyDepth > 0 {
						output.WriteString(prefix)
						lastSpyDepth = spyDepth
					}
					output.WriteString(line.String())
					if trace {
						output.WriteString("[WRAP]")
					}
					output.WriteByte('\n')
					line.Reset()
					continuation = false
					column = 1
					allWS = true
					eatBlanks = true
				}
			}
		case tsEscape:
			switch ch {
			case model.LeAck:
				if spyDepth == 0 {
					if trace {
						output.WriteString("[ACK]")
					}
					if inState == inAckWait {
						unthrottleAt = time.Now().Add(500 * time.Millisecond)
						inState = inThrottle
					}
				}
				state = tsData
			case model.LeSpy:
				state = tsSpyDepth
			case model.LeTrace:
				state = tsTrace
			default:
				state = tsData
			}
		case tsSpyDepth:
			if trace {
				fmt.Fprintf(&output, "[S=%d]", ch)
			}
			prefix = ""
			lastSpyDepth = spyDepth
			spyDepth = int(ch)
			if spyDepth > 0 {
				prefix = strings.Repeat("/", spyDepth) + " "
			}
			state = tsData
		case tsTrace:
			if spyDepth == 0 {
				wasTrace := trace
				trace = (ch != '-')
				if trace || wasTrace {
					fmt.Fprintf(&output, "[T=%t]", trace)
				}
			}
			state = tsData
		}
	}

	if line.Len() > 0 {
		if blankPending {
			if spyDepth == lastSpyDepth {
				output.WriteString(prefix)
			}
			output.WriteByte('\n')
			blankPending = false
		}
		if spyDepth > 0 {
			output.WriteString(prefix)
			lastSpyDepth = spyDepth
		}
		output.WriteString(line.String())
		continuation = true
	}

	if output.Len() > 0 {
		writeOutput([]byte(output.String()))
	}
}

func readInput() {
	buf := make([]byte, 4096)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		if err == unix.EAGAIN {
			return
		}
		log.Printf("read from stdin failed: %v", err)
		flushAllOutput()
		os.Exit(1)
	}

	if n == 0 { // EOF
		flushAllOutput()
		os.Exit(0)
	}

	for i := range n {
		ch := buf[i]
		switch ch {
		case '\n':
			for inputLineLen > 0 && inputLine[inputLineLen-1] == ' ' {
				inputLineLen--
			}
			inputLine[inputLineLen] = '\n'
			inputLineLen++
			line := make([]byte, inputLineLen)
			copy(line, inputLine[:inputLineLen])
			iq.PushBack(line)
			inputLineLen = 0
		default:
			if inputLineLen < len(inputLine) {
				inputLine[inputLineLen] = ch
				inputLineLen++
			}
		}
	}
}

func sigWinch() {
	atomic.StoreInt32(&winch, 1)
}

func unavailable() {
	message := "\n\nFederation is temporarily unavailable. Please try again later.\n\n"

	time.Sleep(2 * time.Second)

	fmt.Fprintln(os.Stdout, message)
	// fflush(stdout);

	time.Sleep(3 * time.Second)
	os.Exit(0)
}

func writeOutput(buf []byte) {
	oq.PushBack(buf)
}
