package db

import "os"

func SetEnvironment() int {
	if err := os.Setenv("DBPATH", path); err != nil {
		return -1
	}
	return 0
}
