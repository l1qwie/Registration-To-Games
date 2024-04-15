package handler

import (
	"Settings/apptype"
	"database/sql"

	_ "github.com/lib/pq"
)

// Tries to find any game with the user
func findUserGame(userId int, f func(error)) bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM GamesWithUsers WHERE userId = $1 AND status != -1", userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

// Tries to find the game that user choosed
func findThatUserGame(gameId, userId int, f func(error)) bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM GamesWithUsers WHERE gameId = $1 AND userId = $2", gameId, userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

// Updates the app's language in the database
func updtLanguage(userId int, lang string, f func(error)) string {
	_, err := apptype.Db.Exec("UPDATE Users SET language = $1, customlanguage = True WHERE userId = $2", lang, userId)
	if err != nil {
		f(err)
	}
	return lang
}

// Prepares the schedule of user's games
func selectUserSchedule(userId, limit, lp int, f func(error), dict map[string]string) []*apptype.Game {
	var (
		rows                 *sql.Rows
		err                  error
		sch                  []*apptype.Game
		i, date, time, statp int
	)
	rows, err = apptype.Db.Query(`SELECT s.gameId, sport, date, time, price, currency, gwu.payment, gwu.seats, gwu.statuspayment
								FROM Schedule s 
								JOIN GamesWithUsers gwu ON s.gameId = gwu.GameId
								WHERE gwu.status != -1 AND gwu.userId = $1
								ORDER BY s.gameId DESC LIMIT $2 OFFSET $3`, userId, limit, lp)
	if err != nil {
		f(err)
	}
	sch = make([]*apptype.Game, limit)
	for rows.Next() {
		sch[i] = &apptype.Game{}
		err = rows.Scan(&sch[i].Id, &sch[i].Sport, &date, &time, &sch[i].Price, &sch[i].Currency, &sch[i].Payment, &sch[i].Seats, &statp)
		if err != nil {
			f(err)
		}
		if statp == 1 {
			sch[i].StatusPayment = dict["Paid"]
		} else {
			sch[i].StatusPayment = dict["WaitForPaid"]
		}
		sch[i].Date = fromIntToStrDate(date)
		sch[i].Time = fromIntToStrTime(time)
		i++
	}
	defer rows.Close()
	return sch
}

// Tries to find some free seats in the game
func findSomeSeats(gameId, userId, wantS int, f func(error)) (int, int, bool) {
	var schs, uss int
	err := apptype.Db.QueryRow("SELECT s.seats, gwu.seats FROM Schedule s JOIN GamesWithUsers gwu ON s.gameId = gwu.gameId WHERE gwu.userId = $1 AND s.gameId = $2 and gwu.status = 1", userId, gameId).Scan(&schs, &uss)
	if err != nil {
		f(err)
	}
	return schs, uss, (schs + uss) >= wantS
}

// Deletes the game in the database
func delTheGame(seats, gameId, userId int, f func(error)) {
	_, err := apptype.Db.Exec("UPDATE GamesWithUsers SET status = -1 WHERE gameId = $1 AND userId = $2", gameId, userId)
	if err != nil {
		f(err)
	}
	_, err = apptype.Db.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", seats, gameId)
	if err != nil {
		f(err)
	}
}

// Select the statistics of pay status
func statPayment(gameId, userId int, f func(error)) bool {
	var stat int
	err := apptype.Db.QueryRow("SELECT statuspayment FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = 1", gameId, userId).Scan(&stat)
	if err != nil {
		f(err)
	}
	return stat != 1
}

// Updates the seats of the user of the user's game
func updtSeats(gameId, userId, genS, oldS, newS int, f func(error)) {
	_, err := apptype.Db.Exec("UPDATE GamesWithUsers SET seats = $1 WHERE gameId = $2 AND userId = $3", newS, gameId, userId)
	if err != nil {
		f(err)
	}
	_, err = apptype.Db.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", (genS+oldS)-newS, gameId)
	if err != nil {
		f(err)
	}
}

// Select the paymethod and compare it with "card"
// then return true if compare is also true
func selPaymethod(gameId, userId int, f func(error)) bool {
	var p string
	err := apptype.Db.QueryRow("SELECT payment FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = 1", gameId, userId).Scan(&p)
	if err != nil {
		f(err)
	}
	return p == "card"
}

// Updates the payment of the user's game
func updtPayment(gameId, userId int, paymeth string, f func(error)) {
	_, err := apptype.Db.Exec("UPDATE GamesWithUsers SET payment = $1 WHERE gameId = $2 AND userId = $3 AND status = 1", paymeth, gameId, userId)
	if err != nil {
		f(err)
	}
}
