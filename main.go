package main

import (
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests"
)

func main() {
	types.Db = types.ConnectToDatabase()
	//database.CreateSchedule()
	//fmtogram.StartWithTelegram()
	//tests.GlobalTest()
	//tests.JustOtherTests()
	//tests.WelcomeTest()
	//tests.RegToGamesTest()
	//tests.SeeTheSchedule()
	tests.MediaTest()
	types.Db.Close()
}
