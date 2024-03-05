package media

import "registrationtogames/fmtogram/formatter"

func ChooseDirection(fm *formatter.Formatter) {
	fm.AssertString("Что вы хотите сделать?", true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Загрузить", "Посмотреть"}, []string{"unload", "upload"}, []string{"cmd", "cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func ChooseMediaGameUnload(fm *formatter.Formatter) {
	fm.AssertString("Выберите игру для просмотра", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Волейбол 12-02-2024 12:00"}, []string{"10"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func ChooseMediaGameUpload(fm *formatter.Formatter) {
	fm.AssertString("Выберите игру для загрузки", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Футбол 12-02-2024 12:00"}, []string{"10"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func WaitingYourMedia(fm *formatter.Formatter) {
	fm.AssertString(`Прекрасно! Вы выбрали игру. У нее осталось 20 свободных мест для медиафайлов (максимум 20). Присылайте ваших 
	фотографий или видео, которые были сделаны на этой игре. За один раз можно прислать не более 10 фалов. Кончено же, желательно, 
	чтобы это число не привышало свободных мест)`, true)
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
	fm.AssertString("Получено файлов: %d...Обрабатываю...\n\nКогда число будет равно количеству файлов, которые вы загрузили нажмите на кнопку ниже", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Все загружено!"}, []string{"End"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}

func CheckUpload(fm *formatter.Formatter) {
	fm.AssertString("Все загружено! Спасибо за ваше участие!\nГлавное Меню", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1}, []string{"Просмотр расписания", "Регистрация на игру", "Наши фото и видео", "Настройки | Мои игры"},
		[]string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)
	fm.AssertEditMessageId(66666, true)
	fm.AssertChatId(499, true)
}
