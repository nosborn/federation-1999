package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

var (
	listener   *net.UnixListener
	sessions   map[ibgames.AccountID]*Session
	sessionsMu sync.Mutex
)

func init() {
	sessions = make(map[ibgames.AccountID]*Session)
}

// StartListener starts the Unix socket listener (caller must hold global lock)
func StartListener(path string) error {
	if listener != nil {
		return fmt.Errorf("listener already started")
	}

	_ = os.Remove(path)

	addr, err := net.ResolveUnixAddr("unix", path)
	if err != nil {
		return err
	}
	ln, err := net.ListenUnix("unix", addr)
	if err != nil {
		return err
	}
	listener = ln

	go acceptLoop()
	return nil
}

// StopListener stops the Unix socket listener (caller must hold global lock, safe to call multiple times)
func StopListener() {
	if listener == nil {
		return
	}

	if err := listener.Close(); err != nil {
		log.Printf("error closing listener: %v", err)
	}
	listener = nil
}

func acceptLoop() {
	for {
		conn, err := listener.AcceptUnix()
		if err != nil {
			log.Printf("accept error: %v", err)
			return
		}
		// debug.Trace("Accepted fd %d", fd)
		go handleConnection(conn)
	}
}

// Legacy function for backward compatibility
func ListenAndServe(path string) error {
	if err := StartListener(path); err != nil {
		return err
	}
	// Block forever since the old function was blocking
	select {}
}

func handleConnection(conn *net.UnixConn) {
	monitoring.ConnectionsTotal.WithLabelValues("fedtpd").Inc()
	monitoring.ConnectionsCurrent.WithLabelValues("fedtpd").Inc()
	defer monitoring.ConnectionsCurrent.WithLabelValues("fedtpd").Dec()

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("error closing conn: %v", err)
		}
	}()

	// r := make(chan string)
	// w := make(chan relay.Output)

	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	if len(parts) != 4 {
		log.Printf("Bad login line: \"%s\"", line)
		return
	}
	uid, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil || uid < ibgames.MinAccountID || uid > ibgames.MaxAccountID {
		log.Printf("Bad login line: \"%s\"", line)
		return
	}
	robobod, err := strconv.Atoi(parts[3])
	if err != nil || (robobod != 0 && robobod != 1) {
		log.Printf("Bad login line: \"%s\"", line)
		return
	}
	remoteHost := parts[1]
	remoteHostname := parts[2]

	s, err := beginSession(conn, ibgames.AccountID(uid), remoteHost, remoteHostname, robobod)
	if err != nil {
		log.Printf("beginSession: %v", err)
		return
	}
	// if err := s.Run(); err != nil {
	// 	if err != io.EOF {
	// 		log.Printf("s.Run: %#v", err)
	// 	}
	// }
	s.Run()
	endSession(s)
}

func beginSession(conn *net.UnixConn, uid ibgames.AccountID, remoteHost, remoteHostname string, robobod int) (*Session, error) {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()

	if session, exists := sessions[uid]; exists {
		log.Printf("Second login for #%d", uid)
		session.Destroy()
		delete(sessions, uid)
	}
	session, err := NewSession(conn, uid, remoteHost, remoteHostname, robobod)
	if err != nil {
		return nil, err
	}
	sessions[uid] = session
	return session, nil
}

func endSession(session *Session) {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()

	session.Destroy()
	if sessions[session.UID()] == session {
		delete(sessions, session.UID())
	}
}
