package handler

import (
	"Settings/api/client"
	"Settings/api/producer"
	"Settings/app/dict"
	"Settings/apptype"
	"fmt"
	"strconv"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	START  int = 0
	LEVEL1 int = 1
	LEVEL2 int = 2
	LEVEL3 int = 3
	LEVEL4 int = 4
	LEVEL5 int = 5
)

// Makes from int date fromat: YYYYMMDD to string format: "DD-MM-YYYY"
func fromIntToStrDate(numberDate int) (date string) {
	var (
		year, month, day       int
		monthString, dayString string
	)
	producer.InterLogs("Start function Settings.fromIntToStrDate()", fmt.Sprintf("numberDate (int): %d", numberDate))
	year = numberDate / 10000
	month = (numberDate - (year * 10000)) / 100
	day = (numberDate - (year * 10000) - (month * 100))
	if month <= 10 {
		monthString = fmt.Sprintf("%d%d", 0, month)
	}
	if day <= 10 {
		dayString = fmt.Sprintf("%d%d", 0, day)
	}
	if monthString != "" && dayString != "" {
		date = fmt.Sprintf("%s-%s-%d", dayString, monthString, year)
	} else if monthString != "" && dayString == "" {
		date = fmt.Sprintf("%d-%s-%d", day, monthString, year)
	} else if monthString == "" && dayString != "" {
		date = fmt.Sprintf("%s-%d-%d", dayString, month, year)
	} else {
		date = fmt.Sprintf("%d-%d-%d", day, month, year)
	}
	return date
}

// Makes from int time fromat: HHMM to string format: "HH-MM"
func fromIntToStrTime(numberTime int) (time string) {
	var (
		hour, minute             int
		hourString, minuteString string
	)
	producer.InterLogs("Start function Settings.fromIntToStrTime()", fmt.Sprintf("numberTime (int): %d", numberTime))
	hour = numberTime / 100
	minute = (numberTime - (hour * 100))
	if hour < 10 {
		hourString = fmt.Sprintf("%d%d", 0, hour)
	}
	if minute < 10 {
		minuteString = fmt.Sprintf("%d%d", 0, minute)
	}
	if hourString != "" && minuteString != "" {
		time = fmt.Sprintf("%s:%s", hourString, minuteString)
	} else if hourString != "" && minuteString == "" {
		time = fmt.Sprintf("%s:%d", hourString, minute)
	} else if hourString == "" && minuteString != "" {
		time = fmt.Sprintf("%d:%s", hour, minuteString)
	} else if hourString == "" && minuteString == "" {
		time = fmt.Sprintf("%d:%d", hour, minute)
	}
	return time
}

// Checkes is there an int number or not
// If yes - returns true and the number
// Else - returns only false (and 0 in var num)
func intCheck(phrase string) (bool, int) {
	producer.InterLogs("Start function Settings.intCheck()",
		fmt.Sprintf("phrase (string): %s", phrase))
	num, err := strconv.Atoi(phrase)
	return err == nil, num
}

// Checkes is it + or - to launch point and then changed if it's true
func increaseLaunchPoit(phrase string) (lp int) {
	producer.InterLogs("Start function Settings.increaseLaunchPoit()",
		fmt.Sprintf("phrase (string): %s", phrase))
	if phrase == "next page" {
		lp = 7
	} else if phrase == "previous page" {
		lp = -7
	}
	return lp
}

// Function directioner to MainMenu
func goToMainMenu(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, mes string) {
	producer.InterLogs("Start function Settings.goToMainMenu()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, mes (string): %s", res.ChatID, res, fm, mes))
	res.Level = 3
	res.Act = "divarication"
	res.Status = true
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
	fm.WriteString(mes)
	producer.ActLogs("The user has completed settings action", res.ChatID)
}

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	producer.InterLogs("Start function Settings.setKb()",
		fmt.Sprintf("fm (*formatter.Formatter): %v, crd ([]int): %v, names ([]string): %v, data ([]string): %v", fm, crd, names, data))
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// Creates the schedule and offers to change app's language or user's game
func schWithlang(res *apptype.Response, req *apptype.Request, fm *formatter.Formatter, uc *UsuallConn, dict map[string]string) {
	producer.InterLogs("Start function Settings.schWithlang()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, req (*apptype.Request): %v, fm (*formatter.Formatter): %v", res.ChatID, res, req, fm))
	res.Level = 1
	mes, _ := createSchedule(req, uc, fm.Error, dict, res.LaunchPoint)
	setKb(fm, []int{1, 1, 1}, []string{dict["Changelang"], dict["ChangRec"], dict["MainMenu"]}, []string{"language", "records", "MainMenu"})
	fm.WriteString(fmt.Sprint(dict["YouHaveGames"], mes, dict["WhatUCanDo"]))
	fm.WriteParseMode(types.HTML)
}

