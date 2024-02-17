package routine

import (
	"database/sql"
	"log"
	"registrationtogames/fmtogram/types"

	_ "github.com/lib/pq"
)

func (user *User) find() (detected bool) {
	var (
		db      *sql.DB
		rows    *sql.Rows
		err     error
		request string
		counter int
	)

	db, err = sql.Open("postgres", types.ConnectTo())
	if err != nil {
		log.Fatal(err)
	}
	request = `SELECT COUNT(*)FROM Users WHERE user_id = $1`
	rows, err = db.Query(request, user.Id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err = rows.Scan(&counter)
		if err != nil {
			log.Fatal(err)
		}
	}
	if counter > 0 {
		detected = true
	}
	return detected
}

func (user *User) createUser() (err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("INSERT INTO Users (user_id, action, language, level) VALUES ($1, $2, $3, $4)", user.Id, "registration", user.Language, 0)
	}
	db.Close()

	return err
}

func (user *User) dbRetrieveUser() (err error) {
	var (
		db      *sql.DB
		rows    *sql.Rows
		gameId  int
		request string
	)
	//Don't forget to change your database for this before you start testing
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		request = `
        SELECT userId, language, gameId, launchPoint, action, level
        sport, seats, payment, 
		timeInterval, direction, limit, mediagroupId, mediagoupCounter, 
		changeable, actGame, willChangeable, newPay,
        FROM Users
        WHERE user_id = $1`
		rows, err = db.Query(request, user.Id)
	}

	if err == nil {
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.Language, &gameId, &user.LaunchPoint, &user.Act, &user.Level,
				&user.Reg.Sport, &user.Reg.Seats, &user.Reg.Payment,
				&user.Media.Interval, &user.Media.Direcrion, &user.Media.Limit, &user.Media.Id, &user.Media.Counter,
				&user.UserRec.Changeable, &user.UserRec.ActGame, &user.UserRec.WillChangeable, &user.UserRec.NewPay)
			if err != nil {
				log.Fatal(err)
			}
		}
		//if level, act, id and etc. {user.Reg.GameId = gameId || user.Media.DelGameId = gameId || user.UserRec.GameId = gameId }
	}
	db.Close()

	return err
}

func (user *User) dbRetainUser() (err error) {
	var (
		db      *sql.DB
		request string
		gameId  int
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	request = `UPDATE Users SET userId = $1, language = $2, gameId = $3, launchPoint = $4, action = $5, level = $6
	sport = $7, seats = $8, payment = $9, 
	timeInterval = $10, direction = $11, limit = $12, mediagroupId = $13, mediagoupCounter = $14, 
	changeable = $15, actGame = $16, willChangeable = $17, newPay = $18 WHERE user_id = $19`
	//if level, act, id and etc. {gameId = user.Reg.GameId || gameId = user.Media.DelGameId || gameId = user.UserRec.GameId }
	if err == nil {
		_, err = db.Exec(request, user.Id, user.Language, gameId, user.LaunchPoint, user.Act, user.Level,
			user.Reg.Sport, user.Reg.Seats, user.Reg.Payment,
			user.Media.Interval, user.Media.Direcrion, user.Media.Limit, user.Media.Id, user.Media.Counter,
			user.UserRec.Changeable, user.UserRec.ActGame, user.UserRec.WillChangeable, user.UserRec.NewPay)
	}
	db.Close()
	return err
}
