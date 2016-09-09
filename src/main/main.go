package main

import (
	"flag"
	"server"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8088, "Port the server listens to")

	flag.Parse()

	server.Run(port)
}
