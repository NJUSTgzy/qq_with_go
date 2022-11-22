package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"net"
)

type WinConfig struct {
	App fyne.App
	Win fyne.Window
}

type Server struct {
	ConnType string
	Ip       string
	Port     string
	isStart  bool
	win      *WinConfig
}

func Init(con, Ip, Port string) *Server {
	server := &Server{
		ConnType: con,
		Ip:       Ip,
		Port:     Port,
		isStart:  false,
		win: &WinConfig{
			App: app.New(),
		},
	}

	fmt.Println("server Init success.....")
	return server
}

func (s *Server) Start() {

	s.win.Win = s.win.App.NewWindow("server")
	makeServer(s.win.Win, s)
	fmt.Println("make serverUI success ....")
	s.win.Win.ShowAndRun()

}

func (s *Server) Handle(conn net.Conn) {

}

func (s *Server) serverStart() {
	listen, err := net.Listen(s.ConnType, fmt.Sprintf("%s:%s", s.Ip, s.Port))
	if err != nil {
		dialog.ShowError(err, s.win.Win)
		fmt.Println("server Listen error", err)
		return
	}
	fmt.Println("server start to accept")
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			dialog.ShowError(err, s.win.Win)
			fmt.Println("accept error", err)
			continue
		}

		go s.Handle(conn)
	}
}
