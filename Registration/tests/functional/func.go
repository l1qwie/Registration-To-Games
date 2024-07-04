package functional

// [{"text":"","callback_data":"","url":""}]
// {"inline_keyboard":[[{"text":"","callback_data":"","url":""}],[{"text":"","callback_data":"","url":""}]]}

import (
	"Registration/apptype"
	"fmt"
	"log"
)

type check struct {
	mes       string
	kb        string
	level     int
	payment   string
	act       string
	lp        int
	gameId    int
	seats     int
	parseMode string
	status    bool
	res       *apptype.Response
}

var (
	truefunc = []func(*check){presentationScheduele, chooseGame, chooseSeats, choosePayment, bestWishes}
)

// Main tester
// Takes data from struct ckeck and ckecks it between Response and what I want
func (ch *check) maintest() {
	if ch.res.ChatID != 477 {
		panic((fmt.Sprintf(`ch.res.ChatID != 477 because %s`, fmt.Sprintf("ch.res.ChatID == %d", ch.res.ChatID))))
	}
	if ch.res.Message != ch.mes {
		panic(fmt.Sprintf(`ch.res.Message != "%s" because %s`, ch.mes, fmt.Sprintf(`ch.res.Message == "%s"`, ch.res.Message)))
	}
	if ch.res.Keyboard != ch.kb {
		panic(fmt.Sprintf(`ch.res.Keyboard != "%s" because %s`, ch.kb, fmt.Sprintf(`ch.res.Keyboard == "%s"`, ch.res.Keyboard)))
	}
	if ch.res.Level != ch.level {
		panic(fmt.Sprintf(`ch.res.Level != %d because %s`, ch.level, fmt.Sprintf("ch.res.level == %d", ch.res.Level)))
	}
	if ch.res.LaunchPoint != ch.lp {
		panic(fmt.Sprintf(`ch.res.LaunchPoint != %d because %s`, ch.lp, fmt.Sprintf("ch.res.LaunchPoint == %d", ch.res.LaunchPoint)))
	}
	if ch.res.GameId != ch.gameId {
		panic(fmt.Sprintf(`ch.res.GameId != %d because %s`, ch.gameId, fmt.Sprintf("ch.res.GameId == %d", ch.res.GameId)))
	}
	if ch.res.Seats != ch.seats {
		panic(fmt.Sprintf(`ch.res.Seats != %d because %s`, ch.seats, fmt.Sprintf("ch.res.Seats == %d", ch.res.Seats)))
	}
	if ch.res.Payment != ch.payment {
		panic(fmt.Sprintf(`ch.res.Payment != "%s" because %s`, ch.payment, fmt.Sprintf(`ch.res.Payment == "%s"`, ch.res.Payment)))
	}
	if ch.res.Act != ch.act {
		panic(fmt.Sprintf(`ch.res.Act != "%s" because %s`, ch.act, fmt.Sprintf(`ch.res.Act == "%s"`, ch.res.Act)))
	}
	if ch.res.ParseMode != ch.parseMode {
		panic(fmt.Sprintf(`ch.res.ParseMode != %s because %s`, ch.parseMode, fmt.Sprintf(`ch.res.ParseMode == "%s"`, ch.res.ParseMode)))
	}
	if ch.res.Status != ch.status {
		panic(fmt.Sprintf(`ch.res.Status != "%v" because %v`, ch.status, fmt.Sprintf(`ch.res.Status == "%v"`, ch.res.Status)))
	}
}

