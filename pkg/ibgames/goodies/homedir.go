package goodies

import (
	"os"
	"os/user"
	"strconv"
	"sync"
)

var (
	homeDir     string
	homeDirOnce sync.Once
)

func HomeDir() string {
	homeDirOnce.Do(func() {
		euid := os.Geteuid()
		u, err := user.LookupId(strconv.Itoa(euid))
		if err != nil {
			panic(err)
		}
		if u.HomeDir == "" {
			panic("user home directory is empty")
		}
		homeDir = u.HomeDir
	})
	return homeDir
}
