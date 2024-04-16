package main

import (
	"Media/enter"
	"Media/tests"
	"time"
)

func main() {
	go enter.Media()
	time.Sleep(time.Second)
	tests.Head()
}
