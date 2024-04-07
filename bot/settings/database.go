package settings

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"database/sql"
	"fmt"
)

func FindUserGames(userId int, fm *formatter.Formatter) bool {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM GamesWithUsers WHERE userId = $1 AND status != -1", userId)
	if err != nil {
		fm.Error(err)
	}
	rows.Next()
	coun := 0
	err = rows.Scan(&coun)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return coun > 0
}

func updateLanguage(userId int, lang string, fm *formatter.Formatter) string {
	_, err := types.Db.Exec("UPDATE Users SET language = $1, customlanguage = True WHERE userId = $2", lang, userId)
	if err != nil {
		fm.Error(err)
	}
	return lang
}

func findUserGame(gameId, userId int, fm *formatter.Formatter) bool {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM GamesWithUsers WHERE gameId = $1 AND userId = $2", gameId, userId)
	if err != nil {
		fm.Error(err)
	}
	cc := 0
	rows.Next()
	err = rows.Scan(&cc)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return cc > 0
}

func selectUserSchedule(userId, limit, lp int, language string, fm *formatter.Formatter) []*bottypes.Game {
	var (
		rows                 *sql.Rows
		err                  error
		sch                  []*bottypes.Game
		i, date, time, statp int
	)
	rows, err = types.Db.Query(`SELECT s.gameId, sport, date, time, price, currency, gwu.payment, gwu.seats, gwu.statuspayment
								FROM Schedule s 
								JOIN GamesWithUsers gwu ON s.gameId = gwu.GameId
								WHERE gwu.status != -1 AND gwu.userId = $1
								ORDER BY s.gameId DESC LIMIT $2 OFFSET $3`, userId, limit, lp)
	if err != nil {
		fm.Error(err)
	}
	sch = make([]*bottypes.Game, limit)
	for rows.Next() {
		sch[i] = &bottypes.Game{}
		err = rows.Scan(&sch[i].Id, &sch[i].Sport, &date, &time, &sch[i].Price, &sch[i].Currency, &sch[i].Payment, &sch[i].Seats, &statp)
		if err != nil {
			fm.Error(err)
		}
		if statp == 1 {
			sch[i].StatusPayment = dictionary.Dictionary[language]["Paid"]
		} else {
			sch[i].StatusPayment = dictionary.Dictionary[language]["WaitForPaid"]
		}
		sch[i].Date = forall.FromIntToStrDate(date)
		sch[i].Time = forall.FromIntToStrTime(time)
		i++
	}
	defer rows.Close()
	return sch
}

func statPayment(gameId, userId int, fm *formatter.Formatter) bool {
	rows, err := types.Db.Query("SELECT statuspayment FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = 1", gameId, userId)
	if err != nil {
		fm.Error(err)
	}
	stat := 0
	rows.Next()
	err = rows.Scan(&stat)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return stat != 1
}

func delTheGame(seats, gameId, userId int, fm *formatter.Formatter) {
	_, err := types.Db.Exec("UPDATE GamesWithUsers SET status = -1 WHERE gameId = $1 AND userId = $2", gameId, userId)
	if err != nil {
		fm.Error(err)
	}
	_, err = types.Db.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", seats, gameId)
	if err != nil {
		fm.Error(err)
	}
}

func selPaymethod(gameId, userId int, fm *formatter.Formatter) bool {
	rows, err := types.Db.Query("SELECT payment FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = 1", gameId, userId)
	if err != nil {
		fm.Error(err)
	}
	p := ""
	rows.Next()
	err = rows.Scan(&p)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return p == "card"
}

func findSomeSeats(gameId, userId, wantS int, fm *formatter.Formatter) (int, int, bool) {
	rows, err := types.Db.Query("SELECT s.seats, gwu.seats FROM Schedule s JOIN GamesWithUsers gwu ON s.gameId = gwu.gameId WHERE gwu.userId = $1 AND s.gameId = $2 and gwu.status = 1", userId, gameId)
	if err != nil {
		fm.Error(err)
	}
	schs := 0
	uss := 0
	rows.Next()
	err = rows.Scan(&schs, &uss)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return schs, uss, (schs + uss) >= wantS
}

func updtPayment(gameId, userId int, paymeth string, fm *formatter.Formatter) {
	fmt.Println(paymeth, gameId, userId)
	_, err := types.Db.Exec("UPDATE GamesWithUsers SET payment = $1 WHERE gameId = $2 AND userId = $3 AND status = 1", paymeth, gameId, userId)
	if err != nil {
		fm.Error(err)
	}
}

func updtSeats(gameId, userId, genS, oldS, newS int, fm *formatter.Formatter) {
	_, err := types.Db.Exec("UPDATE GamesWithUsers SET seats = $1 WHERE gameId = $2 AND userId = $3", newS, gameId, userId)
	if err != nil {
		fm.Error(err)
	}
	_, err = types.Db.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", (genS+oldS)-newS, gameId)
	if err != nil {
		fm.Error(err)
	}
}
