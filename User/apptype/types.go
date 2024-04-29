package apptype

import (
	"User/fmtogram/types"
)

type RegToGames struct {
	GameId  int
	Seats   int
	Sport   string
	Payment string
}

type Media struct {
	Id        string
	Interval  string
	Direction string
	Limit     int
	DelGameId int
	Counter   int
	Photo     string
	Video     string
	MediaF    []types.Media
}

type UserRec struct {
	Changeable     string
	GameId         int
	ActGame        string
	WillChangeable string
	NewPay         string
}

type User struct {
	Id          int
	ExMessageId int
	Photo       string
	Video       string
	MediaF      []types.Media
	Counter     int
	Request     string
	Language    string
	LaunchPoint int
	Act         string
	Level       int
	Reg         RegToGames
	Media       Media
	UserRec     UserRec
	AppErr      chan *error
}

type WelcomeReq struct {
	Id       int    `json:"id"`
	Level    int    `json:"level"`
	Language string `json:"language"`
	Req      string `json:"request"`
	Act      string `json:"action"`
}
type WelcomeRes struct {
	Keyboard string `json:"keyboard"`
	Message  string `json:"message"`
	ChatID   int    `json:"chatid"`
	Act      string `json:"action"`
	Level    int    `json:"level"`
	Error    string `json:"error"`
}

type RegistrationReq struct {
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
}

type RegistrationRes struct {
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
}

type ScheduleReq struct {
	Id       int    `json:"id"`
	Act      string `json:"action"`
	Language string `json:"language"`
}

type ScheduleRes struct {
	Error     string `json:"error"`
	Keyboard  string `json:"keyboard"`
	Message   string `json:"message"`
	ChatID    int    `json:"chaid"`
	Level     int    `json:"level"`
	Act       string `json:"action"`
	ParseMode string `json:"parsemode"`
}

type MediaReq struct {
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
}

type MediaRes struct {
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
}

type SettingsReq struct {
	Id          int    `json:"id"`
	Level       int    `json:"level"`
	Language    string `json:"language"`
	Act         string `json:"action"`
	Limit       int    `json:"limit"`
	LaunchPoint int    `json:"launchpoint"`
	Req         string `json:"request"`
	IsChanged   string `json:"ischanged"`
	GameId      int    `json:"gameid"`
}

type SettingsRes struct {
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
	Id       int    `json:"id"`
	Sport    string `json:"sport"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Seats    int    `json:"seats"`
	Price    int    `json:"price"`
	Currency string `json:"currency"`
	Action   string
}
