package handler

import (
	"Schedule/apptype"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Adds a new client
// Open a connection with Redis DB
func addClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	return client, err
}

// Tries to find any game
// Returns quantity of keys
func findAnyGame(client *redis.Client, f func(error)) int {
	keys, err := client.Keys(context.Background(), "*").Result()
	if err != nil {
		f(err)
	}
	return len(keys)
}

// Selects all games from
func selecGames(client *redis.Client, f func(error), length int) []*apptype.Game {
	var err error
	lsG := make([]*apptype.Game, length)
	for i := 0; i < length && err == nil; i++ {
		g := new(apptype.Game)
		jsonstr, err := client.Get(context.Background(), fmt.Sprintf("gameid:%d", i)).Result()
		if err != nil {
			f(err)
		}
		err = json.Unmarshal([]byte(jsonstr), g)
		if err != nil {
			f(err)
		}
		lsG[i] = &apptype.Game{}
		lsG[i].Id = g.Id
		lsG[i].Sport = g.Sport
		lsG[i].Date = g.Date
		lsG[i].Time = g.Time
		lsG[i].Seats = g.Seats
		lsG[i].Price = g.Price
		lsG[i].Currency = g.Currency
	}
	return lsG
}
