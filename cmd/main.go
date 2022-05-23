package main

import (
	"getir/server"
	"getir/storage"
)

func main() {
	//initialize compare function
	compareFunc := func(a int32, b int32, c int32) bool {
		if c >= a && c <= b {
			return true
		}
		return false
	}
	//initialize server data structure
	server := server.NewServerImpl(server.NewServiceImpl(storage.NewStorageImpl(compareFunc)))
	//start server
	server.StartServer()
}
