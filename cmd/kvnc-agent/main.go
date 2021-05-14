package main

import (
	"flag"

	"fyne.io/fyne/v2"
)

var (
	w              fyne.Window
	config         = kaginawaConfig{}
	configFilePath = flag.String("c", "kaginawa.json", "path to configuration file")
)

func main() {
	flag.Parse()
	var err error
	config, err = loadConfig(*configFilePath)
	w = mainWindow()
	if err != nil {
		showConfigDialog()
	} else if config.AutoStart {
		go handleStart()
	}
	w.ShowAndRun()
}
