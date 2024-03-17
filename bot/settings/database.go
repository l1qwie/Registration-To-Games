package settings

import (
	"RegistrationToGames/fmtogram/types"
)

func FindUserGames(userId int) bool {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM GamesWithUsers WHERE userId = $1 AND status != -1", userId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	coun := 0
	err = rows.Scan(&coun)
	if err != nil {
		panic(err)
	}
	return coun > 0
}

func updateLanguage(userId int, lang string) string {
	_, err := types.Db.Exec("UPDATE Users SET language = $1, customlanguage = True WHERE userId = $2", lang, userId)
	if err != nil {
		panic(err)
	}
	return lang
}
