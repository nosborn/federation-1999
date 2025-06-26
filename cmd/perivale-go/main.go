package main

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/nosborn/federation-1999/internal/ioctl"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/nosborn/federation-1999/pkg/ibgames/goodies"
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

const linkHostnameSize = 64

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

func main() {
	log.SetPrefix(fmt.Sprintf("%s[%d]: ", filepath.Base(os.Args[0]), os.Getpid()))
	log.SetFlags(log.Lmsgprefix)

	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: %s id address robobod\n", os.Args[0])
		os.Exit(1)
	}

	//
	// _term := os.Getenv("TERM")
	// if _term != "" {
	// 	//
	// }

	// Check for a valid UID.
	uid, err := strconv.Atoi(os.Args[1])
	if err != nil || uid < ibgames.MinAccountID || uid >= ibgames.MaxAccountID {
		fmt.Fprintf(os.Stderr, "%s: Bad UID", os.Args[0])
		os.Exit(1)
	}

	// Check for lock out!
	if rules.IsLockedOut(ibgames.AccountID(uid)) {
		lockout()
	}

	remoteHostname := os.Args[2]
	if names, err := net.LookupAddr(os.Args[2]); err == nil && len(names) > 0 {
		remoteHostname = strings.TrimSuffix(names[0], ".")
	}

	for len(remoteHostname) > linkHostnameSize {
		if len(remoteHostname) == 0 || remoteHostname[0] == '.' {
			break
		}
		remoteHostname = remoteHostname[1:]
	}
	if len(remoteHostname) > linkHostnameSize {
		remoteHostname = remoteHostname[:linkHostnameSize]
	}

	robobod, _ := strconv.Atoi(os.Args[3])
	login := fmt.Sprintf("%d %s %s %d\n", uid, os.Args[2], remoteHostname, robobod)

	socketPath := fmt.Sprintf("%s/.fedtpd.socket", goodies.HomeDir())
	unixAddr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		log.Fatalf("socket: %v", err)
	}

	unixConn, err := net.DialUnix("unix", nil, unixAddr)
	if err != nil {
		unavailable()
	}
	defer unixConn.Close()

	if n, err := unixConn.Write([]byte(login)); err != nil || n != len(login) {
		unavailable()
	}

	stdinFd := int(os.Stdin.Fd())
	if err = unix.SetNonblock(stdinFd, true); err != nil {
		log.Fatalf("unix.SetNonblock(stdin): %v", err)
	}

	stdoutFd := int(os.Stdout.Fd())
	if err = unix.SetNonblock(stdoutFd, true); err != nil {
		log.Fatalf("unix.SetNonblock(stdout): %v", err)
	}

	f, err := unixConn.File()
	if err != nil {
		log.Fatalf("unixConn.File: %v", err)
	}
	defer f.Close()

	connFd := int(f.Fd())

	oldTermios, err := unix.IoctlGetTermios(stdinFd, ioctl.GetTerminalAttrs)
	if err != nil {
		log.Printf("IoctlGetTermios: %v", err)
	} else {
		defer func() {
			_ = unix.IoctlSetTermios(stdinFd, ioctl.SetTerminalAttrsNow, oldTermios)
		}()

		newTermios := *oldTermios
		newTermios.Iflag |= unix.ICRNL | unix.IGNBRK | unix.ISTRIP
		newTermios.Lflag |= unix.ICANON
		newTermios.Lflag &^= unix.ECHO | unix.ISIG
		newTermios.Cc[unix.VINTR] = 0
		newTermios.Cc[unix.VQUIT] = 0
		newTermios.Cc[unix.VKILL] = 0

		if err := unix.IoctlSetTermios(stdinFd, ioctl.SetTerminalAttrsNow, &newTermios); err != nil {
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

	if err := selector.AddRead(stdinFd); err != nil {
		log.Fatalf("selector.AddRead(stdin): %v", err)
	}
	if err := selector.AddRead(connFd); err != nil {
		log.Fatalf("selector.AddRead(conn): %v", err)
	}

	for {
		if err := selector.UpdateWrite(stdoutFd, oq.Len() > 0); err != nil && !errors.Is(err, unix.EEXIST) {
			log.Fatalf("selector.UpdateWrite: %v", err)
		}

		timeout := -1 * time.Second
		if inState == inThrottle {
			now := time.Now()
			if now.After(unthrottleAt) {
				inState = inReady
				timeout = 0
			} else {
				timeout = unthrottleAt.Sub(now)
			}
		}

		events, err := selector.Wait(timeout)
		if err != nil {
			log.Fatalf("wait: %v", err)
		}

		if atomic.CompareAndSwapInt32(&winch, 1, 0) {
			cols, rows, err = term.GetSize(stdinFd)
			if err != nil {
				log.Fatal("TIOCGWINSZ")
			}
		}

		if inState == inThrottle && time.Now().After(unthrottleAt) {
			inState = inReady
		}

		if len(events) == 0 && inState == inThrottle && timeout == 0 {
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

				buf := make([]byte, 32768)
				n, err := unixConn.Read(buf)
				if err != nil {
					if opErr, ok := err.(*net.OpError); ok && opErr.Err == unix.EAGAIN {
						break
					}
					log.Printf("read: %v", err)
					flushAllOutput()
					os.Exit(1)
				}

				if n == 0 {
					flushAllOutput()
					os.Exit(0)
				}

				output(buf[:n])
			case stdoutFd:
				flushOutput()
			case stdinFd:
				readInput()
			}
		}

		if iq.Len() > 0 && inState == inReady {
			element := iq.Front()
			line := element.Value.([]byte)
			n, err := unixConn.Write(line)
			if err != nil {
				if opErr, ok := err.(*net.OpError); ok && opErr.Err == unix.EAGAIN {
					continue
				}
				log.Fatalf("write: %v", err)
			}
			if n != len(line) {
				log.Fatalf("write: partial (%d of %d bytes)", n, len(line))
			}
			iq.Remove(element)
			inState = inAckWait
		}
	}
}

