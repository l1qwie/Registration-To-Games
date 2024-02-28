package main

import (
	"registrationtogames/fmtogram"
	"registrationtogames/fmtogram/types"
)

// executer.TelegramTest
// executer.Telegram

func main() {
	types.Db = types.ConnectToDatabase()
	//async.Main()
	//err := database.FirstConnect()
	//database.CreateSchedule()
	//fmtogram.StartWithTelegram()
	//fmtogram.StartTests()
	//fmtogram.JustOtherTests()
	fmtogram.Welcome()
	fmtogram.RegToGames()
	fmtogram.SeeTheSchedule()
	types.Db.Close()
}
