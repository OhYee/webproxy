package proxy

import (
	"context"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/OhYee/webproxy/log"
	"github.com/OhYee/webproxy/utils"
)

func StartTCPServer(src, dst string) error {
	var ipStr, portStr string

	tcpAddrArr := strings.Split(src, ":")
	switch len(tcpAddrArr) {
	case 1:
		ipStr = tcpAddrArr[0]
		portStr = "8000"
	case 2:
		ipStr = tcpAddrArr[0]
		portStr = tcpAddrArr[1]
	default:
		ipStr = "127.0.0.1"
		portStr = "8000"
	}
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP(ipStr),
		Port: int(port),
	})
	if err != nil {
		return err
	}

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Errorf("tcp connection error, due to %s", err)
			continue
		} else {
			go handleTCPConnection(conn, dst)
		}
	}
}

func handleTCPConnection(c *net.TCPConn, redirectAddr string) {
	defer c.Close()

	rconn, err := net.Dial("tcp", redirectAddr)
	if err != nil {
		log.Errorf("redirect tcp connection error, due to %s", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer cancel()

		n, err := redirectIO(ctx, c, rconn)
		if err != nil && err != io.EOF {
			log.Errorf("client => redirect got error %s", err)
		}
		log.Infof("client => redirect closed, transfer %s", utils.FormatBytes(n))
	}()
	go func() {
		defer wg.Done()
		defer cancel()

		n, err := redirectIO(ctx, rconn, c)
		if err != nil && err != io.EOF {
			log.Errorf("redirect => client got error %s", err)
		}
		log.Infof("redirect => client closed, transfer %s", utils.FormatBytes(n))
	}()

	wg.Wait()
}
