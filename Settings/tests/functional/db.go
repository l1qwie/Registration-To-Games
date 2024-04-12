package functional

import (
	"Settings/apptype"
)

// Checks updated language
func checkUpdtLang(lang string) bool {
	count := 0
	rows, err := apptype.Db.Query("SELECT COUNT(*) FROM Users WHERE userId = $1 AND language = $2", 899, lang)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&count)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return count > 0
}
