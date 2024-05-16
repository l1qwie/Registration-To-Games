package main

import (
	"Registration/api/servers"
	"Registration/tests"
)

func main() {
	go servers.Registration()
	go servers.Start()
	tests.Head()
}
