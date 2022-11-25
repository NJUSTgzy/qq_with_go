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
		con := container.NewVBox()
		for k, _ := range s.OnlineMap {
			fmt.Println(k, "is to be send the file")
			con.Add(widget.NewCheck(k, func(b bool) {
				if b {
					s.sendTo[sp] = k
					sp++
				}
			}))

		}

		s.sendP = sp
		btnSendTo := widget.NewButton("send", send(s))
		btnSendTo.Resize(fyne.NewSize(60, 40))
		btnFlush := widget.NewButton("flush", func() {
			win.Close()
			go s.showUsers()
			return
		})

		win.Resize(fyne.NewSize(800, 600))
		con.Add(btnSendTo)
		con.Add(btnFlush)
		win.SetContent(con)
		win.Show()
	}

}

func (s *Server) showUsers() {
	win := s.win.App.NewWindow("Users")
	con := container.NewVBox()
	s.sendP = 0
	for k, _ := range s.OnlineMap {
		chk := widget.NewCheck(k, nil)
		chk.OnChanged = func(b bool) {
			if b {
				s.sendTo[s.sendP] = k
				s.sendP++
				fmt.Println(k, "is to be send the file")
			}
		}
		con.Add(chk)

	}

	btnSendTo := widget.NewButton("send", send(s))
	btnSendTo.Resize(fyne.NewSize(60, 40))
	btnFlush := widget.NewButton("flush", func() {
		win.Close()
		go s.showUsers()
		return
	})

	win.Resize(fyne.NewSize(800, 600))
	con.Add(btnSendTo)
	con.Add(btnFlush)
	win.SetContent(con)
	win.Show()
}

func send(s *Server) func() {
	return func() {

		if s.sendP == 0 {
			dialog.ShowInformation("warning", "please select users", s.win.Win)
		}

		for i := 0; i < s.sendP; i++ {
			value, ok := s.OnlineMap[s.sendTo[i]]
			if ok {
				go sendFile(value, s.file, s)
			} else {
				msg := s.sendTo[i] + "doesn't online"
				dialog.ShowInformation("warning", msg, s.win.Win)
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
