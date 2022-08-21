package main

import (
	"flag"
	"os"

	"github.com/OhYee/webproxy/log"
	"github.com/OhYee/webproxy/proxy"
)

func main() {
	var (
		runType string
	)
	flag.StringVar(&runType, "type", "http", "redirect proxy network: tcp, http, proxy. default type")
	flag.StringVar(&runType, "t", "http", "redirect proxy network: tcp, http. default http")
	flag.Parse()

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8000"
	}

	log.Infof("Server start at %s", addr)
	switch runType {
	case "proxy":
		err := proxy.StartServer(addr)
		if err != nil {
			log.Errorf(err.Error())
		}
	case "tcp":
		redirectAddr := os.Getenv("RADDR")
		err := proxy.StartTCPServer(addr, redirectAddr)
		if err != nil {
			log.Errorf(err.Error())
		}
	case "http":
		redirectAddr := os.Getenv("RADDR")
		err := proxy.StartHTTPServer(addr, redirectAddr)
		if err != nil {
			log.Errorf(err.Error())
		}
	}

}
