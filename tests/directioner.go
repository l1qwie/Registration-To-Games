package tests

import (
	"log"
	"registrationtogames/bot/testdatabase"
	"registrationtogames/fmtogram/errors"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/tests/routine"
)

func DataPreparation() {
	testdatabase.UpdateLanguage("ru", 456)
	testdatabase.UpdateAction("registration", 456)
	testdatabase.UpdateLevel(0, 456)
}

func AcceptanceOfResults(out *formatter.Formatter) {
	if out == nil {
		log.Fatal()
	}
	defer errors.CreateIntestines()
	routine.TestRetrevenUser()
	routine.TestRetainUser()

	//out.AssertString("Добро пожаловать в нашего бота! Этот бот предназначен для регистрации на спортивные игры в Стамбуле, но для начала вам нужно зарегистрироваться у нас!", true)
	//out.AssertChatId(456, true)
	//out.AssertInlineKeyboard([]int{1}, []string{"Зарегистрироваться"}, []string{"GoReg"}, []string{"cmd"}, true)
	//r.AssertPhoto("AgACAgQAAxkDAAIJRGW3rwaLqri1BkTdVQm1VFA8tE4HAAJeszEbEAABvFHW3MOANm9QFQEAAwIAA20AAzQE", true)
	//r.AssertVideo("BAACAgIAAxkDAAIJW2W3sTguaruPGvo722qeKTcOPwvxAAIzPQACy-DASekiOEg76qGiNAQ", true)
	//close(requests)
}

func Receiving() {
}
