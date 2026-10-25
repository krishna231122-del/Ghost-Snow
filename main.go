package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Ghost-Snow/config"
	"github.com/Ghost-Snow/core"
	"github.com/Ghost-Snow/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the Ghost-Snow server")
	flag.IntVar(&config.Port, "port", 7379, "port for the Ghost-Snow server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("running Ghost-Snow")
	server.RunSyncTCPServer()
	value, err := core.Decode([]byte("$5\r\nhello\r\n"))
	fmt.Println(value, err)
}