// Offers to change only app's language
func onlylang(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, beg string) {
	producer.InterLogs("Start function Settings.onlylang()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, beg (string): %s", res.ChatID, res, fm, beg))
	res.Level = 2
	res.IsChanged = "language"
	fm.WriteString(fmt.Sprintf("%s%s", beg, dict["ChangeLang"]))
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["en"], dict["ru"], dict["tur"], dict["MainMenu"]}, []string{"en", "ru", "tur", "MainMenu"})
}

// Makes body response for an option to change or delete
func chooseOptions(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn) {
	producer.InterLogs("Start function Settings.chooseOptions()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Status {
		producer.ActLogs("The user has started settings action", req.Id)
		req.Status = false
	}
	if uc.findUserGame(req.Id, fm.Error) {
		schWithlang(res, req, fm, uc, dict.Dictionary[req.Language])
	} else {
		onlylang(res, fm, dict.Dictionary[req.Language], dict.Dictionary[req.Language]["NoGames"])
	}
}

// Create a schedule of user's games and send it
// Also makes []int with gameIds of these games
func createSchedule(req *apptype.Request, uc *UsuallConn, f func(error), dict map[string]string, lp int) (mes string, ids []int) {
	producer.InterLogs("Start function Settings.createSchedule()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, f (func(error)): %T, lp (int): %d", req.Id, req, f, lp))
	sch := uc.selectUserSchedule(req.Id, req.Limit, lp, f, dict)
	for i := 0; i < len(sch) && sch[i] != nil; i++ {
		ids = append(ids, sch[i].Id)
		mes = fmt.Sprint(mes, fmt.Sprintf(dict["UserSch"], i+1, dict[sch[i].Sport], sch[i].Date, sch[i].Time, sch[i].Seats, sch[i].Price, sch[i].Currency, sch[i].Payment, sch[i].StatusPayment))
		ids[i] = sch[i].Id
	}
	return mes, ids
}

// Make the body-response to the records
func recordsBody(res *apptype.Response, req *apptype.Request, fm *formatter.Formatter, uc *UsuallConn, dict map[string]string) {
	producer.InterLogs("Start function Settings.recordsBody()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, req (*apptype.Request): %v, fm (*formatter.Formatter): %v", req.Id, res, req, fm))
	res.Level = 2
	res.IsChanged = "records"
	mes, ids := createSchedule(req, uc, fm.Error, dict, res.LaunchPoint)
	names := make([]string, len(ids))
	data := make([]string, len(ids))
	crd := make([]int, len(ids))
	for i := 0; i < len(ids); i++ {
		names[i] = fmt.Sprint(i + res.LaunchPoint + 1)
		data[i] = fmt.Sprint(ids[i])
		crd[i] = 1
	}
	fm.WriteString(fmt.Sprint(dict["ChooseGame"], mes))
	fm.WriteParseMode(types.HTML)
	setKb(fm, crd, names, data)
}

// Does a few ckeck before ask the user
func whatWay(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn) {
	producer.InterLogs("Start function Settings.whatWay()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Req == "language" {
		onlylang(res, fm, dict.Dictionary[req.Language], "")
	} else if req.Req == "records" {
		recordsBody(res, req, fm, uc, dict.Dictionary[req.Language])
	} else {
		chooseOptions(req, res, fm, uc)
	}
}

// Changes the app's language and redirect to MainMenu
func changeLanguge(res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn, dict map[string]string, lang string, id int) {
	var (
		ok  bool
		mes string
	)
	producer.InterLogs("Start function Settings.changeLanguge()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, lang (string): %s, id (int): %d", id, res, fm, lang, id))
	res.Language, ok = uc.updtLanguage(id, lang, fm.Error)
	if ok {
		u := new(client.Upd)
		u.UserId = id
		u.Language = lang
		u.Customlang = true
		client.Updates(u, "user", nil)
		mes = dict["Lanchanged"]
	} else {
		mes = dict["Oops!"]
	}
	goToMainMenu(res, fm, dict, mes)
}

