package handler

import (
	"Settings/app/dict"
	"Settings/apptype"
	"Settings/fmtogram/formatter"
	"Settings/fmtogram/types"
	"fmt"
	"strconv"
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
	num, err := strconv.Atoi(phrase)
	return err == nil, num
}

// Checkes is it + or - to launch point and then changed if it's true
func increaseLaunchPoit(phrase string) (lp int) {
	if phrase == "next page" {
		lp = 7
	} else if phrase == "previous page" {
		lp = -7
	}
	return lp
}

// Function directioner to MainMenu
func goToMainMenu(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, mes string) {
	res.Level = 3
	res.Act = "divarication"
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
	fm.WriteString(mes)
}

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// Creates the schedule and offers to change app's language or user's game
func schWithlang(res *apptype.Response, req *apptype.Request, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 1
	mes, _ := createSchedule(req, fm.Error, dict, res.LaunchPoint)
	setKb(fm, []int{1, 1, 1}, []string{dict["Changelang"], dict["ChangRec"], dict["MainMenu"]}, []string{"language", "records", "MainMenu"})
	fm.WriteString(fmt.Sprint(dict["YouHaveGames"], mes, dict["WhatUCanDo"]))
	fm.WriteParseMode(types.HTML)
}

// Offers to change only app's language
func onlylang(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, beg string) {
	res.Level = 2
	res.IsChanged = "language"
	fm.WriteString(fmt.Sprintf("%s%s", beg, dict["ChangeLang"]))
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["en"], dict["ru"], dict["tur"], dict["MainMenu"]}, []string{"en", "ru", "tur", "MainMenu"})
}

// Makes body response for an option to change or delete
func chooseOptions(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if findUserGame(req.Id, fm.Error) {
		schWithlang(res, req, fm, dict.Dictionary[req.Language])
	} else {
		onlylang(res, fm, dict.Dictionary[req.Language], dict.Dictionary[req.Language]["NoGames"])
	}
	fm.WriteChatId(req.Id)
}

// Create a schedule of user's games and send it
// Also makes []int with gameIds of these games
func createSchedule(req *apptype.Request, f func(error), dict map[string]string, lp int) (mes string, ids []int) {
	sch := selectUserSchedule(req.Id, req.Limit, lp, f, dict)
	for i := 0; i < len(sch) && sch[i] != nil; i++ {
		ids = append(ids, sch[i].Id)
		mes = fmt.Sprint(mes, fmt.Sprintf(dict["UserSch"], i+1, dict[sch[i].Sport], sch[i].Date, sch[i].Time, sch[i].Seats, sch[i].Price, sch[i].Currency, sch[i].Payment, sch[i].StatusPayment))
		ids[i] = sch[i].Id
	}
	return mes, ids
}

// Make the body-response to the records
func recordsBody(res *apptype.Response, req *apptype.Request, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 2
	res.IsChanged = "records"
	mes, ids := createSchedule(req, fm.Error, dict, res.LaunchPoint)
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
func whatWay(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Req == "language" {
		onlylang(res, fm, dict.Dictionary[req.Language], "")
	} else if req.Req == "records" {
		recordsBody(res, req, fm, dict.Dictionary[req.Language])
	} else {
		chooseOptions(req, res, fm)
	}
	fm.WriteChatId(req.Id)
}

// Changes the app's language and redirect to MainMenu
func changeLanguge(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, lang string, id int) {
	res.Language = updtLanguage(id, lang, fm.Error)
	//fm.WriteString(dict["Lanchanged"])
	goToMainMenu(res, fm, dict, dict["Lanchanged"])
}

// Would you like change or delete?
func whWouldUCh(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, gid int) {
	res.Level = 3
	res.GameId = gid
	fm.WriteString(dict["ChangeOrDel"])
	setKb(fm, []int{1, 1, 1}, []string{dict["Change"], dict["DelGame"], dict["MainMenu"]}, []string{"change", "del", "MainMenu"})
}

// Directioner of changes
func dirOfChanges(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.IsChanged == "language" {
		if (req.Req == "en") || (req.Req == "ru") || (req.Req == "tur") {
			changeLanguge(res, fm, dict.Dictionary[req.Req], req.Req, req.Id)
		} else {
			onlylang(res, fm, dict.Dictionary[req.Language], "")
		}
	} else if req.IsChanged == "records" {
		det, num := intCheck(req.Req)
		if det {
			if findThatUserGame(num, req.Id, fm.Error) {
				whWouldUCh(res, fm, dict.Dictionary[req.Language], num)
			} else {
				recordsBody(res, req, fm, dict.Dictionary[req.Language])
			}
		} else {
			res.LaunchPoint = increaseLaunchPoit(req.Req)
			recordsBody(res, req, fm, dict.Dictionary[req.Language])
		}
	}
	fm.WriteChatId(req.Id)
}

// Starts the change part of the app
func change() {
	//Body
}

// Deletes the game
func del(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	schs, uss, _ := findSomeSeats(req.GameId, req.Id, 3, fm.Error) //jsut a random number becuse the number influences to the bool var and I don't need it here
	delTheGame(schs+uss, req.GameId, req.Id, fm.Error)
	goToMainMenu(res, fm, dict.Dictionary[req.Language], fmt.Sprint(dict.Dictionary[req.Language]["GameDeleted"], dict.Dictionary[req.Language]["MainMenu"]))
}

// Directioner of records
func dirForRec(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Req == "change" {
		change()
	} else if req.Req == "del" {
		del(req, res, fm)
	} else {
		whWouldUCh(res, fm, dict.Dictionary[req.Language], req.GameId)
	}
	fm.WriteChatId(req.Id)
}

// Directioner
func dir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Level == START {
		chooseOptions(req, res, fm)
	} else if req.Level == LEVEL1 {
		whatWay(req, res, fm)
	} else if req.Level == LEVEL2 {
		dirOfChanges(req, res, fm)
	} else if req.Level == LEVEL3 {
		dirForRec(req, res, fm)
	} else if req.Level == LEVEL4 {
		//chengeable(req, res, fm)
	} else if req.Level == LEVEL5 {
		//confirm(req, res, fm)
	}
}

// The main
// The head
// The directioner
func SettingsAct(req *apptype.Request, res *apptype.Response) {
	fm := new(formatter.Formatter)
	res.Level = req.Level
	res.LaunchPoint = req.LaunchPoint
	res.Act = req.Act
	res.IsChanged = req.IsChanged
	res.Language = req.Language
	res.GameId = req.GameId
	dir(req, res, fm)
	fm.ReadyKB()
	res.Keyboard = fm.Message.ReplyMarkup
	res.Message = fm.Message.Text
	res.ChatID = fm.Message.ChatID
	res.ParseMode = fm.Message.ParseMode
}
