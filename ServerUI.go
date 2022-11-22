package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
)

func makeServer(win fyne.Window, s *Server) {
	//buttons
	btnFile := widget.NewButton("select file", chose(win))
	btnStart := widget.NewButton("Server Start", serverStart(s))
	btnFile.Resize(fyne.NewSize(80, 80))
	
	//container
	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), btnFile, btnStart)
	win.SetContent(container)
	win.Resize(fyne.NewSize(500, 300))
}

func serverStart(s *Server) func() {
	return func() {
		if s.isStart == false {
			fmt.Println("prepared to start Server")
			go s.serverStart()
			s.isStart = true
		} else {
			dialog.ShowInformation("Warning", "Already start ", s.win.Win)
		}
	}
}

func chose(win fyne.Window) func() {
	return func() {
		dia := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if closer == nil {
				return
			}

			defer func(closer fyne.URIReadCloser) {
				err := closer.Close()
				if err != nil {
					dialog.ShowError(err, win)
					return
				}
			}(closer)

			data, err := ioutil.ReadAll(closer)

			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			fmt.Println(len(data))

		}, win)
		dia.Show()
	}
}
