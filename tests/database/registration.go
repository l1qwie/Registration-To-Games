package database

import (
	"database/sql"
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/routine"
	"registrationtogames/fmtogram/types"

	_ "github.com/lib/pq"
)

func AfterPresentationSchedueleCheckDb(userId int) {
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
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "reg to games" {
		panic("user.Act != `reg to games`")
	}
	if user.Level != 1 {
		panic("user.Level != 1")
	}
	if user.Reg.GameId != 0 {
		panic("user.Reg.GameId != 0")
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

func AfterChooseGameCheckDb(userId int) {
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
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "reg to games" {
		panic("user.Act != `reg to games`")
	}
	if user.Level != 2 {
		panic("user.Level != 2")
	}
	if user.Reg.GameId != 2 {
		panic("user.Reg.GameId != 2")
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

func AfterChooseSeatsCheckDb(userId int) {
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
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "reg to games" {
		panic("user.Act != `reg to games`")
	}
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Reg.GameId != 2 {
		panic("user.Reg.GameId != 2")
	}
	if user.Reg.Seats != 2 {
		panic("user.Reg.Seats != 2")
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

func AfterChoosePaymentCheckDb(userId int) {
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
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "reg to games" {
		panic("user.Act != `reg to games`")
	}
	if user.Level != 4 {
		panic("user.Level != 4")
	}
	if user.Reg.GameId != 2 {
		panic("user.Reg.GameId != 2")
	}
	if user.Reg.Seats != 2 {
		panic("user.Reg.Seats != 2")
	}
	if user.Reg.Payment != "card" {
		panic("user.Reg.Payment != `card`")
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

func checkGamesWithUsersTable(userId, gameId, seats int, payment string) (detected bool) {
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
	request = `SELECT COUNT(*) FROM GamesWithUsers WHERE userId = $1 AND gameId = $2 AND seats = $3 AND payment = $4` //AND statuspayment = 0 AND status = 0
	rows, err = db.Query(request, userId, gameId, seats, payment)
	if err != nil {
		panic(err)
	}
	//Try to delete for below
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

func deleteUserFromGamesWithUsers(userId int) {
	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		_, err = db.Exec("DELETE FROM GamesWithUsers WHERE userId = $1", userId)
	}
	if err != nil {
		panic(err)
	}
	db.Close()
}

func AfterBestWishes(userId int) {
	var (
		err  error
		user *bottypes.User
	)
	//defer deleteUserFromGamesWithUsers(userId)
	user = new(bottypes.User)
	user.Id = userId
	err = routine.DbRetrieveUser(user)
	if err != nil {
		panic(err)
	}
	if !checkGamesWithUsersTable(user.Id, 2, user.Reg.Seats, user.Reg.Payment) {
		panic("user doesn't exist in table GamesWithUsers")
	}
	if user.Id != 456 {
		panic("user.Id != 456")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
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
	if user.Reg.GameId != 0 {
		panic("user.Reg.GameId != 0")
	}
	if user.Reg.Seats != 2 {
		panic("user.Reg.Seats != 2")
	}
	if user.Reg.Payment != "card" {
		panic("user.Reg.Payment != `card`")
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
