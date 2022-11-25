package main

import "net"

type user struct {
	name     string
	c        chan string
	Addr     string
	fileList map[string]int
	conn     net.Conn
}

func NewUser(conn net.Conn) *user {
	addr := conn.RemoteAddr().String()
	cl := &user{
		name: addr,
		Addr: addr,
		conn: conn,
		c:    make(chan string),
	}

	return cl
}
