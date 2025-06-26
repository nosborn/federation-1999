package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/internal/workbench"
)

func main() {
	log.SetPrefix(fmt.Sprintf("%s[%d]: ", filepath.Base(os.Args[0]), os.Getpid()))
	log.SetFlags(log.Lmsgprefix)

	if !workbench.ParseArguments() {
		log.Printf("usage: %s [-c] [-n] id", os.Args[0])
		os.Exit(1)
	}

	if workbench.CheckOnly {
		ok := workbench.CheckPlanet()
		if !ok {
			os.Exit(1)
		}
		os.Exit(0)
	}

	fmt.Print(text.Msg(text.Workbench_Hello))
	workbench.MainMenu()
	fmt.Print(text.Msg(text.Workbench_Goodbye))

	time.Sleep(1 * time.Second)

	os.Exit(0)
}
