package login

import (
	"bufio"
	"fmt"
	"log"
	"regexp"

	"github.com/nosborn/ibgames-1999"
	"github.com/nosborn/ibgames-1999/auth"
	"github.com/nosborn/ibgames-1999/db"
)

var alphaOnly = regexp.MustCompile(`^[A-Za-z]+$`)

func signup(reader *bufio.Reader) (ibgames.AccountID, error) {
	for {
		name, err := ReadString("Choose new username: ", reader)
		if err != nil {
			return 0, err
		}
		if len(name) < 3 || len(name) > 15 {
			fmt.Print("Username must be between 3 and 15 characters in length.\n")
			continue
		}
		if !alphaOnly.MatchString(name) {
			fmt.Println("Username must contain only alphabetic characters.")
			continue
		}

		password1, err := ReadPassword("New password: ", reader)
		if err != nil {
			return 0, err
		}
		password2, err := ReadPassword("Retype new password: ", reader)
		if err != nil {
			return 0, err
		}
		if password1 != password2 {
			fmt.Println("Passwords do not match. Please try again.")
			continue
		}
		encrypt, err := auth.PasswordHash(password1)
		if err != nil {
			return 0, err
		}

		log.Printf("creating new account %v", name)
		const insertStmt = `
			INSERT INTO accounts (name, name_key, encrypt, minutes)
                	VALUES(?, ?, ?, ?)`
		result, err := db.Exec(insertStmt, name, auth.UniqueName(name), encrypt, 120)
		if err != nil {
			log.Printf("login.signup: db.Exec: %v", err)
			return 0, err
		}
		uid, err := result.LastInsertId()
		if err != nil {
			log.Printf("login.signup: result.LastInsertId: %v", err)
			return 0, err
		}
		log.Printf("login.signup: created uid = %d", uid)
		return ibgames.AccountID(uid), nil

		// name = strings.ToLower(name)
		// a, err := billing.CreateAccount(ctx, name, hash)
		// if err != nil {
		// 	if err == billing.ErrAccountExists {
		// 		fmt.Println("That name isn't available. Please choose another.")
		// 		continue
		// 	}
		// 	log.Printf("billing.CreateAccount: %#v\n", err)
		// 	return nil, err
		// }
		//
		// return a, nil
	}
}
