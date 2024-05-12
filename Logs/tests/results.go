package tests

import (
	"fmt"
	"log"
	"os"
	"time"
)

func openAndRead(name string) string {
	file, err := os.Open(name)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при открытии файла: %s", err))
	}
	defer file.Close()
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при чтении из файла: %s", err))
	}

	if err := file.Sync(); err != nil {
		panic(fmt.Sprintf("Ошибка при синхронизации данных на диск: %s", err))
	}

	return string(buf[:n])
}

func testInterData() {
	timestampStr := "2024-05-12T10:48:34.670096883Z"
	timestamp, err := time.Parse(time.RFC3339Nano, timestampStr)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при парсинге времени: %s", err))
	}

	res := openAndRead("internal/internal.log")
	log.Print(res)
	if res != fmt.Sprintf("%s | Id: %d | Message: %s | Data of Message: %s", timestamp, 1, "function Receiving in Media microservice was called", "req *types.Request was given") {
		panic(fmt.Sprintf("The result doesn't match the expectation. Result = %s", res))
	}
	log.Print("EVERYTHING IS OKAY! TEST WAS COMPLETED!")
}

func testClData() {

}
