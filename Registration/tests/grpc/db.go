package grpc

import "Registraion/fmtogram/types"

func checkNewGame() bool {
	var count int
	err := types.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND 
		seats = $5 AND price = $6 AND currency = $7 AND status = $8`,
		99, "volleyball", 20240225, 1500, 34, 100, "USDT", 1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangeGame() bool {
	var count int
	err := types.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND 
		seats = $5 AND price = $6 AND currency = $7 AND status = $8`,
		99, "volleyball", 20240225, 1500, 34, 100, "USDT", -1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func deleteGame() {
	_, err := types.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", 99)
	if err != nil {
		panic(err)
	}
}

func createGame() {
	_, err := types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
		VALUES (99, 'volleyball', 20240225,1500, 34, 123.1223232, 55.112399992, 'Tel-Aviv Yafo', 100, 'USDT', 1)`)
	if err != nil {
		panic(err)
	}
}
