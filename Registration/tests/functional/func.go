package functional

import (
	apptype "Registraion/apptype"
	"fmt"
)

type check struct {
	Mes       string
	Kb        string
	Payment   string
	Act       string
	Lp        int
	GameId    int
	Seats     int
	ParseMode string
	Res       *apptype.Response
}

var (
	truefunc = []func(*apptype.Response, *check){presentationScheduele, chooseGame, chooseSeats, choosePayment, bestWishes}
)

// [{"text":"","callback_data":"","url":""}]
// {"inline_keyboard":[[{"text":"","callback_data":"","url":""}],[{"text":"","callback_data":"","url":""}]]}
func (ch *check) maintest() {
	if ch.Res.ChatID != 477 {
		panic(` ch.Res.ChatID != 477`)
	}
	if ch.Res.Message != ch.Mes {
		panic(fmt.Sprintf(`ch.Res.Message != "%s"`, ch.Mes))
	}
	if ch.Res.Keyboard != ch.Kb {
		panic(fmt.Sprintf(`ch.Res.Keyboard != "%s"`, ch.Kb))
	}
	if ch.Res.LaunchPoint != ch.Lp {
		panic(fmt.Sprintf(`ch.Res.LaunchPoint != %d`, ch.Lp))
	}
	if ch.Res.GameId != ch.GameId {
		panic(fmt.Sprintf(`ch.Res.GameId != %d`, ch.GameId))
	}
	if ch.Res.Seats != ch.Seats {
		panic(fmt.Sprintf(`ch.Res.Seats != %d`, ch.Seats))
	}
	if ch.Res.Payment != ch.Payment {
		panic(fmt.Sprintf(`ch.Res.Payment != "%s"`, ch.Payment))
	}
	if ch.Res.Act != ch.Act {
		panic(fmt.Sprintf(`ch.Res.Act != "%s"`, ch.Act))
	}
	if ch.Res.ParseMode != ch.ParseMode {
		panic(fmt.Sprintf(`ch.Res.ParseMode != %s`, ch.ParseMode))
	}
}

func presentationScheduele(res *apptype.Response, ch *check) {
	ch.Mes = "Выберите интересующую вас игру"
	ch.Kb = `{"inline_keyboard":[[{"text":"Волейбол 12-02-2025 12:00 Свободный мест осталось: 55","callback_data":"2","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.Payment = ""
	ch.Lp = 0
	ch.GameId = 0
	ch.Seats = 0
	ch.maintest()
}

func chooseGame(res *apptype.Response, ch *check) {
	ch.Mes = "Выберите или введите желаемое количество мест. Прошу обратить ваше внимание на то, что цена за одно место на эту игру составляет 100 USD"
	ch.Kb = `{"inline_keyboard":[[{"text":"Я буду один","callback_data":"1","url":""}],[{"text":"Нас будет двое","callback_data":"2","url":""}],[{"text":"Нас будет трое","callback_data":"3","url":""}]]}`
	ch.Payment = ""
	ch.Lp = 0
	ch.GameId = 2
	ch.Seats = 0
	ch.maintest()
}

func chooseSeats(res *apptype.Response, ch *check) {
	ch.Mes = "Выберите способ оплаты"
	ch.Kb = `{"inline_keyboard":[[{"text":"Онлайн оплата","callback_data":"card","url":""}],[{"text":"Наличкой администратору","callback_data":"cash","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.Payment = ""
	ch.Lp = 0
	ch.GameId = 2
	ch.Seats = 2
	ch.maintest()
}

func choosePayment(res *apptype.Response, ch *check) {
	ch.Mes = `Вам нужно перевести 200 USD. Выше я выслал вам QR-code, по которому вы можете перейти и оплатить или же нажмите на кнопку ниже "Оплатить"`
	ch.Kb = `{"inline_keyboard":[[{"text":"Оплатить","callback_data":"","url":"https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3"}],[{"text":"Дальше","callback_data":"Next","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.Payment = "card"
	ch.Lp = 0
	ch.GameId = 2
	ch.Seats = 2
	ch.maintest()
}

func bestWishes(res *apptype.Response, ch *check) {
	ch.Mes = `Прекрасно! Теперь вы зарегестрированны на игру\n
	1. Вид спорта: <b>Волейбол</b>\n
	2. Дата: <b>12-02-2025</b>\n
	3. Время: <b>12:00</b>\n
	4. Вы записали <b>2</b> персон на эту игру\n
	5. Оплата: <b>by card</b>\n
	6. Ваша стоимость посещения: <b>200 USD</b>\n\n
	
	***Вы можете изменить некоторые данные Вашей записи\n
	или же удалить ее в Главном Меню нажав на <b>"Настройки | Мои игры"</b>***\n\n
	
	❤️❤️❤️Ждем вас по адресу: Кладбище в Анталии\n
	https://www.google.com/maps?q=36.893444,30.709591`
	ch.Kb = `{"inline_keyboard":[[{"text":"Просмотр расписания","callback_data":"Looking Schedule","url":""}],[{"text":"Регистрация на игру","callback_data":"Reg to games","url":""}],[{"text":"Наши фото и видео","callback_data":"Photo\u0026Video","url":""}],[{"text":"Настройки | Мои игры","callback_data":"My records","url":""}]]}`
	ch.Payment = "card"
	ch.Lp = 0
	ch.GameId = 0 //0 because when switching actions GameID is no longer updated
	ch.Seats = 2

}

func Dir(res *apptype.Response, i, j int) {
	ch := new(check)
	ch.Mes = res.Message
	ch.Kb = res.Keyboard
	ch.Payment = res.Payment
	ch.Act = res.Act
	ch.Lp = res.LaunchPoint
	ch.GameId = res.GameId
	ch.Seats = res.Seats
	ch.ParseMode = res.ParseMode
	ch.Res = res
	if j < 2 {
		if i == 0 {
			truefunc[i](res, ch)
		} else {
			truefunc[i-1](res, ch)
		}
	} else {
		truefunc[i](res, ch)
		if i == 4 {
			afterBestWishes()
		}
	}
}
