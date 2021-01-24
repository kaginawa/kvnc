package main

import (
	"flag"
	"fyne.io/fyne/v2"
)

var (
	w              fyne.Window
	config         = clientConfig{}
	configFilePath = flag.String("c", "kvnc.json", "path to configuration file")
)

func main() {
	flag.Parse()
	var err error
	config, err = loadConfig(*configFilePath)
	w = mainWindow()
	if err != nil {
		showConfigDialog()
	}
	w.ShowAndRun()
}
