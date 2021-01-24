package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/kaginawa/kvnc"
)

func checkTCPPort(host string, port int) error {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), timeout)
	if err != nil {
		return err
	}
	if conn == nil {
		return fmt.Errorf("failed to check TCP port %s:%d", host, port)
	}
	kvnc.SafeClose(conn, "port check")
	return nil
}
