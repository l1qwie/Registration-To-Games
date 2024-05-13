package database

import (
	"Logs/types"
	"time"

	_ "github.com/lib/pq"
)

func ChechDb(t time.Time) bool {
	var count int
	err := types.Db.QueryRow("SELECT COUNT(*) FROM Actions WHERE userId = $1 AND act = $2 AND eventTime = $3", 1283829, "Registration", t).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func DeleteDb(t time.Time) {
	_, err := types.Db.Exec("DELETE FROM Actions WHERE userId = $1 AND act = $2 AND eventTime = $3", 1283829, "Registration", t)
	if err != nil {
		panic(err)
	}
}
