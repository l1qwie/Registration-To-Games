package functional

import (
	"Welcome/apptype"
	"fmt"
	"log"
)

var onetime bool = true

type check struct {
	mes    string
	kb     string
	act    string
	level  int
	status bool
	res    *apptype.Response
}

var truefunc = []func(*check){greetingsToUser, showRules, welcomeToMainMenu}

// Main tester
// Takes data from struct ckeck and ckecks it between Response and what I want
func (ch *check) maintest() {
	if ch.res.ChatID != 456 {
		panic(`ch.res.ChatID != 456`)
	}
	if ch.res.Message != ch.mes {
		panic(fmt.Sprintf(`ch.res.Message != %s %s %s`, ch.mes, "because", fmt.Sprintf(`ch.res.Message == %s`, ch.res.Message)))
	}
	if ch.res.Keyboard != ch.kb {
		panic(fmt.Sprintf(`ch.res.Keyboard != %s %s %s`, ch.kb, "because", fmt.Sprintf(`ch.res.Keyboard == %s`, ch.res.Keyboard)))
	}
	if ch.res.Act != ch.act {
		panic(fmt.Sprintf(`ch.res.Act != %s %s %s`, ch.act, "because", fmt.Sprintf(`ch.res.Act == %s`, ch.res.Act)))
	}
	if ch.res.Level != ch.level {
		panic(fmt.Sprintf(`ch.res.Level != %d %s %s`, ch.level, "because", fmt.Sprintf(`ch.res.Level == %d`, ch.res.Level)))
	}
	if ch.res.Status != ch.status {
		panic(fmt.Sprintf(`ch.res.Status != %v %s %s`, ch.status, "because", fmt.Sprintf(`ch.res.Status == %v`, ch.res.Status)))
	}
}

// Data which I wait after after the function is executed greetingsToUser
func greetingsToUser(ch *check) {
	ch.mes = "Добро пожаловать в нашего бота! Этот бот предназначен для регистрации на спортивные игры в Стамбуле, но для начала вам нужно зарегистрироваться у нас!"
	ch.kb = `{"inline_keyboard":[[{"text":"Зарегистрироваться","callback_data":"GoReg","url":""}]]}`
	ch.level = 1
	ch.act = "registration"
	ch.status = false
	ch.maintest()
}

// Data which I wait after after the function is executed showRules
func showRules(ch *check) {
	ch.mes = `Добро Пожаловать в нашу дружную компанию! Далее я вам напишу некоторые правила, и рекомендации как пользоваться ботом:\n1. Просмотр расписания.\n     
	- Предельно просто. Вы сможете посмтореть полное расписание всех возможных игр.\n2. Регистрация на игру.\n      
	- Бот поможет вам записаться на интересующую вас игру.\n3. Посмотреть фотографии.\n      
	- Тут будут храниться все фотографии, видео, гифки и все остальное, что будут выкладывать другие пользователи с уже прошедших игр.\n4. Посмототреть мои записи.\n      
	- Нажав сюда, вы сможете посмотреть игры, которые вы ожидаете. Так же тут вы сможете отменить запись.\n\n\n
	P.S. Если что, то у бота есть волшебная команда /menu. Эта команда всегда сможет вас вернуть в главное меню. 
	Естесвенно, если вы пришлете ее, когда у вас есть какой то незаконченный процес в боте, то прогресс не сохраниться`
	ch.kb = `{"inline_keyboard":[[{"text":"Все понятно!","callback_data":"GoNext","url":""}]]}`
	ch.level = 2
	ch.act = "registration"
	ch.status = false
	ch.maintest()
}

// Data which I wait after after the function is executed welcomeToMainMenu
func welcomeToMainMenu(ch *check) {
	ch.mes = "Добро пожаловать в гланое меню"
	ch.kb = `{"inline_keyboard":[[{"text":"Просмотр расписания","callback_data":"Looking Schedule","url":""}],[{"text":"Регистрация на игру","callback_data":"Reg to games","url":""}],[{"text":"Наши фото и видео","callback_data":"Photo\u0026Video","url":""}],[{"text":"Настройки | Мои игры","callback_data":"My records","url":""}]]}`
	ch.level = 3
	ch.act = "divarication"
	ch.status = true
	ch.maintest()
}

// Prints the logs
func logs(res *apptype.Response) {
	log.Printf("res.Message = %s", res.Message)
	log.Printf("res.Keyboard = %s", res.Keyboard)
	log.Printf("res.ChatId = %d", res.ChatID)
	log.Printf("res.Act = %s", res.Act)
	log.Printf("res.Level = %d", res.Level)
	log.Printf("res.Status = %v", res.Status)
}

// Just directioner
// Only this function is imported
func Dir(i, j int, res *apptype.Response) {
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
	}
}
