package types

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

type Request struct {
	Id          int     `json:"id"`
	Level       int     `json:"level"`
	Language    string  `json:"language"`
	ExMesId     int     `json:"exmessageid"`
	Act         string  `json:"action"`
	Limit       int     `json:"limit"`
	LaunchPoint int     `json:"launchpoint"`
	Req         string  `json:"request"`
	GameId      int     `json:"gameid"`
	Seats       int     `json:"seats"`
	Payment     string  `json:"payment"`
	Connection  *sql.DB `json:"connection"`
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
	Lattitude     float32
	Longitude     float32
	Payment       string
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
