package types

import (
	"log"
	"os"
	"time"
)

type Common struct {
	Timestamp time.Time `json:"timestamp"`
	LogId     int       `json:"id"`
}

type Internal struct {
	Com     Common
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ClientAct struct {
	Com     Common
	UserId  int    `json:"userid"`
	Message string `json:"message"`
}

// Initialization of logs
func Initlogs(file string) *os.File {
	os.Remove(file)
	lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(lf)
	return lf
}
