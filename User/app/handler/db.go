package handler

import (
	"User/apptype"
	"User/fmtogram/types"
	"log"

	_ "github.com/lib/pq"
)

func find(id int, f func(error)) bool {
	var count int
	err := types.Db.QueryRow("SELECT COUNT(*) FROM Users WHERE userId = $1", id).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func dbRetrieveUser(user *apptype.User, f func(error)) {
	var gameId int
	request := `SELECT userId, language, gameId, launchPoint, action, level,
				seats, payment, 
				timeInterval, direction, mlimit, mediagroupId, mediagoupCounter, 
				changeable, actGame, willChangeable, newPay
				FROM users
				WHERE userId = $1`
	err := types.Db.QueryRow(request, user.Id).Scan(&user.Id, &user.Language, &gameId, &user.LaunchPoint, &user.Act, &user.Level,
		&user.Reg.Seats, &user.Reg.Payment,
		&user.Media.Interval, &user.Media.Direction, &user.Media.Limit, &user.Media.Id, &user.Media.Counter,
		&user.UserRec.Changeable, &user.UserRec.ActGame, &user.UserRec.WillChangeable, &user.UserRec.NewPay)
	if err != nil {
		f(err)
	}
	log.Print("USER ACT: ", user.Act)
	if user.Act == "reg to games" {
		user.Reg.GameId = gameId
	} else if user.Act == "photos and videos" {
		user.Media.DelGameId = gameId
	} else if user.Act == "settings" {
		user.UserRec.GameId = gameId
	}
}

func dbRetainUser(user *apptype.User, f func(error)) {
	var gameId int
	request := `UPDATE Users SET userId = $1, language = $2, gameId = $3, launchPoint = $4, action = $5, level = $6,
	seats = $7, payment = $8, 
	timeInterval = $9, direction = $10, mlimit = $11, mediagroupId = $12, mediagoupCounter = $13, 
	changeable = $14, actGame = $15, willChangeable = $16, newPay = $17 WHERE userId = $1`
	if user.Act == "reg to games" {
		gameId = user.Reg.GameId
	} else if user.Act == "photos and videos" {
		gameId = user.Media.DelGameId
	} else if user.Act == "settings" {
		gameId = user.UserRec.GameId
	}
	if user.Act != "photos and videos" {
		user.Media.Counter = 0
	}
	_, err := types.Db.Exec(request, user.Id, user.Language, gameId, user.LaunchPoint, user.Act, user.Level,
		user.Reg.Seats, user.Reg.Payment,
		user.Media.Interval, user.Media.Direction, user.Media.Limit, user.Media.Id, user.Media.Counter,
		user.UserRec.Changeable, user.UserRec.ActGame, user.UserRec.WillChangeable, user.UserRec.NewPay)
	if err != nil {
		f(err)
	}
}

func createUser(userId int, language string, f func(error)) {
	_, err := types.Db.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", userId, "registration", language, 0) //userId
	if err != nil {
		f(err)
	}
}

func SelectExMessageId(userId int, f func(error)) int {
	var exmid int
	err := types.Db.QueryRow("SELECT ExMessageId FROM Users WHERE userId = $1", userId).Scan(&exmid)
	if err != nil {
		f(err)
	}
	return exmid
}

func updateExMessageId(exmid, userId int, f func(error)) {
	_, err := types.Db.Exec("UPDATE Users SET ExMessageId = $1 WHERE userId = $2", exmid, userId)
	if err != nil {
		f(err)
	}
}

func UpdateTheUser(u *apptype.User, custlang bool) error {
	_, err := types.Db.Exec("UPDATE Users SET language = $1, customlanguage = $2 WHERE userId = $3", u.Language, custlang, u.Id)
	return err
}
