package tests

import (
	"RegistrationToGames/fmtogram"
	"RegistrationToGames/fmtogram/errors"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests/database"
	"RegistrationToGames/tests/media"
	"RegistrationToGames/tests/othertests"
	"RegistrationToGames/tests/registration"
	"RegistrationToGames/tests/schedule"
	"RegistrationToGames/tests/settings"
	"RegistrationToGames/tests/texecuter"
	"RegistrationToGames/tests/welcome"
	"fmt"
)

const (
	userIdT1   int = 456
	userIdT2   int = 477
	userIdT3   int = 488
	userIdT4   int = 499
	userIdT5   int = 899
	gameId     int = 2
	gameIdS    int = 3
	pastgameId int = 10
)

/*
	ts := new(texecuter.TestStuct)
	ts.Round = int
	ts.Name = ""
	ts.Update = []texecuter.Update{}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{}
	ts.ArrFuncMr = []func() *types.MessageResponse{}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{}
	ts.ArrFuncAss = []func(*formatter.Formatter){}
	ts.ArrFuncChDB = []func(int){}
	ts.DoTest()
*/

func UnloadOne() {
	defer database.DeleteUser(userIdT4)
	defer database.DeleteGame(pastgameId)
	defer database.DeleteMedia(pastgameId)
	database.CreateNotFullMediaGame(pastgameId)
	database.CreateUser(userIdT4)
	ts := new(texecuter.TestStuct)
	ts.Round = 3
	ts.Name = "UnloadOneTest"
	ts.Update = []texecuter.Update{{Act: "photos and videos", Lang: "ru", Level: 0, UserId: userIdT4},
		{Act: "photos and videos", Lang: "ru", Level: 1, UserId: userIdT4},
		{Act: "photos and videos", Lang: "ru", Level: 3, UserId: userIdT4}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{media.ChDir, media.ChMUnload, media.UnlOne}
	ts.ArrFuncMr = []func() *types.MessageResponse{media.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{media.Trash, media.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){media.ChooseDirection, media.ChooseMediaGameUnload, media.Unloadone}
	ts.ArrFuncChDB = []func(int){database.AfterChooseDirection, database.AfterChooseMediaGameUnload, database.AfterUnloadone}
	ts.DoTest()
}

func UploadOne() {
	defer database.DeleteUser(userIdT4)
	defer database.DeleteGame(pastgameId)
	defer database.DeleteMedia(pastgameId)
	database.CreateUser(userIdT4)
	database.CreateEmptyMediaGame(pastgameId)
	ts := new(texecuter.TestStuct)
	ts.Round = 3
	ts.Name = "UploadOneTest"
	ts.Update = []texecuter.Update{{Act: "photos and videos", Lang: "ru", Level: 0, UserId: userIdT4},
		{Act: "photos and videos", Lang: "ru", Level: 2, UserId: userIdT4},
		{Act: "photos and videos", Lang: "ru", Level: 3, UserId: userIdT4}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{media.ChDir, media.ChWaitMediaOne, media.UplOne}
	ts.ArrFuncMr = []func() *types.MessageResponse{media.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{media.Trash, media.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){media.NoChoiseOnlyUploadOne, media.WaitingYourMedia, media.Upload}
	ts.ArrFuncChDB = []func(int){database.AfterNoChoiseOnlyUpload, database.AfterWaitingYourMediaOne, database.AfterUploadOne}
	ts.DoTest()
}

func UnloadAfew() {
	defer database.DeleteUser(userIdT4)
	defer database.DeleteSchedule()
	defer database.DeleteMediaSchedule()
	database.CreateMediaShedule()
	database.FillMEdiaSchedule()
	database.CreateUser(userIdT4)
	ts := new(texecuter.TestStuct)
	ts.Round = 3
	ts.Name = "UnloadAfewTest"
	ts.Update = []texecuter.Update{{Act: "photos and videos", Lang: "ru", Level: 0, UserId: userIdT4},
		{Act: "photos and videos", Lang: "ru", Level: 1, UserId: userIdT4},
		{Act: "photos and videos", Lang: "ru", Level: 3, UserId: userIdT4}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{media.ChDir, media.ChMUnload, media.UnlAfew}
	ts.ArrFuncMr = []func() *types.MessageResponse{media.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{media.Trash, media.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){media.ChooseDirection, media.ChooseMediaGamesUnload, media.Unloadall}
	ts.ArrFuncChDB = []func(int){database.AfterChooseDirection, database.AfterChooseMediaGameUnload, database.AfterUnloadAfew}
	ts.DoTest()
}

func UploadAfew() {
	defer database.DeleteUser(userIdT4)
	defer database.DeleteSchedule()
	defer database.DeleteMediaSchedule()
	database.CreateMediaShedule()
	database.CreateUser(userIdT4)

	ts := new(texecuter.TestStuct)
	ts.Round = 3
	ts.Name = "UploadAfewTest"
	ts.Update = []texecuter.Update{{Act: "photos and videos", Lang: "ru", Level: 0, UserId: userIdT4},
		{Act: "photos and videos", Lang: "ru", Level: 2, UserId: userIdT4},
		{Act: "photos and videos", Lang: "ru", Level: 3, UserId: userIdT4}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{media.ChDir, media.ChWaitMediaAfew, media.UplAfew}
	ts.ArrFuncMr = []func() *types.MessageResponse{media.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{media.Trash, media.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){media.NoChoiseOnlyUploadAfew, media.WaitingYourMedia, media.Upload}
	ts.ArrFuncChDB = []func(int){database.AfterNoChOnlyUploadAfew, database.AfterWaitingYourMediaAfew, database.AfterUploadAfew}
	ts.DoTest()
}

func Welcome() {
	defer database.DeleteUser(userIdT1)
	database.CreateUser(userIdT1)
	ts := new(texecuter.TestStuct)
	ts.Round = 3
	ts.Name = "WelcomeTest"
	ts.Update = []texecuter.Update{{Act: "registration", Lang: "ru", Level: 0, UserId: userIdT1}, {Act: "registration", Lang: "ru", Level: 1, UserId: userIdT1}, {Act: "registration", Lang: "ru", Level: 2, UserId: userIdT1}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{welcome.Start, welcome.ShowRules, welcome.ToMainMenu}
	ts.ArrFuncMr = []func() *types.MessageResponse{welcome.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{welcome.Trash, welcome.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){welcome.GreetingsToUser, welcome.ShowTheRules, welcome.WelcomeToMainMenu}
	ts.ArrFuncChDB = []func(int){database.AfterGreetingsToUser, database.AfterShowRules, database.AfterWelcomeMainMenu}
	ts.DoTest()
}

func Schedule() {
	defer database.DeleteUser(userIdT3)
	defer database.DeleteSchedule()
	database.CreateUser(userIdT3)
	database.CreateSchedule()
	ts := new(texecuter.TestStuct)
	ts.Round = 1
	ts.Name = "ScheduleTest"
	ts.Update = []texecuter.Update{{Act: "see schedule", Lang: "ru", Level: 0, UserId: userIdT3}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{schedule.SeeSchedule}
	ts.ArrFuncMr = []func() *types.MessageResponse{schedule.MesResp}
	ts.ArrFuncAss = []func(*formatter.Formatter){schedule.ShowTheSchedule}
	ts.ArrFuncChDB = []func(int){database.AfterShowTheSchedule}
	ts.DoTest()
}

func RegToOneGame() {
	defer database.DeleteUser(userIdT2)
	defer database.DeleteGameWithUser(gameId, userIdT2)
	defer database.DeleteGame(gameId)
	database.CreateUser(userIdT2)
	database.CreateGame(gameId)
	ts := new(texecuter.TestStuct)
	ts.Round = 5
	ts.Name = "RegToGamesTest"
	ts.Update = []texecuter.Update{{Act: "reg to games", Lang: "ru", Level: 0, UserId: userIdT2},
		{Act: "reg to games", Lang: "ru", Level: 1, UserId: userIdT2},
		{Act: "reg to games", Lang: "ru", Level: 2, UserId: userIdT2},
		{Act: "reg to games", Lang: "ru", Level: 3, UserId: userIdT2},
		{Act: "reg to games", Lang: "ru", Level: 4, UserId: userIdT2}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{registration.PresScheduele, registration.ChGame, registration.ChSeats, registration.ChPayment, registration.BWishTR}
	ts.ArrFuncMr = []func() *types.MessageResponse{registration.MesResp, registration.MesResp, registration.MesResp, registration.MesResp, registration.BwishMR}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{registration.Trash, registration.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){registration.PresentationScheduele, registration.ChooseGame, registration.ChooseSeats, registration.ChoosePayment, registration.BestWishes}
	ts.ArrFuncChDB = []func(int){database.AfterPresentationScheduele, database.AfterChooseGameR, database.AfterAfterChooseSeats, database.AfterChoosePayment, database.AfterBestWishes}
	ts.DoTest()
}

func ChangeLanguage() {
	defer database.DeleteUser(userIdT5)
	database.CreateUser(userIdT5)
	ts := new(texecuter.TestStuct)
	ts.Round = 2
	ts.Name = "ChangeLanguageTest"
	ts.Update = []texecuter.Update{{Act: "settings", Lang: "ru", Level: 0, UserId: userIdT5}, {Act: "settings", Lang: "ru", Level: 1, UserId: userIdT5}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{settings.ChOpt, settings.ChLang}
	ts.ArrFuncMr = []func() *types.MessageResponse{settings.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{settings.Trash, settings.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){settings.ChooseOptionOnlyLang, settings.ChooseLanguage}
	ts.ArrFuncChDB = []func(int){database.AfterChOptLang, database.AfterChooseLanguage}
	ts.DoTest()
}

func DeleteGame() {
	defer database.DeleteUser(userIdT5)
	defer database.DeleteUserSchedule(userIdT5)
	defer database.DeleteSchedule()
	database.CreateScheduleForUser()
	database.CreateUserScehdule(userIdT5)
	database.CreateUser(userIdT5)
	ts := new(texecuter.TestStuct)
	ts.Round = 4
	ts.Name = "DeleteGameTest"
	ts.Update = []texecuter.Update{{Act: "settings", Lang: "ru", Level: 0, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 1, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 2, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 3, UserId: userIdT5}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{settings.ChOpt, settings.WhatOpt, settings.DeleteGame, settings.Del}
	ts.ArrFuncMr = []func() *types.MessageResponse{settings.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{settings.Trash, settings.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){settings.ChooseOptionTwo, settings.ChoGame, settings.ChOrDelGame, settings.DelGame}
	ts.ArrFuncChDB = []func(int){database.AfterChooseOpt, database.AfterChGame, database.AfterChOrDelGameDel, database.AfterDelGame}
	ts.DoTest()
}

func ChangePayment() {
	defer database.DeleteUser(userIdT5)
	defer database.DeleteUserSchedule(userIdT5)
	defer database.DeleteSchedule()
	database.CreateScheduleForUser()
	database.CreateUserScehdule(userIdT5)
	database.CreateUser(userIdT5)
	ts := new(texecuter.TestStuct)
	ts.Round = 5
	ts.Name = "ChangePaymentTest"
	ts.Update = []texecuter.Update{{Act: "settings", Lang: "ru", Level: 0, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 1, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 2, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 3, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 4, UserId: userIdT5}} // {Act: "settings", Lang: "ru", Level: 5, UserId: userIdT5}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{settings.ChOpt, settings.WhatOpt, settings.ChDelGame, settings.Change, settings.Payment}
	ts.ArrFuncMr = []func() *types.MessageResponse{settings.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{settings.Trash, settings.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){settings.ChooseOptionTwo, settings.ChoGame, settings.ChOrDelGame, settings.ChangeWay, settings.PayByCash} // settings.WhichPayment level 4
	ts.ArrFuncChDB = []func(int){database.AfterChooseOpt, database.AfterChGame, database.AfterChOrDelGame, database.AfterChangeWay, database.AfterPBCash}  //  database.AfterWhPayment level 4
	ts.DoTest()
}

func ChangeSeats() {
	defer database.DeleteUser(userIdT5)
	defer database.DeleteUserSchedule(userIdT5)
	defer database.DeleteSchedule()
	database.CreateScheduleForUser()
	database.CreateUserScehdule(userIdT5)
	database.CreateUser(userIdT5)
	ts := new(texecuter.TestStuct)
	ts.Round = 6
	ts.Name = "ChangeSeatsTest"
	ts.Update = []texecuter.Update{{Act: "settings", Lang: "ru", Level: 0, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 1, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 2, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 3, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 4, UserId: userIdT5},
		{Act: "settings", Lang: "ru", Level: 5, UserId: userIdT5}}
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.Fmt = make(chan *formatter.Formatter, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{settings.ChOpt, settings.WhatOpt, settings.ChDelGame, settings.Change, settings.Seats, settings.NumOfSeast}
	ts.ArrFuncMr = []func() *types.MessageResponse{settings.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{settings.Trash, settings.Trash2}
	ts.ArrFuncAss = []func(*formatter.Formatter){settings.ChooseOptionTwo, settings.ChoGame, settings.ChOrDelGame, settings.ChangeWay, settings.WriteSeats, settings.ChanSeats}
	ts.ArrFuncChDB = []func(int){database.AfterChooseOpt, database.AfterChGame, database.AfterChOrDelGame, database.AfterChangeWay, database.AfterWrSeats, database.AfterChanSeats}
	ts.DoTest()
}

func RegToGamesTest() {
	defer errors.MakeIntestines()
	RegToOneGame() //with one game in database
	fmt.Print("All RegToGamesTest were alright!\n")
}

func WelcomeTest() {
	defer errors.MakeIntestines()
	Welcome()
	fmt.Print("All WelcomeTests were alright!\n")
}

func SeeTheSchedule() {
	Schedule()
	fmt.Print("All SeeTheScheduleTests were alright!\n")
}

func MediaTest() {
	defer errors.MakeIntestines()
	//UnloadOne() //with two games (to unload and upload)
	//UploadOne() //with one games (to upload)
	//UnloadAfew() //with four games (only to unload)
	UploadAfew() //with four games (only to upload)
	fmt.Print("All MediaTests were alright!\n")
}

func SettingsTest() {
	defer errors.MakeIntestines()
	ChangeLanguage() //test "change language" functionality // no games
	DeleteGame()     //test "delete user game" functionality
	ChangeSeats()    //test "change seats on user games" functionality
	ChangePayment()  //test "change payment on user games" functionality
	fmt.Print("All SettingsTest were alright!\n")
}

func delEnv() {
	database.DeleteSchedule()
	database.DeleteEmptyMediaGame(10)
}

func createEnv() {
	delEnv()
	database.CreateSchedule()
	database.CreateEmptyMediaGame(10)
}

func TelegramTests() {
	createEnv()
	fmtogram.StartWithTelegram()
}

func SendPhoto() {
	fm := new(formatter.Formatter)
	fm.AddPhotoFromStorage("red.jpg")
	fm.WriteChatId(738070596)
	_, err := fm.Send()
	if err != nil {
		panic(err)
	}
}
func SendMediaGroup() {
	fm := new(formatter.Formatter)
	mg := make([]types.Media, 3)

	mg[0].Type = "photo"
	mg[0].Media = "AgACAgIAAxkDAAIN12YECI0n-nb0tvzwNOpQpPUPQXr-AAJI1TEbTcHZSqUQNYC0VYW-AQADAgADcwADNAQ"
	mg[1].Type = "photo"
	mg[1].Media = "AgACAgIAAxkDAAIN12YECI0n-nb0tvzwNOpQpPUPQXr-AAJI1TEbTcHZSqUQNYC0VYW-AQADAgADcwADNAQ"
	mg[2].Type = "photo"
	mg[2].Media = "AgACAgIAAxkDAAIN12YECI0n-nb0tvzwNOpQpPUPQXr-AAJI1TEbTcHZSqUQNYC0VYW-AQADAgADcwADNAQ"
	/*
		mg[0].Type = "photo"
		mg[0].Media = "red.jpg"
		mg[1].Type = "photo"
		mg[1].Media = "redd.jpg"
		mg[2].Type = "photo"
		mg[2].Media = "reddd.jpg"
	*/
	fm.AddMapOfMedia(mg)
	fm.WriteChatId(738070596)
	_, err := fm.Send()
	if err != nil {
		panic(err)
	}
}

func JustOtherTests() {
	defer errors.MakeIntestines()
	if !othertests.TestFromIntToStrDate() {
		panic("Date isn't a date")
	}
	fmt.Println("Date is correct")
	if !othertests.TestFromIntToStrTime() {
		panic("Time isn't a time")
	}
	fmt.Println("Time is correct")
}
