package billing

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func OpenDB(path string) error {
	var err error
	db, err = sql.Open("sqlite", fmt.Sprintf("file:%s?_pragma=foreign_keys(1)&_time_format=sqlite", path))
	if err != nil {
		return err
	}
	_, err = db.ExecContext(context.TODO(), "PRAGMA journal_mode = WAL;")
	if err != nil {
		log.Panicf("db.Exec: %v", err)
		_ = db.Close()
		db = nil
		return err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)
	return nil
}

func CloseDB() error {
	err := db.Close()
	if err != nil {
		return err
	}
	db = nil
	return nil
}
