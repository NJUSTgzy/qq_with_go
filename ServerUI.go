package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func makeServer(win fyne.Window, s *Server) {
	//buttons
	btnFile := widget.NewButton("select file", chose(s, win))
	btnStart := widget.NewButton("Server Start", serverStart(s))
	btnFile.Resize(fyne.NewSize(80, 80))
	fileLabel := widget.NewLabel("")
	btnCheckUser := widget.NewButton("Check User", showUser(s))
	s.btnSend = widget.NewButton("send", send(s))
	s.btnSend.Disabled()
	//container
	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), btnFile, btnStart, fileLabel, s.btnSend, btnCheckUser)
	win.SetContent(container)
	win.Resize(fyne.NewSize(500, 300))
}

func showUser(s *Server) func() {
	return func() {
		win := s.win.App.NewWindow("Users")
		sp := 0
		for k, _ := range s.OnlineMap {
			check := widget.NewCheck(k, func(b bool) {
				if b {
					s.sendTo[sp] = k
					sp++
				}
			})
			win.SetContent(container.NewVBox(check))
		}
		s.sendP = sp
		btnSendTo := widget.NewButton("send", send(s))
		btnSendTo.Resize(fyne.NewSize(60, 40))
		btnFlush := widget.NewButton("flush", func() {
			go a(s)
			win.Close()
		})

		win.Resize(fyne.NewSize(800, 600))
		win.SetContent(container.NewVBox(btnSendTo, btnFlush))

		win.Show()
	}

}

func a(s *Server) {
	win := s.win.App.NewWindow("Users")
	sp := 0
	for k, _ := range s.OnlineMap {
		check := widget.NewCheck(k, func(b bool) {
			if b {
				s.sendTo[sp] = k
				sp++
			}
		})
		fmt.Println(k)
		win.SetContent(container.NewVBox(check))
	}
	s.sendP = sp
	btnSendTo := widget.NewButton("send", send(s))
	btnSendTo.Resize(fyne.NewSize(60, 40))
	btnFlush := widget.NewButton("flush", func() {
		go showUser(s)
		win.Close()
	})

	win.Resize(fyne.NewSize(800, 600))
	win.SetContent(container.NewVBox(btnSendTo, btnFlush))

	win.Show()
}

func send(s *Server) func() {
	return func() {

		for i := 0; i < len(s.sendTo); i++ {
			fmt.Println(s.sendTo[i])
			value, ok := s.OnlineMap[s.sendTo[i]]
			if ok {
				go sendFile(value, s.file, s)
			}

		}

	}
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

func chose(s *Server, win fyne.Window) func() {
	return func() {
		dia := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if closer == nil {
				dialog.ShowInformation("err", "closer is nil", win)
				return
			}

			fileName := closer.URI().Path()
			defer closer.Close()
			s.file = fileName
			s.btnSend.Enable()

		}, win)
		dia.Show()
	}
}
