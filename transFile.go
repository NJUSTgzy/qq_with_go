package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, fileName string, s *Server) {
	fmt.Println(conn.RemoteAddr().String(), " is recving")
	ofile, err := os.Open(fileName)
	fileSize, _ := ofile.Seek(0, os.SEEK_END)
	ofile.Seek(0, os.SEEK_SET)
	fmt.Println(fileSize)
	defer ofile.Close()

	if err != nil {
		fmt.Println("failed to open the file : ", fileName, err)
		return
	}

	buf := make([]byte, 4096)
	for {
		n, err := ofile.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("send success")
			} else {
				fmt.Println("file read err: ", err)
			}
			return
		}

		_, err = conn.Write(buf[:n])

	}

}

func recvFile(conn net.Conn, fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("file create err: ", err)
		return
	}

	buf := make([]byte, 4096)
	for {

		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println(fileName, "accept success")
			} else {
				fmt.Println("recv error: ", err)
			}
			return
		}

		_, err = file.Write(buf[:n])

		if err != nil {

			return
		}

	}

}
