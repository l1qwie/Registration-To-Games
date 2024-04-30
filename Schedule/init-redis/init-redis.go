package initredis

import (
	"Schedule/apptype"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Adds data for schedule in array
// Returns this data
func addData() []*apptype.Game {
	lsG := make([]*apptype.Game, 4)
	for k := 0; k < 4; k++ {
		lsG[k] = &apptype.Game{}
		lsG[k].Id = k
		if k == 0 {
			lsG[k].Sport = "volleyball"
			lsG[k].Date = "12-02-2025"
			lsG[k].Time = "12:00"
			lsG[k].Seats = 44
			lsG[k].Price = 100
			lsG[k].Currency = "POUNDS"
		} else if k == 1 {
			lsG[k].Sport = "volleyball"
			lsG[k].Date = "12-06-2026"
			lsG[k].Time = "11:00"
			lsG[k].Seats = 34
			lsG[k].Price = 10
			lsG[k].Currency = "POUNDS"
		} else if k == 2 {
			lsG[k].Sport = "football"
			lsG[k].Date = "12-04-2025"
			lsG[k].Time = "18:00"
			lsG[k].Seats = 14
			lsG[k].Price = 1000
			lsG[k].Currency = "USD"
		} else if k == 3 {
			lsG[k].Sport = "volleyball"
			lsG[k].Date = "02-02-2025"
			lsG[k].Time = "08:00"
			lsG[k].Seats = 77
			lsG[k].Price = 100
			lsG[k].Currency = "RUB"
		}
	}
	return lsG
}

// Put data into Redis DB
func putInRedis(lsg []*apptype.Game, cl *redis.Client) {
	for k := 0; k < 4; k++ {
		key := fmt.Sprintf("gameid:%d", lsg[k].Id)
		jsonbyt, err := json.Marshal(lsg[k])
		if err != nil {
			panic(err)
		}
		err = cl.Set(context.Background(), key, jsonbyt, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

// The main function
// Main logic is here
func createSchRedis(cl *redis.Client) {
	lsG := addData()
	putInRedis(lsG, cl)
}

// deletes all schedule from Redis DB
func Del() {
	cl := apptype.AddCleint()
	err := cl.FlushAll(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}

func Start() {
	cl := apptype.AddCleint()
	createSchRedis(cl)
}
