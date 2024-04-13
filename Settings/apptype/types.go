package apptype

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Request struct {
	Id          int     `json:"id"`
	Level       int     `json:"level"`
	Language    string  `json:"language"`
	Act         string  `json:"action"`
	Limit       int     `json:"limit"`
	LaunchPoint int     `json:"launchpoint"`
	Req         string  `json:"request"`
	IsChanged   string  `json:"ischanged"`
	GameId      int     `json:"gameid"`
	Connection  *sql.DB `json:"connection"`
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
