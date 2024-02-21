package welcome

import (
	"database/sql"
	"registrationtogames/fmtogram/types"

	_ "github.com/lib/pq"
)

func CompletionOfRegistration(userId int) (err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET setup_reg = 'completed' WHERE userId = $1", userId)
	}
	db.Close()

	return err
}
