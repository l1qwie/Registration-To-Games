package grpc

import (
	"Registration/apptype"
)

func checkNewGame() bool {
	var count int
	err := apptype.Db.QueryRow(`
		SELECT COUNT(*) FROM Schedule WHERE 
		gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND 
		seats = $5 AND price = $6 AND currency = $7 AND address = $8 AND
		link = $9 AND status = $10`,
		99, "volleyball", 20240225, 1500, 34, 100, "USDT", "New Zeland", "https://www.google.com/maps/@30.1111727,31.5916288,14z?entry=ttu", 1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangeGame() bool {
	var count int
	err := apptype.Db.QueryRow(`
		SELECT COUNT(*) FROM Schedule WHERE 
		gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND 
		seats = $5 AND link = $6 AND address = $7 AND price = $8 AND currency = $9 AND status = $10`,
		99, "volleyball", 20240225, 1500, 34, "https://www.google.com/maps/@30.1111727,31.5916288,14z?entry=ttu", "New Zeland", 100, "USDT", -1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func deleteGame() {
	_, err := apptype.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", 99)
	if err != nil {
		panic(err)
	}
}

func createGame() {
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, address, link, price, currency, status) 
		VALUES (99, 'volleyball', 20240225, 1500, 34, 'New Zeland', 'https://www.google.com/maps/@30.1111727,31.5916288,14z?entry=ttu', 100, 'USDT', 1)`)
	if err != nil {
		panic(err)
	}
}
