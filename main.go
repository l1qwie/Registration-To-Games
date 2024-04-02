package main

import (
	"RegistrationToGames/fmtogram/types"
)

func main() {
	types.Db = types.ConnectToDatabase(true)
	//tests.TelegramTests()
	//tests.SendMediaGroup()
	//tests.SendPhoto()
	//fmtogram.StartWithTelegram()
	//tests.JustOtherTests()
	//tests.WelcomeTest()
	//tests.RegToGamesTest()
	//tests.SeeTheSchedule()
	//tests.MediaTest()
	//tests.SettingsTest()
	types.Db.Close()
}
