package tests

import (
	"Logs/tests/database"
	"fmt"
	"io"
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
	content, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при чтении из файла: %s", err))
	}

	if err := file.Sync(); err != nil {
		panic(fmt.Sprintf("Ошибка при синхронизации данных на диск: %s", err))
	}

	return string(content)
}

func testInterData() {
	timestampStr := "2024-05-12T10:48:34.670096883Z"
	timestamp, err := time.Parse(time.RFC3339Nano, timestampStr)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при парсинге времени: %s", err))
	}

	res := openAndRead("internal/internal.log")
	if res != fmt.Sprintf("Time: %v | Message: function Receiving in Media microservice was called | Data of Message: req *apptype.Request was given\n", timestamp) {
		panic(fmt.Sprintf("The result doesn't match the expectation. \n\nResult = %s", res))
	}
	log.Print("EVERYTHING IS OKAY! TEST WAS COMPLETED!")
}

func testClData() {
	timestampStr := "2024-05-12T10:48:34.670096883Z"
	timestamp, err := time.Parse(time.RFC3339Nano, timestampStr)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при парсинге времени: %s", err))
	}
	defer database.DeleteDb(timestamp)
	res := openAndRead("clientact/client-activities.log")
	log.Print(res)
	if res != fmt.Sprintf("Time: %v | User-ID: %d | Action: %s | Message: %s\n", timestamp, 1283829, "Registration", "The user has started registration action") {
		panic(fmt.Sprintf("The result doesn't match the expectation. \n\nResult = %s", res))
	}
	if !database.ChechDb(timestamp) {
		panic("The data in the database isn't match the expectation")
	}
	log.Print("EVERYTHING IS OKAY! TEST WAS COMPLETED!")
}
