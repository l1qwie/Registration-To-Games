package main

import (
	"Settings/enter"
	"Settings/tests"
	"time"
)

func main() {
	go enter.Settings()
	time.Sleep(time.Second * 3)
	tests.Head()
}
