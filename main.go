package main

import (
	"fmt"
	"os"

	"github.com/dickeyxxx/ploy/builder"
	"github.com/dickeyxxx/ploy/server"
)

func main() {
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "server":
		server.Run()
	case "build":
		builder.Run()
	default:
		fmt.Println("unknown command")
		server.Run()
	}
}
