package functional

import (
	"Welcome/apptype"
	"fmt"
)

// [{"text":"","callback_data":"","url":""}]
// {"inline_keyboard":[[{"text":"","callback_data":"","url":""}],[{"text":"","callback_data":"","url":""}]]}

var ch = new(check)

type check struct {
	mes    string
	kb     string
	lvl    int
	act    string
	status bool
	res    *apptype.Response
}

func (ch *check) maintest() {
	if ch.res.ChatID != 1999 {
		panic(fmt.Sprintf(`ch.res.ChatID != 1999 because ch.res.ChatID == %d`, ch.res.ChatID))
	}
	if ch.res.Message != ch.mes {
		panic(fmt.Sprintf(`ch.res.Message != ch.mes because ch.res.Messag == %s and ch.mes == %s`, ch.res.Message, ch.mes))
	}
	if ch.res.Keyboard != ch.kb {
		panic(fmt.Sprintf(`ch.res.Keyboard != ch.kb beacuse ch.res.Keyboard == %s and ch.kb == %s`, ch.res.Keyboard, ch.kb))
	}
	if ch.res.Level != ch.lvl {
		panic(fmt.Sprintf(`ch.res.Level != ch.lvl because ch.res.Level == %d and ch.lvl == %d`, ch.res.Level, ch.lvl))
	}
	if ch.res.Act != ch.act {
		panic(fmt.Sprintf(`ch.res.Act != ch.act because ch.res.Act == %s and ch.act == %s`, ch.res.Act, ch.act))
	}
	if ch.res.Status != ch.status {
		panic(fmt.Sprintf(`ch.res.Status != ch.status because ch.res.Status == %v and ch.status == %v`, ch.res.Status, ch.status))
	}
}

func SayHello(res *apptype.Response) {
	ch.mes = "Добро пожаловать в нашего бота! Этот бот может пригодиться вам, если вы приобрели его первую часть и вам требуется административная часть. Если это ваш случай, то нажмите на кнопку снизу"
	ch.kb = `{"inline_keyboard":[[{"text":"Зарегистрироваться","callback_data":"GoReg","url":""}]]}`
	ch.lvl = 1
	ch.act = "registration"
	ch.status = false
	ch.res = res
	ch.maintest()
}

func SendTheRules(res *apptype.Response) {
	ch.mes = `Вот пара рекомендаций для того, как пользоваться этим ботом:
	
	1. <b>Игры</b>         
			-Тут вы сможете сделать все что угодно с вашими играми (создание, настройка, удаление и тд.)
	2. <b>Клиенты</b>         
			-Тут вы сможете сделать все со своими клиентами (создать, поменять, удалить)
	3. <b>Активность</b>
	        -Тут вы сможете сделать своеобразную игру для чата ваших клиентов 
			 (суть игры в том, что ваши клиенты должны будут наперегонки регистрироваться на игру). 
			 А так же тут вы сможете посмотреть куда и кто записался, а так же найти контакты нужного вам человека
	4. <b>Деньги</b>
	        -Здесь вы сможете посмотреть кто вам должен за игры, от кого ожидается оплата, а кто уже оплатил. 
			 Так же именно тут вы сможете увидеть небольшую статистику по всей финансовым оборотам в боте
	5. <b>Настройки</b>          
			-Тут вы сможете изменить настрйоки бота. Например язык`
	ch.kb = `{"inline_keyboard":[[{"text":"Начнем же!","callback_data":"Start","url":""}]]}`
	ch.lvl = 2
	ch.act = "registration"
	ch.status = false
	ch.res = res
	ch.maintest()
}

func SendMainMenu(res *apptype.Response) {
	ch.mes = "Главное Меню"
	ch.kb = `{"inline_keyboard":[[{"text":"Игры","callback_data":"Games","url":""}],[{"text":"Клиенты","callback_data":"Clients","url":""}],[{"text":"Активность","callback_data":"Activity","url":""}],[{"text":"Деньги","callback_data":"Finances","url":""}],[{"text":"Настройки","callback_data":"Settings","url":""}]]}`
	ch.lvl = 3
	ch.act = "divarication"
	ch.status = true
	ch.res = res
	ch.maintest()
}
