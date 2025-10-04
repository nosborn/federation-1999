package login

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nosborn/federation-1999/internal/ioctl"
	"golang.org/x/sys/unix"
)

func ReadPassword(prompt string, reader *bufio.Reader) (string, error) {
	termios, err := unix.IoctlGetTermios(int(os.Stdin.Fd()), ioctl.GetTerminalAttrs)
	if err != nil {
		log.Printf("readPassword: unix.IoctlGetTermios: %v", err)
	}
	defer func() {
		err = unix.IoctlSetTermios(int(os.Stdin.Fd()), ioctl.SetTerminalAttrsNow, termios)
		if err != nil {
			log.Printf("readPassword: unix.IoctlSetTermios: %v", err)
		}
	}()

	tempTermios := *termios
	tempTermios.Lflag &^= unix.ECHO
	err = unix.IoctlSetTermios(int(os.Stdin.Fd()), ioctl.SetTerminalAttrsNow, &tempTermios)
	if err != nil {
		log.Printf("readPassword: unix.IoctlSetTermios: %v", err)
	}

	// Flush any buffered input before prompting
	for reader.Buffered() > 0 {
		_, _ = reader.ReadByte()
	}

	password, err := ReadString(prompt, reader)
	if err != nil {
		return "", err
	}
	_, _ = fmt.Print("\n")

	return password, nil
}

func ReadString(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}
