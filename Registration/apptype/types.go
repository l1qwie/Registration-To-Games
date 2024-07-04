package apptype

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Request struct {
	Id          int    `json:"id"`
	Level       int    `json:"level"`
	Language    string `json:"language"`
	ExMesId     int    `json:"exmessageid"`
	Act         string `json:"action"`
	Limit       int    `json:"limit"`
	LaunchPoint int    `json:"launchpoint"`
	Req         string `json:"request"`
	GameId      int    `json:"gameid"`
	Seats       int    `json:"seats"`
	Payment     string `json:"payment"`
	Status      bool   `json:"status"`
}

type Response struct {
	Error       string `json:"error"`
	Keyboard    string `json:"keyboard"`
	Message     string `json:"message"`
	ChatID      int    `json:"chatid"`
	Level       int    `json:"level"`
	LaunchPoint int    `json:"launchpoint"`
	GameId      int    `json:"gameid"`
	Seats       int    `json:"seats"`
	Payment     string `json:"payment"`
	Act         string `json:"action"`
	ParseMode   string `json:"parsemode"`
	Status      bool   `json:"status"`
}

type Game struct {
	Id            int
	Sport         string
	Date          string
	Time          string
	Seats         int
	Price         int
	Currency      string
	Address       string
	Link          string
	Payment       string
	StatusPayment string
}

type Updates struct {
	Id      int
	GameId  int
	UserId  int
	Seats   int
	Payment string
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
