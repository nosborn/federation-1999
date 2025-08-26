package login

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/nosborn/federation-1999/internal/billing"
	"github.com/nosborn/federation-1999/internal/config"
	"golang.org/x/crypto/bcrypt"
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

	err := billing.OpenDB(config.DatabasePath)
	if err != nil {
		log.Panic("failed to open database: ", err)
	}
	defer billing.CloseDB()

	cnt := 0
	ctx := context.Background()

	// fmt.Print(text.Msg(text.LoginMOTD))

	var a *billing.Account

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
		if strings.EqualFold(username, "new") {
			a, err = Signup(ctx, reader)
			if err != nil {
				if err == io.EOF {
					os.Exit(0)
				}
				continue
			}
			break
		}

		password, err := ReadPassword("Password: ", reader)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			log.Panic("login.ReadPassword: ", err)
		}
		a, err = billing.GetAccountByName(ctx, username)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Login incorrect")
				continue
			}
			fmt.Printf("Database error: %v\n", err)
			continue
		}
		if a == nil || bcrypt.CompareHashAndPassword([]byte(a.Encrypt), []byte(password)) != nil {
			if a != nil && a.Encrypt == "*" && CheckHoneypot(username, password) {
				log.Printf("Honeypot login for %q", username)
				time.Sleep(time.Duration(5+rand.IntN(5)) * time.Second)
				break
			}
			log.Printf("Wrong password for %s", username)

			err = a.BadLogin(ctx)
			if err != nil {
				log.Printf("%v", err)
			}
			fmt.Println("Login incorrect")

			// We allow up to 'retry' (10) tries, but after
			// 'backoff' (3) we start backing off. The alarm
			// timeout will probably fire before this reaches max
			// attempts.
			cnt++
			if cnt > 3 /*BACKOFF*/ {
				if cnt >= 10 /*RETRIES*/ {
					sleepExit(1 /*EXIT_FAILURE*/)
				}
				time.Sleep(time.Duration((cnt-3 /*BACKOFF*/)*5) * time.Second)
			}

			continue
		}
		break
	}

	// Committed to login -- turn off timeout.
	timeout.Cancel()

	// Disconnect from the database.
	_ = billing.CloseDB()

	// Preserve TERM from the current environment.
	term := os.Getenv("TERM")
	if term == "" {
		term = "dumb"
	}

	// Set up the new environment.
	env := []string{
		"TERM=" + term,
	}

	if a.SLogin != nil {
		doLastLog("  successful", a.SLogin.String(), "*")
	}
	if a.ULogin != nil {
		doLastLog("unsuccessful", a.ULogin.String(), "*")
	}

	// Check for disabled logins.
	checkNoLogin()

	// Show any message-of-the-day file.
	// doFile(MOTDFILE)

	// Reset signal handlers to defaults.
	// signal(SIGALRM, SIG_DFL)
	// signal(SIGINT, SIG_DFL)
	// signal(SIGQUIT, SIG_DFL)

	err = a.GoodLogin(ctx)
	if err != nil {
		log.Printf("%v", err)
	}

	// Launch the 'shell' program (perivale).
	args := []string{"perivale", strconv.FormatInt(a.UID, 10), remoteHostname, "0"}
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
// func disconnect(message string)

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
