//go:build linux

package ioctl

import (
	"golang.org/x/sys/unix"
)

const (
	GetTerminalAttrs    = unix.TCGETS
	SetTerminalAttrsNow = unix.TCSETS
)