func flushAllOutput() {
	stdoutFd := int(os.Stdout.Fd())
	_ = unix.SetNonblock(stdoutFd, false)
	defer unix.SetNonblock(stdoutFd, true)

	for oq.Len() > 0 {
		flushOutput()
		if oq.Len() > 0 {
			time.Sleep(1 * time.Second)
		}
	}
}

func flushOutput() {
	if oq.Len() == 0 {
		return
	}

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
	if len(buf) == 0 {
		return
	}

	var out bytes.Buffer
	line := make([]byte, 0, len(buf))

	if trace {
		fmt.Fprintf(&out, "[LEN=%d]", len(buf))
	}

	for i := 0; i < len(buf); i++ {
		ch := buf[i]

		switch state {
		case tsData:
			switch ch {
			case model.DLE:
				state = tsEscape
			case '\a':
				line = append(line, ch)
			case '\n':
				if allWS {
					if !lastBlank {
						lastBlank = true
					}
					line = line[:0]
					blankPending = true
				} else {
					lastBlank = false
				}

				if len(line) > 0 || continuation {
					line = trimTrailingSpaces(line)
					if blankPending {
						if trace {
							out.WriteString("[PEND]")
						}
						if spyDepth == lastSpyDepth {
							out.WriteString(prefix)
						}
						out.WriteByte('\n')
						blankPending = false
					}
					if spyDepth > 0 {
						out.WriteString(prefix)
						lastSpyDepth = spyDepth
					}
					out.Write(line)
					if trace {
						out.WriteString("[PARA]")
					}
					out.WriteByte('\n')
					line = line[:0]
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
					line = append(line, '^', ch^0x40)
				} else {
					line = append(line, ch)
				}

				wrapLimit := cols - len(prefix)
				if cols > 0 && wrapLimit > 0 && column == wrapLimit {
					if lastSpace := bytes.LastIndexByte(line, ' '); lastSpace != -1 {
						chop := len(line) - lastSpace - 1
						if chop > 0 {
							i -= chop
							if i < 0 {
								i = -1
							}
							line = line[:lastSpace+1]
						}
					}

					line = trimTrailingSpaces(line)

					if blankPending {
						if trace {
							out.WriteString("[PEND]")
						}
						if spyDepth == lastSpyDepth {
							out.WriteString(prefix)
						}
						out.WriteByte('\n')
						blankPending = false
					}
					if spyDepth > 0 {
						out.WriteString(prefix)
						lastSpyDepth = spyDepth
					}
					out.Write(line)
					if trace {
						out.WriteString("[WRAP]")
					}
					out.WriteByte('\n')
					line = line[:0]
					continuation = false
					column = 1
					allWS = true
					eatBlanks = true
				} else {
					column++
				}
			}
		case tsEscape:
			switch ch {
			case model.LeAck:
				if spyDepth == 0 {
					if trace {
						out.WriteString("[ACK]")
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
				fmt.Fprintf(&out, "[S=%d]", ch)
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
					fmt.Fprintf(&out, "[T=%d]", boolToInt(trace))
				}
			}
			state = tsData
		}
	}

	if len(line) > 0 {
		if blankPending {
			if spyDepth == lastSpyDepth {
				out.WriteString(prefix)
			}
			out.WriteByte('\n')
			blankPending = false
		}
		if spyDepth > 0 {
			out.WriteString(prefix)
			lastSpyDepth = spyDepth
		}
		out.Write(line)
		continuation = true
	}

	if out.Len() > 0 {
		data := make([]byte, out.Len())
		copy(data, out.Bytes())
		writeOutput(data)
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

	for _, ch := range buf[:n] {
		switch ch {
		case '\n':
			for inputLineLen > 0 && inputLine[inputLineLen-1] == ' ' {
				inputLineLen--
			}
			if inputLineLen >= len(inputLine) {
				inputLineLen = len(inputLine) - 1
			}
			inputLine[inputLineLen] = '\n'
			inputLineLen++
			line := make([]byte, inputLineLen)
			copy(line, inputLine[:inputLineLen])
			iq.PushBack(line)
			inputLineLen = 0
		case 0x12: // Ctrl-R
			if inputLineLen > 0 {
				line := make([]byte, inputLineLen)
				copy(line, inputLine[:inputLineLen])
				output(line)
			}
		case 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
			0x08, 0x09, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
			0x11, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
			0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x7F:
			// Ignored control characters.
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

	fmt.Fprint(os.Stdout, message)

	time.Sleep(3 * time.Second)
	os.Exit(0)
}

func writeOutput(buf []byte) {
	if len(buf) == 0 {
		return
	}

	if oq.Len() == 0 {
		n, err := os.Stdout.Write(buf)
		if err != nil {
			if errors.Is(err, unix.EAGAIN) {
				n = 0
			} else {
				log.Fatalf("write to stdout failed: %v", err)
			}
		}
		if n == len(buf) {
			return
		}
		buf = buf[n:]
	}

	data := make([]byte, len(buf))
	copy(data, buf)
	oq.PushBack(data)
}

func trimTrailingSpaces(b []byte) []byte {
	i := len(b)
	for i > 0 && b[i-1] == ' ' {
		i--
	}
	return b[:i]
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
