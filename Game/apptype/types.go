package apptype

import (
	"database/sql"
	"fmt"
	"time"
)

var Db *sql.DB

type Request struct {
	Id          int    `json:"id"`
	Level       int    `json:"level"`
	Req         string `json:"request"`
	Language    string `json:"language"`
	Act         string `json:"action"`
	LaunchPoint int    `json:"launchpoint"`
	Limit       int    `json:"limit"`
	Sport       string `json:"sport"`
	Date        int    `json:"date"`
	Time        int    `json:"time"`
	Seats       int    `json:"seats"`
	Price       int    `json:"price"`
	Currency    string `json:"currency"`
	Link        string `json:"link"`
	Address     string `json:"address"`
	Status      bool   `json:"status"`
}

type Response struct {
	Keyboard    string `json:"keyboard"`
	Message     string `json:"message"`
	ChatID      int    `json:"chatid"`
	Level       int    `json:"level"`
	Act         string `json:"action"`
	LaunchPoint int    `json:"launchpoint"`
	Sport       string `json:"sport"`
	Date        int    `json:"date"`
	Time        int    `json:"time"`
	Seats       int    `json:"seats"`
	Price       int    `json:"price"`
	Currency    string `json:"currency"`
	Link        string `json:"link"`
	Address     string `json:"address"`
	Error       string `json:"error"`
	Status      bool   `json:"status"`
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

func ConnectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", docConnect())
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
