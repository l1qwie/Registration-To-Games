package database

import (
	"database/sql"
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/routine"
	"registrationtogames/fmtogram/types"

	_ "github.com/lib/pq"
)

func UpdateLanguage(language string, userId int) {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET language = $1 WHERE userId = $2", language, userId)
	}
	if err != nil {
		panic(err)
	}
	db.Close()

}

func UpdateAction(action string, userId int) {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET action = $1 WHERE userId = $2", action, userId)
	}
	if err != nil {
		panic(err)
	}
	db.Close()

}

func UpdateLevel(level, userId int) {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET level = $1 WHERE userId = $2", level, userId)
	}
	if err != nil {
		panic(err)
	}
	db.Close()

}

func InsertGameId(gameId, userId int) {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET gameId = $1 WHERE userId = $2", gameId, userId)
	}
	if err != nil {
		panic(err)
	}
	db.Close()
}

func InsertSeats(seats, userId int) {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET seats = $1 WHERE userId = $2", seats, userId)
	}
	if err != nil {
		panic(err)
	}
	db.Close()
}

func InsertPayment(payment string, userId int) {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("UPDATE Users SET payment = $1 WHERE userId = $2", payment, userId)
	}
	if err != nil {
		panic(err)
	}
	db.Close()
}

func CreateGame() (err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
		VALUES (2, 'volleyball', 20250212, 1200, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', 1)`)
	}
	if err != nil {
		panic(err)
	}
	db.Close()

	return err
}

func DeleteGame(gameId int) {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("DELETE FROM Schedule WHERE gameId = $1", gameId)
	}
	if err != nil {
		panic(err)
	}
	db.Close()
}

func AfterMainMenuCheckDb(userId int) {
	var (
		err  error
		user *bottypes.User
	)
	user = new(bottypes.User)
	user.Id = userId
	err = routine.DbRetrieveUser(user)
	if err != nil {
		panic(err)
	}

	if user.Id != 456 {
		panic("user.Id != 456")
	} else if user.Language != "ru" {
		panic("user.Language != `ru`")
		//Change it in the future!!
		//} else if gameId != 0 {
		//panic("gameId != 0")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
	}
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Reg.Seats != 0 {
		panic("user.Reg.Seats != 0")
	}
	if user.Reg.Payment != "" {
		panic("user.Reg.Payment != ``")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direcrion != "" {
		panic("user.Media.Direcrion != ``")
	}
	if user.Media.Limit != 7 {
		panic("user.Media.Limit != 7")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
	if user.UserRec.Changeable != "" {
		panic("user.UserRec.Changeable != ``")
	}
	if user.UserRec.ActGame != "" {
		panic("user.UserRec.ActGame != ``")
	}
	if user.UserRec.WillChangeable != "" {
		panic("user.UserRec.WillChangeable != ``")
	}
	if user.UserRec.NewPay != "" {
		panic("user.UserRec.NewPay != ``")
	}
}
