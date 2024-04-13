package tests

import (
	"Settings/apptype"
)

// Updates column "action" in table "Users" in the database
func UpdateAction() {
	_, err := apptype.Db.Exec("UPDATE Users SET action = $1 WHERE userId = $2", "reg to games", 899)
	if err != nil {
		panic(err)
	}
}

// Updates column "language" in the table "Users" in the database
func UpdateLanguage() {
	_, err := apptype.Db.Exec("UPDATE Users SET language = $1 WHERE userId = $2", "ru", 899)
	if err != nil {
		panic(err)
	}
}

// Updates column "level" in the table "Users" in the database
func UpdateLevel(level int) {
	_, err := apptype.Db.Exec("UPDATE Users SET level = $1 WHERE userId = $2", level, 899)
	if err != nil {
		panic(err)
	}

}

// Deletes a user from the table "Users"
func DeleteUser() {
	_, err := apptype.Db.Exec("DELETE FROM Users WHERE userId = $1", 899)
	if err != nil {
		panic(err)
	}
}

// Creates a user in the table "Users"
func CreateUser() {
	_, err := apptype.Db.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", 899, "reg to games", "ru", 0)
	if err != nil {
		panic(err)
	}
}

// Creates schedule for testing one user
// Changes the table "Schedule"
func CreateScheduleForUser() {
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (0, 'volleyball', 20250212, 1200, 44, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'POUNDS', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = apptype.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (1, 'volleyball', 20251212, 1100, 44, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = apptype.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (2, 'volleyball', 20250612, 1000, 44, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'RUB', 1)`)
	if err != nil {
		panic(err)
	}
}

// Creates a few games with the user
// Changes the table "GamesWithUsers"
func CreateUserScehdule() {
	_, err := apptype.Db.Exec("INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, statuspayment, status) VALUES (nextval('gameswithusers_id_seq'), $1, 0, 3, 'card', 1, 1)", 899)
	if err != nil {
		panic(err)
	}
	_, err = apptype.Db.Exec("INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, statuspayment, status) VALUES (nextval('gameswithusers_id_seq'), $1, 1, 6, 'card', 0, 1)", 899)
	if err != nil {
		panic(err)
	}
	_, err = apptype.Db.Exec("INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, statuspayment, status) VALUES (nextval('gameswithusers_id_seq'), $1, 2, 7, 'cash', 0, 1)", 899)
	if err != nil {
		panic(err)
	}
}

// Deletes the user's records to games
// Changes the table "GamesWithUsers"
func DeleteUserSchedule() {
	_, err := apptype.Db.Exec("DELETE FROM GamesWithUsers WHERE userId = $1", 899)
	if err != nil {
		panic(err)
	}
}

// Deletes the schedule's games
// Changes the table "Schedule"
func DeleteSchedule() {
	_, err := apptype.Db.Exec("DELETE FROM Schedule WHERE gameId = 0 OR gameId = 1 OR gameId = 2 OR gameId = 3 OR gameId = 4")
	if err != nil {
		panic(err)
	}
}
