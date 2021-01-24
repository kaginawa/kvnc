package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/kaginawa/kaginawa-sdk-go"
	"github.com/kaginawa/kvnc"
	"golang.org/x/crypto/ssh"
)

func listen(tunnel *kaginawa.SSHServer, port int) {
	tunnelConfig, err := createSSHConfig(tunnel.User, tunnel.Key, tunnel.Password)
	if err != nil {
		showError(fmt.Errorf("failed to create SSH config: %v", err))
		return
	}
	// Connect to SSH tunneling server
	tConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", tunnel.Host, tunnel.Port), tunnelConfig)
	if err != nil {
		showError(fmt.Errorf("failed to connect SSH tunneling server: %v", err))
		return
	}

	// Listen local port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		showError(fmt.Errorf("failed to listen local port: %v", err))
		return
	}
	addr := listener.Addr().String()
	localPort := addr[strings.LastIndex(addr, ":")+1:]
	log.Printf("listening local port %s", localPort)
	go startViewer(localPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			showError(fmt.Errorf("failed to accept a connection: %v", err))
			return
		}
		go handleLocalConn(conn, tConn, port)
	}
}

func handleLocalConn(local net.Conn, tunnel *ssh.Client, port int) {
	defer kvnc.SafeClose(local, "local connection")
	tConn, err := tunnel.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		showError(fmt.Errorf("failed to connect: %v", err))
		return
	}
	chDone := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(tConn, local)
		if err != nil {
			log.Printf("error while copy remote->local: %s", err)
		}
		chDone <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(local, tConn)
		if err != nil {
			log.Printf("error while copy local->remote: %s", err)
		}
		chDone <- true
	}()
	<-chDone
}
