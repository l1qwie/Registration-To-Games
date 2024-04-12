package tests

import "Settings/apptype"

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
