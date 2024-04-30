package main

import (
	"Schedule/api/servers"
	initredis "Schedule/init-redis"
)

func main() {
	defer initredis.Del()
	initredis.Start()
	go servers.Schedule()
	servers.Start()
}
