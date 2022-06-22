package main

import (
	"os"

	"github.com/OhYee/webproxy/log"
	"github.com/OhYee/webproxy/proxy"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8000"
	}

	log.Infof("Server start at %s", addr)
	err := proxy.StartServer(addr)
	if err != nil {
		log.Errorf(err.Error())
	}
}
