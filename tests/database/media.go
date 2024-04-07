package database

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/media"
	"RegistrationToGames/bot/routine"
	"RegistrationToGames/fmtogram/formatter"
)

var fm = new(formatter.Formatter)

// check defalt columns without Media type
// only user.Media.Limit could be here cause it is const
func withoutM(userId int) *bottypes.User {
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
	if user.ExMessageId != 66666 {
		panic("ser.ExMessageId != 66666")
	}
	if user.Id != 499 {
		panic("user.Id != 499")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "photos and videos" {
		panic("user.Act != `photos and video`")
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
	return user
}

func AfterChooseDirection(userId int) {
	user := withoutM(userId)
	if user.Level != 1 {
		panic("user.Level != 1")
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
}

func AfterNoChOnlyUploadAfew(userId int) {
	user := withoutM(userId)
	if user.Level != 2 {
		panic("user.Level != 2")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "upload" {
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
}

func AfterChooseMediaGameUnload(userId int) {
	user := withoutM(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "unload" {
		panic("user.Media.Direction != `unload`")
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
}

func AfterUnloadone(userId int) {
	user := withoutM(userId)
	if user.Level != 4 {
		panic("user.Level != 4")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "unload" {
		panic("user.Media.Direction != `unload`")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.DelGameId != 10 {
		panic("user.Media.DelGameId != 10")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
}

func AfterUnloadAfew(userId int) {
	user := withoutM(userId)
	if user.Level != 4 {
		panic("user.Level != 4")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "unload" {
		panic("user.Media.Direction != `unload`")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.DelGameId != 1 {
		panic("user.Media.DelGameId != 1")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
}

func AfterNoChoiseOnlyUpload(userId int) {
	user := withoutM(userId)
	if user.Level != 2 {
		panic("user.Level != 2")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "upload" {
		panic("user.Media.Direction != `upload`")
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
}

func AfterWaitingYourMediaOne(userId int) {
	user := withoutM(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "upload" {
		panic("user.Media.Direction != `upload`")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.DelGameId != 10 {
		panic("user.Media.DelGameId != 10")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
}

func AfterWaitingYourMediaAfew(userId int) {
	user := withoutM(userId)
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "upload" {
		panic("user.Media.Direction != `upload`")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.DelGameId != 1 {
		panic("user.Media.DelGameId != 1")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
}

func AfterUploadOne(userId int) {
	user := withoutM(userId)
	if user.Level != 4 {
		panic("user.Level != 4")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "upload" {
		panic("user.Media.Direction != `upload`")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.DelGameId != 10 {
		panic("user.Media.DelGameId != 10")
	}
	if user.Media.Counter != 1 {
		panic("user.Media.Counter != 1")
	}
	if !media.FindMediaGame(10, fm) {
		panic("don't have the MediaGame")
	}
}

func AfterUploadAfew(userId int) {
	user := withoutM(userId)
	if user.Level != 4 {
		panic("user.Level != 4")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "upload" {
		panic("user.Media.Direction != `upload`")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.DelGameId != 1 {
		panic("user.Media.DelGameId != 1")
	}
	if user.Media.Counter != 4 {
		panic("user.Media.Counter != 4")
	}
	if !findMediaSchedule() {
		panic("Files haven't added yet")
	}
}
