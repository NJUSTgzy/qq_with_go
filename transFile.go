package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, fileName string, size int64) {
	file, err := os.Open(fileName)
	size, _ = file.Seek(0, os.SEEK_END)
	defer file.Close()

	if err != nil {
		fmt.Println("failed to open the file :%s ,error %v", fileName, err)
		return
	}

	buf := make([]byte, 4096)

	for {
		n, err := file.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Read success")
			} else {
				fmt.Println("file read err: %v", err)
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
		fmt.Println("file create err: %v ", err)
		return
	}

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("%s accept success", fileName)
			} else {
				fmt.Println("recv error: %v", err)
			}
			return
		}

		file.Write(buf[:n])

	}

}
