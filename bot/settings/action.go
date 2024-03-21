package settings

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/formatter"
	"fmt"
)

func createSchedule(user *bottypes.User) (string, []int) {
	var (
		sch  []*bottypes.Game
		dict map[string]string
		mes  string
		ids  []int
	)
	dict = dictionary.Dictionary[user.Language]
	sch = selectUserSchedule(user.Id, user.Media.Limit, user.LaunchPoint, user.Language)
	for i := 0; i < len(sch) && sch[i] != nil; i++ {
		ids = append(ids, sch[i].Id)
		mes = fmt.Sprint(mes, fmt.Sprintf(dict["UserSch"], i+1, dict[sch[i].Sport], sch[i].Date, sch[i].Time, sch[i].Seats, sch[i].Price, sch[i].Currency, sch[i].Payment, sch[i].StatusPayment))
		ids[i] = sch[i].Id
	}
	return mes, ids
}

func ChooseOptions(user *bottypes.User, fm *formatter.Formatter) {
	var (
		dict         map[string]string
		mes          string
		crd          []int
		names, datas []string
	)
	dict = dictionary.Dictionary[user.Language]
	if FindUserGames(user.Id) {
		mes, _ = createSchedule(user)
		crd = []int{1, 1, 1}
		names = []string{dict["Changelang"], dict["ChangRec"], dict["MainMenu"]} //[]string{"Изменить язык", "Изменить записи", "Главное Меню"}
		datas = []string{"language", "records", "MainMenu"}
		fm.WriteParseMode("HTML")
	} else {
		user.Request = "language"
		mes = dict["NoGamesChangeLang"]
		crd = []int{1, 1, 1, 1}
		names = []string{dict["en"], dict["ru"], dict["tur"], dict["MainMenu"]}
		datas = []string{"en", "ru", "tur", "MainMenu"}
	}
	user.Level = 1
	fm.WriteString(mes)
	forall.SetTheKeyboard(fm, crd, names, datas)
	fm.WriteChatId(user.Id)
}

func changeLanguge(user *bottypes.User, fm *formatter.Formatter) {
	if user.Request == "en" || user.Request == "ru" || user.Request == "tur" {
		user.Language = updateLanguage(user.Id, user.Request)
		user.Level = 3
		user.Act = "divarication"
		dict := dictionary.Dictionary[user.Language]
		forall.GoToMainMenu(user, fm, dict["Lanchanged"])
		fm.WriteChatId(user.Id)
	} else {
		ChooseOptions(user, fm)
	}
}

func changeRecords(user *bottypes.User, fm *formatter.Formatter) {
	user.Level = 2
	dict := dictionary.Dictionary[user.Language]
	mes, ids := createSchedule(user)
	names := make([]string, len(ids))
	datas := make([]string, len(ids))
	crd := make([]int, len(ids))
	for i := 0; i < len(ids); i++ {
		names[i] = fmt.Sprint(i + user.LaunchPoint + 1)
		datas[i] = fmt.Sprint(ids[i])
		crd[i] = 1
	}
	fm.WriteString(fmt.Sprint(dict["ChooseGame"], mes))
	forall.SetTheKeyboard(fm, crd, names, datas)
	fm.WriteChatId(user.Id)
}

func WhatWay(user *bottypes.User, fm *formatter.Formatter) {
	if user.Request == "language" {
		changeLanguge(user, fm)
	} else if user.Request == "records" {
		changeRecords(user, fm)
	} else {
		ChooseOptions(user, fm)
	}
}

func ChangeOrDel(user *bottypes.User, fm *formatter.Formatter) {
	dict := dictionary.Dictionary[user.Language]
	det, num := forall.IntCheck(user.Request)
	if det {
		if findUserGame(num, user.Id) {
			user.Level = 3
			user.UserRec.GameId = num
			fm.WriteString(dict["ChangeOrDel"])
			forall.SetTheKeyboard(fm, []int{1, 1, 1}, []string{dict["Change"], dict["DelGame"], dict["MainMenu"]}, []string{"change", "del", "MainMenu"})
			fm.WriteChatId(user.Id)
		} else {
			user.Request = "records"
			WhatWay(user, fm)
		}
	} else {
		user.LaunchPoint = forall.IncreaseLaunchPoit(user.Request)
		user.Request = "records"
		WhatWay(user, fm)
	}
}

