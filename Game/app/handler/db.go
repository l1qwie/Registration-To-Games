package handler

import (
	"Game/apptype"
	"database/sql"
	"fmt"
)

type Conn struct {
	Db *sql.DB
}

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
		length, d, t int
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
			err = rows.Scan(&gameids[i], &d, &t)
			if err != nil {
				f(err)
			}
			dates[i] = fromIntToStrDate(d)
			times[i] = fromIntToStrTime(t)
			i++
		}
		err = rows.Close()
		if err != nil {
			f(err)
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

func updateDBStr(column, input string, gameid int, f func(error)) {
	_, err := apptype.Db.Exec(fmt.Sprintf("UPDATE Schedule SET %s = $1 WHERE gameid = $2", column), input, gameid)
	if err != nil {
		f(err)
	}
}

func updateDBInt(column string, input, gameid int, f func(error)) {
	_, err := apptype.Db.Exec(fmt.Sprintf("UPDATE Schedule SET %s = $1 WHERE gameid = $2", column), input, gameid)
	if err != nil {
		f(err)
	}
}

func selecAllGameInf(gameid int, f func(error)) (date, time, seats, price int, sport, currency, link, address string) {
	err := apptype.Db.QueryRow("SELECT sport, date, time, seats, price, currency, link, address FROM Schedule WHERE gameid = $1",
		gameid).Scan(&sport, &date, &time, &seats, &price, &currency, &link, &address)
	if err != nil {
		f(err)
	}
	return date, time, seats, price, sport, currency, link, address
}

func selectDate(gameid int, f func(error)) int {
	var date int
	err := apptype.Db.QueryRow("SELECT date FROM Schedule WHERE gameid = $1 and status != -1", gameid).Scan(&date)
	if err != nil {
		f(err)
	}
	return date
}

func deleteGameDB(gameid int, f func(error)) {
	_, err := apptype.Db.Exec("UPDATE Schedule SET status = -1 WHERE gameid = $1", gameid)
	if err != nil {
		f(err)
	}
}

func (c *Conn) SaveTheKey(key []byte, ip string) error {
	_, err := c.Db.Exec("INSERT INTO SymKeys (key, ip) VALUES ($1, $2)", key)
	return err
}

func (c *Conn) GetTheKey(ip string) ([]byte, error) {
	var key []byte
	err := c.Db.QueryRow("SELECT key FROM SymKeys WHERE ip = $1", ip).Scan(&key)
	return key, err
}
