package grpc

import "User/fmtogram/types"

func checkChangeLanguage() bool {
	var count int
	err := types.Db.QueryRow("SELECT COUNT(*) FROM Users WHERE userId = $1 AND language = $2 AND customlanguage = $3", 999, "en", true).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func createUser(userId int) {
	_, err := types.Db.Exec("INSERT INTO Users (userId, language, customlanguage) VALUES ($1, $2, $3)", userId, "ru", false)
	if err != nil {
		panic(err)
	}
}

func deleteUser(userId int) {
	_, err := types.Db.Exec("DELETE FROM Users WHERE userId = $1", userId)
	if err != nil {
		panic(err)
	}
}
