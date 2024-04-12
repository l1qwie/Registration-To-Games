package handler

import (
	"Welcome/app/dict"
	"Welcome/fmtogram/formatter"
	"Welcome/types"
)

const (
	START  int = 0
	LEVEL1 int = 1
	LEVEL2 int = 2
)

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// Greet to user
// Says hello
func greetingsToUser(res *types.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 1
	setKb(fm, []int{1}, []string{dict["reg"]}, []string{"GoReg"})
	fm.WriteString(dict["WelcomeToBot"])
}

// The body of showRules
// Prepares all data
func srulBody(res *types.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 2
	setKb(fm, []int{1}, []string{dict["allright"]}, []string{"GoNext"})
	fm.WriteString(dict["BotRules"])
}

// Sows the rules of the app
func showRules(req *types.Request, res *types.Response, fm *formatter.Formatter) {
	if req.Req == "GoReg" {
		srulBody(res, fm, dict.Dictionary[req.Language])
	} else {
		greetingsToUser(res, fm, dict.Dictionary[req.Language])
	}
}

// The body of welcomeToMainMenu
// Prepares all data
func wtmmBody(res *types.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 3
	res.Act = "divarication"
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
	fm.WriteString(dict["WelcomeToMainMenu"])
}

// Redirect to Main Menu
// This is the first enter of Main Menu
func welcomeToMainMenu(req *types.Request, res *types.Response, fm *formatter.Formatter) {
	if req.Req == "GoNext" {
		wtmmBody(res, fm, dict.Dictionary[req.Language])
	} else {
		srulBody(res, fm, dict.Dictionary[req.Language])
	}
}

// Directioner
func dir(req *types.Request, res *types.Response, fm *formatter.Formatter) {
	if req.Level == START {
		greetingsToUser(res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL1 {
		showRules(req, res, fm)
	} else if req.Level == LEVEL2 {
		welcomeToMainMenu(req, res, fm)
	}
}

// The main
// The head
// The directioner
func WelcomeAct(req *types.Request, res *types.Response) {
	fm := new(formatter.Formatter)
	res.Level = req.Level
	res.Act = req.Act
	dir(req, res, fm)
	fm.ReadyKB()
	res.Keyboard = fm.Message.ReplyMarkup
	res.Message = fm.Message.Text
	res.ChatID = req.Id
}
