package main

import (
	"Registration/api/servers"
	"Registration/tests/grpc"
)

func main() {
	go servers.Registration()
	go servers.Start()
	//tests.Head()
	grpc.Start()
}
