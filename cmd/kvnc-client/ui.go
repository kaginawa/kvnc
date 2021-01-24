package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/kaginawa/kaginawa-sdk-go"
)

func mainWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow("kvnc-client")
	w.Resize(fyne.Size{Width: 450, Height: 300})
	cid := widget.NewEntry()
	cid.Text = config.CustomID
	mac := widget.NewEntry()
	trueColor := widget.NewCheck("8bit true color", nil)
	fullScreen := widget.NewCheck("Full screen", nil)
	viewOnly := widget.NewCheck("View only", nil)
	connectButton := widget.NewButton("Connect", func() {})
	tabs := container.NewAppTabs(
		container.NewTabItem(
			"Custom ID",
			container.NewVBox(widget.NewLabel("Custom ID:"), cid, trueColor, fullScreen, viewOnly, connectButton),
		),
		container.NewTabItem(
			"MAC Address",
			container.NewVBox(widget.NewLabel("MAC Address:"), mac, trueColor, fullScreen, viewOnly, connectButton),
		),
	)
	connectButton.OnTapped = func() {
		params := viewerParams{
			trueColor:  trueColor.Checked,
			fullScreen: fullScreen.Checked,
			viewOnly:   viewOnly.Checked,
		}
		switch tabs.CurrentTabIndex() {
		case 0:
			handleConnectByCID(connectButton, cid, params)
		case 1:
			handleConnectByMAC(connectButton, mac, params)
		default:

		}
	}
	w.SetContent(tabs)
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
		if err := config.Save(*configFilePath); err != nil {
			println(err)
		}
	}, w)
	d.Show()
}

func handleConnectByMAC(button *widget.Button, mac *widget.Entry, params viewerParams) {
	button.SetText("Connecting...")
	button.Disable()
	defer func() {
		button.SetText("Connect")
		button.Enable()
	}()
	client, err := newKaginawaClient()
	if err != nil {
		showError(err)
		return
	}
	report, err := client.FindNode(context.Background(), mac.Text)
	if err != nil {
		showError(err)
		return
	}
	if report == nil {
		showError(errors.New("target not found"))
		return
	}
	connect(client, *report, params)
}

func handleConnectByCID(button *widget.Button, cid *widget.Entry, params viewerParams) {
	button.SetText("Connecting...")
	button.Disable()
	defer func() {
		button.SetText("Connect")
		button.Enable()
	}()
	client, err := newKaginawaClient()
	if err != nil {
		showError(err)
		return
	}
	reports, err := client.ListNodesByCustomID(context.Background(), cid.Text)
	if err != nil {
		showError(err)
		return
	}
	if len(reports) == 0 {
		showError(errors.New("target not found"))
		return
	}
	if len(reports) == 1 {
		connect(client, reports[0], params)
	}
	entries := make([]string, len(reports))
	for i, r := range reports {
		if len(r.LocalIPv4) == 0 && len(r.LocalIPv6) > 0 {
			entries[i] = fmt.Sprintf("%s@%s %s\n", r.LocalIPv6, r.Adapter, r.Hostname)
		} else {
			entries[i] = fmt.Sprintf("%s@%s %s\n", r.LocalIPv4, r.Adapter, r.Hostname)
		}
	}
	reportSelect := widget.NewSelectEntry(entries)
	reportSelect.SetText(entries[0])
	dialog.NewForm("Multiple choices", "OK", "Cancel", []*widget.FormItem{
		widget.NewFormItem("Target:", reportSelect),
	}, func(ok bool) {
		if !ok {
			return
		}
		for i, entry := range entries {
			if entry == reportSelect.SelectedText() {
				connect(client, reports[i], params)
				return
			}
		}
		os.Exit(1)
	}, w).Show()
}

func connect(client *kaginawa.Client, report kaginawa.Report, params viewerParams) {
	if report.SSHRemotePort == 0 {
		showError(errors.New("ssh not connected"))
		return
	}
	tunnel, err := client.FindSSHServerByHostname(context.Background(), report.SSHServerHost)
	if err != nil {
		showError(fmt.Errorf("failed to get ssh server information: %v", err))
		return
	}
	if tunnel == nil {
		showError(fmt.Errorf("unknown ssh server: %s", report.SSHServerHost))
		return
	}
	config.CustomID = report.CustomID
	if err := config.Save(*configFilePath); err != nil {
		log.Println(err)
	}
	listen(tunnel, report.SSHRemotePort, params)
}
