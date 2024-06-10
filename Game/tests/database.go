package tests

import "Game/apptype"

func deleteTGame() {
	_, err := apptype.Db.Exec(`DELETE FROM Schedule WHERE 
		sport = 'volleyball' AND date = 20241209 AND time = 1900 AND seats = 15 AND price = 1000 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`)
	if err != nil {
		panic(err)
	}
}

func deleteChGame() {
	_, err := apptype.Db.Exec("DELETE FROM Schedule WHERE gameid = $1", 6667)
	if err != nil {
		panic(err)
	}
}

func createTestGame() {
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (gameid, sport, date, time, seats, price, currency, link, address, status)
		VALUES (6667, 'volleyball', 20241209, 1900, 15, 1000, 'RUB', 'https://www.google.com/maps?q=36.893445,30.709591', 'Игровая Площадка', 1)`)
	if err != nil {
		panic(err)
	}
}
