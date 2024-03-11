package schedule

import (
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/types"
	"database/sql"

	_ "github.com/lib/pq"
)

func FindGame() bool {
	var (
		exist bool
		rows  *sql.Rows
		cn    int
		err   error
		req   string
	)
	req = "SELECT gameId FROM Schedule WHERE status != -1"
	rows, err = types.Db.Query(req)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&cn)
		if err != nil {
			panic(err)
		}
	}
	if cn > 0 {
		exist = true
	}
	return exist
}

func selectSchedule(language string) (schedule []*forall.Game) {
	var (
		rows           *sql.Rows
		err            error
		request, sport string
		i, date, time  int
	)
	request = `SELECT gameId, sport, date, time, seats, price, currency FROM Schedule WHERE status != -1`
	rows, err = types.Db.Query(request)
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
