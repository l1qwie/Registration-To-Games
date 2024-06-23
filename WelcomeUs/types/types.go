package types

import (
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
