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
	"github.com/nosborn/ibgames-1999"
	"github.com/nosborn/ibgames-1999/auth"
	"github.com/nosborn/ibgames-1999/db"
)

const (
	msgNoCredit    = "\nSorry, your ibgames account has no time credits so you\ncannot use it to play Federation.\n\n"
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

	err := db.Connect(config.DatabasePath, false)
	if err != nil {
		log.Panic("failed to open database: ", err)
	}
	defer db.Exit()

	cnt := 0
	// ctx := context.Background()

	// fmt.Print(text.Msg(text.LoginMOTD))

	var session auth.Session

	reader := bufio.NewReader(os.Stdin)

	for {
		username, err := ReadString("fed99 login: ", reader)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			log.Panic("login.ReadString: ", err)
		}
		if username == "" {
			continue
		}
		// if strings.EqualFold(username, "new") {
		// 	session.UID, err = signup(reader)
		// 	if err != nil {
		// 		if err == io.EOF {
		// 			log.Printf("login: EOF from Signup") // FIXME
		// 			os.Exit(0)
		// 		}
		// 		fmt.Printf("Name in use.\n")               // FIXME
		// 		log.Printf("login: continue after Signup") // FIXME
		// 		continue
		// 	}
		// 	fmt.Printf("OK\n")                      // FIXME
		// 	log.Printf("login: break after Signup") // FIXME
		// 	break
		// }

		password, err := ReadPassword("Password: ", reader)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			log.Panic("login.ReadPassword: ", err)
		}

		result := auth.Login(username, password, "TODO", &session)
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
			disconnect(msgSystemError)
		}

		if err := db.Commit(); err != nil {
			disconnect(msgSystemError)
		}
		if err := db.Disconnect(); err != nil {
			disconnect(msgSystemError)
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

		// a, err = billing.GetAccountByName(ctx, username)
		// if err != nil {
		// 	if err == sql.ErrNoRows {
		// 		fmt.Println("Login incorrect")
		// 		continue
		// 	}
		// 	fmt.Printf("Database error: %v\n", err)
		// 	continue
		// }
		// if a == nil || bcrypt.CompareHashAndPassword([]byte(a.Encrypt), []byte(password)) != nil {
		// 	if a != nil && a.Encrypt == "*" && CheckHoneypot(username, password) {
		// 		log.Printf("Honeypot login for %q", username)
		// 		time.Sleep(time.Duration(5+rand.IntN(5)) * time.Second)
		// 		break
		// 	}
		// 	log.Printf("Wrong password for %s", username)
		//
		// 	err = a.BadLogin(ctx)
		// 	if err != nil {
		// 		log.Printf("%v", err)
		// 	}
		// 	fmt.Println("Login incorrect")
		//
		// 	// We allow up to 'retry' (10) tries, but after
		// 	// 'backoff' (3) we start backing off. The alarm
		// 	// timeout will probably fire before this reaches max
		// 	// attempts.
		// 	cnt++
		// 	if cnt > 3 /*BACKOFF*/ {
		// 		if cnt >= 10 /*RETRIES*/ {
		// 			sleepExit(1 /*EXIT_FAILURE*/)
		// 		}
		// 		time.Sleep(time.Duration((cnt-3 /*BACKOFF*/)*5) * time.Second)
		// 	}
		//
		// 	continue
		// }
		// break
	}
	log.Print("login: committed to login")

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
	args := []string{"perivale", fmt.Sprintf("%d", session.UID), remoteHostname, "0"}
	err = syscall.Exec(filepath.Join(config.BinDir, "perivale"), args, env)
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
		names, err := net.LookupAddr(ip) //nolint:noctx
		if err != nil {
			log.Printf("doLastLog: %v", err)
		}
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
