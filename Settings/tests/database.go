package tests

import (
	"Settings/apptype"
)

// Updates column "action" in table "Users" in the database
func UpdateAction() {
	_, err := apptype.TestDb.Exec("UPDATE Users SET action = $1 WHERE userId = $2", "reg to games", 899)
	if err != nil {
		panic(err)
	}
}

// Updates column "language" in the table "Users" in the database
func UpdateLanguage() {
	_, err := apptype.TestDb.Exec("UPDATE Users SET language = $1 WHERE userId = $2", "ru", 899)
	if err != nil {
		panic(err)
	}
}

// Updates column "level" in the table "Users" in the database
func UpdateLevel(level int) {
	_, err := apptype.TestDb.Exec("UPDATE Users SET level = $1 WHERE userId = $2", level, 899)
	if err != nil {
		panic(err)
	}

}

// Deletes a user from the table "Users"
func DeleteUser() {
	_, err := apptype.TestDb.Exec("DELETE FROM Users WHERE userId = $1", 899)
	if err != nil {
		panic(err)
	}
}

// Creates a user in the table "Users"
func CreateUser() {
	_, err := apptype.TestDb.Exec("INSERT INTO Users (userId, language, customlanguage) VALUES ($1, $2, $3)", 899, "ru", false)
	if err != nil {
		panic(err)
	}
}

// Creates schedule for testing one user
// Changes the table "Schedule"
func CreateScheduleForUser() {
	_, err := apptype.TestDb.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, price, currency, status) 
		VALUES (0, 'volleyball', 20250212, 1200, 44, 100, 'POUNDS', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = apptype.TestDb.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, price, currency, status) 
		VALUES (1, 'volleyball', 20251212, 1100, 44, 100, 'USD', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = apptype.TestDb.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, price, currency, status) 
		VALUES (2, 'volleyball', 20250612, 1000, 44, 100, 'RUB', 1)`)
	if err != nil {
		panic(err)
	}
}

// Creates a few games with the user
// Changes the table "GamesWithUsers"
func CreateUserScehdule() {
	_, err := apptype.TestDb.Exec("INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, statuspayment, status) VALUES (2323, $1, 0, 3, 'card', 1, 1)", 899)
	if err != nil {
		panic(err)
	}
	_, err = apptype.TestDb.Exec("INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, statuspayment, status) VALUES (123154, $1, 1, 6, 'card', 0, 1)", 899)
	if err != nil {
		panic(err)
	}
	_, err = apptype.TestDb.Exec("INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, statuspayment, status) VALUES (222, $1, 2, 7, 'cash', 0, 1)", 899)
	if err != nil {
		panic(err)
	}
}

// Deletes the user's records to games
// Changes the table "GamesWithUsers"
func DeleteUserSchedule() {
	_, err := apptype.TestDb.Exec("DELETE FROM GamesWithUsers WHERE userId = $1", 899)
	if err != nil {
		panic(err)
	}
}

// Deletes the schedule's games
// Changes the table "Schedule"
func DeleteSchedule() {
	_, err := apptype.TestDb.Exec("DELETE FROM Schedule WHERE gameId = 0 OR gameId = 1 OR gameId = 2 OR gameId = 3 OR gameId = 4")
	if err != nil {
		panic(err)
	}
}
