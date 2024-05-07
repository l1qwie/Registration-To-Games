package main

import (
	"Welcome/api/server"
	"Welcome/tests"
)

func main() {
	go server.Welcome()
	tests.Start()
}
