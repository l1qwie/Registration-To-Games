package main

import (
	"Settings/api/servers"
	"Settings/tests"
)

func main() {
	go servers.Start()
	go servers.Settings()
	tests.Head()
}
