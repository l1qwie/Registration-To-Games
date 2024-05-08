package main

import (
	"Logs/app/consumer"
	"Logs/tests"
)

func main() {
	go consumer.StartConsumer()
	tests.StartTest()
}
