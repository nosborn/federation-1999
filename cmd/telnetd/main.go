package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/telnet"
)

func main() {
	log.SetPrefix(fmt.Sprintf("%s[%d]: ", filepath.Base(os.Args[0]), os.Getpid()))
	log.SetFlags(log.Lmsgprefix)

	var proxyProto bool
	flag.BoolVar(&proxyProto, "proxy-proto", false, "enable PROXY protocol")
	flag.Parse()

	monitoring.StartServer(":8081")

	_ = telnet.ListenAndServe(":23", proxyProto)
}
