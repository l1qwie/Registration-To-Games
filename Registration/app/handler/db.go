package handler

import (
	"Registration/api/client"
	"Registration/api/producer"
	"Registration/app/dict"
	"Registration/apptype"
	"database/sql"
	"fmt"
	"log"
)

// Selects the schedule of the app
func selectTheSchedule(limit, offset int, language string, f func(error)) []*apptype.Game {
	var (
		rows           *sql.Rows
		err            error
		request, sport string
		i, date, time  int
	)
	producer.InterLogs("Start function Registration.selectTheSchedule()",
		fmt.Sprintf("limit (int): %d, offset (int): %d, langauge (string): %s, f (func(error)): %T", limit, offset, language, f))
	schedule := make([]*apptype.Game, limit)
	request = `SELECT gameId, sport, date, time, seats FROM Schedule WHERE (status != -1) ORDER BY Schedule DESC LIMIT $1 OFFSET $2`
	rows, err = apptype.Db.Query(request, limit, offset)
	if err != nil {
		f(err)
	}
	for rows.Next() {
		schedule[i] = &apptype.Game{}
		err = rows.Scan(&schedule[i].Id, &sport, &date, &time, &schedule[i].Seats)
		if err != nil {
			f(err)
		}
		schedule[i].Sport = dict.Dictionary[language][sport]
		schedule[i].Date = fromIntToStrDate(date)
		schedule[i].Time = fromIntToStrTime(time)
		i++
	}
	defer rows.Close()
	return schedule
}

// Tries to find any active(status != -1) game
func findGame(f func(error)) bool {
	var count int
	producer.InterLogs("Start function Registration.findGame()",
		fmt.Sprintf("f (func(error)): %T", f))
	rows, err := apptype.Db.Query("SELECT gameId FROM Schedule WHERE status != -1")
	if err != nil {
		f(err)
	}
	for rows.Next() {
		count++
	}
	defer rows.Close()
	log.Print("count =", count)
	return count > 0
}

// Tries to find the game that user choosed
func findThatGame(gameId int, f func(error)) bool {
	var count int
	producer.InterLogs("Start function Registration.findThatGame()",
		fmt.Sprintf("gameId (int): %d, f (func(error)): %T", gameId, f))
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE status != -1 AND gameId = $1", gameId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

// Selects some information about the game
func selectThePrice(gameId int, f func(error)) (price, space int, currency string) {
	producer.InterLogs("Start function Registration.selectThePrice()",
		fmt.Sprintf("gameId (int): %d, f (func(error)): %T", gameId, f))
	err := apptype.Db.QueryRow("SELECT price, seats, currency FROM Schedule WHERE gameId = $1", gameId).Scan(&price, &space, &currency)
	if err != nil {
		f(err)
	}
	return price, space, currency
}

// Checks is there any free seats (space) to registraite the user
func howManyIsLeft(gameId int, wishfulseats int, f func(error)) bool {
	var seats int
	producer.InterLogs("Start function Registration.howManyIsLeft()",
		fmt.Sprintf("gameId (int): %d, wishfulseats (int): %d, f (func(error)): %T", gameId, wishfulseats, f))
	err := apptype.Db.QueryRow("SELECT seats FROM Schedule WHERE gameId = $1", gameId).Scan(&seats)
	if err != nil {
		f(err)
	}
	return wishfulseats < seats
}

func endOfTransaction(err error, f func(error)) {
	var res string
	if err != nil {
		res = "ROLLBACK"
	} else {
		res = "COMMIT"
	}
	_, err = apptype.Db.Exec(res)
	if err != nil {
		f(err)
	}
}

// Adds a new string in the table "GamesWithUsers"
// Updates seats of the game in the table "Schedule"
func completeRegistration(userId, gameId, seats int, payment string, f func(error)) bool {
	var gameSeats int
	producer.InterLogs("Start function Registration.completeRegistration()",
		fmt.Sprintf("userId (int): %d, gameId (int): %d, seats (int): %d, payment (string): %s, f (func(error)): %T", userId, gameId, seats, payment, f))
	_, err := apptype.Db.Exec("START TRANSACTION")
	if err == nil {
		_, err = apptype.Db.Exec("INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, status) VALUES (nextval('gameswithusers_id_seq'), $1, $2, $3, $4, 1)", userId, gameId, seats, payment)
	}
	if err == nil {
		err = apptype.Db.QueryRow("SELECT seats FROM Schedule WHERE gameId = $1", gameId).Scan(&gameSeats)
	}
	if err == nil {
		_, err = apptype.Db.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", gameSeats-seats, gameId)
	}
	defer endOfTransaction(err, f)
	return err == nil
}

// Selects all game data
func selectDetailOfGame(gameId int, language string, f func(error)) *apptype.Game {
	var date, time int
	producer.InterLogs("Start function Registration.selectDetailOfGame()",
		fmt.Sprintf("gameId (int): %d, language (string): %s, f (func(error)): %T", gameId, language, f))
	details := new(apptype.Game)
	request := `SELECT gameId, sport, date, time, seats, latitude, longitude, address, price, currency FROM Schedule WHERE gameId = $1`
	err := apptype.Db.QueryRow(request, gameId).Scan(&details.Id, &details.Sport, &date, &time, &details.Seats, &details.Lattitude, &details.Longitude, &details.Address, &details.Price, &details.Currency)
	if err != nil {
		f(err)
	}
	details.Sport = dict.Dictionary[language][details.Sport]
	details.Date = fromIntToStrDate(date)
	details.Time = fromIntToStrTime(time)
	return details
}

func UpdateTheSchedule(date, time, status int, g *apptype.Game, act string) error {
	var request string
	producer.InterLogs("Start function Registration.UpdateTheSchedule()",
		fmt.Sprintf("date (int): %d, time (int): %d, status (int): %d, g (*apptype.Game): %v, act (string): %s", date, time, status, g, act))
	if act == "new" {
		request = `INSERT INTO Schedule (gameId, sport, date, time, seats, price, currency, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	} else if act == "change" || act == "del" {
		request = `UPDATE Schedule SET gameId = $1, sport = $2, date = $3, time = $4, seats = $5, price = $6, currency = $7, status = $8 WHERE gameId = $1`
	}
	_, err := apptype.Db.Exec(request, g.Id, g.Sport, date, time, g.Seats, g.Price, g.Currency, status)
	if err != nil {
		panic(err)
	}
	return err
}

func fill(gameId, userId int) (*client.Upd, error) {
	producer.InterLogs("Start function Registration.fill()",
		fmt.Sprintf("gameId (int): %d, userId (int): %d", gameId, userId))
	u := new(client.Upd)
	err := apptype.Db.QueryRow("SELECT gameId, sport, date, time, price, currency, seats, status FROM Schedule WHERE gameId = $1",
		gameId).Scan(&u.GameId, &u.Sport, &u.Date, &u.Time, &u.Price, &u.Currency, &u.Seats, &u.Status)
	if err == nil {
		u.Datestr = fromIntToStrDate(u.Date)
		u.Timestr = fromIntToStrTime(u.Time)
		err = apptype.Db.QueryRow("SELECT id, userId, gameId, seats, payment, statuspayment, status FROM GamesWithUsers WHERE userId = $1 AND gameId = $2",
			userId, gameId).Scan(&u.Id, &u.UserId, &u.GameIdGWU, &u.SeatsGWU, &u.Payment, &u.Statpay, &u.StatusGWU)
	}
	return u, err
}
