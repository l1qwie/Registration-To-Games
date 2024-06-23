package apptype

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Internal struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      string    `json:"data"`
}

type ClientAct struct {
	Timestamp time.Time `json:"timestamp"`
	UserId    int       `json:"userid"`
	Action    string    `json:"action"`
	Message   string    `json:"message"`
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

func ConnectToDatabase() *sql.DB {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", docConnect())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func docConnect() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		docHost,
		docPort,
		docUsername,
		docPass,
		docDbname,
		docSslmode)
}
