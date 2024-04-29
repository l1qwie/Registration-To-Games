package main

import (
	"Schedule/api/servers"
)

func main() {
	go servers.Schedule()
	servers.Start()
}
