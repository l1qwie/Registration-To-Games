package functional

import (
	"Game/apptype"
	"fmt"
)

// [{"text":"","callback_data":"","url":""}]
// `{"inline_keyboard":[[{"text":"","callback_data":"","url":""}],[{"text":"","callback_data":"","url":""}]]}`

/*
	ch.kb = ``
	ch.mes = ""
	ch.level = 0
	ch.act = ""
	ch.launchpoint = 0
	ch.sport = ""
	ch.date = 0
	ch.time = 0
	ch.seats = 0
	ch.price = 0
	ch.currency = ""
	ch.link = ""
	ch.address = ""
	ch.res = res
	ch.maintest()

*/

var ch = new(check)

type check struct {
	kb          string
	mes         string
	level       int
	act         string
	launchpoint int
	sport       string
	date        int
	time        int
	seats       int
	price       int
	currency    string
	link        string
	address     string

	res *apptype.Response
}

func (ch *check) maintest() {
	if ch.res.ChatID != 1111 {
		panic(fmt.Sprintf("ch.res.ChatID != 1111 because ch.res.ChatID = %d", ch.res.ChatID))
	}
	if ch.res.Keyboard != ch.kb {
		panic(fmt.Sprintf("ch.res.Keyboard != ch.kb because ch.res.Keyboard = %s but ch.kb = %s", ch.res.Keyboard, ch.kb))
	}
	if ch.res.Message != ch.mes {
		panic(fmt.Sprintf("ch.res.Message != ch.mes because ch.res.Message = %s but ch.mes = %s", ch.res.Message, ch.mes))
	}
	if ch.res.Level != ch.level {
		panic(fmt.Sprintf("ch.res.Level != ch.level because ch.res.Level = %d but ch.level = %d", ch.res.Level, ch.level))
	}
	if ch.res.Act != ch.act {
		panic(fmt.Sprintf("ch.res.Act != ch.act because ch.res.Act = %s but ch.act = %s", ch.res.Act, ch.act))
	}
	if ch.res.LaunchPoint != ch.launchpoint {
		panic(fmt.Sprintf("ch.res.LaunchPoint != ch.launchpoint because ch.res.LaunchPoint = %d but ch.launchpoint = %d", ch.res.LaunchPoint, ch.launchpoint))
	}
	if ch.res.Sport != ch.sport {
		panic(fmt.Sprintf("ch.res.Sport != ch.sport because ch.res.Sport = %s but ch.sport = %s", ch.res.Sport, ch.sport))
	}
	if ch.res.Date != ch.date {
		panic(fmt.Sprintf("ch.res.Date != ch.date because ch.res.Date = %d but ch.date = %d", ch.res.Date, ch.date))
	}
	if ch.res.Time != ch.time {
		panic(fmt.Sprintf("ch.res.Time != ch.time because ch.res.Time = %d but ch.time = %d", ch.res.Time, ch.time))
	}
	if ch.res.Seats != ch.seats {
		panic(fmt.Sprintf("ch.res.Seats != ch.seats because ch.res.Seats = %d but ch.seats = %d", ch.res.Seats, ch.seats))
	}
	if ch.res.Price != ch.price {
		panic(fmt.Sprintf("ch.res.Price != ch.price because ch.res.Price = %d but ch.price = %d", ch.res.Price, ch.price))
	}
}

func ChooseOneOfThree(res *apptype.Response) {
}

func SelectSport(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Волейбол","callback_data":"volleyball","url":""}],[{"text":"Футбол","callback_data":"football","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "Выберите вид спорта"
	ch.level = 2
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = ""
	ch.date = 0
	ch.time = 0
	ch.seats = 0
	ch.price = 0
	ch.currency = ""
	ch.link = ""
	ch.address = ""
	ch.res = res
	ch.maintest()
}

func SelectDate(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "<b>Вид спорта:</b> Волейбол\n\n\nВведите дату проведения игры в формате ДДММГГГГ используя любой разделитель"
	ch.level = 3
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 0
	ch.time = 0
	ch.seats = 0
	ch.price = 0
	ch.currency = ""
	ch.link = ""
	ch.address = ""
	ch.res = res
	ch.maintest()
}

func SelectTime(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "<b>Вид спорта:</b> Волейбол\n<b>Дата:</b> 09-12-2024\n\n\nВведите время проведения игры в формате ЧЧММ используя любой разделитель"
	ch.level = 4
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 20241209
	ch.time = 0
	ch.seats = 0
	ch.price = 0
	ch.currency = ""
	ch.link = ""
	ch.address = ""
	ch.res = res
	ch.maintest()
}

