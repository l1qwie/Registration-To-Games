package apptype

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

type Request struct {
	Id       int    `json:"id"`
	Act      string `json:"action"`
	Language string `json:"language"`
}

type Response struct {
	Keyboard  string `json:"keyboard"`
	Message   string `json:"message"`
	ChatId    int    `json:"chaid"`
	Level     int    `json:"level"`
	Act       string `json:"action"`
	ParseMode string `json:"parsemode"`
	Error     string `json:"error"`
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

// Opens connection with Redis DB
func AddCleint() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return client
}
