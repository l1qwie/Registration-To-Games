package welcome

import (
	"RegistrationToGames/fmtogram/formatter"
)

func TestGreetingsToUser(fm *formatter.Formatter) {
	fm.AssertChatId(456, true)
	fm.AssertString("Добро пожаловать в нашего бота! Этот бот предназначен для регистрации на спортивные игры в Стамбуле, но для начала вам нужно зарегистрироваться у нас!", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Зарегистрироваться"}, []string{"GoReg"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(8888, true)
}

func TestShowRules(fm *formatter.Formatter) {
	fm.AssertChatId(456, true)
	fm.AssertString(`Добро Пожаловать в нашу дружную компанию! Далее я вам напишу некоторые правила, и рекомендации как пользоваться ботом:\n1. Просмотр расписания.\n     
	- Предельно просто. Вы сможете посмтореть полное расписание всех возможных игр.\n2. Регистрация на игру.\n      
	- Бот поможет вам записаться на интересующую вас игру.\n3. Посмотреть фотографии.\n      
	- Тут будут храниться все фотографии, видео, гифки и все остальное, что будут выкладывать другие пользователи с уже прошедших игр.\n4. Посмототреть мои записи.\n      
	- Нажав сюда, вы сможете посмотреть игры, которые вы ожидаете. Так же тут вы сможете отменить запись.\n\n\n
	P.S. Если что, то у бота есть волшебная команда /menu. Эта команда всегда сможет вас вернуть в главное меню. 
	Естесвенно, если вы пришлете ее, когда у вас есть какой то незаконченный процес в боте, то прогресс не сохраниться`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Все понятно!"}, []string{"GoNext"}, []string{"cmd"}, true)
	fm.AssertEditMessageId(8888, true)
}

func TestWelcomeToMainMenu(fm *formatter.Formatter) {
	fm.AssertChatId(456, true)
	fm.AssertString("Добро пожаловать в гланое меню", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1}, []string{"Просмотр расписания", "Регистрация на игру", "Наши фото и видео", "Настройки | Мои игры"},
		[]string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)
	fm.AssertEditMessageId(8888, true)

}
