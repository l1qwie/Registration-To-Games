package apptype

import (
	"context"

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

// Opens connection with Redis DB
func AddCleint() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return client
}
