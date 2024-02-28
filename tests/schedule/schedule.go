package schedule

import (
	"registrationtogames/fmtogram/formatter"
)

func ShowTheSchedule(fm *formatter.Formatter) {
	var mes string
	fm.AssertChatId(488, true)
	mes += "<b>1.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-02-2025</b>, <em>Время:</em> <b>12:00</b>, <em>Осталось свободных мест:</em> <b>44</b>, <em>Цена за одно место:</em> <b>100 POUNDS</b>\n\n\n"
	mes += "<b>2.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-02-2026</b>, <em>Время:</em> <b>11:00</b>, <em>Осталось свободных мест:</em> <b>34</b>, <em>Цена за одно место:</em> <b>10 POUNDS</b>\n\n\n"
	mes += "<b>3.</b> <em>Спорт:</em> <b>Футбол</b>, <em>Дата:</em> <b>12-04-2025</b>, <em>Время:</em> <b>18:00</b>, <em>Осталось свободных мест:</em> <b>14</b>, <em>Цена за одно место:</em> <b>1000 USD</b>\n\n\n"
	mes += "<b>4.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>02-02-2025</b>, <em>Время:</em> <b>08:00</b>, <em>Осталось свободных мест:</em> <b>77</b>, <em>Цена за одно место:</em> <b>100 RUB</b>\n\n\n"
	fm.AssertParseMode("HTML", true)
	fm.AssertString(mes, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(9999, true)
}
