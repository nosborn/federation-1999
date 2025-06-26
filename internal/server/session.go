package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"golang.org/x/sys/unix"
)

type SessionState int

const (
	SessionContinue SessionState = iota
	SessionQuit
	SessionSwitchToPlay
	SessionSwitchToShipOrder
	SessionSwitchToTour
	SessionSwitchToWorkbench
)

type SessionDriver interface {
	Dispatch(string) bool
	Destroy()
}

type Session struct {
	conn           *net.UnixConn
	driver         SessionDriver
	isRoboBod      bool
	logFile        *os.File
	outputMu       sync.Mutex
	player         *Player
	remoteHost     string
	remoteHostname string
	scanner        *bufio.Scanner
	state          SessionState
	timeout        int32      // in minutes
	_              time.Timer // TODO: timeoutTimer
	uid            ibgames.AccountID
}

func NewSession(conn *net.UnixConn, uid ibgames.AccountID, remoteHost, remoteHostname string, robobod int) (*Session, error) {
	log.Printf("Login #%d from %s", uid, remoteHost)

	s := &Session{
		conn:           conn,
		remoteHost:     remoteHost,
		remoteHostname: remoteHostname,
		scanner:        bufio.NewScanner(conn),
		state:          SessionContinue,
		timeout:        20, // default is 20 minutes
		uid:            uid,
	}
	if robobod == 1 {
		s.isRoboBod = true
	}

	// Start session logging immediately if we're testing.
	if global.TestFeaturesEnabled {
		s.Logging(true)
	}

	// Tcl_CreateChannelHandler(m_channel, TCL_READABLE, channelHandler, this);

	if err := s.Output(text.Msg(text.Banner)); err != nil {
		log.Printf("s.Output: %v", err)
		_ = s.conn.Close()
		return nil, err
	}

	// Start the timeout timer.
	// m_timeoutTimer = Tcl_CreateTimerHandler(m_timeout,
	//                                           timeoutTimerHandler,
	//                                           this);
	// s.timeoutTimer = ... TODO

	var ok bool
	s.player, ok = FindPlayerByID(uid)
	if !ok {
		s.beginSetup()
		return s, nil
	}
	if s.player.Session() != nil {
		log.Panic("s.player.Session() is not nil")
	}

	// Start logging for staff accounts.
	if s.player.Rank() >= model.RankSenator && s.player.Rank() < model.RankDeity {
		s.Logging(true)
	} else if s.player.IsPromoCharacter() {
		s.Logging(true)
	}

	// Say hello...
	err := s.Output(text.Msg(text.Linking))
	if err != nil {
		log.Printf("s.Output: %v", err)
		_ = s.conn.Close()
		return nil, err
	}
	s.beginPlay()
	return s, nil
}

func (s *Session) BilledTime() int {
	// TODO
	return 0
}

func (s *Session) BillingTick() bool {
	// TODO
	return true
}

func (s *Session) Destroy() {
	log.Printf("Session.Destroy(%d)", s.uid)

	// Kill the driver.
	if s.driver != nil {
		s.driver.Destroy()
		s.driver = nil
	}

	// End the billing session.
	// charge := billing.EndSession()
	// if charge > 0 {
	// 	output(message(mnBillingSummary, charge, (charge) == 1 ? "" : "s"));
	// }

	// Turn off logging.
	s.Logging(false)

	// Detach the player record.
	if s.player != nil { //nolint:staticcheck // SA9003: empty branch
		// debug.Check(m_player->m_session == NULL);
	}
	s.player = nil

	// Kill the timeout timer.
	// Tcl_DeleteTimerHandler(m_timeoutTimer);

	//
	if err := s.conn.Close(); err != nil {
		log.Printf("Session.Destroy(%d): s.conn.Close: %v", s.uid, err)
	}
	// s.conn = nil

	//
	log.Printf("Logout #%d", s.uid)
}

func (s *Session) EndSetup(player *Player) error {
	if player == nil {
		log.Panicf("player is nil")
	}
	if player.UID() != s.uid {
		log.Panic("player.UID does not match s.uid")
	}

	// Stash the player pointer.
	s.player = player
	Players[s.player.Name()] = s.player

	// Resume charging.
	s.StartBilling()

	// Say hello...
	err := s.Output(text.Msg(text.Linking))
	if err != nil {
		return err
	}

	s.beginPlay()
	return nil
}

func (s *Session) IsRoboBod() bool {
	return s.isRoboBod
}

