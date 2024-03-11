package registration

import (
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/types"
	"database/sql"
)

func selectTheSchedule(limit, offset int, language string) []*forall.Game {
	var (
		rows           *sql.Rows
		err            error
		request, sport string
		i, date, time  int
	)
	schedule := make([]*forall.Game, limit)
	//for j := 0; j < limit; j++ {
	//		schedule[j] = &forall.Game{}
	//	}
	request = `SELECT gameId, sport, date, time, seats FROM Schedule WHERE (status != -1) ORDER BY Schedule DESC LIMIT $1 OFFSET $2`
	rows, err = types.Db.Query(request, limit, offset)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		schedule[i] = &forall.Game{}
		err = rows.Scan(&schedule[i].Id, &sport, &date, &time, &schedule[i].Seats)
		if err != nil {
			panic(err)
		}
		schedule[i].Sport = dictionary.Dictionary[language][sport]
		schedule[i].Date = forall.FromIntToStrDate(date)
		schedule[i].Time = forall.FromIntToStrTime(time)
		i++
	}
	return schedule
}

func FindAGame(gameId int) (detected bool) {
	var (
		rows    *sql.Rows
		err     error
		request string
		counter int
	)
	request = `SELECT COUNT(*) FROM Schedule WHERE status != -1 AND gameId = $1`
	rows, err = types.Db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&counter)
	if err != nil {
		panic(err)
	}

	if counter > 0 {
		detected = true
	}
	return detected
}

func selectThePrice(gameId int) (price, space int, currency string) {
	var (
		rows    *sql.Rows
		err     error
		request string
	)
	request = `SELECT price, seats, currency FROM Schedule WHERE gameId = $1`
	rows, err = types.Db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&price, &space, &currency)
	if err != nil {
		panic(err)
	}

	return price, space, currency
}

func howManyIsLeft(gameId int, wishfulseats int) (thereArePlaces bool) {
	var (
		rows    *sql.Rows
		err     error
		request string
		seats   int
	)
	request = `SELECT seats FROM Schedule WHERE gameId = $1`
	rows, err = types.Db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&seats)
	if err != nil {
		panic(err)
	}

	if wishfulseats < seats {
		thereArePlaces = true
	}
	return thereArePlaces
}

func completeRegistration(userId, gameId, seats int, payment string) (err error) {
	var (
		rows      *sql.Rows
		gameSeats int
	)
	_, err = types.Db.Exec("INSERT INTO GamesWithUsers (userId, gameId, seats, payment) VALUES ($1, $2, $3, $4)", userId, gameId, seats, payment)
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
	return err
}

func selectDetailOfGame(gameId int, language string) (details *forall.Game) {
	var (
		rows       *sql.Rows
		err        error
		request    string
		date, time int
	)
	details = new(forall.Game)
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
	return details
}
