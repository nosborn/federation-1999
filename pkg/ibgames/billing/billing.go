package billing

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/nosborn/federation-1999/pkg/ibgames/db"
)

type Product int

const (
	NoProduct Product = iota
	Federation
	AgeOfAdventure
)

const (
	minimumCharge = 1
)

type Session struct {
	uid           ibgames.AccountID // Account being billed
	complimentary bool              // Is account complimentary?
	lastCharge    int               // Billed minutes at last database update
	sid           int32             // Key of session record
	nextWrite     int64             // Unix timestamp at which current minute ends
	lastTick      int64             // Unix timestamp this session last ticked
	ticking       bool              // Accumulate billable seconds?
	seconds       int64             // Billable seconds
}

const (
	insertQuery = `
		INSERT INTO sessions (product, uid, ip_address, minutes)
		VALUES(%d, ?, ?, ?)`
	selectQuery = `
		SELECT complimentary
		FROM accounts
		WHERE uid = ?`
	update1Query = `
		UPDATE accounts
		SET minutes = minutes - ?
		WHERE uid = ?`
	update2Query = `
		UPDATE sessions
		SET end = CURRENT_TIMESTAMP, minutes = ?
		WHERE sid = ?`
)

var (
	insertStmt  *sql.Stmt
	selectStmt  *sql.Stmt
	update1Stmt *sql.Stmt
	update2Stmt *sql.Stmt
)

var (
	autoCommit bool
	freePeriod bool
	prepared   bool
)

func Init(product Product) error {
	if prepared {
		return nil
	}

	var err error
	query := fmt.Sprintf(insertQuery, product)
	insertStmt, err = db.Prepare(query)
	if err != nil {
		return err
	}
	selectStmt, err = db.Prepare(selectQuery)
	if err != nil {
		return err
	}
	update1Stmt, err = db.Prepare(update1Query)
	if err != nil {
		return err
	}
	update2Stmt, err = db.Prepare(update2Query)
	if err != nil {
		return err
	}

	prepared = true
	return nil
}

func AutoCommit(on bool) {
	autoCommit = on
	_ = autoCommit // keep linter happy
}

func FreePeriod(on bool) {
	freePeriod = on
}

func BeginSession(uid ibgames.AccountID, addr string) (*Session, error) { // FIXME: struct in_addr addr
	// Initialize the session record.
	s := &Session{
		uid: uid,
	}

	var complimentary string
	err := selectStmt.QueryRow(uid).Scan(&complimentary)
	if err != nil {
		return nil, err
	}
	s.complimentary = (complimentary == "Y")

	var minutes int
	if s.complimentary || freePeriod {
		minutes = 0
	} else {
		s.lastCharge = minimumCharge
		minutes = s.lastCharge
	}

	result, err := insertStmt.Exec(uid, addr, minutes)
	if err != nil {
		return nil, err
	}

	sid, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	s.sid = int32(sid)

	if minutes > 0 {
		result, err := update1Stmt.Exec(minutes, uid)
		if err != nil {
			return nil, err
		}
		rows, err := result.RowsAffected()
		if err != nil || rows != 1 {
			return nil, fmt.Errorf("expected to update 1 row, updated %d rows", rows)
		}
	}

	if err := db.Commit(); err != nil {
		log.Print("billing.BeginSession: db.Commit() failed")
		return nil, err
	}

	s.nextWrite = time.Now().Unix() + 60

	// Start the clock.
	s.StartClock()

	return s, nil
}

func (s *Session) End() int {
	return s.Time()
}

func (s *Session) StartClock() {
	if !s.ticking {
		s.lastTick = time.Now().Unix()
		s.ticking = true // Start the clock
	}
}

func (s *Session) StopClock() {
	if s.ticking {
		s.Tick()          // Accumulate time since last tick
		s.ticking = false // Stop the clock
	}
}

func (s *Session) Tick() int {
	// assert(session != NULL);

	now := time.Now().Unix()

	if !s.complimentary && !freePeriod && s.ticking {
		s.seconds += now - s.lastTick
		s.lastTick = now
	}

	charge := int(s.seconds / 60) // Convert seconds to minutes

	if charge > s.lastCharge || now >= s.nextWrite {
		result, err := update2Stmt.Exec(charge, s.sid)
		if err != nil {
			log.Printf("UPDATE: %v", err)
			return 0
		}
		rows, err := result.RowsAffected()
		if err != nil || rows != 1 {
			log.Printf("expected to update 1 row, updated %d rows", rows)
			return 0
		}

		if charge > s.lastCharge {
			result, err := update1Stmt.Exec(charge, s.uid)
			if err != nil {
				log.Printf("UPDATE: %v", err)
				return 0
			}
			rows, err := result.RowsAffected()
			if err != nil || rows != 1 {
				log.Printf("expected to update 1 row, updated %d rows", rows)
				return 0
			}
		}

		if err := db.Commit(); err != nil {
			log.Printf("COMMIT: %v", err)
			return 0
		}

		s.nextWrite = now + 60
		s.lastCharge = charge
	}

	return 1
}

func (s *Session) Time() int {
	s.Tick()            // Accumulate time since the last tick
	return s.lastCharge // Return the charged time so far
}
