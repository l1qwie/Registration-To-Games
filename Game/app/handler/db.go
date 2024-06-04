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
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (gameid, sport, date, time, seats, price, currency, link, address)
		VALUES (nextval(gameswithusers_id_seq), $1, $2, $3, $4, $5, $6, $7, $8, $9 )`,
		res.Sport, res.Date, res.Time, res.Sport, res.Price, res.Currency, res.Level, res.Address)
	if err != nil {
		f(err)
	}
}
