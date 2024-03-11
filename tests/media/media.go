package media

import "RegistrationToGames/fmtogram/formatter"

func ChooseDirection(fm *formatter.Formatter) {
	fm.AssertString("Вы хотите загрузить или посмотреть?", true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Посмотреть", "Загрузить"}, []string{"unload", "upload"}, []string{"cmd", "cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func NoChoiseOnlyUpload(fm *formatter.Formatter) {
	fm.AssertString("Игр с уже загруженными медиафайлами еще нет, но вы можете стать первым!\nВыберите интересующую вас игру", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Волейбол 12-02-2024 12:00"}, []string{"10"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func ChooseGame(fm *formatter.Formatter) {
	fm.AssertString("Выберите интересующую вас игру", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Волейбол 12-02-2024 12:00"}, []string{"10"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func ChooseMediaGameUnload(fm *formatter.Formatter) {
	fm.AssertString("Выберите интересующую вас игру", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Футбол 12-02-2024 12:00"}, []string{"10"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func ChooseMediaGameUpload(fm *formatter.Formatter) {
	fm.AssertString("Выберите интересующую вас игру", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Волейбол 12-02-2024 12:00"}, []string{"10"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func WaitingYourMedia(fm *formatter.Formatter) {
	fm.AssertString("Прекрасно! Вы выбрали игру. У нее осталось 20 свободных мест для медиафайлов (максимум 20). Присылайте ваших фотографий или видео, которые были сделаны на этой игре. За один раз можно прислать не более 10 фалов. Кончено же, желательно, чтобы это число не привышало свободных мест", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func Unload(fm *formatter.Formatter) {
	fm.AssertString("Вот все медиа по это игре", true)
	fm.AssertPhoto("!@#IOJSIOJE!@#**()!@#$*()SIOPE!@()#", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func Upload(fm *formatter.Formatter) {
	fm.AssertString("Все файлы успешно загружены", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
	fm.AssertDeleteMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func CheckUpload(fm *formatter.Formatter) {
	fm.AssertString("Все загружено! Спасибо за ваше участие!\nГлавное Меню", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1}, []string{"Просмотр расписания", "Регистрация на игру", "Наши фото и видео", "Настройки | Мои игры"},
		[]string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}
