package registration

import "registrationtogames/fmtogram/formatter"

func PresentationScheduele(fm *formatter.Formatter) {
	fm.AssertChatId(456, true)
	fm.AssertString("Выберите интересующую вас игру", true)
	//fm.AssertInlineKeyboard([]int{1}, []string{"Зарегистрироваться"}, []string{"GoReg"}, []string{"cmd"}, true)
}

func ChooseGame(fm *formatter.Formatter) {
	fm.AssertChatId(456, true)
	fm.AssertString("Выберите или введите желаемое количество мест. Прошу обратить ваше внимание на то, что цена за одно место на эту игру составляет 100 USD", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Я буду один", "Нас будет двое", "Нас будет трое"}, []string{"1", "2", "3"}, []string{"cmd", "cmd", "cmd"}, true)
}

func ChooseSeats(fm *formatter.Formatter) {
	fm.AssertChatId(456, true)
	fm.AssertString("Выберите способ оплаты", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Онлайн оплата", "Наличкой администратору", "Главное Меню"}, []string{"card", "cash", "MainMenu"}, []string{"cmd", "cmd", "cmd"}, true)
}

func ChoosePayment(fm *formatter.Formatter) {
	fm.AssertChatId(456, true)
	fm.AssertString(`Вам нужно перевести 200 USD. Выше я выслал вам QR-code, по которому вы можете перейти и оплатить или же нажмите на кнопку ниже "Оплатить"`, true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Оплатить", "Дальше"},
		[]string{"https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3", "Next"}, []string{"url", "cmd"}, true)
	fm.AssertPhoto("qr.jpg", true)
}

func BestWishes(fm *formatter.Formatter) {
	fm.AssertChatId(456, true)
	fm.AssertString(`Прекрасно! Теперь вы зарегестрированны на игру\n
	1. Вид спорта: <b>Волейбол</b>\n
	2. Дата: <b>12.02.2025</b>\n
	3. Время: <b>12:00</b>\n
	4. Вы записали <b>2</b> персон на эту игру\n
	5. Оплата: <b>Карта</b>\n
	6. Ваша стоимость посещения: <b>200 USD</b>\n\n
	
	***Вы можете изменить некоторые данные Вашей записи\n
	или же удалить ее в Главном Меню нажав на "Посмотреть мои записи"***\n\n
	
	❤️❤️❤️Ждем вас по адресу:\nhttps://www.google.com/maps?q=36.893445,30.709591`, true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1}, []string{"Просмотр расписания", "Регистрация на игру", "Наши фото и видео", "Настройки | Мои игры"},
		[]string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)

}
