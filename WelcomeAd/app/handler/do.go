package handler

import (
	"Welcome/app/dict"
	"Welcome/apptype"
	"Welcome/fmtogram/formatter"
	"fmt"
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

// Greet to admin
// Says hello
func greetingsToAdmin(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 1
	setKb(fm, []int{1}, []string{dict["reg"]}, []string{"GoReg"})
	fm.WriteString(dict["SayHello"])
}

// Makes response-body with Rules of the Admin part
func makeRules(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 2
	setKb(fm, []int{1}, []string{dict["LetsStart"]}, []string{"Start"})
	fm.WriteString(dict["AdminRules"])
}

// Checks correct request and then redirect for making body or sending back
func showRules(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Req == "GoReg" {
		makeRules(res, fm, dict.Dictionary[req.Language])
	} else {
		greetingsToAdmin(res, fm, dict.Dictionary[req.Language])
	}
}

// Redirects to Main Menu
func goToMainMenu(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 3
	res.Act = "divarication"
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
	fm.WriteString(dict["MainMenu"])
}

// Checks correct request and then redirect for going to Main Menu or sending back
func welcomeToMainMenu(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Req == "Start" {
		goToMainMenu(res, fm, dict.Dictionary[req.Language])
	} else {
		makeRules(res, fm, dict.Dictionary[req.Language])
	}
}

// Directioner
func dir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Level == START {
		greetingsToAdmin(res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL1 {
		showRules(req, res, fm)
	} else if req.Level == LEVEL2 {
		welcomeToMainMenu(req, res, fm)
	}
}

// Start of the Welcome Act for Admin part.
// Only this function is imported
func WelcomeAct(req *apptype.Request, res *apptype.Response) {
	fm := new(formatter.Formatter)
	res.Level = req.Level
	res.Act = req.Act
	dir(req, res, fm)
	fm.ReadyKB()
	res.Keyboard = fm.Message.ReplyMarkup
	res.Message = fm.Message.Text
	res.ChatID = req.Id
	if fm.Err != nil {
		res.Error = fmt.Sprint(fm.Err)
	}
}