func (s *Session) Output(text string) error {
	s.outputMu.Lock()
	defer s.outputMu.Unlock()

	if s.logFile != nil {
		_, _ = fmt.Fprint(s.logFile, text)
	}

	if _, err := s.conn.Write([]byte(text)); err != nil {
		log.Printf("s.conn.Write: %v", err)
		_ = s.conn.Close()
		return err
	}
	return nil
}

// func (s *Session) Player() *Player {
// 	return s.player
// }

func (s *Session) Quit() {
	s.state = SessionQuit
}

func (s *Session) State() SessionState {
	return s.state
}

func (s *Session) RemoteHost() string {
	return s.remoteHost
}

func (s *Session) RemoteHostname() string {
	return s.remoteHostname
}

func (s *Session) Run() {
	for s.channelHandler() {
	}
}

func (s *Session) channelHandler() bool {
	defer database.CommitDatabase()

	if !s.channelProc() {
		endSession(s) // (s.uid)
		return false
	}
	return true
}

func (s *Session) channelProc() bool {
	if !s.scanner.Scan() {
		_ = s.scanner.Err()
		// if (Tcl_InputBlocked(m_channel)) {
		// 	return true;
		// }
		// if err == nil {
		// 	// debug.Trace("EOF [%d]", s.uid)
		// 	err = io.EOF
		// } else {
		// 	log.Printf("s.scanner.Scan: %v", err)
		// }
		// _ = s.conn.Close()
		return false
	}
	line := s.scanner.Text()
	log.Printf("Session.Run(%d): line=\"%v\"", s.uid, line)

	// Tcl_DStringAppend(&lineRead, "\n", 1);
	//
	// Tcl_DeleteTimerHandler(m_timeoutTimer);
	// m_timeoutTimer = NULL;

	// FIXME: needs to use s.outputMu
	_, err := s.conn.Write([]byte{model.DLE, model.LeSpy, 0, model.DLE, model.LeAck})
	if err != nil {
		log.Printf("s.conn.Write: %v", err)
		return false
	}

	if !s.driver.Dispatch(line) {
		debug.Trace("Run: Failed session [%d] after dispatch", s.uid)
		return false // io.EOF // FIXME: made up error!
	}

	// Tcl_Flush(m_channel);
	//
	// if (m_timeout > 0) {
	// 	m_timeoutTimer = Tcl_CreateTimerHandler(m_timeout, timeoutTimerHandler, this)
	// }

	switch s.state {
	case SessionContinue:
		// do nothing
	case SessionQuit:
		return false
	case SessionSwitchToPlay:
		s.beginPlay()
	case SessionSwitchToShipOrder:
		s.beginShipOrder()
	case SessionSwitchToTour:
		s.beginTour()
	case SessionSwitchToWorkbench:
		s.beginWorkbench()
	}

	return true
}

func (s *Session) SwitchToPlay() {
	s.state = SessionSwitchToPlay
}

func (s *Session) SwitchToTour() {
	s.state = SessionSwitchToTour
}

func (s *Session) SwitchToWorkbench() {
	s.state = SessionSwitchToWorkbench
}

func (s *Session) UID() ibgames.AccountID {
	return s.uid
}

func (s *Session) beginPlay() {
	if s.uid < ibgames.MinAccountID || s.uid > ibgames.MaxAccountID {
		log.Panicf("s.uid out of range")
	}
	if s.player == nil {
		log.Panicf("s.player is nil")
	}
	if s.player.Session() != nil {
		log.Panicf("s.player.Session() is not nil")
	}

	// Switch to the play driver.

	if s.driver != nil {
		s.driver.Destroy()
		s.driver = nil
	}

	s.driver = NewPlayDriver(s, s.player)
	if s.driver == nil {
		log.Panicf("s.driver is nil")
	}

	s.state = SessionContinue
}

func (s *Session) beginSetup() {
	if s.uid < ibgames.MinAccountID || s.uid > ibgames.MaxAccountID {
		log.Panicf("s.uid out of range")
	}
	if s.player != nil {
		log.Panicf("s.player is not nil")
	}

	// Don't charge for setting up a persona.
	s.StopBilling()

	// Start the setup driver.
	if s.driver != nil {
		log.Panicf("s.driver is not nil")
	}
	s.driver = NewSetupDriver(s, s.uid)
	if s.driver == nil {
		log.Panicf("s.driver is nil")
	}

	s.state = SessionContinue
}

