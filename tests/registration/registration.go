package registration

import (
	"registrationtogames/fmtogram/formatter"
)

func PresentationScheduele(fm *formatter.Formatter) {
	//fmt.Println("ASIOKLDJOUIHJKSADHJKASD")
	fm.AssertChatId(477, true)
	fm.AssertString("Выберите интересующую вас игру", true)
	fm.AssertInlineKeyboard([]int{1, 2, 1}, []string{"Волейбол 12-02-2025 12:00 Свободный мест осталось: 55", "<<", ">>", "Главное Меню"}, []string{"2", "previous page", "next page", "MainMenu"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)
	fm.AssertEditMessageId(8883, true)
}

func ChooseGame(fm *formatter.Formatter) {
	fm.AssertChatId(477, true)
	fm.AssertString("Выберите или введите желаемое количество мест. Прошу обратить ваше внимание на то, что цена за одно место на эту игру составляет 100 USD", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Я буду один", "Нас будет двое", "Нас будет трое"}, []string{"1", "2", "3"}, []string{"cmd", "cmd", "cmd"}, true)
	fm.AssertEditMessageId(8883, true)
}

func ChooseSeats(fm *formatter.Formatter) {
	fm.AssertChatId(477, true)
	fm.AssertString("Выберите способ оплаты", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Онлайн оплата", "Наличкой администратору", "Главное Меню"}, []string{"card", "cash", "MainMenu"}, []string{"cmd", "cmd", "cmd"}, true)
	fm.AssertEditMessageId(8883, true)
}

func ChoosePayment(fm *formatter.Formatter) {
	fm.AssertChatId(477, true)
	fm.AssertString(`Вам нужно перевести 200 USD. Выше я выслал вам QR-code, по которому вы можете перейти и оплатить или же нажмите на кнопку ниже "Оплатить"`, true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Оплатить", "Дальше", "Главное Меню"},
		[]string{"https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3", "Next", "MainMenu"}, []string{"url", "cmd", "cmd"}, true)
	fm.AssertPhoto("qr.jpg", true)
	fm.AssertDeleteMessageId(8883, true)
}

func BestWishes(fm *formatter.Formatter) {
	fm.AssertChatId(477, true)
	fm.AssertString(`Прекрасно! Теперь вы зарегестрированны на игру\n
	1. Вид спорта: <b>Волейбол</b>\n
	2. Дата: <b>12-02-2025</b>\n
	3. Время: <b>12:00</b>\n
	4. Вы записали <b>2</b> персон на эту игру\n
	5. Оплата: <b>by card</b>\n
	6. Ваша стоимость посещения: <b>200 USD</b>\n\n
	
	***Вы можете изменить некоторые данные Вашей записи\n
	или же удалить ее в Главном Меню нажав на <b>"Настройки | Мои игры"</b>***\n\n
	
	❤️❤️❤️Ждем вас по адресу: Кладбище в Анталии\n
	https://www.google.com/maps?q=36.893444,30.709591`, true)
	fm.AssertParseMode("HTML", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1}, []string{"Просмотр расписания", "Регистрация на игру", "Наши фото и видео", "Настройки | Мои игры"},
		[]string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)
	fm.AssertDeleteMessageId(8883, true)
}
