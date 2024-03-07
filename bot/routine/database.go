package routine

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/fmtogram/types"
	"database/sql"

	_ "github.com/lib/pq"
)

func Find(id int) (detected bool) {
	var (
		rows    *sql.Rows
		err     error
		request string
		counter int
	)
	request = `SELECT COUNT(*) FROM Users WHERE userId = $1`
	rows, err = types.Db.Query(request, id)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&counter)
	if err != nil {
		panic(err)
	}
	if counter > 0 {
		detected = true
	}
	return detected
}

func CreateUser(userId int, language string) (err error) {
	_, err = types.Db.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", userId, "registration", language, 0)
	return err
}

func DbRetrieveUser(user *bottypes.User) (err error) {
	var (
		rows    *sql.Rows
		gameId  int
		request string
	)
	request = `SELECT userId, language, gameId, launchPoint, action, level,
				seats, payment, 
				timeInterval, direction, mlimit, mediagroupId, mediagoupCounter, 
				changeable, actGame, willChangeable, newPay
				FROM users
				WHERE userId = $1`
	rows, err = types.Db.Query(request, user.Id)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.Language, &gameId, &user.LaunchPoint, &user.Act, &user.Level,
				&user.Reg.Seats, &user.Reg.Payment,
				&user.Media.Interval, &user.Media.Direction, &user.Media.Limit, &user.Media.Id, &user.Media.Counter,
				&user.UserRec.Changeable, &user.UserRec.ActGame, &user.UserRec.WillChangeable, &user.UserRec.NewPay)
			if err != nil {
				panic(err)
			}
		}
		if user.Act == "reg to games" {
			user.Reg.GameId = gameId
		}
	}

	return err
}

func DbRetainUser(user *bottypes.User) (err error) {
	var (
		request string
		gameId  int
	)
	request = `UPDATE Users SET userId = $1, language = $2, gameId = $3, launchPoint = $4, action = $5, level = $6,
	seats = $7, payment = $8, 
	timeInterval = $9, direction = $10, mlimit = $11, mediagroupId = $12, mediagoupCounter = $13, 
	changeable = $14, actGame = $15, willChangeable = $16, newPay = $17 WHERE userId = $1`
	if user.Act == "reg to games" {
		gameId = user.Reg.GameId
	}
	if err == nil {
		_, err = types.Db.Exec(request, user.Id, user.Language, gameId, user.LaunchPoint, user.Act, user.Level,
			user.Reg.Seats, user.Reg.Payment,
			user.Media.Interval, user.Media.Direction, user.Media.Limit, user.Media.Id, user.Media.Counter,
			user.UserRec.Changeable, user.UserRec.ActGame, user.UserRec.WillChangeable, user.UserRec.NewPay)
	}
	return err
}

func SelectExMessageId(userId int) (exMessageId int, err error) {
	var (
		rows *sql.Rows
	)
	if err == nil {
		rows, err = types.Db.Query("SELECT ExMessageId FROM Users WHERE userId = $1", userId)
	}
	if err == nil {
		rows.Next()
		err = rows.Scan(&exMessageId)
		if err != nil {
			panic(err)
		}

	}
	return exMessageId, err
}

func updateExMessageId(exMessageId, userId int) (err error) {
	_, err = types.Db.Exec("UPDATE Users SET ExMessageId = $1 WHERE userId = $2", exMessageId, userId)
	return err
}
