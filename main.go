package main

import (
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests"
)

// executer.TelegramTest
// executer.Telegram

func main() {
	types.Db = types.ConnectToDatabase()
	//database.CreateSchedule()
	//fmtogram.StartWithTelegram()
	//tests.GlobalTest()
	tests.MediaTest()
	//tests.JustOtherTests()
	//tests.WelcomeTest()
	//tests.RegToGamesTest()
	//tests.SeeTheSchedule()
	types.Db.Close()
}
