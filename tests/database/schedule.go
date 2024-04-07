package database

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/routine"
)

func withoutS(userId int) *bottypes.User {
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

	if user.Id != 488 {
		panic("user.Id != 488")
	}
	if user.ExMessageId != 9999 {
		panic("user.ExMessageId != 9999")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "see schedule" {
		panic("user.Act != `see schedule`")
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
	if user.Media.Direction != "" {
		panic("user.Media.Direction != ``")
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
	return user
}

func AfterShowTheSchedule(userId int) {
	user := withoutS(userId)
	if user.Level != 0 {
		panic("user.Level != 0")
	}
}
