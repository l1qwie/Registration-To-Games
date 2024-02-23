package routine

import (
	"database/sql"
	"log"
	"registrationtogames/bot/bottypes"
	"registrationtogames/fmtogram/types"
	"runtime"

	_ "github.com/lib/pq"
)

func Find(id int) (detected bool) {
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
	request = `SELECT COUNT(*) FROM Users WHERE userId = $1`
	rows, err = db.Query(request, id)
	if err != nil {
		panic(err)
	}
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

func CreateUser(userId int, language string) (err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", userId, "registration", language, 0)
	}
	db.Close()

	return err
}

func DbRetrieveUser(user *bottypes.User) (err error) {
	var (
		db      *sql.DB
		rows    *sql.Rows
		gameId  int
		request string
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		request = `
        SELECT userId, language, gameId, launchPoint, action, level,
        seats, payment, 
		timeInterval, direction, mlimit, mediagroupId, mediagoupCounter, 
		changeable, actGame, willChangeable, newPay
        FROM users
        WHERE userId = $1`
		rows, err = db.Query(request, user.Id)
	}

	if err == nil {
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.Language, &gameId, &user.LaunchPoint, &user.Act, &user.Level,
				&user.Reg.Seats, &user.Reg.Payment,
				&user.Media.Interval, &user.Media.Direcrion, &user.Media.Limit, &user.Media.Id, &user.Media.Counter,
				&user.UserRec.Changeable, &user.UserRec.ActGame, &user.UserRec.WillChangeable, &user.UserRec.NewPay)
			if err != nil {
				_, file, line, _ := runtime.Caller(0)
				log.Fatalf("Error at %s:%d: %v", file, line, err)
			}
		}
		if user.Act == "reg to games" {
			user.Reg.GameId = gameId
		}
		//if level, act, id and etc. {user.Reg.GameId = gameId || user.Media.DelGameId = gameId || user.UserRec.GameId = gameId }
	}
	db.Close()

	return err
}

func DbRetainUser(user *bottypes.User) (err error) {
	var (
		db      *sql.DB
		request string
		gameId  int
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	request = `UPDATE Users SET userId = $1, language = $2, gameId = $3, launchPoint = $4, action = $5, level = $6,
	seats = $7, payment = $8, 
	timeInterval = $9, direction = $10, mlimit = $11, mediagroupId = $12, mediagoupCounter = $13, 
	changeable = $14, actGame = $15, willChangeable = $16, newPay = $17 WHERE userId = $1`
	if user.Act == "reg to games" {
		gameId = user.Reg.GameId
	}
	//if level, act, id and etc. {gameId = user.Reg.GameId || gameId = user.Media.DelGameId || gameId = user.UserRec.GameId }
	if err == nil {
		_, err = db.Exec(request, user.Id, user.Language, gameId, user.LaunchPoint, user.Act, user.Level,
			user.Reg.Seats, user.Reg.Payment,
			user.Media.Interval, user.Media.Direcrion, user.Media.Limit, user.Media.Id, user.Media.Counter,
			user.UserRec.Changeable, user.UserRec.ActGame, user.UserRec.WillChangeable, user.UserRec.NewPay)
	}
	db.Close()
	return err
}
