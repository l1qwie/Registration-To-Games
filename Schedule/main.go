package main

import (
	"Schedule/enter"
	"Schedule/tests"
	"time"
)

func main() {
	go enter.Schedule()
	time.Sleep(time.Second)
	tests.Head()
}
