package main

import (
	"github.com/ebladrocher/smtp-client/server"
)

func main() {
	start()
}

func start() {
	srv := server.NewServer()
	srv.Start()
}
