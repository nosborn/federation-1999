//go:build darwin

package main

import (
	"time"

	"golang.org/x/sys/unix"
)

type selector struct {
	kfd    int
	events []unix.Kevent_t
}

func NewSelector() (Selector, error) {
	kfd, err := unix.Kqueue()
	if err != nil {
		return nil, err
	}
	return &selector{
		kfd:    kfd,
		events: make([]unix.Kevent_t, 10),
	}, nil
}

func (p *selector) AddRead(fd int) error {
	event := unix.Kevent_t{
		Ident:  uint64(fd),
		Filter: unix.EVFILT_READ,
		Flags:  unix.EV_ADD,
	}
	_, err := unix.Kevent(p.kfd, []unix.Kevent_t{event}, nil, nil)
	return err
}

func (p *selector) UpdateWrite(fd int, add bool) error {
	event := unix.Kevent_t{
		Ident:  uint64(fd),
		Filter: unix.EVFILT_WRITE,
	}
	if add {
		event.Flags = unix.EV_ADD
	} else {
		event.Flags = unix.EV_DELETE
	}
	// We can ignore errors here, as the event might already be in the desired state.
	_, err := unix.Kevent(p.kfd, []unix.Kevent_t{event}, nil, nil)
	return err
}

func (p *selector) Wait(timeout time.Duration) ([]Event, error) {
	var ts *unix.Timespec
	if timeout >= 0 {
		t := unix.NsecToTimespec(timeout.Nanoseconds())
		ts = &t
	}

	nevents, err := unix.Kevent(p.kfd, nil, p.events, ts)
	if err != nil {
		if err == unix.EINTR {
			return nil, nil // Interrupted, no events
		}
		return nil, err
	}

	if nevents == 0 {
		return nil, nil
	}

	result := make([]Event, nevents)
	for i := range nevents {
		ev := p.events[i]
		result[i] = Event{
			Fd:    int(ev.Ident),
			IsEOF: ev.Flags&unix.EV_EOF != 0,
		}
	}
	return result, nil
}

func (p *selector) Close() error {
	return unix.Close(p.kfd)
}
