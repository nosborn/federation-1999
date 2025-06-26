package server

import (
	"log"
	"time"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
)

type Transaction struct {
	now time.Time
}

var transaction *Transaction

func BeginTransaction() {
	global.Lock()

	if transaction != nil {
		log.Panic("Nested transaction")
	}
	transaction = &Transaction{
		now: time.Now(),
	}
}

func EndTransaction() {
	if transaction == nil {
		log.Panic("No active transaction")
	}
	database.CommitDatabase()
	transaction = nil

	global.Unlock()
}

func DayNumber() int {
	return int(transaction.now.Unix() / model.SECS_IN_A_DAY)
}

func TimeOfDay() time.Time {
	return transaction.now
}