func SelectSeats(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "<b>Вид спорта:</b> Волейбол\n<b>Дата:</b> 09-12-2024\n<b>Время:</b> 19:00\n\n\nВведите количество свободных мест на эту игру"
	ch.level = 5
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 20241209
	ch.time = 1900
	ch.seats = 0
	ch.price = 0
	ch.currency = ""
	ch.link = ""
	ch.address = ""
	ch.res = res
	ch.maintest()
}

func SelectPrice(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "<b>Вид спорта:</b> Волейбол\n<b>Дата:</b> 09-12-2024\n<b>Время:</b> 19:00\n<b>Всего свободных мест:</b> 15\n\n\nВведите цену за одно место в формате цифры на эту игру"
	ch.level = 6
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 20241209
	ch.time = 1900
	ch.seats = 15
	ch.price = 0
	ch.currency = ""
	ch.link = ""
	ch.address = ""
	ch.res = res
	ch.maintest()
}

func SelectCurrency(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "<b>Вид спорта:</b> Волейбол\n<b>Дата:</b> 09-12-2024\n<b>Время:</b> 19:00\n<b>Всего свободных мест:</b> 15\n\n\nВведите имя валюты. Я не никак не контралирую то название, которое вы введете, так что советую вводить так, чтобы все понимали. Пример: USD EURO TL и тд"
	ch.level = 7
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 20241209
	ch.time = 1900
	ch.seats = 15
	ch.price = 1000
	ch.currency = ""
	ch.link = ""
	ch.address = ""
	ch.res = res
	ch.maintest()
}

func SelectLink(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = fmt.Sprint("<b>Вид спорта:</b> Волейбол\n<b>Дата:</b> 09-12-2024\n<b>Время:</b> 19:00\n<b>Всего свободных мест:</b> 15\n<b>Цена на одно место:</b> 1000 RUB\n\n\n", `Пршлите ссылку с Google Maps с тем местом, где будет проходить игра`)
	ch.level = 8
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 20241209
	ch.time = 1900
	ch.seats = 15
	ch.price = 1000
	ch.currency = "RUB"
	ch.link = ""
	ch.address = ""
	ch.res = res
	ch.maintest()
}

func SelectAddress(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "<b>Вид спорта:</b> Волейбол\n<b>Дата:</b> 09-12-2024\n<b>Время:</b> 19:00\n<b>Всего свободных мест:</b> 15\n<b>Цена на одно место:</b> 1000 RUB\n<b>Ссылка на место проведения:</b> https://www.google.com/maps?q=36.893445,30.709591\n\n\nВведите название адреса"
	ch.level = 9
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 20241209
	ch.time = 1900
	ch.seats = 15
	ch.price = 1000
	ch.currency = "RUB"
	ch.link = "https://www.google.com/maps?q=36.893445,30.709591"
	ch.address = ""
	ch.res = res
	ch.maintest()
}

func SemiFinal(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Сохранить","callback_data":"save","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "<b>Вид спорта:</b> Волейбол\n<b>Дата:</b> 09-12-2024\n<b>Время:</b> 19:00\n<b>Всего свободных мест:</b> 15\n<b>Цена на одно место:</b> 1000 RUB\n<b>Ссылка на место проведения:</b> https://www.google.com/maps?q=36.893445,30.709591\n<b>Название адреса:</b> Игровая Площадка\n\n\nВы закончили заполнять информацию для создания новой игры. Сохраните эту игру, если все данные верны"
	ch.level = 10
	ch.act = "game"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 20241209
	ch.time = 1900
	ch.seats = 15
	ch.price = 1000
	ch.currency = "RUB"
	ch.link = "https://www.google.com/maps?q=36.893445,30.709591"
	ch.address = "Игровая Площадка"
	ch.res = res
	ch.maintest()
}

func Final(res *apptype.Response) {
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.mes = "Игра сохранена и теперь доступна вашим клиентам для регистрации"
	ch.level = 10
	ch.act = "divarication"
	ch.launchpoint = 0
	ch.sport = "volleyball"
	ch.date = 20241209
	ch.time = 1900
	ch.seats = 15
	ch.price = 1000
	ch.currency = "RUB"
	ch.link = "https://www.google.com/maps?q=36.893445,30.709591"
	ch.address = "Игровая Площадка"
	ch.res = res
	ch.maintest()
	if !gameSaved() {
		panic("The game was't saved in the database")
	}
}