// Would you like change or delete?
func whWouldUCh(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, gid int) {
	producer.InterLogs("Start function Settings.whWouldUCh()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, gid (int): %d", res.ChatID, res, fm, gid))
	res.Level = 3
	res.GameId = gid
	fm.WriteString(dict["ChangeOrDel"])
	setKb(fm, []int{1, 1, 1}, []string{dict["Change"], dict["DelGame"], dict["MainMenu"]}, []string{"change", "del", "MainMenu"})
}

// Directioner of changes
func dirOfChanges(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn) {
	producer.InterLogs("Start function Settings.dirOfChanges()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.IsChanged == "language" {
		if (req.Req == "en") || (req.Req == "ru") || (req.Req == "tur") {
			changeLanguge(res, fm, uc, dict.Dictionary[req.Req], req.Req, req.Id)
		} else {
			onlylang(res, fm, dict.Dictionary[req.Language], "")
		}
	} else if req.IsChanged == "records" {
		det, num := intCheck(req.Req)
		if det {
			if uc.findThatUserGame(num, req.Id, fm.Error) {
				whWouldUCh(res, fm, dict.Dictionary[req.Language], num)
			} else {
				recordsBody(res, req, fm, uc, dict.Dictionary[req.Language])
			}
		} else {
			res.LaunchPoint = increaseLaunchPoit(req.Req)
			recordsBody(res, req, fm, uc, dict.Dictionary[req.Language])
		}
	}
}

// Starts the change part of the app
func change(res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn, dict map[string]string, id int) {
	var (
		crd         []int
		names, data []string
	)
	producer.InterLogs("Start function Settings.change()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, id (int): %d", id, res, fm, id))
	if uc.statPayment(res.GameId, id, fm.Error) {
		crd = []int{1, 1, 1}
		names = []string{dict["Payment"], dict["Seats"], dict["MainMenu"]}
		data = []string{"payment", "myseats", "MainMenu"}
	} else {
		crd = []int{1, 1}
		names = []string{dict["Seats"], dict["MainMenu"]}
		data = []string{"myseats", "MainMenu"}
	}
	res.Level = 4
	fm.WriteString(dict["WhatUWhantToCh"])
	setKb(fm, crd, names, data)
}

func upd(uc *UsuallConn, gameId int) {
	producer.InterLogs("Start function Settings.upd()",
		fmt.Sprintf("gameId (int): %d", gameId))
	u, err := uc.fill(gameId)
	u.Action = "change"
	client.Updates(u, "other", err)
}

// Deletes the game
func del(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn) {
	var mes string
	producer.InterLogs("Start function Settings.del()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	schs, uss, _ := uc.findSomeSeats(req.GameId, req.Id, 3, fm.Error) //jsut a random number becuse the number influences to the bool var and I don't need it here
	ok := uc.delTheGame(schs+uss, req.GameId, req.Id, fm.Error)
	if ok {
		upd(uc, req.GameId)
		mes = dict.Dictionary[req.Language]["GameDeleted"]
	} else {
		mes = dict.Dictionary[req.Language]["Oops!"]
	}
	goToMainMenu(res, fm, dict.Dictionary[req.Language], fmt.Sprint(mes, "\n\n", dict.Dictionary[req.Language]["MainMenu"]))
}

// Directioner of records
func dirForRec(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn) {
	producer.InterLogs("Start function Settings.dirForRec()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Req == "change" {
		change(res, fm, uc, dict.Dictionary[req.Language], req.Id)
	} else if req.Req == "del" {
		del(req, res, fm, uc)
	} else {
		whWouldUCh(res, fm, dict.Dictionary[req.Language], req.GameId)
	}
}

// Makes the body of a response and changes the paymethod
func changeToCard(res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn, dict map[string]string, pm string, id int) {
	producer.InterLogs("Start function Settings.changeToCard()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, pm (string): %s, id (int): %d", res.ChatID, res, fm, pm, id))
	res.Level = 5
	uc.updtPayment(res.GameId, id, pm, fm.Error)
	fm.SetIkbdDim([]int{1, 1})
	fm.WriteInlineButtonUrl(dict["pay"], "https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3")
	fm.WriteInlineButtonCmd(dict["MainMenu"], "MainMenu")
	fm.WriteString(dict["ThxForChange"])
}

// Dicrectioner to wich way would you change your paymethod
func chPayment(res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn, dict map[string]string, id int) {
	producer.InterLogs("Start function Settings.chPayment()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, id (int): %d", res.ChatID, res, fm, id))
	res.IsChanged = "payment"
	if uc.selPaymethod(res.GameId, id, fm.Error) {
		changeToCard(res, fm, uc, dict, "card", id)
	} else {
		uc.updtPayment(res.GameId, id, "cash", fm.Error)
		goToMainMenu(res, fm, dict, fmt.Sprint(dict["ThxForChange"], "\n\n", dict["MainMenu"]))
	}
}