func change(user *bottypes.User, fm *formatter.Formatter) {
	user.Level = 4
	dict := dictionary.Dictionary[user.Language]
	crd := make([]int, 3)
	names := make([]string, 3)
	datas := make([]string, 3)
	mes := ""
	if statPayment(user.UserRec.GameId, user.Id) {
		mes = dict["WhatUWhantToCh"]
		crd = []int{1, 1, 1}
		names = []string{dict["Payment"], dict["Seats"], dict["MainMenu"]}
		datas = []string{"payment", "myseats", "MainMenu"}
	} else {
		mes = dict["WhatUWhantToCh"]
		crd = []int{1, 1}
		names = []string{dict["Seats"], dict["MainMenu"]}
		datas = []string{"myseats", "MainMenu"}
	}
	fm.WriteString(mes)
	forall.SetTheKeyboard(fm, crd, names, datas)
	fm.WriteChatId(user.Id)
}

func del(user *bottypes.User, fm *formatter.Formatter) {
	schs, uss, _ := findSomeSeats(user.UserRec.GameId, user.Id, 3) //jsut a random number
	delTheGame(schs+uss, user.UserRec.GameId, user.Id)
	forall.GoToMainMenu(user, fm, dictionary.Dictionary[user.Language]["GameDeleted"])
}

func DirForRec(user *bottypes.User, fm *formatter.Formatter) {
	if user.Request == "change" {
		change(user, fm)
	} else if user.Request == "del" {
		del(user, fm)
	} else {
		user.Request = fmt.Sprint(user.UserRec.GameId)
		ChangeOrDel(user, fm)
	}
}

func chPayment(user *bottypes.User, fm *formatter.Formatter) {
	var (
		dict         map[string]string
		names, datas []string
	)
	user.Level = 5
	user.UserRec.Changeable = user.Request
	dict = dictionary.Dictionary[user.Language]
	if selPaymethod(user.UserRec.GameId, user.Id) {
		names = []string{dict["payByCard"], dict["MainMenu"]}
		datas = []string{"card", "MainMenu"}
	} else {
		names = []string{dict["payByCash"], dict["MainMenu"]}
		datas = []string{"cash", "MainMenu"}
	}
	forall.SetTheKeyboard(fm, []int{1, 1}, names, datas)
	fm.WriteString(dict["ChoosePaymethod"])
}

func chMySeats(user *bottypes.User, fm *formatter.Formatter) {
	user.Level = 5
	user.UserRec.Changeable = user.Request
	schs, uss, _ := findSomeSeats(user.UserRec.GameId, user.Id, 3) //jsut a random number
	s := 0
	if schs > 3 {
		s = 3
	} else {
		s = schs
	}
	forall.Seats(s, fm)
	fm.WriteString(fmt.Sprintf(dictionary.Dictionary[user.Language]["ChooseSeat"], schs, uss, schs+uss))
	fm.WriteChatId(user.Id)
}

func Chengeable(user *bottypes.User, fm *formatter.Formatter) {
	if user.Request == "payment" {
		chPayment(user, fm)
	} else if user.Request == "myseats" {
		chMySeats(user, fm)
	} else {
		change(user, fm)
	}
}

func confPayment(user *bottypes.User, fm *formatter.Formatter) {
	dict := dictionary.Dictionary[user.Language]
	if user.Request == "card" {
		updtPayment(user.UserRec.GameId, user.Id, user.Request)
		fm.SetIkbdDim([]int{1, 1})
		fm.WriteInlineButtonUrl(dict["pay"], "https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3")
		fm.WriteInlineButtonCmd(dict["MainMenu"], "MainMenu")
		fm.WriteString(dict["ThxForChange"])
		fm.WriteChatId(user.Id)
	} else if user.Request == "cash" {
		updtPayment(user.UserRec.GameId, user.Id, user.Request)
		forall.GoToMainMenu(user, fm, dict["ThxForChange"])
	} else {
		user.Request = user.UserRec.Changeable
		chPayment(user, fm)
	}
}

func confMySeats(user *bottypes.User, fm *formatter.Formatter) {
	det, num := forall.IntCheck(user.Request)
	if det {
		schs, uss, freeS := findSomeSeats(user.UserRec.GameId, user.Id, num)
		if freeS {
			updtSeats(user.UserRec.GameId, user.Id, schs, uss, num)
			forall.GoToMainMenu(user, fm, dictionary.Dictionary[user.Language]["ThxForChange"])
		}
	} else {
		user.Request = user.UserRec.Changeable
		chMySeats(user, fm)
	}
}

func Confirm(user *bottypes.User, fm *formatter.Formatter) {
	if user.UserRec.Changeable == "payment" {
		confPayment(user, fm)
	} else {
		confMySeats(user, fm)
	}
}
