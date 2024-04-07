package schedule

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"database/sql"

	_ "github.com/lib/pq"
)

func FindGame(fm *formatter.Formatter) bool {
	rows, err := types.Db.Query("SELECT gameId FROM Schedule WHERE status != -1")
	if err != nil {
		fm.Error(err)
	}
	cn := 0
	for rows.Next() {
		err = rows.Scan(&cn)
		if err != nil {
			fm.Error(err)
		}
	}
	defer rows.Close()
	return cn > 0
}

func selectSchedule(language string, fm *formatter.Formatter) (schedule []*bottypes.Game) {
	var (
		rows           *sql.Rows
		err            error
		request, sport string
		i, date, time  int
	)
	request = `SELECT gameId, sport, date, time, seats, price, currency FROM Schedule WHERE status != -1`
	rows, err = types.Db.Query(request)
	if err != nil {
		fm.Error(err)
	}
	for rows.Next() {
		schedule = append(schedule, &bottypes.Game{})
		err = rows.Scan(&schedule[i].Id, &sport, &date, &time, &schedule[i].Seats, &schedule[i].Price, &schedule[i].Currency)
		if err != nil {
			fm.Error(err)
		}
		schedule[i].Sport = dictionary.Dictionary[language][sport]
		schedule[i].Date = forall.FromIntToStrDate(date)
		schedule[i].Time = forall.FromIntToStrTime(time)
		i++
	}
	defer rows.Close()
	return schedule
}
