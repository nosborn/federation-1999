package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nosborn/federation-1999/internal/login"
)

func main() {
	log.SetPrefix(fmt.Sprintf("%s[%d]: ", filepath.Base(os.Args[0]), os.Getpid()))
	log.SetFlags(log.Lmsgprefix)

	remoteHostname := flag.String("h", "", "hostname")
	_ = flag.Bool("p", false, "")
	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatalf("usage: %s [-p] [-h hostname]", os.Args[0])
	}

	hostname := os.Getenv("TCPREMOTEIP")
	if hostname == "" {
		hostname = *remoteHostname
	}

	login.Login(hostname)
}
