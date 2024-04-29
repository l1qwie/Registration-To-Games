package main

import "Registraion/api/servers"

func main() {
	go servers.Registration()
	servers.Start()
}
