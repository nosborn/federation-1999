package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/nosborn/federation-1999/internal/perivale"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

func main() {
	log.SetPrefix(fmt.Sprintf("%s[%d]: ", filepath.Base(os.Args[0]), os.Getpid()))
	log.SetFlags(log.Lmsgprefix)

	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: %s id address robobod\n", os.Args[0])
		os.Exit(1)
	}

	// Check for a valid UID.
	uid, err := strconv.Atoi(os.Args[1])
	if err != nil || uid < ibgames.MinAccountID || uid >= ibgames.MaxAccountID {
		fmt.Fprintf(os.Stderr, "%s: Bad UID", os.Args[0])
		os.Exit(1)
	}

	address := os.Args[2] // TODO
	robobod, _ := strconv.Atoi(os.Args[3])

	perivale.Perivale(ibgames.AccountID(uid), address, robobod)
}
