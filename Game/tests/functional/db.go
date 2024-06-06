package functional

import "Game/apptype"

func gameSaved() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'volleyball' AND date = 20241209 AND time = 1900 AND seats = 15 AND price = 1000 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangedGame() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'football' AND date = 20241209 AND time = 1900 AND seats = 15 AND price = 1000 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
