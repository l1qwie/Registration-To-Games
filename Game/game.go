package main

import (
	"Game/api/servers"
	"Game/tests"
	"time"
)

func main() {
	go servers.Game()
	time.Sleep(time.Second / 20)
	tests.Start()
}
