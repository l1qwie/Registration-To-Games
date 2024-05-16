package main

import (
	"Schedule/api/servers"
	initredis "Schedule/init-redis"
	"Schedule/tests"
)

func main() {
	defer initredis.Del()
	initredis.Start()
	go servers.Schedule()
	go servers.Start()
	tests.Head()
}
