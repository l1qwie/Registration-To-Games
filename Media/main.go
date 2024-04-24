package main

import (
	"Media/api/servers"
)

func main() {
	go servers.Start()
	servers.Media()
}
