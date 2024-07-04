package tests

import "Registration/apptype"

// Updates column "action" in table "Users" in the database
func UpdateAction() {
	_, err := apptype.Db.Exec("UPDATE Users SET action = $1 WHERE userId = $2", "reg to games", 477)
	if err != nil {
		panic(err)
	}
}

// Updates column "language" in the table "Users" in the database
func UpdateLanguage() {
	_, err := apptype.Db.Exec("UPDATE Users SET language = $1 WHERE userId = $2", "ru", 477)
	if err != nil {
		panic(err)
	}
}

// Updates column "level" in the table "Users" in the database
func UpdateLevel() {
	_, err := apptype.Db.Exec("UPDATE Users SET level = $1 WHERE userId = $2", i, 477)
	if err != nil {
		panic(err)
	}

}

// Deletes a user from the table "Users"
func DeleteUser() {
	_, err := apptype.Db.Exec("DELETE FROM Users WHERE userId = $1", 477)
	if err != nil {
		panic(err)
	}
}

// Creates a user in the table "Users"
func CreateUser() {
	_, err := apptype.Db.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", 477, "reg to games", "ru", 0)
	if err != nil {
		panic(err)
	}
}

// Deletes a string from the table "GamesWithUsers" in the database
func DeleteGameWithUser() {
	_, err := apptype.Db.Exec("DELETE FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = 1", 477, 2)
	if err != nil {
		panic(err)
	}
}

// Deletes a game from the table "Schedule" in the database
func DeleteGame() {
	_, err := apptype.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", 2)
	if err != nil {
		panic(err)
	}
}

// Creates only one game in the table "Schedule" in the database
func CreateGame() {
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, link, address, price, currency, status) 
		VALUES ($1, 'volleyball', 20250212, 1200, 55, 'https://www.google.com/maps/place/31%C2%B051''47.5%22N+34%C2%B051''50.8%22E/@31.863181,34.8626321,17', 'Кладбище в Анталии', 100, 'USD', 1)`, 2)
	if err != nil {
		panic(err)
	}
}
