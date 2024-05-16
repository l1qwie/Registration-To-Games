package handler

import (
	"Schedule/api/producer"
	"Schedule/apptype"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// Adds a new client
// Open a connection with Redis DB
func addClient() (*redis.Client, error) {
	producer.InterLogs("Start function Schedule.addClient()", "no enter data")
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	return client, err
}

// Tries to find any game
// Returns quantity of keys
func findAnyGame(client *redis.Client, f func(error)) int {
	producer.InterLogs("Start function Schedule.findAnyGame()", fmt.Sprintf("client (*redis.Client): %v, f (func(error)): %T", client, f))
	keys, err := client.Keys(context.Background(), "*").Result()
	if err != nil {
		f(err)
	}
	return len(keys)
}

// Selects all games from
func selecGames(client *redis.Client, f func(error), length int) []*apptype.Game {
	producer.InterLogs("Start function Schedule.selecGames()", fmt.Sprintf("client (*redis.Client): %v, f (func(error)): %T, length (int): %d", client, f, length))
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

// Creates a new game
func newGame(client *redis.Client, g *apptype.Game) error {
	producer.InterLogs("Start function Schedule.newGame()", fmt.Sprintf("client (*redis.Client): %v, g (*apptype.Game): %v", client, g))
	key := fmt.Sprintf("gameid:%d", g.Id)
	jsonbyt, err := json.Marshal(g)
	if err != nil {
		log.Print(err)
	} else {
		err = client.Set(context.Background(), key, jsonbyt, 0).Err()
		if err != nil {
			log.Print(err)
		}
	}
	return err
}

// Deletes a game
func delGame(client *redis.Client, gid int) error {
	producer.InterLogs("Start function Schedule.delGame()", fmt.Sprintf("client (*redis.Client): %v, gid (int): %d", client, gid))
	key := fmt.Sprintf("gameid:%d", gid)
	err := client.Del(context.Background(), key).Err()
	if err != nil {
		log.Print(err)
	}
	return err
}

// "Change a game", but honestly this function just calls [delGame] and [newGame] then
func changeGame(client *redis.Client, g *apptype.Game) error {
	producer.InterLogs("Start function Schedule.changeGame()", fmt.Sprintf("client (*redis.Client): %v, g (*apptype.Game): %v", client, g))
	err := delGame(client, g.Id)
	if err == nil {
		err = newGame(client, g)
	}
	return err
}
