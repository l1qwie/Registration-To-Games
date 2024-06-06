package handler

import "Game/apptype"

func isThereAnyGame(f func(error)) bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE status != -1").Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func saveToDatabase(res *apptype.Response, f func(error)) {
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (sport, date, time, seats, price, currency, link, address, status)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		res.Sport, res.Date, res.Time, res.Seats, res.Price, res.Currency, res.Link, res.Address, 1)
	if err != nil {
		f(err)
	}
}

func selectDateTime(f func(error)) ([]string, []string, []int) {
	var (
		length       int
		dates, times []string
		gameids      []int
	)
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE status != -1").Scan(&length)
	if err != nil {
		f(err)
	} else {
		dates = make([]string, length)
		times = make([]string, length)
		gameids = make([]int, length)
		i := 0
		rows, err := apptype.Db.Query("SELECT gameid, date, time FROM Schedule WHERE status != -1")
		if err != nil {
			f(err)
		}
		for rows.Next() && err == nil {
			err = rows.Scan(&dates[i], &times[i], &gameids[i])
			if err != nil {
				f(err)
			}
			i++
		}
	}
	return dates, times, gameids
}

func findGame(gameid int, f func(error)) bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE gameid = $1 AND status != -1", gameid).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}
