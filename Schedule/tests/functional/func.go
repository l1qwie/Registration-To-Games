package functional

import (
	"Schedule/apptype"
	"log"
)

// makes sure that was the real schedule
func IsItSch(res *apptype.Response) {
	var mes string
	mes += "<b>1.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-02-2025</b>, <em>Время:</em> <b>12:00</b>, <em>Осталось свободных мест:</em> <b>44</b>, <em>Цена за одно место:</em> <b>100 POUNDS</b>\n\n\n"
	mes += "<b>2.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>12-06-2026</b>, <em>Время:</em> <b>11:00</b>, <em>Осталось свободных мест:</em> <b>34</b>, <em>Цена за одно место:</em> <b>10 POUNDS</b>\n\n\n"
	mes += "<b>3.</b> <em>Спорт:</em> <b>Футбол</b>, <em>Дата:</em> <b>12-04-2025</b>, <em>Время:</em> <b>18:00</b>, <em>Осталось свободных мест:</em> <b>14</b>, <em>Цена за одно место:</em> <b>1000 USD</b>\n\n\n"
	mes += "<b>4.</b> <em>Спорт:</em> <b>Волейбол</b>, <em>Дата:</em> <b>02-02-2025</b>, <em>Время:</em> <b>08:00</b>, <em>Осталось свободных мест:</em> <b>77</b>, <em>Цена за одно место:</em> <b>100 RUB</b>\n\n\n"

	if res.ChatId != 488 {
		panic("res.ChatId != 488")
	}
	if res.Message != mes {
		panic("res.Message != mes")
	}
	if res.Keyboard != `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}` {
		panic(`res.Keyboard != "{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}"`)
	}
	if res.ParseMode != "HTML" {
		panic(`res.ParseMode != "HTML"`)
	}

	log.Printf("res.Message = %s", res.Message)
	log.Printf("res.Keyboard = %s", res.Keyboard)
	log.Printf("res.ChatId = %d", res.ChatId)
	log.Printf("res.ParseMode = %s", res.ParseMode)
	log.Printf("res.Act = %s", res.Act)
	log.Printf("res.Level = %d", res.Level)
}
