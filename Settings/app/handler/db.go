package handler

import (
	"Settings/api/client"
	"Settings/api/producer"
	"Settings/apptype"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type UsuallConn struct {
	DB *sql.DB
}

type ApiConn struct {
	DB *sql.DB
}

// Tries to find any game with the user
func (uc *UsuallConn) findUserGame(userId int, f func(error)) bool {
	var count int
	producer.InterLogs("Start function Settings.findUserGame()",
		fmt.Sprintf("UserId: %d, userId (int): %d, f (func(error))): %T", userId, userId, f))
	err := uc.DB.QueryRow("SELECT COUNT(*) FROM GamesWithUsers WHERE userId = $1 AND status != -1", userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

// Tries to find the game that user choosed
func (uc *UsuallConn) findThatUserGame(gameId, userId int, f func(error)) bool {
	var count int
	producer.InterLogs("Start function Settings.findThatUserGame()",
		fmt.Sprintf("UserId: %d, gameId (int): %d, userId (int): %d, f (func(error))): %T", userId, gameId, userId, f))
	err := uc.DB.QueryRow("SELECT COUNT(*) FROM GamesWithUsers WHERE gameId = $1 AND userId = $2", gameId, userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

// Updates the app's language in the database
func (uc *UsuallConn) updtLanguage(userId int, lang string, f func(error)) (string, bool) {
	producer.InterLogs("Start function Settings.updtLanguage()",
		fmt.Sprintf("UserId: %d, userId (int): %d, lang (string): %s, f (func(error))): %T", userId, userId, lang, f))
	_, err := uc.DB.Exec("UPDATE Users SET language = $1, customlanguage = True WHERE userId = $2", lang, userId)
	if err != nil {
		f(err)
	}
	return lang, err == nil
}

// Prepares the schedule of user's games
func (uc *UsuallConn) selectUserSchedule(userId, limit, lp int, f func(error), dict map[string]string) []*apptype.Game {
	var (
		i, date, time, statp int
	)
	producer.InterLogs("Start function Settings.selectUserSchedule()",
		fmt.Sprintf("UserId: %d, userId (int): %d, limit (int): %d, lp (int): %d, f (func(error))): %T", userId, userId, limit, lp, f))
	rows, err := uc.DB.Query(`SELECT s.gameId, sport, date, time, price, currency, gwu.payment, gwu.seats, gwu.statuspayment
								FROM Schedule s 
								JOIN GamesWithUsers gwu ON s.gameId = gwu.GameId
								WHERE gwu.status != -1 AND gwu.userId = $1
								ORDER BY s.gameId DESC LIMIT $2 OFFSET $3`, userId, limit, lp)
	if err != nil {
		f(err)
	}
	sch := make([]*apptype.Game, limit)
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
func (uc *UsuallConn) findSomeSeats(gameId, userId, wantS int, f func(error)) (int, int, bool) {
	var schs, uss int
	producer.InterLogs("Start function Settings.findSomeSeats()",
		fmt.Sprintf("UserId: %d, gameId (int): %d, userId (int): %d, wantS (int): %d, f (func(error))): %T", userId, gameId, userId, wantS, f))
	err := uc.DB.QueryRow("SELECT s.seats, gwu.seats FROM Schedule s JOIN GamesWithUsers gwu ON s.gameId = gwu.gameId WHERE gwu.userId = $1 AND s.gameId = $2 and gwu.status = 1", userId, gameId).Scan(&schs, &uss)
	if err != nil {
		f(err)
	}
	return schs, uss, (schs + uss) >= wantS
}

func endOfTransaction(err error, DB *sql.DB, f func(error)) {
	var res string
	if err != nil {
		res = "ROLLBACK"
	} else {
		res = "COMMIT"
	}
	_, err = DB.Exec(res)
	if err != nil {
		f(err)
	}
}

// Deletes the game in the database
func (uc *UsuallConn) delTheGame(seats, gameId, userId int, f func(error)) bool {
	var err error
	producer.InterLogs("Start function Settings.delTheGame()",
		fmt.Sprintf("UserId: %d, seats (int): %d, gameId (int): %d, userId (int): %d, f (func(error))): %T", userId, seats, gameId, userId, f))
	_, err = uc.DB.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	if err == nil {
		_, err = uc.DB.Exec("UPDATE GamesWithUsers SET status = -1 WHERE gameId = $1 AND userId = $2", gameId, userId)
	}
	if err == nil {
		_, err = uc.DB.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", seats, gameId)
	}
	defer endOfTransaction(err, uc.DB, f)
	return err == nil
}

// Select the statistics of pay status
func (uc *UsuallConn) statPayment(gameId, userId int, f func(error)) bool {
	var stat int
	producer.InterLogs("Start function Settings.statPayment()",
		fmt.Sprintf("UserId: %d, gameId (int): %d, userId (int): %d, f (func(error))): %T", userId, gameId, userId, f))
	err := uc.DB.QueryRow("SELECT statuspayment FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = 1", gameId, userId).Scan(&stat)
	if err != nil {
		f(err)
	}
	return stat != 1
}

// Updates the seats of the user of the user's game
func (uc *UsuallConn) updtSeats(gameId, userId, genS, oldS, newS int, f func(error)) bool {
	var err error
	producer.InterLogs("Start function Settings.updtSeats()",
		fmt.Sprintf("UserId: %d, gameId (int): %d, userId (int): %d, genS (int): %d, oldS (int): %d, newS (int): %d, f (func(error))): %T", userId, gameId, userId, genS, oldS, newS, f))
	_, err = uc.DB.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	if err == nil {
		_, err = uc.DB.Exec("UPDATE GamesWithUsers SET seats = $1 WHERE gameId = $2 AND userId = $3", newS, gameId, userId)
	}
	if err == nil {
		_, err = uc.DB.Exec("UPDATE Schedule SET seats = $2 WHERE gameId = $1;", gameId, (genS+oldS)-newS)
	}
	defer endOfTransaction(err, uc.DB, f)
	return err == nil
}

// Select the paymethod and compare it with "card"
// then return true if compare is also true
func (uc *UsuallConn) selPaymethod(gameId, userId int, f func(error)) bool {
	var p string
	producer.InterLogs("Start function Settings.selPaymethod()",
		fmt.Sprintf("UserId: %d, gameId (int): %d, userId (int): %d, f (func(error))): %T", userId, gameId, userId, f))
	err := uc.DB.QueryRow("SELECT payment FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = 1", gameId, userId).Scan(&p)
	if err != nil {
		f(err)
	}
	return p == "card"
}

// Updates the payment of the user's game
func (uc *UsuallConn) updtPayment(gameId, userId int, paymeth string, f func(error)) {
	producer.InterLogs("Start function Settings.updtPayment()",
		fmt.Sprintf("UserId: %d, gameId (int): %d, userId (int): %d, paymeth (string): %s, f (func(error))): %T", userId, gameId, userId, paymeth, f))
	_, err := uc.DB.Exec("UPDATE GamesWithUsers SET payment = $1 WHERE gameId = $2 AND userId = $3 AND status = 1", paymeth, gameId, userId)
	if err != nil {
		f(err)
	}
}

func (ac *ApiConn) UpdateTheSchedule(date, time, status int, g *apptype.Game, act string) error {
	var request string
	producer.InterLogs("Start function Settings.UpdateTheSchedule()",
		fmt.Sprintf("date (int): %d, time (int): %d, status (int): %d, g (*apptype.Game): %v, act (string): %s", date, time, status, g, act))
	if act == "new" {
		request = "INSERT INTO Schedule (gameId, sport, date, time, seats, price, currency, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	} else if act == "change" {
		request = "UPDATE Schedule SET gameId = $1, sport = $2, date = $3, time = $4, price = $5, currency = $6, seats = $7, status = $8 WHERE gameId = $1"
	}
	_, err := ac.DB.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	if err == nil {
		_, err = ac.DB.Exec(request, g.Id, g.Sport, date, time, g.Price, g.Currency, g.Seats, status)
	}
	defer endOfTransaction(err, ac.DB, func(err error) { log.Print(err) })
	return err
}

func (ac *ApiConn) UpdateGWU(g *apptype.GWU, act string) error {
	var request string
	producer.InterLogs("Start function Settings.UpdateGWU()",
		fmt.Sprintf("g (*apptype.GWU): %v, act (string): %s", g, act))
	if act == "new" {
		request = "INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, statuspayment, status) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	} else if act == "change" {
		request = "UPDATE GamesWithUsers SET id = $1, userId = $2, gameId = $3, seats = $4, payment = $5, statuspayment = $6, status = $7 WHERE gameId = $3"
	}
	_, err := ac.DB.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	if err == nil {
		_, err = ac.DB.Exec(request, g.Id, g.UserId, g.GameId, g.Seats, g.Payment, g.Statpay, g.Status)
	}
	defer endOfTransaction(err, ac.DB, func(err error) { log.Print(err) })
	return err
}

func (ac *ApiConn) UpdateUsers(userId int, lang string, custlang bool) error {
	producer.InterLogs("Start function Settings.UpdateUsers()",
		fmt.Sprintf("userId (int): %d, lang (string): %s, custlang (bool): %v", userId, lang, custlang))
	_, err := ac.DB.Exec("INSERT INTO Users (userId, language, customlanguage) VALUES ($1, $2, $3)", userId, lang, custlang)
	return err
}

func (uc *UsuallConn) fill(gameId int) (*client.Upd, error) {
	producer.InterLogs("Start function Settings.fill()",
		fmt.Sprintf("gameId (int): %d", gameId))
	u := new(client.Upd)
	err := uc.DB.QueryRow("SELECT gameId, sport, date, time, price, currency, seats, status FROM Schedule WHERE gameId = $1",
		gameId).Scan(&u.GameId, &u.Sport, &u.Date, &u.Time, &u.Price, &u.Currency, &u.Seats, &u.Status)
	if err == nil {
		u.Datestr = fromIntToStrDate(u.Date)
		u.Timestr = fromIntToStrTime(u.Time)
	}
	return u, err
}
