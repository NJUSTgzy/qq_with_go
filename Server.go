package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"net"
	"strconv"
	"sync"
)

type WinConfig struct {
	App fyne.App
	Win fyne.Window
}

type Server struct {
	ConnType  string
	Ip        string
	Port      string
	isStart   bool
	OnlineMap map[string]*user
	Message   string
	file      string
	btnSend   *widget.Button
	sendTo    []string
	nowPeople int
	UserLock  sync.RWMutex
	win       *WinConfig
}

func Init(con, Ip, Port string) *Server {
	server := &Server{
		ConnType:  con,
		Ip:        Ip,
		Port:      Port,
		isStart:   false,
		OnlineMap: make(map[string]*user),
		nowPeople: 0,
		sendTo:    make([]string, 16),

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
	User := NewUser(conn)
	fmt.Println(conn.RemoteAddr().String(), " online...")
	s.UserLock.Lock()
	s.OnlineMap[conn.RemoteAddr().String()] = User
	s.sendTo[s.nowPeople] = conn.RemoteAddr().String()
	s.nowPeople++
	s.UserLock.Unlock()

}

func (s *Server) boardCast() {

	port, _ := strconv.Atoi(s.Port)
	laddr := net.UDPAddr{
		IP:   net.IP(s.Ip),
		Port: port,
	}

	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 3000,
	}

	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		println(err.Error())
		return
	}

	conn.Write([]byte(s.Message))
	conn.Close()

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
