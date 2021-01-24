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

type viewerParams struct {
	trueColor  bool
	fullScreen bool
	viewOnly   bool
}

func startViewer(port string, params viewerParams) {
	format := defaultCommandFormat
	if len(config.ViewerCmd) > 0 {
		format = config.ViewerCmd
	}
	command := strings.Replace(format, hostPlaceholder, "localhost", 1)
	command = strings.Replace(command, portPlaceholder, port, 1)
	if params.trueColor {
		command += " -8bit"
	}
	if params.fullScreen {
		command += " -fullscreen"
	}
	if params.viewOnly {
		command += " -viewonly"
	}
	tokens := strings.Split(command, " ")
	res, err := exec.Command(tokens[0], tokens[1:]...).Output()
	if err != nil {
		println(string(res))
		os.Exit(1)
	}
	fmt.Println(string(res))
	os.Exit(0)
}
