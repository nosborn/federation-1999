package login

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/nosborn/federation-1999/internal/billing"
	"golang.org/x/crypto/bcrypt"
)

var alphaOnly = regexp.MustCompile(`^[A-Za-z]+$`)

func Signup(ctx context.Context, reader *bufio.Reader) (*billing.Account, error) {
	for {
		name, err := ReadString("Choose new username: ", reader)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		password2, err := ReadPassword("Retype new password: ", reader)
		if err != nil {
			return nil, err
		}
		if password1 != password2 {
			fmt.Println("Passwords do not match. Please try again.")
			continue
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		name = strings.ToLower(name)
		a, err := billing.CreateAccount(ctx, name, string(hash))
		if err != nil {
			if err == billing.ErrAccountExists {
				fmt.Println("That name isn't available. Please choose another.")
				continue
			}
			log.Printf("billing.CreateAccount: %#v\n", err)
			return nil, err
		}

		return a, nil
	}
}
