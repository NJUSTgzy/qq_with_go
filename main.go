package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args[1])

	if os.Args == nil || os.Args[1] == "client" || os.Args[1] == "" {
		client := cInit()
		client.cStart()
	} else {
		server := Init("tcp", "127.0.0.2", "9090")
		server.Start()
	}

}
