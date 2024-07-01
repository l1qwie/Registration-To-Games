package dbsync

import (
	"Settings/apptype"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

var tDB *sql.DB

type Game struct {
	gameid   int
	sport    string
	date     int
	time     int
	seats    int
	price    int
	currency string
	status   int
}

func endOfTransaction(err error, db *sql.DB) {
	var res string
	if err != nil {
		res = "ROLLBACK"
	} else {
		res = "COMMIT"
	}
	_, err = db.Exec(res)
	if err != nil {
		panic(err)
	}
}

func clean() {
	_, err := apptype.TestDb.Exec("TRUNCATE TABLE Schedule")
	if err != nil {
		panic(err)
	}
}

func createTestGame() {
	_, err := apptype.TestDb.Exec(`
			INSERT INTO Schedule 
			(gameid, sport, date, time, seats, price, currency, status)
			VALUES 
			(88, 'volleyball', 20241212, 1300, 78, 200, 'USDT', 0)`)
	if err != nil {
		panic(err)
	}
}

func checkSelInTrans(g *Game) {
	if g.gameid != 99 {
		panic(fmt.Sprintf("gameid != 99. gameid == %d", g.gameid))
	}
	if g.sport != "football" {
		panic(fmt.Sprintf("sport != football. sport == %s", g.sport))
	}
	if g.date != 20241212 {
		panic(fmt.Sprintf("date != 20241212. date == %d", g.date))
	}
	if g.time != 1200 {
		panic(fmt.Sprintf("time != 1200. time == %d", g.time))
	}
	if g.seats != 88 {
		panic(fmt.Sprintf("seats != 88. seats == %d", g.seats))
	}
	if g.price != 199 {
		panic(fmt.Sprintf("price != 199. price == %d", g.price))
	}
	if g.currency != "RUB" {
		panic(fmt.Sprintf("currency != RUB. currency == %s", g.currency))
	}
	if g.status != 0 {
		panic(fmt.Sprintf("status != 0. status == %d", g.status))
	}
}

func checkTestGame(g *Game) {
	if g.gameid != 88 {
		panic(fmt.Sprintf("gameid != 88. gameid == %d", g.gameid))
	}
	if g.sport != "volleyball" {
		panic(fmt.Sprintf("sport != volleyball. sport == %s", g.sport))
	}
	if g.date != 20241212 {
		panic(fmt.Sprintf("date != 20241212. date == %d", g.date))
	}
	if g.time != 1300 {
		panic(fmt.Sprintf("time != 1300. time == %d", g.time))
	}
	if g.seats != 78 {
		panic(fmt.Sprintf("seats != 78. seats == %d", g.seats))
	}
	if g.price != 200 {
		panic(fmt.Sprintf("price != 200. price == %d", g.price))
	}
	if g.currency != "USDT" {
		panic(fmt.Sprintf("currency != USDT. currency == %s", g.currency))
	}
	if g.status != 0 {
		panic(fmt.Sprintf("status != 0. status == %d", g.status))
	}
}

func checkTestGameUpd(g *Game) {
	if g.gameid != 88 {
		panic(fmt.Sprintf("gameid != 88. gameid == %d", g.gameid))
	}
	if g.sport != "football" {
		panic(fmt.Sprintf("sport != football. sport == %s", g.sport))
	}
	if g.date != 20241212 {
		panic(fmt.Sprintf("date != 20241212. date == %d", g.date))
	}
	if g.time != 1300 {
		panic(fmt.Sprintf("time != 1300. time == %d", g.time))
	}
	if g.seats != 78 {
		panic(fmt.Sprintf("seats != 78. seats == %d", g.seats))
	}
	if g.price != 200 {
		panic(fmt.Sprintf("price != 200. price == %d", g.price))
	}
	if g.currency != "USDT" {
		panic(fmt.Sprintf("currency != USDT. currency == %s", g.currency))
	}
	if g.status != 0 {
		panic(fmt.Sprintf("status != 0. status == %d", g.status))
	}
}

func checkTestGameIns(g *Game) {
	if g.gameid != 199 {
		panic(fmt.Sprintf("gameid != 199. gameid == %d", g.gameid))
	}
	if g.sport != "football" {
		panic(fmt.Sprintf("sport != football. sport == %s", g.sport))
	}
	if g.date != 20241212 {
		panic(fmt.Sprintf("date != 20241212. date == %d", g.date))
	}
	if g.time != 2000 {
		panic(fmt.Sprintf("time != 2000. time == %d", g.time))
	}
	if g.seats != 88 {
		panic(fmt.Sprintf("seats != 88. seats == %d", g.seats))
	}
	if g.price != 199 {
		panic(fmt.Sprintf("price != 199. price == %d", g.price))
	}
	if g.currency != "RUB" {
		panic(fmt.Sprintf("currency != RUB. currency == %s", g.currency))
	}
	if g.status != 0 {
		panic(fmt.Sprintf("status != 0. status == %d", g.status))
	}
}

func startTransaction() {
	var (
		rows  *sql.Rows
		count int
	)
	_, err := apptype.TestDb.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	defer endOfTransaction(err, apptype.TestDb)
	if err == nil {
		err = apptype.TestDb.QueryRow("SELECT COUNT(*) FROM Schedule").Scan(&count)
	}
	if err == nil {
		log.Printf("Rows in Schedule from transaction 1 = %d", count)
		_, err = apptype.TestDb.Exec(`
			INSERT INTO Schedule 
			(gameid, sport, date, time, seats, price, currency, status)
			VALUES 
			(99, 'football', 20241212, 1200, 88, 199, 'RUB', 0)`)
	}
	if err == nil {
		time.Sleep(time.Second)
		rows, err = apptype.TestDb.Query("SELECT * FROM Schedule WHERE status != -1")
		if err == nil {
			defer rows.Close()
			i := 0
			for rows.Next() && err == nil {
				g := new(Game)
				err = rows.Scan(&g.gameid, &g.sport, &g.date, &g.time, &g.seats, &g.price, &g.currency, &g.status)
				if err == nil {
					if i == 0 {
						checkTestGame(g)
					} else {
						checkSelInTrans(g)
					}
				}
				i++
			}
		}
	}
}

func startSelect() {
	var count int
	g := new(Game)
	_, err := tDB.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	if err == nil {
		err = tDB.QueryRow("SELECT COUNT(*) FROM Schedule").Scan(&count)
	}
	if err == nil {
		log.Printf("Rows in Schedule from transaction 2 = %d", count)
		err = tDB.QueryRow("SELECT * FROM Schedule WHERE status != -1").Scan(&g.gameid, &g.sport, &g.date, &g.time, &g.seats, &g.price, &g.currency, &g.status)
	}
	defer endOfTransaction(err, tDB)
	checkTestGame(g)
}

func checkSel() {
	var wg sync.WaitGroup
	clean()
	createTestGame()
	//defer clean()
	wg.Add(2)
	go func() {
		defer wg.Done()
		startTransaction()
	}()
	go func() {
		defer wg.Done()
		startSelect()
	}()
	wg.Wait()
}

func startUpdate() {
	var count int
	g := new(Game)
	_, err := tDB.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	if err == nil {
		err = tDB.QueryRow("SELECT COUNT(*) FROM Schedule").Scan(&count)
	}
	if err == nil {
		log.Printf("Rows in Schedule from transaction 2 = %d", count)
		_, err = tDB.Exec("UPDATE Schedule SET sport = 'football' WHERE gameid = 88")
	}
	if err == nil {
		err = tDB.QueryRow("SELECT * FROM Schedule WHERE status != -1").Scan(&g.gameid, &g.sport, &g.date, &g.time, &g.seats, &g.price, &g.currency, &g.status)
	}
	defer endOfTransaction(err, tDB)
	checkTestGameUpd(g)
}

func checkUpd() {
	var wg sync.WaitGroup
	clean()
	createTestGame()
	defer clean()
	wg.Add(2)
	go func() {
		defer wg.Done()
		startTransaction()
	}()
	go func() {
		defer wg.Done()
		startUpdate()
	}()
	wg.Wait()
}

func startInsert() {
	var (
		count, i int
		rows     *sql.Rows
	)
	_, err := tDB.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	defer endOfTransaction(err, tDB)
	if err == nil {
		err = tDB.QueryRow("SELECT COUNT(*) FROM Schedule").Scan(&count)
	}
	if err == nil {
		log.Printf("Rows in Schedule from transaction 2 = %d", count)
		_, err = tDB.Exec(`
			INSERT INTO Schedule 
			(gameid, sport, date, time, seats, price, currency, status)
			VALUES 
			(199, 'football', 20241212, 2000, 88, 199, 'RUB', 0)`)
	}
	if err == nil {
		rows, err = tDB.Query("SELECT * FROM Schedule WHERE status != -1")
		if err == nil {
			defer rows.Close()
			for rows.Next() && err == nil {
				g := new(Game)
				err = rows.Scan(&g.gameid, &g.sport, &g.date, &g.time, &g.seats, &g.price, &g.currency, &g.status)
				if err == nil {
					if i == 0 {
						checkTestGame(g)
					} else {
						checkTestGameIns(g)
					}
					i++
				}
			}
		}
	}
}

func checkIns() {
	var wg sync.WaitGroup
	clean()
	createTestGame()
	defer clean()
	wg.Add(2)
	go func() {
		defer wg.Done()
		startTransaction()
	}()
	go func() {
		defer wg.Done()
		startInsert()
	}()
	wg.Wait()
}

// Tests for concurrent use of the database
func Start() {
	apptype.TestDb = apptype.ConnectToDatabase()
	tDB = apptype.ConnectToDatabase()
	checkSel()
	checkUpd()
	checkIns()
}
