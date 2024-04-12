package main

import (
	"Welcome/enter"
	"Welcome/tests"
	"time"
)

func main() {
	go enter.Welcome()
	time.Sleep(time.Second * 3)
	tests.Head()
}
