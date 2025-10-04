package login

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/nosborn/federation-1999/internal/config"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/nosborn/federation-1999/pkg/ibgames/auth"
	"github.com/nosborn/federation-1999/pkg/ibgames/db"
)

const banner = `

Welcome to ibgames, the home of Federation.

Please type your ibgames account ID at the login prompt, then type in your
password.

If you do not have an ibgames account, you will need to go to our web page
at <URL:http://federation-1999.fly.dev/index.html> to set an account up.

`

const (
	msgNoCredit    = "\nSorry, your ibgames account has no time credits so you\ncannot use it to play Federation.\n\n" // #nosec G101
	msgSuspended   = "\nYour account has been suspended and you cannot use it\nto play Federation.\n\n"
	msgSystemError = "System error. Unable to authenticate.\n\n"
)

func Login(remoteHostname string) {
	timeout := SetTimeout()

	signal.Ignore(syscall.SIGINT, syscall.SIGQUIT)

	// Set up some sane terminal modes.
	if err := SetInputTermios(int(os.Stdin.Fd())); err != nil {
		log.Fatalf("setInputTermios: %v", err)
	}
	if err := SetOutputTermios(int(os.Stdout.Fd())); err != nil {
		log.Fatalf("setOutputTermios: %v", err)
	}

	fmt.Print(banner + "\n")

	if db.SetEnvironment() == -1 {
		log.Printf("db.SetEnvironment() failed") // FIXME
		disconnect(msgSystemError)
	}

	cnt := 0
	var session auth.Session

	reader := bufio.NewReader(os.Stdin)

	for {
		name, err := ReadString("ibgames login: ", reader)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			log.Panic("login.ReadString: ", err)
		}
		if name == "" {
			continue
		}
		if len(name) > auth.NameSize {
			os.Exit(1)
		}

		password, err := ReadPassword("Password: ", reader)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			log.Panic("login.ReadPassword: ", err)
		}

		session.UID = 0

		if len(name) <= auth.NameSize && password != "" && len(password) <= auth.PasswordSize {
			err = db.Connect(false)
			if err != nil {
				log.Printf("db.Connect() failed") // FIXME
				disconnect(msgSystemError)
			}

			result := auth.Login(name, password, "TODO", &session)
			switch result {
			case auth.LoginOK:
				//
			case auth.LoginIncorrect:
				session.UID = 0
			case auth.LoginNoCredit:
				db.Commit()
				db.Disconnect()
				disconnect(msgNoCredit)
			case auth.LoginSuspended:
				db.Commit()
				db.Disconnect()
				disconnect(msgSuspended)
			default:
				db.Disconnect()
				log.Printf("auth.Login() failed") // FIXME
				disconnect(msgSystemError)
			}

			if err := db.Commit(); err != nil {
				log.Printf("db.Commit() failed") // FIXME
				disconnect(msgSystemError)
			}
			if err := db.Disconnect(); err != nil {
				log.Printf("db.Disconnect() failed") // FIXME
				disconnect(msgSystemError)
			}
		}

		if session.UID >= ibgames.MinAccountID && session.UID <= ibgames.MaxAccountID {
			break
		}

		fmt.Println("Login incorrect")

		// We allow up to 'retry' (10) tries, but after 'backoff' (3)
		// we start backing off. The alarm timeout will probably fire
		// before this reaches max attempts.
		cnt++
		if cnt > 3 /*BACKOFF*/ {
			if cnt >= 10 /*RETRIES*/ {
				sleepExit(1 /*EXIT_FAILURE*/)
			}
			time.Sleep(time.Duration((cnt-3 /*BACKOFF*/)*5) * time.Second)
		}
	}

	// Committed to login -- turn off timeout.
	timeout.Cancel()

	// Disconnect from the database.
	if err := db.Exit(); err != nil {
		sleepExit(1)
	}

	// Preserve TERM from the current environment.
	term := os.Getenv("TERM")
	if term == "" {
		term = "dumb"
	}

	// Set up the new environment.
	env := []string{
		"TERM=" + term,
	}

	if session.SLogin != "" {
		doLastLog("  successful", session.SLogin, "*")
	}
	if session.ULogin != "" {
		doLastLog("unsuccessful", session.ULogin, "*")
	}

	// Check for disabled logins.
	checkNoLogin()

	// Show any message-of-the-day file.
	// doFile(MOTDFILE)

	// Reset signal handlers to defaults.
	// signal(SIGALRM, SIG_DFL)
	// signal(SIGINT, SIG_DFL)
	// signal(SIGQUIT, SIG_DFL)

	// err = a.GoodLogin(ctx)
	// if err != nil {
	// 	log.Printf("%v", err)
	// }

	// Launch the 'shell' program (perivale).
	args := []string{filepath.Join(config.BinDir, "perivale"), fmt.Sprintf("%d", session.UID), remoteHostname, "0"}
	// #nosec G204 -- remoteHostname comes from reverse DNS lookup, content is constrained
	err := syscall.Exec(args[0], args, env)
	if err != nil {
		panic(err)
	}
}

func checkNoLogin() {
	// if doFile(NOLOGINFILE) {
	// 	sleepExit(0 /*EXIT_SUCCESS*/)
	// }
}

// func doFile(path string)

func disconnect(message string) {
	fmt.Print(message)
	time.Sleep(1 * time.Second)
	os.Exit(0)
}

func doLastLog(narrative string, dtime string, ip string) {
	fmt.Printf("Last %s login: %s", narrative, dtime)
	if ip != "" {
		names, _ := net.LookupAddr(ip) //nolint:noctx
		// if err != nil {
		// 	log.Printf("doLastLog: %v", err)
		// }
		if len(names) == 0 {
			fmt.Printf(" from %s", ip)
		} else {
			fmt.Printf(" from %s", names[0])
		}
	}
	fmt.Print("\n")
}

// func getName() string
// func getPassword() string
// func setEnv(name, value string, clobber bool)

func sleepExit(status int) {
	time.Sleep(5 * time.Second)
	os.Exit(status)
}

// func timedOut(_ int)
