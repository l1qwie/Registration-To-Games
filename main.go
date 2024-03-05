package main

import (
	"registrationtogames/fmtogram/types"
	"registrationtogames/tests"
)

// executer.TelegramTest
// executer.Telegram

func main() {
	types.Db = types.ConnectToDatabase()
	//async.Main()
	//err := database.FirstConnect()
	//database.CreateSchedule()
	//fmtogram.StartWithTelegram()
	//tests.GlobalTest()
	tests.MediaTest()
	//tests.JustOtherTests()
	//tests.Welcome()
	//tests.RegToGames()
	//tests.SeeTheSchedule()
	types.Db.Close()
}
