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

func checkChangedGameSport() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'football' AND date = 20241209 AND time = 1900 AND seats = 15 AND price = 1000 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangedGameDate() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'volleyball' AND date = 20241212 AND time = 1900 AND seats = 15 AND price = 1000 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangedGameTime() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'volleyball' AND date = 20241209 AND time = 2000 AND seats = 15 AND price = 1000 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangedGameSeats() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'volleyball' AND date = 20241209 AND time = 1900 AND seats = 99 AND price = 1000 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangedGamePrice() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'volleyball' AND date = 20241209 AND time = 1900 AND seats = 15 AND price = 199 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangedGameCurrency() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'volleyball' AND date = 20241209 AND time = 1900 AND seats = 15 AND price = 1000 AND currency = 'USDT' AND 
		link = 'https://www.google.com/maps?q=36.893445,30.709591' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangedGameLink() bool {
	var count int
	err := apptype.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		sport = 'volleyball' AND date = 20241209 AND time = 1900 AND seats = 15 AND price = 1000 AND currency = 'RUB' AND 
		link = 'https://www.google.com/maps/place/31%C2%B051''47.5%22N+34%C2%B051''50.8%22E/@31.863181,34.8626321,17' AND address = 'Игровая Площадка' AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
