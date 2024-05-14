package writer

import (
	"Logs/app/handler/database"
	"Logs/types"
	"fmt"
	"log"
	"os"
)

func initfile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("You have an error when creating/opening file %s: %s", name, err)
	}
	return file
}

func write(text string, file *os.File) {
	defer file.Close()
	_, err := file.WriteString(text + "\n")
	if err != nil {
		log.Print("You have an error when writing to file:", err)
	}
}

func Internal(msg *types.Internal) {
	file := initfile("internal/internal.log")
	write(fmt.Sprintf("Time: %s | Id: %d | Message: %s | Data of Message: %s", msg.Com.Timestamp, msg.Com.LogId, msg.Message, msg.Data), file)
}

func ClientActivities(msg *types.ClientAct) {
	file := initfile("clientact/client-activities.log")
	write(fmt.Sprintf("Time: %s | Id: %d | User-ID: %d | Action: %s | Message: %s", msg.Com.Timestamp, msg.Com.LogId, msg.UserId, msg.Action, msg.Message), file)
	database.Add(msg)
}
