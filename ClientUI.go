package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"net"
	"strings"
)

func makeClientUI(win fyne.Window, c *Client) {
	WayEntry := widget.NewEntry()
	IPEntry := widget.NewEntry()
	PortEntry := widget.NewEntry()

	AddrForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "TCP/UDP", Widget: WayEntry},
			{Text: "IP", Widget: IPEntry},
			{Text: "Port", Widget: PortEntry},
		},
		OnSubmit: func() {
			conn, err := net.Dial(strings.ToLower(WayEntry.Text), fmt.Sprintf("%s:%s", IPEntry, PortEntry))
			if err != nil {
				dialog.ShowError(err, c.win.Win)
			} else {
				c.conn = conn
			}
		},
	}

	win.SetContent(AddrForm)

}
