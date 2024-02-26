package registration

import (
	"database/sql"
	"registrationtogames/bot/dictionary"
	"registrationtogames/bot/forall"
	"registrationtogames/fmtogram/types"

	_ "github.com/lib/pq"
)

func selectTheSchedule(limit, offset int, language string) (schedule []*Game) {
	var (
		db             *sql.DB
		rows           *sql.Rows
		err            error
		request, sport string
		i, date, time  int
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err != nil {
		panic(err)
	}
	request = `SELECT gameId, sport, date, time, seats FROM Schedule WHERE (status != -1) ORDER BY Schedule DESC LIMIT $1 OFFSET $2`
	rows, err = db.Query(request, limit, offset)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		schedule = append(schedule, &Game{})
		err = rows.Scan(&schedule[i].id, &sport, &date, &time, &schedule[i].seats)
		if err != nil {
			panic(err)
		}
		schedule[i].sport = dictionary.Dictionary[language][sport]
		schedule[i].date = forall.FromIntToStrDate(date)
		schedule[i].time = forall.FromIntToStrTime(time)
		i++
	}
	return schedule
}

func findAGame(gameId int) (detected bool) {
	var (
		db      *sql.DB
		rows    *sql.Rows
		err     error
		request string
		counter int
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err != nil {
		panic(err)
	}
	request = `SELECT COUNT(*) FROM Schedule WHERE status != -1 AND gameId = $1`
	rows, err = db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	//Try to delete for below
	for rows.Next() {
		err = rows.Scan(&counter)
		if err != nil {
			panic(err)
		}
	}
	if counter > 0 {
		detected = true
	}
	return detected
}

func selectThePrice(gameId int) (price, space int, currency string) {
	var (
		db      *sql.DB
		rows    *sql.Rows
		err     error
		request string
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err != nil {
		panic(err)
	}
	request = `SELECT price, seats, currency FROM Schedule WHERE gameId = $1`
	rows, err = db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	//Try to delete for below
	for rows.Next() {
		err = rows.Scan(&price, &space, &currency)
		if err != nil {
			panic(err)
		}
	}
	return price, space, currency
}

func howManyIsLeft(gameId int, wishfulseats int) (thereArePlaces bool) {
	var (
		db      *sql.DB
		rows    *sql.Rows
		err     error
		request string
		seats   int
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err != nil {
		panic(err)
	}
	request = `SELECT seats FROM Schedule WHERE gameId = $1`
	rows, err = db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	//Try to delete for below
	for rows.Next() {
		err = rows.Scan(&seats)
		if err != nil {
			panic(err)
		}
	}
	if wishfulseats < seats {
		thereArePlaces = true
	}
	return thereArePlaces
}

func completeRegistration(userId, gameId, seats int, payment string) (err error) {
	var (
		db        *sql.DB
		rows      *sql.Rows
		gameSeats int
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("INSERT INTO GamesWithUsers (userId, gameId, seats, payment) VALUES ($1, $2, $3, $4)", userId, gameId, seats, payment)
	}
	if err == nil {
		rows, err = db.Query("SELECT seats FROM Schedule WHERE gameId = $1", gameId)
	}
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&gameSeats)
			if err != nil {
				panic(err)
			}
		}
	}
	if err == nil {
		_, err = db.Exec("UPDATE Schedule SET seats = $1 WHERE gameId = $2", gameSeats-seats, gameId)
	}

	if err != nil {
		panic(err)
	}
	db.Close()

	return err
}

func selectDetailOfGame(gameId int, language string) (details *Game) {
	var (
		db         *sql.DB
		rows       *sql.Rows
		err        error
		request    string
		date, time int
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err != nil {
		panic(err)
	}
	details = new(Game)
	request = `SELECT gameId, sport, date, time, seats, latitude, longitude, address, price, currency FROM Schedule WHERE gameId = $1`
	rows, err = db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	//Try to delete for below
	for rows.Next() {
		err = rows.Scan(&details.id, &details.sport, &date, &time, &details.seats, &details.lattitude, &details.longitude, &details.address, &details.price, &details.currency)
		if err != nil {
			panic(err)
		}
		details.sport = dictionary.Dictionary[language][details.sport]
		details.date = forall.FromIntToStrDate(date)
		details.time = forall.FromIntToStrTime(time)
	}
	return details
}
