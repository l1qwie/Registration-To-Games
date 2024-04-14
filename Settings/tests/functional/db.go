package functional

import (
	"Settings/apptype"
)

// Checks updated language
func checkUpdtLang(lang string) bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Users WHERE userId = $1 AND language = $2", 899, lang).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

// Checks the deleted game
func checkDelGame() bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM GamesWithUsers WHERE gameId = $1 and userId = $2 and status = -1", 1, 899).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

// Chechs the change seats
func checkSeatsWereChange() bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND seats = $3 AND status = 1", 2, 899, 6).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
