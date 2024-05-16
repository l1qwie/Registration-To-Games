package types

import (
	"database/sql"
	"fmt"
	"time"
)

type Response struct {
	Keyboard string `json:"keyboard"`
	Message  string `json:"message"`
	ChatID   int    `json:"chatid"`
	Act      string `json:"action"`
	Level    int    `json:"level"`
	Error    string `json:"error"`
	Status   bool   `json:"status"`
}

type Request struct {
	Id       int    `json:"id"`
	Level    int    `json:"level"`
	Language string `json:"language"`
	Req      string `json:"request"`
	Act      string `json:"action"`
	Status   bool   `json:"status"`
}

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

func connectData() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmode)
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

func ConnectToDatabase(doc bool) *sql.DB {
	var (
		db  *sql.DB
		err error
	)
	if doc {
		db, err = sql.Open("postgres", docConnect())
	} else {
		db, err = sql.Open("postgres", connectData())
	}
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
