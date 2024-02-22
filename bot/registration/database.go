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
		err = rows.Scan(&schedule[i].id, &sport, &schedule[i].date, &schedule[i].time, &schedule[i].seats)
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

func selectDetailOfGame(gameId int) (details *Game) {

	return details
}
