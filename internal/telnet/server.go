package telnet

import (
	"context"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"

	"github.com/nosborn/federation-1999/internal/config"
	"github.com/nosborn/federation-1999/internal/ioctl"
	"github.com/nosborn/federation-1999/internal/modem"
	proxyproto "github.com/pires/go-proxyproto"
	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

func ListenAndServe(addr string, proxyProto bool) error {
	if addr == "" {
		addr = ":telnet"
	}
	ln, err := net.Listen("tcp", addr) //nolint:noctx
	if err != nil {
		return err
	}
	if proxyProto {
		ln = &proxyproto.Listener{Listener: ln}
	}
	return serve(ln)
}

func serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		log.Printf("connection from %s on %s", conn.RemoteAddr().String(), conn.LocalAddr().String())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	modemConn := modem.NewV34_28K(conn) // Emulate 28.8K V.34 modem
	defer func() {
		_ = modemConn.Close()
	}()

	tconn := NewConn(modemConn)
	defer func() {
		_ = tconn.Close()
	}()

	remoteIP, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	handler(tconn, remoteIP)
}

func handler(conn *Conn, remoteIP string) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}()

	master, slave, err := termios.Pty()
	if err != nil {
		log.Printf("pty error: %v", err)
		return
	}
	defer func() {
		if err := master.Close(); err != nil {
			log.Printf("error closing master: %v", err)
		}
	}()
	unix.CloseOnExec(int(master.Fd()))

	conn.OnEchoChange = func(enabled bool) {
		termios, err := unix.IoctlGetTermios(int(master.Fd()), ioctl.GetTerminalAttrs)
		if err != nil {
			log.Printf("echo change: get termios error: %v", err)
			return
		}
		if enabled {
			termios.Lflag |= unix.ECHO
		} else {
			termios.Lflag &^= unix.ECHO
		}
		if err := unix.IoctlSetTermios(int(master.Fd()), ioctl.SetTerminalAttrsNow, termios); err != nil {
			log.Printf("echo change: set termios error: %v", err)
		}
	}

	conn.OnTerminalSpeed = func(transmit, receive int) {
		// log.Printf("terminal speed is %d,%d\n", transmit, receive)
		// TODO
	}

	conn.OnWindowSizeChange = func(rows uint16, cols uint16) {
		ws := &unix.Winsize{Row: rows, Col: cols}
		if err := unix.IoctlSetWinsize(int(master.Fd()), unix.TIOCSWINSZ, ws); err != nil {
			log.Printf("resize error: %v", err)
		}
	}

	cmd := exec.CommandContext(context.TODO(), config.BinDir+"/login", "-h", remoteIP) //nolint:gosec
	cmd.Stdin = slave
	cmd.Stdout = slave
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid:  true,
		Setctty: true,
	}

	if unix.Getuid() == 0 {
		const unprivilegedUser = "fed"
		u, err := user.Lookup(unprivilegedUser)
		if err != nil {
			log.Panicf("failed to look up user %q: %v", unprivilegedUser, err)
		}
		uid, err := strconv.ParseUint(u.Uid, 10, 32)
		if err != nil {
			log.Panicf("failed to parse uid %q for user %q: %v", u.Uid, unprivilegedUser, err)
		}
		gid, err := strconv.ParseUint(u.Gid, 10, 32)
		if err != nil {
			log.Panicf("failed to parse gid %q for user %q: %v", u.Gid, unprivilegedUser, err)
		}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
		cmd.Dir = u.HomeDir
	}

	if err := cmd.Start(); err != nil {
		log.Printf("failed to start login: %v", err)
		return
	}

	if err := slave.Close(); err != nil {
		log.Printf("error closing slave pty in parent: %v", err)
	}

	connToPtyDone := make(chan error, 1)
	ptyToConnDone := make(chan error, 1)
	childDone := make(chan error, 1)

	go func() {
		_, err := master.ReadFrom(conn)
		connToPtyDone <- err
	}()

	go func() {
		_, err := master.WriteTo(conn)
		ptyToConnDone <- err
	}()

	go func() {
		err := cmd.Wait()
		// log.Printf("child process finished: %#v", err)
		childDone <- err
	}()

	select {
	case err := <-connToPtyDone:
		_ = err
	case err := <-ptyToConnDone:
		_ = err
	case err := <-childDone:
		_ = err
	}

	if err := conn.Close(); err != nil {
		log.Printf("error closing network connection: %v", err)
	}
	if err := master.Close(); err != nil {
		log.Printf("error closing PTY master: %v", err)
	}
}
