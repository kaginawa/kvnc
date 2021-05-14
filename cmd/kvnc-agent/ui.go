package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/kaginawa/kvnc"
)

var (
	startButton   *widget.Button
	customIDEntry = widget.NewEntry()
)

func mainWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow("kvnc-agent")
	w.Resize(fyne.Size{Width: 450, Height: 300})
	customIDEntry.Text = config.CustomID
	autoStart := widget.NewCheck("Auto start", func(b bool) {
		config.AutoStart = b
	})
	autoStart.SetChecked(config.AutoStart)
	startButton = widget.NewButton("Start", func() { handleStart() })
	w.SetContent(container.NewVBox(widget.NewLabel("Custom ID:"), customIDEntry, autoStart, startButton))
	return w
}

func showError(err error) {
	dialog.NewError(err, w).Show()
}

func showConfigDialog() {
	serverEntry := widget.NewEntry()
	serverEntry.PlaceHolder = "foo.com"
	serverEntry.Validator = func(s string) error {
		if len(s) == 0 {
			return errors.New("please input server URL")
		}
		if _, err := url.Parse("https://" + s); err != nil {
			return errors.New("invalid URL")
		}
		return nil
	}
	apiKeyEntry := widget.NewEntry()
	apiKeyEntry.PlaceHolder = "bar"
	apiKeyEntry.Validator = func(s string) error {
		if len(s) == 0 {
			return errors.New("please input API key")
		}
		return nil
	}
	d := dialog.NewForm("Configuration", "OK", "Cancel", []*widget.FormItem{
		widget.NewFormItem("Server:", serverEntry),
		widget.NewFormItem("API Key: ", apiKeyEntry),
	}, func(ok bool) {
		if !ok {
			os.Exit(0)
		}
		config.Server = serverEntry.Text
		config.APIKey = apiKeyEntry.Text
		if err := config.save(*configFilePath); err != nil {
			println(err)
		}
	}, w)
	d.Show()
}

func handleStart() {
	startButton.SetText("Checking...")
	startButton.Disable()
	defer func() {
		startButton.SetText("Start")
		startButton.Enable()
	}()
	config.CustomID = customIDEntry.Text
	if err := config.save(*configFilePath); err != nil {
		log.Println(err)
	}
	if err := checkTCPPort("localhost", 5900); err != nil {
		if runtime.GOOS == "windows" {
			winVNC := exec.Command("WinVNC.exe")
			kvnc.PrepareBackgroundCommand(winVNC)
			if err := winVNC.Start(); err != nil {
				showError(err)
				return
			}
		} else {
			showError(err)
			return
		}
	}
	kaginawa := exec.Command("./kaginawa")
	kvnc.PrepareBackgroundCommand(kaginawa)
	stdout, err := kaginawa.StdoutPipe()
	if err != nil {
		showError(err)
		return
	}
	if err := kaginawa.Start(); err != nil {
		showError(fmt.Errorf("failed to start kaginawa: %v", err))
		return
	}
	startButton.SetText("Working...")
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := kaginawa.Wait(); err != nil {
		showError(fmt.Errorf("kaginawa exited: %v", err))
		return
	}
}
