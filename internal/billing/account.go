package billing

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var ErrAccountExists = errors.New("account already exists")

type Account struct {
	Name          string     `json:"name"`
	UID           int64      `json:"-"`
	Encrypt       string     `json:"encrypt"`
	SChange       time.Time  `json:"schange"`
	AcctExpire    *time.Time `json:"acct_expire,omitempty"`
	SLogin        *time.Time `json:"slogin,omitempty"`
	ULogin        *time.Time `json:"ulogin,omitempty"`
	SucIP         string     `json:"sucip,omitempty"`
	NUnsucLog     int        `json:"nunsuclog,omitempty"`
	UnsucIP       string     `json:"unsucip,omitempty"`
	Email         string     `json:"email,omitempty"`
	Signup        time.Time  `json:"signup"`
	Status        string     `json:"status"`        // FIXME
	Complimentary string     `json:"complimentary"` // FIXME
	Minutes       string     `json:"minutes,omitempty"`
	AdminUID      int64      `json:"admin_uid,omitempty"`
	AdminDate     time.Time  `json:"admin_date,omitempty"`
	BulkEmail     string     `json:"bulk_email,omitempty"` // FIXME
	BadCheck      string     `json:"bad_check,omitempty"`  // FIXME
	BuddyUID      int64      `json:"buddy_uid,omitempty"`
	BuddyPayment  string     `json:"buddy_payment,omitempty"` // FIXME
}

func CreateAccount(ctx context.Context, name string, encrypt string) (*Account, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				//
			}
		}
	}()

	var existingID int64
	err = tx.QueryRowContext(ctx, "SELECT id FROM accounts WHERE name = ?", name).Scan(&existingID)
	if err == nil {
		return nil, ErrAccountExists
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	const query = `
		INSERT INTO accounts (data)
		VALUES (?)`

	a := &Account{
		Name:          name,
		Encrypt:       encrypt,
		SChange:       time.Now(),
		Signup:        time.Now(),
		Status:        "A", // FIXME
		Complimentary: "N", // FIXME
	}

	data, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	result, err := tx.ExecContext(ctx, query, string(data))
	if err != nil {
		return nil, err
	}

	a.UID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return a, nil
}

func GetAccountByName(ctx context.Context, name string) (*Account, error) {
	const query = `
		SELECT uid, data
		FROM accounts
		WHERE name = ?`

	var a Account
	var data string

	err := db.QueryRowContext(ctx, query, name).Scan(&a.UID, &data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func GetAccountByUID(ctx context.Context, uid uint32) (*Account, error) {
	const query = `
		SELECT uid, data
		FROM accounts
		WHERE uid = ?`

	var a Account
	var data string

	err := db.QueryRowContext(ctx, query, uid).Scan(&a.UID, &data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *Account) BadLogin(ctx context.Context) error {
	const query = `
		UPDATE accounts
		SET data = json_set(data, '$.nunsuclog', ?, '$.ulogin', ?)
		WHERE uid = ?`

	// now := time.Now().UTC()
	a.NUnsucLog++
	// a.ULogin = &now

	_, err := db.ExecContext(ctx, query, a.NUnsucLog, a.ULogin, a.UID)
	if err != nil {
		return fmt.Errorf("account.BadLogin: %v", err)
	}
	return nil
}

func (a *Account) GoodLogin(ctx context.Context) error {
	// TODO

	// const query = `
	// 	UPDATE accounts
	// 	SET data = json_set(data, '$.slogin', ?)
	// 	WHERE uid = ?`
	//
	// now := time.Now().UTC()
	// a.LastGoodLogin = &now
	//
	// _, err := db.Exec(query, a.LastGoodLogin, a.UID)
	// if err != nil {
	// 	return fmt.Errorf("account.GoodLogin: %v", err)
	// }
	return nil
}
