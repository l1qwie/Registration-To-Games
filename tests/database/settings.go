package database

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/routine"
)

func withOutS(userId int) *bottypes.User {
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
	if user.Id != 899 {
		panic("user.Id != 899")
	}
	if user.ExMessageId != 1111 {
		panic("user.ExMessageId != 1111")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
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
	return user
}

func AfterChOptLang(userId int) {
	user := withOutS(userId)
	if user.Level != 1 {
		panic("user.Level != 1")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
	}
	if user.UserRec.Changeable != "language" {
		panic("user.UserRec.Changeable != `language`")
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
	if user.UserRec.GameId != 0 {
		panic("user.UserRec.GameId != 0")
	}
}

func AfterChooseOpt(userId int) {
	user := withOutS(userId)
	if user.Level != 1 {
		panic("user.Level != 1")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
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
	if user.UserRec.GameId != 0 {
		panic("user.UserRec.GameId != 0")
	}
}

func AfterChGame(userId int) {
	user := withOutS(userId)
	if user.Level != 2 {
		panic("user.Level != 2")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
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
	if user.UserRec.GameId != 0 {
		panic("user.UserRec.GameId != 0")
	}
}

func AfterChooseLanguage(userId int) {
	user := withOutS(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Language != "en" {
		panic("user.Language != `en`")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
	}
	if user.UserRec.Changeable != "language" {
		panic("user.UserRec.Changeable != `language`")
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
	if user.UserRec.GameId != 0 {
		panic("user.UserRec.GameId != 0")
	}
}

func AfterChOrDelGame(userId int) {
	user := withOutS(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
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
	if user.UserRec.GameId != 2 {
		panic("user.UserRec.GameId != 2")
	}
}

func AfterChOrDelGameDel(userId int) {
	user := withOutS(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
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
	if user.UserRec.GameId != 1 {
		panic("user.UserRec.GameId != 1")
	}
}

func AfterDelGame(userId int) {
	user := withOutS(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
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
	if user.UserRec.GameId != 0 {
		panic("user.UserRec.GameId != 0")
	}
	if !checkDelGame(1, user.Id) {
		panic("The game wasn't deleted")
	}
}

func AfterChangeWay(userId int) {
	user := withOutS(userId)
	if user.Level != 4 {
		panic("user.Level != 4")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
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
	if user.UserRec.GameId != 2 {
		panic("user.UserRec.GameId != 2")
	}
}

func AfterWhPayment(userId int) {
	user := withOutS(userId)
	if user.Level != 5 {
		panic("user.Level != 5")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
	}
	if user.UserRec.Changeable != "payment" {
		panic("user.UserRec.Changeable != `payment`")
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
	if user.UserRec.GameId != 1 {
		panic("user.UserRec.GameId != 1")
	}
}

func AfterPBCash(userId int) {
	user := withOutS(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
	}
	if user.UserRec.Changeable != "payment" {
		panic("user.UserRec.Changeable != `payment`")
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
	if user.UserRec.GameId != 0 {
		panic("user.UserRec.GameId != 0")
	}
	if !checkChangePaymethodCash(userId, 2) {
		panic("Paymethod wasn't changed to cash")
	}
}

func AfterWrSeats(userId int) {
	user := withOutS(userId)
	if user.Level != 5 {
		panic("user.Level != 5")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
	}
	if user.UserRec.Changeable != "myseats" {
		panic("user.UserRec.Changeable != `myseats`")
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
	if user.UserRec.GameId != 2 {
		panic("user.UserRec.GameId != 2")
	}
}

func AfterPBCard(userId int) {
	user := withOutS(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
	}
	if user.UserRec.Changeable != "payment" {
		panic("user.UserRec.Changeable != `payment`")
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
	if user.UserRec.GameId != 0 {
		panic("user.UserRec.GameId != 0")
	}
	if !checkChangePaymethodCard(userId, 1) {
		panic("Paymethod wasn't changed to card")
	}
}

func AfterChanSeats(userId int) {
	user := withOutS(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
	}
	if user.UserRec.Changeable != "myseats" {
		panic("user.UserRec.Changeable != `myseats`")
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
	if user.UserRec.GameId != 0 {
		panic("user.UserRec.GameId != 0")
	}
	if !ckeckChangedSeats(userId, 2, 6) {
		panic("Seats haven't changed yet")
	}
}
