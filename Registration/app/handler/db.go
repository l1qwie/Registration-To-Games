package handler

import (
	"Registraion/app/dict"
	apptype "Registraion/apptype"
	"Registraion/fmtogram/types"
	"database/sql"
)

// Selects the schedule of the app
func selectTheSchedule(limit, offset int, language string, f func(error)) []*apptype.Game {
	var (
		rows           *sql.Rows
		err            error
		request, sport string
		i, date, time  int
	)
	schedule := make([]*apptype.Game, limit)
	request = `SELECT gameId, sport, date, time, seats FROM Schedule WHERE (status != -1) ORDER BY Schedule DESC LIMIT $1 OFFSET $2`
	rows, err = types.Db.Query(request, limit, offset)
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
	rows, err := types.Db.Query("SELECT gameId FROM Schedule WHERE status != -1")
	if err != nil {
		f(err)
	}
	cn := 0
	for rows.Next() {
		err = rows.Scan(&cn)
		if err != nil {
			f(err)
		}
	}
	defer rows.Close()
	return cn > 0
}

// Tries to find the game that user choosed
func findThatGame(gameId int, f func(error)) bool {
	counter := 0
	rows, err := types.Db.Query("SELECT COUNT(*) FROM Schedule WHERE status != -1 AND gameId = $1", gameId)
	if err != nil {
		f(err)
	} else {
		rows.Next()
		err = rows.Scan(&counter)
		if err != nil {
			f(err)
		}
	}
	defer rows.Close()
	return counter > 0
}

// Selects some information about the game
func selectThePrice(gameId int, f func(error)) (price, space int, currency string) {
	rows, err := types.Db.Query("SELECT price, seats, currency FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		f(err)
	}
	rows.Next()
	err = rows.Scan(&price, &space, &currency)
	if err != nil {
		f(err)
	}
	defer rows.Close()
	return price, space, currency
}

// Checks is there any free seats (space) to registraite the user
func howManyIsLeft(gameId int, wishfulseats int, f func(error)) bool {
	rows, err := types.Db.Query("SELECT seats FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		f(err)
	}
	rows.Next()
	seats := 0
	err = rows.Scan(&seats)
	if err != nil {
		f(err)
	}
	defer rows.Close()
	return wishfulseats < seats
}

// Adds a new string in the table "GamesWithUsers"
// Updates seats of the game in the table "Schedule"
func completeRegistration(userId, gameId, seats int, payment string, f func(error)) {
	var (
		rows      *sql.Rows
		gameSeats int
	)
	_, err := types.Db.Exec("INSERT INTO GamesWithUsers (id, userId, gameId, seats, payment, status) VALUES (nextval('gameswithusers_id_seq'), $1, $2, $3, $4, 1)", userId, gameId, seats, payment)
	if err == nil {
		rows, err = types.Db.Query("SELECT seats FROM Schedule WHERE gameId = $1", gameId)
	}
	if err == nil {
		rows.Next()
		err = rows.Scan(&gameSeats)
		if err != nil {
			f(err)
		}
	}
	_, err = types.Db.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", gameSeats-seats, gameId)
	if err != nil {
		f(err)
	}
	defer rows.Close()
}

// Selects all game data
func selectDetailOfGame(gameId int, language string, f func(error)) (details *apptype.Game) {
	var (
		rows       *sql.Rows
		err        error
		request    string
		date, time int
	)
	details = new(apptype.Game)
	request = `SELECT gameId, sport, date, time, seats, latitude, longitude, address, price, currency FROM Schedule WHERE gameId = $1`
	rows, err = types.Db.Query(request, gameId)
	if err != nil {
		f(err)
	}
	rows.Next()
	err = rows.Scan(&details.Id, &details.Sport, &date, &time, &details.Seats, &details.Lattitude, &details.Longitude, &details.Address, &details.Price, &details.Currency)
	if err != nil {
		f(err)
	}
	details.Sport = dict.Dictionary[language][details.Sport]
	details.Date = fromIntToStrDate(date)
	details.Time = fromIntToStrTime(time)
	defer rows.Close()
	return details
}