func (s *Session) beginShipOrder() {
	if s.uid < ibgames.MinAccountID || s.uid > ibgames.MaxAccountID {
		log.Panicf("s.uid out of range")
	}
	if s.player == nil {
		log.Panicf("s.player is nil")
	}

	if s.driver == nil {
		log.Panicf("s.driver is nil")
	}
	s.driver.Destroy()
	s.driver = nil
	if s.player.Session() != nil {
		log.Panicf("s.player.Session() is not nil")
	}

	s.driver = NewShipOrderDriver(s, s.player)
	if s.driver == nil {
		log.Panicf("s.driver is nil")
	}

	s.state = SessionContinue
}

func (s *Session) beginTour() {
	if s.uid < ibgames.MinAccountID || s.uid > ibgames.MaxAccountID {
		log.Panicf("s.uid out of range")
	}
	if s.player == nil {
		log.Panicf("s.player is nil")
	}

	if s.driver == nil {
		log.Panicf("s.driver is nil")
	}
	s.driver.Destroy()
	s.driver = nil
	if s.player.Session() != nil {
		log.Panicf("s.player.Session() is not nil")
	}

	s.driver = NewTourDriver(s, s.player)
	if s.driver == nil {
		log.Panicf("s.driver is nil")
	}

	s.state = SessionContinue
}

func (s *Session) beginWorkbench() {
	if s.uid < ibgames.MinAccountID || s.uid > ibgames.MaxAccountID {
		log.Panicf("s.uid out of range")
	}
	if s.player == nil {
		log.Panicf("s.player is nil")
	}

	if s.driver == nil {
		log.Panicf("s.driver is nil")
	}
	s.driver.Destroy()
	s.driver = nil
	if s.player.Session() != nil {
		log.Panicf("s.player.Session() is not nil")
	}

	s.driver = NewWorkbenchDriver(s, s.player)
	if s.driver == nil {
		log.Panicf("s.driver is nil")
	}

	s.state = SessionContinue
}

func (s *Session) GetTimeout() int32 {
	return s.timeout
}

func (s *Session) Logging(on bool) {
	if on {
		datestamp := time.Now().Format("20060102")
		filename := fmt.Sprintf("%d-%s", s.uid, datestamp)
		pathname := filepath.Join("log", "session", filename)

		//nolint:gosec // G304: uid is validated, datestamp is controlled
		file, err := os.OpenFile(pathname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
		if err != nil {
			log.Printf("Session::logging: os.OpenFile(%s) failed: %v", pathname, err)
			return
		}
		s.logFile = file
		unix.CloseOnExec(int(s.logFile.Fd()))

		timestamp := time.Now().Format("Mon Jan 2 15:04:05 2006 MDT")
		_, _ = fmt.Fprintf(s.logFile, "==== Session log opened: %s ====\n\n", timestamp /*timestamp*/)
	} else if s.logFile != nil {
		timestamp := time.Now().Format("Mon Jan 2 15:04:05 2006 MDT")
		_, _ = fmt.Fprintf(s.logFile, "\n\n==== Session log closed: %s ====\n", timestamp /*timestamp*/)

		_ = s.logFile.Close()
		s.logFile = nil
	}
}

// func (s *Session) ReadLine(prompt string) (string, error) {
// 	if err := s.Output(prompt); err != nil {
// 		return "", err
// 	}
// 	if !s.scanner.Scan() {
// 		err := s.conn.Write([]byte{link.DLE, link.Spy, 0, link.DLE, link.Ack})
// 		if err := s.scanner.Err(); err == nil {
// 			err := io.EOF
// 		}
// 		log.Printf("scanner.Scan: %v", err)
// 		_ = s.conn.Close() // FIXME: probably don't do this here
// 		return err
// 	}
// 	line := s.scanner.Text()
// 	return strings.TrimSpace(line), nil
// }

func (s *Session) SetTimeout(minutes int32) {
	// TODO
	s.timeout = minutes
}

func (s *Session) StartBilling() {
	// TODO
}

func (s *Session) StopBilling() {
	// TODO
}

// func (s *Session) readPassword(prompt string) (string, error) {
// 	if err := s.Output(prompt, 0); err != nil {
// 		return "", nil
// 	}
// 	line := <-s.readChan
// 	return strings.TrimSpace(line), nil
// }

// func (s *Session) timeoutTimerProc() { -- TODO
// 	log.Printf("Idle timeout for #%d", s.uid)
//
// 	// TODO:
// 	// m_timeoutTimer = NULL;
// 	//
// 	// char buf[BUFSIZ];
// 	// int len = sprintf(buf, "\n\nNo input for %d minutes.\n\n", getTimeout());
// 	//
// 	// output(buf, len);
// 	//
// 	// deleteSession(m_uid);
// }
