package telnet

import (
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Conn struct {
	net.Conn

	buf     []byte
	sbBuf   []byte
	state   int
	writeMu sync.Mutex

	OnEchoChange       func(enabled bool)
	OnTerminalSpeed    func(transmit, receive int)
	OnWindowSizeChange func(rows, cols uint16)
}

const (
	STATE_DATA = iota
	STATE_IAC
	STATE_CR
	STATE_CMD
	STATE_SB
	STATE_SB_IAC
)

func NewConn(c net.Conn) *Conn {
	t := &Conn{
		Conn:  c,
		buf:   make([]byte, 0, 1024),
		sbBuf: make([]byte, 0, 64),
		// echoEnabled: false,
	}

	// file, err := c.(*net.TCPConn).File()
	// fd := file.Fd()
	//
	// termios, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
	// if err != nil {
	// 	log.Printf("unix.IoctlGetTermios, %v", err)
	// }
	// termios.Lflag |= unix.ICANON
	// termios.Lflag |= unix.ECHO
	// termios.Oflag |= unix.ONLCR
	// err = unix.IoctlSetTermios(int(fd), unix.TCSETS, termios)
	// if err != nil {
	// 	log.Printf("unix.IoctlSetTermios: %v", err)
	// }

	_, err := c.Write([]byte{
		IAC, WILL, TELOPT_SGA,
		IAC, WONT, TELOPT_LINEMODE,
		IAC, WILL, TELOPT_ECHO,
		IAC, DO, TELOPT_NAWS,
		IAC, DO, TELOPT_TSPEED,
	})
	if err != nil {
		return nil // FIXME
	}

	_, err = c.Write([]byte("\r\n\r\n" +
		"Welcome to ibgames, the home of Federation.\r\n\r\n" +
		"Please type your ibgames account ID at the login prompt, then type in your\r\n" +
		"password.\r\n\r\n" +
		"If you do not have an ibgames account, you will need to go to our web page\r\n" +
		"at <URL:https://federation-1999.fly.dev/index.html> to set an account up.\r\n\r\n\r\n",
	))
	if err != nil {
		return nil // FIXME
	}

	return t
}

func (t *Conn) Read(p []byte) (int, error) {
	buf := make([]byte, 1)
	for {
		_, err := t.Conn.Read(buf)
		if err != nil {
			return 0, err
		}
		b := buf[0]
		switch t.state {
		case STATE_CR:
			t.state = STATE_DATA
			if b == 0 || b == '\n' {
				break
			}
			fallthrough

		case STATE_DATA:
			switch b {
			case IAC:
				t.state = STATE_IAC
			case '\r':
				p[0] = '\n'
				// if t.echoEnabled {
				// 	_, _ = t.Write([]byte{b})    // FIXME
				// 	_, _ = t.Write([]byte{'\n'}) // FIXME
				// }
				t.state = STATE_CR
				return 1, nil
			default:
				p[0] = b
				// if t.echoEnabled {
				// 	_, _ = t.Write([]byte{b}) // FIXME
				// }
				return 1, nil
			}

		case STATE_IAC:
			switch b {
			case IAC:
				p[0] = IAC
				t.state = STATE_DATA
				return 1, nil
			case DO, DONT, WILL, WONT:
				t.state = STATE_CMD
				t.buf = append(t.buf[:0], b)
			case SB:
				t.state = STATE_SB
				t.sbBuf = t.sbBuf[:0]
			default:
				t.state = STATE_DATA
			}
		case STATE_CMD:
			cmd := t.buf[0]
			opt := b
			t.handleOption(cmd, opt)
			t.state = STATE_DATA
		case STATE_SB:
			if b == IAC {
				t.state = STATE_SB_IAC
			} else {
				t.sbBuf = append(t.sbBuf, b)
			}
		case STATE_SB_IAC:
			switch b {
			case SE:
				if len(t.sbBuf) > 0 {
					t.handleSubnegotiation(t.sbBuf[0], t.sbBuf[1:])
				}
				t.state = STATE_DATA
			case IAC:
				t.sbBuf = append(t.sbBuf, IAC)
				t.state = STATE_SB
			default:
				t.state = STATE_SB
			}
		default:
			log.Fatal("invalid state")
		}
	}
}

func (t *Conn) Write(p []byte) (int, error) {
	out := make([]byte, 0, len(p)*2)
	for _, b := range p {
		switch b {
		// case '\n':
		// 	out = append(out, '\r')
		// 	out = append(out, '\n')
		case IAC:
			out = append(out, IAC, IAC)
		default:
			out = append(out, b)
		}
	}
	n, err := t.write(out)
	if err != nil {
		return n, err // FIXME: n might be invalid
	}
	return len(p), nil
}

func (t *Conn) handleOption(cmd, opt byte) {
	switch cmd {
	case DO:
		switch opt {
		case TELOPT_ECHO:
			if t.OnEchoChange != nil {
				t.OnEchoChange(true)
			}
			_, _ = t.write([]byte{IAC, WILL, opt})
		case TELOPT_SGA:
			_, _ = t.write([]byte{IAC, WILL, opt}) // FIXME
		// case TELOPT_STATUS:
		// 	t.write([]byte{IAC, WILL, opt})
		// case TELOPT_LINEMODE, TELOPT_LFLOW, TELOPT_NAWS, TELOPT_TSPEED, TELOPT_TTYPE:
		// 	break // do nothing
		default:
			_, _ = t.write([]byte{IAC, WONT, opt}) // FIXME
		}
	case DONT:
		switch opt {
		case TELOPT_ECHO:
			if t.OnEchoChange != nil {
				t.OnEchoChange(false)
			}
			_, _ = t.write([]byte{IAC, WONT, opt})
		case TELOPT_SGA:
			_, _ = t.write([]byte{IAC, WONT, opt}) // FIXME
		default:
			_, _ = t.write([]byte{IAC, WONT, opt}) // FIXME
		}

	case WILL:
		switch opt {
		case TELOPT_ECHO:
			_, _ = t.write([]byte{IAC, DO, opt}) // FIXME
			// t.echoEnabled = true
		// case TELOPT_LFLOW:
		// 	// TODO
		case TELOPT_NAWS:
			_, _ = t.write([]byte{IAC, DO, opt}) // FIXME
		// case TELOPT_STATUS:
		// 	t.write([]byte{IAC, DO, opt})
		case TELOPT_TSPEED:
			_, _ = t.write([]byte{IAC, SB, TELOPT_TSPEED, TELQUAL_SEND, IAC, SE})
		// case TELOPT_TTYPE:
		// 	t.write([]byte{IAC, SB, TELOPT_TTYPE, TELQUAL_SEND, IAC, SE})
		default:
			_, _ = t.write([]byte{IAC, DONT, opt}) // FIXME
		}
	case WONT:
		switch opt {
		case TELOPT_ECHO:
			// t.echoEnabled = false
		default:
			_, _ = t.write([]byte{IAC, DONT, opt}) // FIXME
		}
	}
}

func (t *Conn) handleSubnegotiation(opt byte, data []byte) {
	switch opt {
	case TELOPT_NAWS:
		if len(data) >= 4 {
			width := uint16(data[0])<<8 | uint16(data[1])
			height := uint16(data[2])<<8 | uint16(data[3])
			if t.OnWindowSizeChange != nil {
				t.OnWindowSizeChange(height, width)
			}
		}
	case TELOPT_TSPEED:
		if len(data) > 1 && data[0] == TELQUAL_IS {
			speed := string(data[1:])
			before, after, found := strings.Cut(speed, ",")
			if !found {
				break
			}
			transmit, err := strconv.Atoi(before)
			if err != nil {
				break
			}
			receive, err := strconv.Atoi(after)
			if err != nil {
				break
			}
			if t.OnTerminalSpeed != nil {
				t.OnTerminalSpeed(transmit, receive)
			}
		}
		// case TELOPT_TTYPE:
		// 	if len(data) > 1 && data[0] == TELQUAL_IS {
		// 		termType := string(data[1:])
		// 		log.Printf("terminal type is %q\n", termType)
		// 	}
	}
}

func (t *Conn) write(p []byte) (int, error) {
	t.writeMu.Lock()
	defer t.writeMu.Unlock()
	return t.Conn.Write(p)
}
