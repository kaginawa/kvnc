package main

import (
	"flag"
	"fyne.io/fyne/v2"

	"github.com/kaginawa/kvnc"
)

var (
	w              fyne.Window
	config         = kvnc.Config{}
	configFilePath = flag.String("c", "kvnc.json", "path to configuration file")
)

func main() {
	flag.Parse()
	var err error
	config, err = kvnc.LoadConfig(*configFilePath)
	w = mainWindow()
	if err != nil {
		showConfigDialog()
	}
	w.ShowAndRun()
}
