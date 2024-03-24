package registration

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/types"
	"database/sql"
	"fmt"
)

func selectTheSchedule(limit, offset int, language string) []*bottypes.Game {
	var (
		rows           *sql.Rows
		err            error
		request, sport string
		i, date, time  int
	)
	schedule := make([]*bottypes.Game, limit)
	request = `SELECT gameId, sport, date, time, seats FROM Schedule WHERE (status != -1) ORDER BY Schedule DESC LIMIT $1 OFFSET $2`
	rows, err = types.Db.Query(request, limit, offset)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		schedule[i] = &bottypes.Game{}
		err = rows.Scan(&schedule[i].Id, &sport, &date, &time, &schedule[i].Seats)
		if err != nil {
			panic(err)
		}
		schedule[i].Sport = dictionary.Dictionary[language][sport]
		schedule[i].Date = forall.FromIntToStrDate(date)
		schedule[i].Time = forall.FromIntToStrTime(time)
		i++
	}
	defer rows.Close()
	return schedule
}

func FindAGame(gameId int) bool {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM Schedule WHERE status != -1 AND gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	counter := 0
	err = rows.Scan(&counter)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return counter > 0
}

func selectThePrice(gameId int) (price, space int, currency string) {
	rows, err := types.Db.Query("SELECT price, seats, currency FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&price, &space, &currency)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return price, space, currency
}

func howManyIsLeft(gameId int, wishfulseats int) bool {
	rows, err := types.Db.Query("SELECT seats FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	seats := 0
	err = rows.Scan(&seats)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return wishfulseats < seats
}

func completeRegistration(userId, gameId, seats int, payment string) (err error) {
	var (
		rows      *sql.Rows
		gameSeats int
	)
	fmt.Println(userId, gameId, seats, payment)
	_, err = types.Db.Exec("INSERT INTO GamesWithUsers (userId, gameId, seats, payment, status) VALUES ($1, $2, $3, $4, 1)", userId, gameId, seats, payment)
	if err == nil {
		rows, err = types.Db.Query("SELECT seats FROM Schedule WHERE gameId = $1", gameId)
	}
	if err == nil {
		rows.Next()
		err = rows.Scan(&gameSeats)
		if err != nil {
			panic(err)
		}
	}
	_, err = types.Db.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", gameSeats-seats, gameId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return err
}

func selectDetailOfGame(gameId int, language string) (details *bottypes.Game) {
	var (
		rows       *sql.Rows
		err        error
		request    string
		date, time int
	)
	details = new(bottypes.Game)
	request = `SELECT gameId, sport, date, time, seats, latitude, longitude, address, price, currency FROM Schedule WHERE gameId = $1`
	rows, err = types.Db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&details.Id, &details.Sport, &date, &time, &details.Seats, &details.Lattitude, &details.Longitude, &details.Address, &details.Price, &details.Currency)
	if err != nil {
		panic(err)
	}
	details.Sport = dictionary.Dictionary[language][details.Sport]
	details.Date = forall.FromIntToStrDate(date)
	details.Time = forall.FromIntToStrTime(time)
	defer rows.Close()
	return details
}
