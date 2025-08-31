//go:build darwin

package ioctl

import (
	"golang.org/x/sys/unix"
)

const (
	GetTerminalAttrs    = unix.TIOCGETA
	SetTerminalAttrsNow = unix.TIOCSETA
)
