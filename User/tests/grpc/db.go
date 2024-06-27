package grpc

import (
	"context"

	"github.com/l1qwie/Fmtogram/types"

	"github.com/go-redis/redis/v8"
)

func checkChangeLanguage() bool {
	var count int
	err := types.Db.QueryRow("SELECT COUNT(*) FROM Users WHERE userId = $1 AND language = $2 AND customlanguage = $3", 999, "en", true).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func createUser(userId int) {
	_, err := types.Db.Exec("INSERT INTO Users (userId, language, customlanguage) VALUES ($1, $2, $3)", userId, "ru", false)
	if err != nil {
		panic(err)
	}
}

func deleteUser(userId int) {
	_, err := types.Db.Exec("DELETE FROM Users WHERE userId = $1", userId)
	if err != nil {
		panic(err)
	}
}

func DeleteSchedule() {
	_, err := types.Db.Exec("DELETE FROM Schedule WHERE gameId = 0 OR gameId = 1 OR gameId = 2 OR gameId = 3 OR gameId = 4")
	if err != nil {
		panic(err)
	}
}

func DeleteEmptyMediaGame(gameId int) {
	_, err := types.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
}

func CreateSchedule() (err error) {
	var request string
	for i := 0; i < 4; i++ {
		if i == 0 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (0, 'volleyball', 20250212, 1200, 44, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'POUNDS', 1)`
		} else if i == 1 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (1, 'volleyball', 20260212, 1100, 34, 36.893445, 30.709591, 'Кладбище в Анталии', 10, 'POUNDS', 1)`
		} else if i == 2 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (2, 'football', 20250412, 1800, 14, 36.893445, 30.709591, 'Кладбище в Анталии', 1000, 'USD', 1)`
		} else if i == 3 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (3, 'volleyball', 20250202, 0800, 77, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'RUB', 1)`
		}
		_, err = types.Db.Exec(request)
		if err != nil {
			panic(err)
		}
	}
	return err
}

func CreateEmptyMediaGame(gameId int) (err error) {
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES ($1, 'volleyball', 20240212, 1200, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', -1)`, gameId)
	return err
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
