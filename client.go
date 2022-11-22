package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"net"
)

type cWinConfig struct {
	App fyne.App
	Win fyne.Window
}

type Client struct {
	conn net.Conn
	win  *cWinConfig
}

func cInit() *Client {
	client := &Client{
		win: &cWinConfig{
			App: app.New(),
		},
	}

	fmt.Println("client Init success.....")
	return client
}

func (c *Client) cStart() {

	c.win.Win = c.win.App.NewWindow("client")
	makeClientUI(c.win.Win, c)
	fmt.Println("make clientUI success ....")
	c.win.Win.ShowAndRun()

}
