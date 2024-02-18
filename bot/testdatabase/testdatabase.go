package testdatabase

import (
	"database/sql"
	"registrationtogames/fmtogram/types"

	_ "github.com/lib/pq"
)

func UpdateLanguage(language string, userId int) (err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET language = $1 WHERE userId = $2", language, userId)
	}
	db.Close()

	return err
}

func UpdateAction(action string, userId int) (err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET action = $1 WHERE userId = $2", action, userId)
	}
	db.Close()

	return err
}

func UpdateLevel(level, userId int) (err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET level = $1 WHERE userId = $2", level, userId)
	}
	db.Close()

	return err
}
