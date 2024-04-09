package tests

import (
	"Registraion/fmtogram/types"
)

func UpdateAction() {
	_, err := types.Db.Exec("UPDATE Users SET action = $1 WHERE userId = $2", "reg to games", 477)
	if err != nil {
		panic(err)
	}
}

func UpdateLanguage() {
	_, err := types.Db.Exec("UPDATE Users SET language = $1 WHERE userId = $2", "ru", 477)
	if err != nil {
		panic(err)
	}
}

func UpdateLevel() {
	_, err := types.Db.Exec("UPDATE Users SET level = $1 WHERE userId = $2", i, 477)
	if err != nil {
		panic(err)
	}

}

func DeleteUser() {
	_, err := types.Db.Exec("DELETE FROM Users WHERE userId = $1", 477)
	if err != nil {
		panic(err)
	}
}

func CreateUser() {
	_, err := types.Db.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", 477, "reg to games", "ru", 0)
	if err != nil {
		panic(err)
	}
}

func DeleteGameWithUser() {
	_, err := types.Db.Exec("DELETE FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = 1", 477, 2)
	if err != nil {
		panic(err)
	}
}

func DeleteGame() {
	_, err := types.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", 2)
	if err != nil {
		panic(err)
	}
}

func CreateGame() {
	_, err := types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
		VALUES ($1, 'volleyball', 20250212, 1200, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', 1)`, 2)
	if err != nil {
		panic(err)
	}
}
