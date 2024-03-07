package database

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/routine"
)

func AfterGreetingsToUserCheckDb(userId int) {
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
	user.ExMessageId, err = routine.SelectExMessageId(user.Id)
	if err != nil {
		panic(err)
	}
	if user.Id != 456 {
		panic("user.Id != 456")
	}
	if user.ExMessageId != 8888 {
		panic("user.ExMessageId != 8888")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "registration" {
		panic("user.Act != `registration`")
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
}

func AfterShowRulesCheckDb(userId int) {
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
	user.ExMessageId, err = routine.SelectExMessageId(user.Id)
	if err != nil {
		panic(err)
	}
	if user.ExMessageId != 8888 {
		panic("user.ExMessageId != 8888")
	}
	if user.Id != 456 {
		panic("user.Id != 456")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Reg.GameId != 0 {
		panic("user.Reg.GameId != 0")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "registration" {
		panic("user.Act != `registration`")
	}
	if user.Level != 2 {
		panic("user.Level != 2")
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
}
