package apptype

import (
	"Media/fmtogram/types"
	"database/sql"
	"fmt"
	"time"
)

var Db *sql.DB

type Request struct {
	Id           int           `json:"id"`
	Level        int           `json:"level"`
	Language     string        `json:"language"`
	Act          string        `json:"action"`
	Limit        int           `json:"limit"`
	LaunchPoint  int           `json:"launchpoint"`
	Req          string        `json:"request"`
	MediaDir     string        `json:"mediadir"`
	GameId       int           `json:"gameid"`
	FileId       string        `json:"fileid"`
	TypeOffile   string        `json:"typeoffile"`
	MediaCounter int           `json:"mediacounter"`
	Media        []types.Media `json:"media"`
	Status       bool          `json:"status"`
}

type Response struct {
	Message     string        `json:"message"`
	Keyboard    string        `json:"keyboard"`
	ChatID      int           `json:"chatid"`
	Level       int           `json:"level"`
	Error       string        `json:"error"`
	LaunchPoint int           `json:"launchpoint"`
	Act         string        `json:"action"`
	MediaDir    string        `json:"mediadir"`
	GameId      int           `json:"gameid"`
	FileId      string        `json:"fileid"`
	TypeOffile  string        `json:"typeoffile"`
	Media       []types.Media `json:"media"`
	ParseMode   string        `json:"parsemode"`
	Status      bool          `json:"status"`
}

type Game struct {
	Id    int
	Sport string
	Date  string
	Time  string
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
