package main

import (
	"Logs/app/consumer"
	"Logs/tests"
)

func main() {
	go consumer.Start()
	tests.StartTest()
}
