package handler

import (
	"Settings/app/dict"
	"Settings/apptype"
	"Settings/fmtogram/formatter"
	"fmt"
)

const (
	START  int = 0
	LEVEL1 int = 1
	LEVEL2 int = 2
	LEVEL3 int = 3
	LEVEL4 int = 4
	LEVEL5 int = 5
)

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// Creates the schedule and offers to change app's language or user's game
func schWithlang(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	//body
}

// Offers to change only app's language
func onlylang(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, beg string) {
	res.Level = 2
	res.IsChanged = "language"
	fm.WriteString(fmt.Sprintf("%s%s", beg, dict["ChangeLang"]))
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["en"], dict["ru"], dict["tur"], dict["MainMenu"]}, []string{"en", "ru", "tur", "MainMenu"})
}

func chooseOptions(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if findUserGame(req.Id, fm.Error) {
		schWithlang(res, fm, dict.Dictionary[req.Language])
	} else {
		onlylang(res, fm, dict.Dictionary[req.Language], dict.Dictionary[req.Language]["NoGames"])
	}
	fm.WriteChatId(req.Id)
}

func changeLanguge(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, lang string, id int) {
	res.Language = updtLanguage(id, lang, fm.Error)
	res.Level = 3
	res.Act = "divarication"
	fm.WriteString(dict["Lanchanged"])
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
}

func whatWay(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Req == "language" {
		onlylang(res, fm, dict.Dictionary[req.Language], "")
	} else if req.Req == "records" {
		//body
	} else {
		chooseOptions(req, res, fm)
	}
}

func dirOfChanges(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.IsChanged == "language" {
		if req.Req == "en" || req.Req == "ru" || req.Req == "tur" {
			changeLanguge(res, fm, dict.Dictionary[req.Req], req.Req, req.Id)
		} else {
			onlylang(res, fm, dict.Dictionary[req.Language], "")
		}
	} else if req.IsChanged == "records" {
		//
	}
	fm.WriteChatId(req.Id)
}

// Directioner
func dir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Level == START {
		chooseOptions(req, res, fm)
	} else if req.Level == LEVEL1 {
		//whatWay(req, res, fm)
	} else if req.Level == LEVEL2 {
		dirOfChanges(req, res, fm)
		//changeOrDel(req, res, fm)
	} else if req.Level == LEVEL3 {
		//dirForRec(req, res, fm)
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
