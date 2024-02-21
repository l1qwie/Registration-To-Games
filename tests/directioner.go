package tests

import (
	"log"
	root "registrationtogames/bot/routine"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/tests/database"
	"registrationtogames/tests/routine"
	"registrationtogames/tests/welcome"
)

func DataPreparation(counter int) {
	if !root.Find(456) {
		err := root.CreateUser(456, "ru")
		if err != nil {
			panic(err)
		}
	}
	if counter == 0 {
		database.UpdateLanguage("ru", 456)
		database.UpdateAction("registration", 456)
		database.UpdateLevel(0, 456)
	} else if counter == 1 {
		database.UpdateLanguage("ru", 456)
		database.UpdateAction("registration", 456)
		database.UpdateLevel(1, 456)
	} else if counter == 2 {
		database.UpdateLanguage("ru", 456)
		database.UpdateAction("registration", 456)
		database.UpdateLevel(2, 456)
	}
}

func def() {
	routine.TestRetrevenUser()
	routine.TestRetainUser()
}

func AcceptanceOfResults(output *formatter.Formatter, counter int) {
	if output == nil {
		log.Fatal()
	}
	//def()
	if counter == 0 {
		welcome.TestGreetingsToUser(output)
		database.AfterGreetingsToUserCheckDb(456)
	} else if counter == 1 {
		welcome.TestShowRules(output)
		database.AfterShowRulesCheckDb(456)
	} else if counter == 2 {
		welcome.TestWelcomeToMainMenu(output)
		database.AfterWelcomeToMainMenuCheckDb(456)
	}

	//out.AssertString("Добро пожаловать в нашего бота! Этот бот предназначен для регистрации на спортивные игры в Стамбуле, но для начала вам нужно зарегистрироваться у нас!", true)
	//out.AssertChatId(456, true)
	//out.AssertInlineKeyboard([]int{1}, []string{"Зарегистрироваться"}, []string{"GoReg"}, []string{"cmd"}, true)
	//r.AssertPhoto("AgACAgQAAxkDAAIJRGW3rwaLqri1BkTdVQm1VFA8tE4HAAJeszEbEAABvFHW3MOANm9QFQEAAwIAA20AAzQE", true)
	//r.AssertVideo("BAACAgIAAxkDAAIJW2W3sTguaruPGvo722qeKTcOPwvxAAIzPQACy-DASekiOEg76qGiNAQ", true)
	//close(requests)
}

func Receiving() {
}
