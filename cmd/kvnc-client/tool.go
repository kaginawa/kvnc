package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	hostPlaceholder      = "{HOST}"
	portPlaceholder      = "{PORT}"
	defaultCommandFormat = "./vncviewer " + hostPlaceholder + "::" + portPlaceholder
)

func startViewer(port string) {
	format := defaultCommandFormat
	if len(config.ViewerCmd) > 0 {
		format = config.ViewerCmd
	}
	command := strings.Replace(format, hostPlaceholder, "localhost", 1)
	command = strings.Replace(command, portPlaceholder, port, 1)
	tokens := strings.Split(command, " ")
	res, err := exec.Command(tokens[0], tokens[1:]...).Output()
	if err != nil {
		println(string(res))
		os.Exit(1)
	}
	fmt.Println(string(res))
	os.Exit(0)
}
