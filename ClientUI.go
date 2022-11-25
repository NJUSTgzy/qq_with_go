package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"net"
	"strings"
)

func makeClientUI(win fyne.Window, c *Client) {
	WayEntry := widget.NewEntry()
	IPEntry := widget.NewEntry()
	PortEntry := widget.NewEntry()
	fileName := "/home/njustgzy/Desktop/first2.md"
	btnRecv := widget.NewButton("recvFile", recv(c, fileName))

	AddrForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "TCP/UDP", Widget: WayEntry},
			{Text: "IP", Widget: IPEntry},
			{Text: "Port", Widget: PortEntry},
		},
		OnSubmit: func() {
			conn, err := net.Dial(strings.ToLower(WayEntry.Text), fmt.Sprintf("%s:%s", IPEntry.Text, PortEntry.Text))
			if err != nil {
				dialog.ShowError(err, c.win.Win)
			} else {
				fmt.Println("client no err in connection")
				c.conn = conn
			}
		},
	}

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), AddrForm, btnRecv)

	win.SetContent(container)

}

func recv(c *Client, name string) func() {
	return func() {
		go recvFile(c.conn, name)
	}
}
