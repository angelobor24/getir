package main

import (
	"getir/server"
)

func main() {
	server := server.NewServerImpl(server.NewServiceImpl())
	server.StartServer()
}
