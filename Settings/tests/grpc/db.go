package grpc

import (
	"Settings/apptype"
)

func createGame(gameId int) {
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, price, currency, seats, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		gameId, "football", 20240909, 1900, 33, "RUB", 55, 1)
	if err != nil {
		panic(err)
	}
}

func createGWUGame(gameId int) {
	_, err := apptype.Db.Exec(`INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, statuspayment, status) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		884, 33333, gameId, 5, "card", 1, 1)
	if err != nil {
		panic(err)
	}
}

func deleteGame(gameId int) {
	_, err := apptype.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
}

func deleteGWUGame(gameId int) {
	_, err := apptype.Db.Exec("DELETE FROM GamesWithUsers WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
}

func checkANewGameSch() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND price = $5 AND currency = $6 AND seats = $7 AND status = $8`,
		199, "volleyball", 20240909, 1800, 33, "USD", 55, 1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangeGameSch() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND price = $5 AND currency = $6 AND seats = $7 AND status = $8`,
		99, "football", 20240909, 1900, 33, "RUB", 55, -1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

//SELECT COUNT(*) FROM GamesWithUsers WHERE
//id = $1 AND userId = $2 AND gameId = $3 AND seats = $4 AND payment = $5 AND statuspayment = $6 AND status = $7

func ckeckAddGWU() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM GamesWithUsers WHERE
		id = $1 AND userId = $2 AND gameId = $3 AND seats = $4 AND payment = $5 AND statuspayment = $6 AND status = $7`,
		44, 233, 3, 5, "card", 0, 1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func ckeckChangeGWU() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM GamesWithUsers WHERE
		id = $1 AND userId = $2 AND gameId = $3 AND seats = $4 AND payment = $5 AND statuspayment = $6 AND status = $7`,
		884, 33333, 8, 5, "cash", 1, -1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
