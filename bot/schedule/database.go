package schedule

import (
	"database/sql"
	"registrationtogames/bot/dictionary"
	"registrationtogames/bot/forall"
	"registrationtogames/fmtogram/types"

	_ "github.com/lib/pq"
)

func selectSchedule(language string) (schedule []*forall.Game) {
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
	request = `SELECT gameId, sport, date, time, seats, price, currency FROM Schedule WHERE status != -1`
	rows, err = db.Query(request)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		schedule = append(schedule, &forall.Game{})
		err = rows.Scan(&schedule[i].Id, &sport, &date, &time, &schedule[i].Seats, &schedule[i].Price, &schedule[i].Currency)
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
