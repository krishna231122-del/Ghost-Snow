package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dicedb/dice/config"
	"github.com/dicedb/dice/server"
	"github.com/ghost-snow/core"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the dice server")
	flag.IntVar(&config.Port, "port", 7379, "port for the dice server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("rolling the dice 🎲")
	server.RunSyncTCPServer()
	value, err := core.Decode([]byte("$5\r\nhello\r\n"))
	fmt.Println(value, err)
}
