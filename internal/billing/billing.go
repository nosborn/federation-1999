package billing

import (
	"context"
	"encoding/json"
	"time"

	"github.com/nosborn/federation-1999/internal/ibgames"
)

type Product int

const (
	NoProduct Product = iota
	Federation
	AgeOfAdventure
)

const (
	MINIMUM_CHARGE = 1
)

type Session struct {
	uid           ibgames.AccountID // Account being billed
	complimentary bool              // Is account complimentary?
	lastCharge    int               // Billed minutes at last database update
	sid           int32             // Key of session record
	nextWrite     time.Time         // Time at which current minute ends
	lastTick      time.Time         // Time this session last ticked
	ticking       bool              // Accumulate billable seconds?
	seconds       time.Duration     // Billable seconds
}

var (
	autoCommit bool
	freePeriod bool
)

func Init(product Product) int {
	// TODO
	return 0
}

func AutoCommit(on bool) {
	autoCommit = on
}

func FreePeriod(on bool) {
	freePeriod = on
}

func BeginSession(uid ibgames.AccountID, addr string) (*Session, error) { // FIXME: struct in_addr addr
	ctx := context.TODO()

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

	// Initialize the session record.
	s := &Session{
		uid: uid,
	}

	var complimentary string
	err = tx.QueryRowContext(ctx, "SELECT json_extract(data, '$.complimentary') FROM accounts WHERE uid = ?", uid).Scan(&complimentary)
	if err != nil {
		return nil, err
	}
	s.complimentary = (complimentary == "Y")

	var minutes int
	if s.complimentary || freePeriod {
		minutes = 0
	} else {
		s.lastCharge = MINIMUM_CHARGE
		minutes = s.lastCharge
	}

	// EXEC SQL execute insert_p using :uid_p, :ip_address, :minutes;
	// if (strncmp(SQLSTATE, "00", 2) != 0) {
	// 	fprintf(stderr, "INSERT %s %ld\n", SQLSTATE, sqlca.sqlcode);
	// 	free(session);
	// 	return NULL;
	// }
	const query = `
                INSERT INTO sessions (data)
                VALUES (?)`

	// a := &Account{
	// 	Name:          name,
	// 	Encrypt:       encrypt,
	// 	SChange:       time.Now(),
	// 	Signup:        time.Now(),
	// 	Status:        "A", // FIXME
	// 	Complimentary: "N", // FIXME
	// }

	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	result, err := tx.ExecContext(ctx, query, string(data))
	if err != nil {
		return nil, err
	}

	sid, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	s.sid = int32(sid)

	if minutes > 0 {
		// EXEC SQL execute update1_p using :minutes, :uid_p;
		// if strncmp(SQLSTATE, "00", 2) != 0 {
		// 	log.Printf("UPDATE %s %d", SQLSTATE, sqlca.sqlcode)
		// 	return nil
		// }
	}

	if autoCommit {
		// if dbCommit() == -1 {
		// 	log.Print("billing.BeginSession: dbCommit() failed")
		// 	return nil
		// }
	}

	s.nextWrite = time.Now().Add(60 * time.Second)

	// Start the clock.
	s.StartClock()

	return s, nil
}

func (s *Session) End() int {
	return s.Time()
}

func (s *Session) StartClock() {
	if !s.ticking {
		s.lastTick = time.Now()
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
	// time_t now;
	// time_t charge;

	// assert(session != NULL);

	now := time.Now()

	if !s.complimentary && !freePeriod && s.ticking {
		s.seconds += now.Sub(s.lastTick)
		s.lastTick = now
	}

	charge := int(s.seconds.Minutes())

	if charge > s.lastCharge || now.After(s.nextWrite) {
		// EXEC SQL BEGIN DECLARE SECTION;
		// int sid_p = session->sid;
		// int charge_p = charge;
		// EXEC SQL END DECLARE SECTION;
		//
		// EXEC SQL execute update2_p using :charge_p, :sid_p;

		// if strncmp(SQLSTATE, "00", 2) != 0 {
		// 	log.Printf("UPDATE %s %d", SQLSTATE, sqlca.sqlcode)
		// 	return 0
		// }

		if charge > s.lastCharge {
			// EXEC SQL BEGIN DECLARE SECTION;
			// int uid_p = session->uid;
			// int this_charge = charge - session->last_charge;
			// EXEC SQL END DECLARE SECTION;
			//
			// EXEC SQL execute update1_p using :this_charge, :uid_p;

			// if strncmp(SQLSTATE, "00", 2) != 0 {
			// 	fprintf(stderr, "UPDATE %s %ld\n", SQLSTATE, sqlca.sqlcode)
			// 	return 0
			// }
		}

		if autoCommit {
			// if dbCommit() == -1 {
			// 	log.Printf("billingTick: dbCommit() failed")
			// 	return 0
			// }
		}

		s.nextWrite = now.Add(60 * time.Second)
		s.lastCharge = charge
	}

	return 1
}

func (s *Session) Time() int {
	s.Tick()            // Accumulate time since the last tick
	return s.lastCharge // Return the charged time so far
}
