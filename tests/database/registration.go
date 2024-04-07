package database

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/routine"

	_ "github.com/lib/pq"
)

func withoutR(userId int) *bottypes.User {
	var (
		err  error
		user *bottypes.User
	)
	user = new(bottypes.User)
	user.Id = userId
	err = routine.DbRetrieveUser(user, fm)
	if err != nil {
		panic(err)
	}
	user.ExMessageId, err = routine.SelectExMessageId(user.Id, fm)
	if err != nil {
		panic(err)
	}
	if user.ExMessageId != 8883 {
		panic("ser.ExMessageId != 8883")
	}
	if user.Id != 477 {
		panic("user.Id != 477")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
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
	if user.Media.Limit != 7 {
		panic("user.Media.Limit != 7")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "" {
		panic("user.Media.Direction != ``")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.DelGameId != 0 {
		panic("user.Media.DelGameId != 0")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
	return user
}

func AfterPresentationScheduele(userId int) {
	user := withoutR(userId)
	if user.Level != 1 {
		panic("user.Level != 1")
	}
	if user.Act != "reg to games" {
		panic("user.Act != `reg to games`")
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
}

func AfterChooseGameR(userId int) {
	user := withoutR(userId)
	if user.Level != 2 {
		panic("user.Level != 2")
	}
	if user.Act != "reg to games" {
		panic("user.Act != `reg to games`")
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

}

func AfterAfterChooseSeats(userId int) {
	user := withoutR(userId)
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
}

func AfterChoosePayment(userId int) {
	user := withoutR(userId)
	if user.Level != 4 {
		panic("user.Level != 4")
	}
	if user.Act != "reg to games" {
		panic("user.Act != `reg to games`")
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
}

func AfterBestWishes(userId int) {
	user := withoutR(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
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
	if !checkGamesWithUsersTable(user.Id, 2, user.Reg.Seats, user.Reg.Payment) {
		panic("user doesn't exist in table GamesWithUsers")
	}
}
