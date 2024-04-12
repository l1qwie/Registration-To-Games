package handler

import (
	"Settings/apptype"
)

// Tries find any game with the user
func findUserGame(userId int, f func(error)) bool {
	coun := 0
	rows, err := apptype.Db.Query("SELECT COUNT(*) FROM GamesWithUsers WHERE userId = $1 AND status != -1", userId)
	if err != nil {
		f(err)
	}
	rows.Next()
	err = rows.Scan(&coun)
	if err != nil {
		f(err)
	}
	defer rows.Close()
	return coun > 0
}

func updtLanguage(userId int, lang string, f func(error)) string {
	_, err := apptype.Db.Exec("UPDATE Users SET language = $1, customlanguage = True WHERE userId = $2", lang, userId)
	if err != nil {
		f(err)
	}
	return lang
}