// Data which I wait after after the function is executed presentationScheduele
func presentationScheduele(ch *check) {
	ch.mes = "Выберите интересующую вас игру"
	ch.kb = `{"inline_keyboard":[[{"text":"Волейбол 12-02-2025 12:00 Свободный мест осталось: 55","callback_data":"2","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.level = 1
	ch.payment = ""
	ch.lp = 0
	ch.gameId = 0
	ch.seats = 0
	ch.act = "reg to games"
	ch.parseMode = ""
	ch.status = false
	ch.maintest()
}

// Data which I wait after after the function is executed chooseGame
func chooseGame(ch *check) {
	ch.mes = "Выберите или введите желаемое количество мест. Прошу обратить ваше внимание на то, что цена за одно место на эту игру составляет 100 USD"
	ch.kb = `{"inline_keyboard":[[{"text":"Я буду один","callback_data":"1","url":""}],[{"text":"Нас будет двое","callback_data":"2","url":""}],[{"text":"Нас будет трое","callback_data":"3","url":""}]]}`
	ch.level = 2
	ch.payment = ""
	ch.lp = 0
	ch.gameId = 2
	ch.seats = 0
	ch.act = "reg to games"
	ch.parseMode = ""
	ch.status = false
	ch.maintest()
}

// Data which I wait after after the function is executed chooseSeats
func chooseSeats(ch *check) {
	ch.mes = "Выберите способ оплаты"
	ch.kb = `{"inline_keyboard":[[{"text":"Онлайн оплата","callback_data":"card","url":""}],[{"text":"Наличкой администратору","callback_data":"cash","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.level = 3
	ch.payment = ""
	ch.lp = 0
	ch.gameId = 2
	ch.seats = 2
	ch.act = "reg to games"
	ch.parseMode = ""
	ch.status = false
	ch.maintest()
}

// Data which I wait after after the function is executed choosePayment
func choosePayment(ch *check) {
	ch.mes = `Вам нужно перевести 200 USD. Выше я выслал вам QR-code, по которому вы можете перейти и оплатить или же нажмите на кнопку ниже "Оплатить"`
	ch.kb = `{"inline_keyboard":[[{"text":"Оплатить","callback_data":"","url":"https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3"}],[{"text":"Дальше","callback_data":"Next","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.level = 4
	ch.payment = "card"
	ch.lp = 0
	ch.gameId = 2
	ch.seats = 2
	ch.act = "reg to games"
	ch.parseMode = ""
	ch.status = false
	ch.maintest()
}

// Data which I wait after after the function is executed bestWishes
func bestWishes(ch *check) {
	ch.mes = `Прекрасно! Теперь вы зарегестрированны на игру\n
	1. Вид спорта: <b>Волейбол</b>\n
	2. Дата: <b>12-02-2025</b>\n
	3. Время: <b>12:00</b>\n
	4. Вы записали <b>2</b> персон на эту игру\n
	5. Оплата: <b>by card</b>\n
	6. Ваша стоимость посещения: <b>200 USD</b>\n\n
	
	***Вы можете изменить некоторые данные Вашей записи\n
	или же удалить ее в Главном Меню нажав на <b>"Настройки | Мои игры"</b>***\n\n
	
	❤️❤️❤️Ждем вас по адресу: Кладбище в Анталии\n
	https://www.google.com/maps/place/31%C2%B051'47.5%22N+34%C2%B051'50.8%22E/@31.863181,34.8626321,17`
	ch.kb = `{"inline_keyboard":[[{"text":"Просмотр расписания","callback_data":"Looking Schedule","url":""}],[{"text":"Регистрация на игру","callback_data":"Reg to games","url":""}],[{"text":"Наши фото и видео","callback_data":"Photo\u0026Video","url":""}],[{"text":"Настройки | Мои игры","callback_data":"My records","url":""}]]}`
	ch.level = 3
	ch.payment = "card"
	ch.lp = 0
	ch.gameId = 0 //0 because when switching actions GameID is no longer updated
	ch.seats = 2
	ch.act = "divarication"
	ch.parseMode = "HTML"
	ch.status = true
	ch.maintest()
}

// Prints the logs
func logs(res *apptype.Response) {
	log.Printf("res.Message = %s", res.Message)
	log.Printf("res.Keyboard = %s", res.Keyboard)
	log.Printf("res.ChatId = %d", res.ChatID)
	log.Printf("res.level = %d", res.Level)
	log.Printf("res.LaunchPoint = %d", res.LaunchPoint)
	log.Printf("res.GameId = %d", res.GameId)
	log.Printf("res.Seats = %d", res.Seats)
	log.Printf("res.Payment = %s", res.Payment)
	log.Printf("res.Act = %s", res.Act)
	log.Printf("res.ParseMode = %s", res.ParseMode)
}

// Just directioner
// Only this function is imported
func Dir(res *apptype.Response, i, j int) {
	logs(res)
	ch := new(check)
	ch.res = res
	if j < 2 {
		if i == 0 {
			truefunc[i](ch)
		} else {
			truefunc[i-1](ch)
		}
	} else {
		truefunc[i](ch)
		if i == 4 {
			afterBestWishes()
		}
	}
}
