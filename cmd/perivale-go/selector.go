package main

import "time"

type Event struct {
	Fd    int
	IsEOF bool
}

type Selector interface {
	AddRead(fd int) error
	UpdateWrite(fd int, add bool) error
	Wait(timeout time.Duration) ([]Event, error)
	Close() error
}
