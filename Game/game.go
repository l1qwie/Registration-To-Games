package main

import (
	"Game/api/servers"
	"Game/tests"
)

func main() {
	go servers.Game()
	tests.Start()
}
