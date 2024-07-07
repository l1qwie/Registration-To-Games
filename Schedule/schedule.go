package main

import (
	"Schedule/api/servers"
)

func main() {
	//defer initredis.Del()
	//initredis.Start()
	go servers.Schedule()
	servers.Start()
	//tests.Head()
}
