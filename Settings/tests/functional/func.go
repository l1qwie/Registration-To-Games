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

// Data which I wait after after the function is executed "chooseOptions"
// With two things that user can change - language or records to games
func ChTwoOpt(res *apptype.Response) {
	ch.mes = "Вот игры на которые у вас есть запись:\n\n"
	ch.mes += "<b>1.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-06-2025</b>, <em>Время:</em> <b>10:00</b>, <em>Включая вас с вами будет:</em> <b>7</b>, <em>Цена за одно место:</em> <b>100 RUB</b>, <em>Способ оплаты:</em> <b>cash</b>, <em>Статус оплаты:</em> <b>Не оплачено</b>\n\n\n"
	ch.mes += "<b>2.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-12-2025</b>, <em>Время:</em> <b>11:00</b>, <em>Включая вас с вами будет:</em> <b>6</b>, <em>Цена за одно место:</em> <b>100 USD</b>, <em>Способ оплаты:</em> <b>card</b>, <em>Статус оплаты:</em> <b>Не оплачено</b>\n\n\n"
	ch.mes += "<b>3.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-02-2025</b>, <em>Время:</em> <b>12:00</b>, <em>Включая вас с вами будет:</em> <b>3</b>, <em>Цена за одно место:</em> <b>100 POUNDS</b>, <em>Способ оплаты:</em> <b>card</b>, <em>Статус оплаты:</em> <b>Оплачено</b>\n\n\n"
	ch.mes += "Вы можете изменить информацию по вашей записи или же изменить язык"
	ch.kb = `{"inline_keyboard":[[{"text":"Изменить язык","callback_data":"language","url":""}],[{"text":"Изменить бронь на игру","callback_data":"records","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 1
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = ""
	ch.lang = "ru"
	ch.gameid = 0
	ch.prmode = "HTML"
	ch.res = res
	ch.maintest()
}

// Data which I wait after after the function is executed "whatWay"
func ChGame(res *apptype.Response) {
	ch.mes = "Выберите вашу игру\n\n"
	ch.mes += "<b>1.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-06-2025</b>, <em>Время:</em> <b>10:00</b>, <em>Включая вас с вами будет:</em> <b>7</b>, <em>Цена за одно место:</em> <b>100 RUB</b>, <em>Способ оплаты:</em> <b>cash</b>, <em>Статус оплаты:</em> <b>Не оплачено</b>\n\n\n"
	ch.mes += "<b>2.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-12-2025</b>, <em>Время:</em> <b>11:00</b>, <em>Включая вас с вами будет:</em> <b>6</b>, <em>Цена за одно место:</em> <b>100 USD</b>, <em>Способ оплаты:</em> <b>card</b>, <em>Статус оплаты:</em> <b>Не оплачено</b>\n\n\n"
	ch.mes += "<b>3.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-02-2025</b>, <em>Время:</em> <b>12:00</b>, <em>Включая вас с вами будет:</em> <b>3</b>, <em>Цена за одно место:</em> <b>100 POUNDS</b>, <em>Способ оплаты:</em> <b>card</b>, <em>Статус оплаты:</em> <b>Оплачено</b>\n\n\n"
	ch.kb = `{"inline_keyboard":[[{"text":"1","callback_data":"2","url":""}],[{"text":"2","callback_data":"1","url":""}],[{"text":"3","callback_data":"0","url":""}]]}`
	ch.lvl = 2
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = "records"
	ch.lang = "ru"
	ch.gameid = 0
	ch.prmode = "HTML"
	ch.res = res
	ch.maintest()
}

// Data which I wait after after the function is executed "dirOfChanges"
func ChOrDel(res *apptype.Response) {
	ch.mes = "Вы хотите изменить или удалить?"
	ch.kb = `{"inline_keyboard":[[{"text":"Изменить игру","callback_data":"change","url":""}],[{"text":"Удалить игру","callback_data":"del","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 3
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = "records"
	ch.lang = "ru"
	//ch.gameid = 1 // if you want to delete
	ch.gameid = 2 // if you want to change
	ch.prmode = ""
	ch.res = res
	ch.maintest()
}

// Data which I wait after after the function is executed "dirForRec"
func DelGame(res *apptype.Response) {
	ch.mes = "Ваша бронь на игру успешно удалена\n\nГлавное Меню"
	ch.kb = `{"inline_keyboard":[[{"text":"Просмотр расписания","callback_data":"Looking Schedule","url":""}],[{"text":"Регистрация на игру","callback_data":"Reg to games","url":""}],[{"text":"Наши фото и видео","callback_data":"Photo\u0026Video","url":""}],[{"text":"Настройки | Мои игры","callback_data":"My records","url":""}]]}`
	ch.lvl = 3
	ch.lp = 0
	ch.act = "divarication"
	ch.isCh = "records"
	ch.lang = "ru"
	ch.gameid = 2
	ch.prmode = ""
	ch.res = res
	ch.maintest()
	if !checkDelGame() {
		panic("The game wasn't deleted")
	}
}

// Data which I wait after after the function is executed "dirOfChanges".
func ChThwWay(res *apptype.Response) {
	ch.mes = "Что вы хотите изменить?"
	ch.kb = `{"inline_keyboard":[[{"text":"Способ оплаты","callback_data":"payment","url":""}],[{"text":"Количество человек со мной","callback_data":"myseats","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 4
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = "records"
	ch.lang = "ru"
	ch.gameid = 2
	ch.prmode = ""
	ch.res = res
	ch.maintest()
}

// Data which I wait after after the function is executed "chengeable"
// The seats part
func WrtSeats(res *apptype.Response) {
	ch.mes = "Выберите или напишите мне количество мест на игру, которые вы хоите занять. На эту игру есть свободных мест 44. У вас уже есть бронь на 7. После ввода ваша бронь на 7 мест будет снята. Всего мест, если не учитывать вашу бронь 51"
	ch.kb = `{"inline_keyboard":[[{"text":"1","callback_data":"1","url":""}],[{"text":"2","callback_data":"2","url":""}],[{"text":"3","callback_data":"3","url":""}]]}`
	ch.lvl = 5
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = "myseats"
	ch.lang = "ru"
	ch.gameid = 2
	ch.prmode = ""
	ch.res = res
	ch.maintest()
}

// End of all changes
func endOfChanges(res *apptype.Response) {
	ch.mes = "Все успешно изменено\n\nГлавное Меню"
	ch.kb = `{"inline_keyboard":[[{"text":"Просмотр расписания","callback_data":"Looking Schedule","url":""}],[{"text":"Регистрация на игру","callback_data":"Reg to games","url":""}],[{"text":"Наши фото и видео","callback_data":"Photo\u0026Video","url":""}],[{"text":"Настройки | Мои игры","callback_data":"My records","url":""}]]}`
	ch.lvl = 3
	ch.lp = 0
	ch.act = "divarication"
	ch.lang = "ru"
	ch.gameid = 2
	ch.prmode = ""
	ch.res = res
}

// Data which I wait after after the function is executed "confirm".
// Change seats part
func ChangeSeats(res *apptype.Response) {
	endOfChanges(res)
	ch.isCh = "myseats"
	ch.maintest()
	if !checkSeatsWereChanged() {
		panic("The seats weren't changed")
	}
}

// Data which I wait after after the function is executed "confirm".
// Change paymethod part (only cash)
func PaybyCash(res *apptype.Response) {
	endOfChanges(res)
	ch.isCh = "payment"
	ch.maintest()
	if !checkPaymethodWasChanged("cash") {
		panic("The paymethod wasn't changed to cash")
	}
}

func PaybyCard(res *apptype.Response) {
	ch.mes = "Все успешно изменено\n\n"
	ch.kb = `{"inline_keyboard":[[{"text":"Оплатить","callback_data":"","url":"https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3"}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 5
	ch.lp = 0
	ch.act = "settings"
	ch.isCh = "payment"
	ch.lang = "ru"
	ch.gameid = 2
	ch.prmode = ""
	ch.res = res
	ch.maintest()
	if !checkPaymethodWasChanged("card") {
		panic("The paymethod wasn't changed to card")
	}
}
