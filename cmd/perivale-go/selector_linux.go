//go:build linux

package main

import (
	"time"

	"golang.org/x/sys/unix"
)

type selector struct {
	efd    int
	events []unix.EpollEvent
}

func NewSelector() (Selector, error) {
	efd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &selector{
		efd:    efd,
		events: make([]unix.EpollEvent, 10),
	}, nil
}

func (p *selector) AddRead(fd int) error {
	event := &unix.EpollEvent{
		Events: unix.EPOLLIN,
		Fd:     int32(fd),
	}
	return unix.EpollCtl(p.efd, unix.EPOLL_CTL_ADD, fd, event)
}

func (p *selector) UpdateWrite(fd int, add bool) error {
	event := &unix.EpollEvent{
		Events: unix.EPOLLOUT,
		Fd:     int32(fd),
	}
	if add {
		return unix.EpollCtl(p.efd, unix.EPOLL_CTL_ADD, fd, event)
	}
	return unix.EpollCtl(p.efd, unix.EPOLL_CTL_DEL, fd, event)
}

func (p *selector) Wait(timeout time.Duration) ([]Event, error) {
	timeoutMillis := -1
	if timeout >= 0 {
		timeoutMillis = int(timeout.Milliseconds())
	}

	nevents, err := unix.EpollWait(p.efd, p.events, timeoutMillis)
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
			Fd:    int(ev.Fd),
			IsEOF: ev.Events&unix.EPOLLHUP != 0,
		}
	}
	return result, nil
}

func (p *selector) Close() error {
	return unix.Close(p.efd)
}
