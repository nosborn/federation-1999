package login

import (
	"github.com/nosborn/federation-1999/internal/ioctl"
	"golang.org/x/sys/unix"
)

func SetInputTermios(fd int) error {
	termios, err := unix.IoctlGetTermios(fd, ioctl.GetTerminalAttrs)
	if err != nil {
		return err
	}
	termios.Lflag &^= unix.ISIG
	termios.Lflag |= unix.ICANON
	termios.Cc[unix.VINTR] = 0 // _POSIX_VDISABLE
	termios.Cc[unix.VQUIT] = 0 // _POSIX_VDISABLE
	// termios.Cc[unix.VERASE] = '\b' -- this seems to be broken and disagrees with perivale
	termios.Cc[unix.VKILL] = 0 // _POSIX_VDISABLE
	return unix.IoctlSetTermios(fd, ioctl.SetTerminalAttrsNow, termios)
}

func SetOutputTermios(fd int) error {
	termios, err := unix.IoctlGetTermios(fd, ioctl.GetTerminalAttrs)
	if err != nil {
		return err
	}
	termios.Oflag |= unix.OPOST | unix.ONLCR
	return unix.IoctlSetTermios(fd, ioctl.SetTerminalAttrsNow, termios)
}
