package functional

import (
	"Registraion/fmtogram/types"
	"database/sql"
)

func checkGamesWithUsersTable(userId, gameId, seats int, payment string) bool {
	var (
		rows    *sql.Rows
		err     error
		request string
		counter int
	)
	request = `SELECT COUNT(*) FROM GamesWithUsers WHERE userId = $1 AND gameId = $2 AND seats = $3 AND payment = $4` //AND statuspayment = 0 AND status = 0
	rows, err = types.Db.Query(request, userId, gameId, seats, payment)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&counter)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return counter > 0
}

func checkUpdatedSchedule(gameId, seats int) bool {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM Schedule WHERE gameId = $1 AND seats = $2", gameId, 55-seats)
	if err != nil {
		panic(err)
	}
	counter := 0
	rows.Next()
	err = rows.Scan(&counter)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return counter > 0
}

func afterBestWishes() {
	if !checkGamesWithUsersTable(477, 2, 2, "card") {
		panic("user doesn't exist in table GamesWithUsers")
	}
	if !checkUpdatedSchedule(2, 2) {
		panic("schedule wasn't updated")
	}
}
