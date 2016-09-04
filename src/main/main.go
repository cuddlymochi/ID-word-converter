package main

import (
	"db"
	"flag"
	"server"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8081, "Port the server listens to")

	flag.Parse()

	db.Open()
	defer db.Db.Close()
	server.Run(port)
}