// The body of response is created here
func chMySeats(res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn, dict map[string]string, id int) {
	producer.InterLogs("Start function Settings.chMySeats()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, id (int): %d", res.ChatID, res, fm, id))
	res.Level = 5
	res.IsChanged = "myseats"
	schs, uss, _ := uc.findSomeSeats(res.GameId, id, 3, fm.Error) //jsut a random number becuse the number influences to the bool var and I don't need it here
	s := 0
	if schs > 3 {
		s = 3
	} else {
		s = schs
	}
	crd := make([]int, s)
	names := make([]string, s)
	data := make([]string, s)
	for i := range crd {
		crd[i] = 1
		names[i] = fmt.Sprint(i + 1)
		data[i] = fmt.Sprint(i + 1)
	}
	setKb(fm, crd, names, data)
	fm.WriteString(fmt.Sprintf(dict["ChooseSeat"], schs, uss, uss, (schs + uss)))
}

// What will we change?
// This is the question that the func asks to user
func chengeable(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn) {
	producer.InterLogs("Start function Settings.chengeable()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Req == "payment" {
		chPayment(res, fm, uc, dict.Dictionary[req.Language], req.Id)
	} else if req.Req == "myseats" {
		chMySeats(res, fm, uc, dict.Dictionary[req.Language], req.Id)
	} else {
		change(res, fm, uc, dict.Dictionary[req.Language], req.Id)
	}
}

// Confirm the changes of seats
func confMySeats(res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn, dict map[string]string, phrase string, id int) {
	var mes string
	producer.InterLogs("Start function Settings.confMySeats()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v,  phrase (string): %s, id (int): %d", res.ChatID, res, fm, phrase, id))
	det, num := intCheck(phrase)
	if det {
		schs, uss, freeS := uc.findSomeSeats(res.GameId, id, num, fm.Error)
		if freeS {
			ok := uc.updtSeats(res.GameId, id, schs, uss, num, fm.Error)
			if ok {
				upd(uc, res.GameId)
				mes = dict["ThxForChange"]
			} else {
				mes = dict["Oops!"]
			}
			goToMainMenu(res, fm, dict, fmt.Sprint(mes, "\n\n", dict["MainMenu"]))
		}
	} else {
		chMySeats(res, fm, uc, dict, id)
	}
}

// Starts the act "Confirms the changes"
// Directs to confirm changes of seats
func confirm(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn) {
	producer.InterLogs("Start function Settings.confirm()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.IsChanged == "myseats" {
		confMySeats(res, fm, uc, dict.Dictionary[req.Language], req.Req, req.Id)
	} else {
		changeToCard(res, fm, uc, dict.Dictionary[req.Language], "card", req.Id)
	}

}

// Directioner
func dir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, uc *UsuallConn) {
	producer.InterLogs("Start function Settings.dir()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Level == START {
		chooseOptions(req, res, fm, uc)
	} else if req.Level == LEVEL1 {
		whatWay(req, res, fm, uc)
	} else if req.Level == LEVEL2 {
		dirOfChanges(req, res, fm, uc)
	} else if req.Level == LEVEL3 {
		dirForRec(req, res, fm, uc)
	} else if req.Level == LEVEL4 {
		chengeable(req, res, fm, uc)
	} else if req.Level == LEVEL5 {
		confirm(req, res, fm, uc)
	}
	fm.WriteChatId(req.Id)
}

// The main
// The head
// The directioner
func SettingsAct(req *apptype.Request, res *apptype.Response) {
	producer.InterLogs("Start function Settings.SettingsAct()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v", req.Id, req, res))
	uc := new(UsuallConn)
	uc.DB = apptype.ConnectToDatabase()
	fm := new(formatter.Formatter)
	res.Level = req.Level
	res.LaunchPoint = req.LaunchPoint
	res.Act = req.Act
	res.IsChanged = req.IsChanged
	res.Language = req.Language
	res.GameId = req.GameId
	res.ChatID = req.Id
	res.Status = req.Status
	dir(req, res, fm, uc)
	fm.ReadyKB()
	res.Keyboard = fm.Message.ReplyMarkup
	res.Message = fm.Message.Text
	res.ParseMode = fm.Message.ParseMode
	if fm.Err != nil {
		res.Error = fmt.Sprint(fm.Err)
	}
}
