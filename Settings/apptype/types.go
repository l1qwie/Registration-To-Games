package apptype

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var TestDb *sql.DB

type Request struct {
	Id          int    `json:"id"`
	Level       int    `json:"level"`
	Language    string `json:"language"`
	Act         string `json:"action"`
	Limit       int    `json:"limit"`
	LaunchPoint int    `json:"launchpoint"`
	Req         string `json:"request"`
	IsChanged   string `json:"ischanged"`
	GameId      int    `json:"gameid"`
	Status      bool   `json:"status"`
}

type Response struct {
	Message     string `json:"message"`
	Keyboard    string `json:"keyboard"`
	ChatID      int    `json:"chatid"`
	Level       int    `json:"level"`
	LaunchPoint int    `json:"launchpoint"`
	Act         string `json:"action"`
	IsChanged   string `json:"ischanged"`
	Language    string `json:"language"`
	GameId      int    `json:"gameid"`
	ParseMode   string `json:"parsemode"`
	Error       string `json:"error"`
	Status      bool   `json:"status"`
}

type Game struct {
	Id            int
	Sport         string
	Date          string
	Time          string
	Price         int
	Currency      string
	Payment       string
	Seats         int
	StatusPayment string
}

type GWU struct {
	Id      int
	UserId  int
	GameId  int
	Seats   int
	Payment string
	Statpay int
	Status  int
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

func docConnect() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		docHost,
		docPort,
		docUsername,
		docPass,
		docDbname,
		docSslmode)
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
