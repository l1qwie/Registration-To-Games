package main

import (
	"Registraion/enter"
	"Registraion/tests"
	"time"
)

func main() {
	go enter.Registration()
	time.Sleep(time.Second * 3)
	tests.Head()
}
