package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/link"
	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/server"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/pkg/ibgames/billing"
	"github.com/nosborn/federation-1999/pkg/ibgames/db"
	"golang.org/x/sys/unix"
)

const (
	lockFilePath = ".fedtpd.pid"
)

var lockFile *os.File

func main() {
	log.SetPrefix(fmt.Sprintf("%s[%d]: ", filepath.Base(os.Args[0]), os.Getpid()))
	log.SetFlags(log.Lmsgprefix)

	var debugCheck, freePeriod, debugPrecondition, debugTrace bool
	flag.BoolVar(&debugCheck, "check", false, "enable debug checks")
	flag.BoolVar(&freePeriod, "free-period", false, "free period")
	flag.BoolVar(&debugPrecondition, "precondition", false, "enable debug preconditions")
	flag.BoolVar(&global.TestFeaturesEnabled, "test-features", false, "enable test features")
	flag.BoolVar(&debugTrace, "trace", false, "enable debug tracing")
	flag.Parse()

	if debugCheck {
		debug.EnableCheck()
	}
	if debugPrecondition {
		debug.EnablePrecondition()
	}
	if debugTrace {
		debug.EnableTrace()
	}

	monitoring.StartServer(":8082")

	var err error
	lockFile, err = os.Create(lockFilePath)
	if err != nil {
		log.Fatalf("os.Create(%s) failed: %v", lockFilePath, err)
	}
	if err = syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		if errors.Is(err, syscall.EACCES) || errors.Is(err, syscall.EAGAIN) {
			log.Fatal("Process is already running")
		}
		log.Fatalf("syscall.Flock(%s) failed: %v", lockFilePath, err)
	}
	if _, err = fmt.Fprintf(lockFile, "%d\n", os.Getpid()); err != nil {
		log.Fatalf("fmt.Fprintf(%s) failed: %v", lockFilePath, err)
	}
	unix.CloseOnExec(int(lockFile.Fd()))
	defer func() { _ = lockFile.Close() }() // make sure the file isn't garbage-collected

	log.Print(text.Msg(text.Dashs))
	log.Print(text.Msg(text.Starting, getServerHostName()))

	var rlim unix.Rlimit
	if err = unix.Getrlimit(unix.RLIMIT_NOFILE, &rlim); err != nil {
		log.Printf("main: unix.Getrlimit() failed: %s", err)
		os.Exit(1)
	}
	if rlim.Cur < rlim.Max {
		rlim.Cur = rlim.Max
		if err = unix.Setrlimit(unix.RLIMIT_NOFILE, &rlim); err != nil {
			log.Printf("main: unix.Setrlimit() failed: %s", err)
			os.Exit(1)
		}
	}
	log.Printf("Open file limit is %d", rlim.Cur)

	err = db.Connect(false)
	if err != nil {
		log.Panic(err)
	}
	defer db.Exit()

	// billing.AutoCommit(true)
	billing.FreePeriod(freePeriod)

	if !server.InitializeGame(startListener) {
		log.Fatal("main: initializeGame() failed")
	}
	log.Printf("Waiting for duchies to load")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

sigLoop:
	for {
		sig := <-sigCh
		switch sig {
		case syscall.SIGTERM:
			log.Print("Received SIGTERM")
			break sigLoop
		case syscall.SIGUSR1:
			log.Print("Issuing 5 minute shutdown warning")
			time := time.Now().Format("Mon Jan _2 15:04:05 2006 MDT")
			_ = text.Msg(text.GameClosing_5Minutes, time)
			// TODO: send to sessions
			stopListener()
		case syscall.SIGUSR2:
			log.Print("Issuing 1 minute shutdown warning")
			time := time.Now().Format("Mon Jan _2 15:04:05 2006 MDT")
			_ = text.Msg(text.GameClosing_1Minute, time)
			// TODO: send to sessions
			stopListener()
		}
	}
	database.CommitDatabase()
}

func getServerHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("getServerHostName: os.Hostname() failed")
		hostname = "UNKNOWN"
	}
	return hostname
}

func startListener() {
	// global.Lock()
	// defer global.Unlock()

	if err := server.StartListener(link.SocketPath); err != nil {
		log.Panic("server start failed: ", err)
	}
	log.Print("Listener started")
}

func stopListener() {
	global.Lock()
	defer global.Unlock()

	server.StopListener()
	log.Print("Listener stopped")
}
