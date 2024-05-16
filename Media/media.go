package main

import (
	"Media/api/servers"
	"Media/tests"
)

func main() {
	go servers.Start()
	go servers.Media()
	tests.Head()
}
