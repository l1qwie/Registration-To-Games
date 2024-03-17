package settings

import "RegistrationToGames/fmtogram/formatter"

func everyTime(fm *formatter.Formatter) {
	fm.AssertChatId(899, true)
	fm.AssertEditMessageId(1111, true)
}

// One option - change language
func ChooseOptionOnlyLang(fm *formatter.Formatter) {
	fm.AssertString("У вас пока нет ни одной записи об игре! Вам доступна только смена языка! Пожалуйста, выберите предпочтительный язык", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1}, []string{"Английский язык", "Русский язык", "Турецкий язык", "Главное Меню"}, []string{"en", "ru", "tur", "MainMenu"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)
	everyTime(fm)
}

// Two options - change language and change my records
func ChooseOptionTwo(fm *formatter.Formatter) {
	mes := "Вот игры на которые у вас есть запись:\n\n"
	mes += "<b>1.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-02-2025</b>, <em>Время:</em> <b>12:00</b>, <em>Включая вас с вами будет:</em> <b>3</b>, <em>Цена за одно место:</em> <b>100 POUNDS</b>, <em>Способ оплаты:</em> <b>Онлайн оплата</b>, <em>Статус оплаты:</em> <b>Оплачено</b>\n\n\n"
	mes += "<b>2.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-12-2025</b>, <em>Время:</em> <b>11:00</b>, <em>Включая вас с вами будет:</em> <b>6</b>, <em>Цена за одно место:</em> <b>100 USD</b>, <em>Способ оплаты:</em> <b>Онлайн оплата</b>, <em>Статус оплаты:</em> <b>Неоплачено</b>\n\n\n"
	mes += "<b>3.</b> <em>Спорт:</em> <b>Футбол</b>, <em>Дата:</em> <b>12-06-2025</b>, <em>Время:</em> <b>10:00</b>, <em>Включая вас с вами будет:</em> <b>7</b>, <em>Цена за одно место:</em> <b>100 RUB</b>, <em>Способ оплаты:</em> <b>Наличкой администратору</b>, <em>Статус оплаты:</em> <b>Неоплачено</b>\n\n\n"
	mes += "Вы можете изменить информацию по вашей записи или же изменить язык"
	fm.AssertString(mes, true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Изменить язык", "Изменить записи", "Главное Меню"}, []string{"changeLang", "changRec", "MainMenu"}, []string{"cmd", "cmd", "cmd"}, true)
	everyTime(fm)
}

func ChooseLanguage(fm *formatter.Formatter) {
	fm.AssertString("The bot language has been successfully changed", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1}, []string{"View schedule", "Game registration", "Our photos and videos", "Settings | My games"}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)
	everyTime(fm)
}
