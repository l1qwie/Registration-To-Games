package main

import (
	"Settings/api/servers"
)

func main() {
	go servers.Start()
	servers.Settings()
}
