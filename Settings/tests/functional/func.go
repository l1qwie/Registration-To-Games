package functional

import (
	"Settings/apptype"
	"fmt"
)

// [{"text":"","callback_data":"","url":""}]
// {"inline_keyboard":[[{"text":"","callback_data":"","url":""}],[{"text":"","callback_data":"","url":""}]]}
/*
	ch.mes = ""
	ch.kb = ``
	ch.lvl = 0
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = ""
	ch.lang = "ru"
	ch.gameid = 0
	ch.prmode = ""
	ch.res = res
*/

var ch = new(check)

type check struct {
	mes    string
	kb     string
	lvl    int
	lp     int
	act    string
	isCh   string
	lang   string
	gameid int
	prmode string
	res    *apptype.Response
}

// Main tester
// Takes data from struct ckeck and ckecks it between Response and what I want
func (ch *check) maintest() {
	if ch.res.ChatID != 899 {
		panic((fmt.Sprintf(`ch.res.ChatID != 899 because %s`, fmt.Sprintf("ch.res.ChatID == %d", ch.res.ChatID))))
	}
	if ch.res.Message != ch.mes {
		panic(fmt.Sprintf(`ch.res.Message != "%s" because %s`, ch.mes, fmt.Sprintf(`ch.res.Message == "%s"`, ch.res.Message)))
	}
	if ch.res.Keyboard != ch.kb {
		panic(fmt.Sprintf(`ch.res.Keyboard != "%s" because %s`, ch.kb, fmt.Sprintf(`ch.res.Keyboard == "%s"`, ch.res.Keyboard)))
	}
	if ch.res.Level != ch.lvl {
		panic(fmt.Sprintf(`ch.res.Level != %d because %s`, ch.lvl, fmt.Sprintf("ch.res.level == %d", ch.res.Level)))
	}
	if ch.res.LaunchPoint != ch.lp {
		panic(fmt.Sprintf(`ch.res.LaunchPoint != %d because %s`, ch.lp, fmt.Sprintf("ch.res.LaunchPoint == %d", ch.res.LaunchPoint)))
	}
	if ch.res.Act != ch.act {
		panic(fmt.Sprintf(`ch.res.Act != "%s" because %s`, ch.act, fmt.Sprintf(`ch.res.Act == "%s"`, ch.res.Act)))
	}
	if ch.res.IsChanged != ch.isCh {
		panic(fmt.Sprintf(`ch.res.IsChanged != "%s" because %s`, ch.isCh, fmt.Sprintf(`ch.res.IsChanged == "%s"`, ch.res.IsChanged)))
	}
	if ch.res.Language != ch.lang {
		panic(fmt.Sprintf(`ch.res.Language != "%s" because %s`, ch.lang, fmt.Sprintf(`ch.res.Language == "%s"`, ch.res.Language)))
	}
	if ch.res.GameId != ch.gameid {
		panic(fmt.Sprintf(`ch.res.GameId != %d because %s`, ch.gameid, fmt.Sprintf("ch.res.GameId == %d", ch.res.GameId)))
	}
	if ch.res.ParseMode != ch.prmode {
		panic(fmt.Sprintf(`ch.res.ParseMode != %s because %s`, ch.prmode, fmt.Sprintf(`ch.res.ParseMode == "%s"`, ch.res.ParseMode)))
	}
}

// Data which I wait after after the function is executed "chooseOptions"
// With only one thing that user can change - language
func ChOptOnlyLang(res *apptype.Response) {
	ch.mes = "У вас пока нет ни одной записи об игре! Вам доступна только смена языка! Пожалуйста, выберите предпочтительный язык"
	ch.kb = `{"inline_keyboard":[[{"text":"Английский язык","callback_data":"en","url":""}],[{"text":"Русский язык","callback_data":"ru","url":""}],[{"text":"Турецкий язык","callback_data":"tur","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 2
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = "language"
	ch.lang = "ru"
	ch.gameid = 0
	ch.prmode = ""
	ch.res = res
	ch.maintest()
}

// Data which I wait after after the function is executed "chooseOptions" after the first try
func ChOptOnlyLangTr(res *apptype.Response) {
	ch.mes = "Пожалуйста, выберите предпочтительный язык"
	ch.kb = `{"inline_keyboard":[[{"text":"Английский язык","callback_data":"en","url":""}],[{"text":"Русский язык","callback_data":"ru","url":""}],[{"text":"Турецкий язык","callback_data":"tur","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 2
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = "language"
	ch.lang = "ru"
	ch.gameid = 0
	ch.prmode = ""
	ch.res = res
	ch.maintest()
}

// Data which I wait after after the function is executed "changeLanguge"
func ChooseLang(res *apptype.Response) {
	ch.mes = "The bot language has been successfully changed"
	ch.kb = `{"inline_keyboard":[[{"text":"View schedule","callback_data":"Looking Schedule","url":""}],[{"text":"Game registration","callback_data":"Reg to games","url":""}],[{"text":"Our photos and videos","callback_data":"Photo\u0026Video","url":""}],[{"text":"Settings | My games","callback_data":"My records","url":""}]]}`
	ch.lvl = 3
	ch.lp = 0
	ch.act = "divarication"
	ch.isCh = "language"
	ch.lang = "en"
	ch.gameid = 0
	ch.prmode = ""
	ch.res = res
	ch.maintest()
	if !checkUpdtLang(ch.lang) {
		panic("language wasn't changed")
	}
}
