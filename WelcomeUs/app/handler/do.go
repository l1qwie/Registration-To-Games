package handler

import (
	client "Welcome/api/client"
	"Welcome/api/producer"
	"Welcome/app/dict"
	"Welcome/fmtogram/formatter"
	"Welcome/types"
	"fmt"
)

const (
	START  int = 0
	LEVEL1 int = 1
	LEVEL2 int = 2
)

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	producer.InterLogs("Start function WelcomeUs.setKb()",
		fmt.Sprintf("fm (*formatter.Formatter): %v, crd ([]int): %v, names ([]string): %v, data ([]string): %v", fm, crd, names, data))
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// Greet to user
// Says hello
func greetingsToUser(res *types.Response, fm *formatter.Formatter, dict map[string]string) {
	producer.InterLogs("Start function WelcomeUs.greetingsToUser()",
		fmt.Sprintf("UserId: %d, res (*types.Response): %v, fm (*formatter.Formatter): %v", res.ChatID, res, fm))
	res.Level = 1
	setKb(fm, []int{1}, []string{dict["reg"]}, []string{"GoReg"})
	fm.WriteString(dict["WelcomeToBot"])
}

// The body of showRules
// Prepares all data
func srulBody(res *types.Response, fm *formatter.Formatter, dict map[string]string) {
	producer.InterLogs("Start function WelcomeUs.srulBody()",
		fmt.Sprintf("UserId: %d, res (*types.Response): %v, fm (*formatter.Formatter): %v", res.ChatID, res, fm))
	res.Level = 2
	setKb(fm, []int{1}, []string{dict["allright"]}, []string{"GoNext"})
	fm.WriteString(dict["BotRules"])
}

// Sows the rules of the app
func showRules(req *types.Request, res *types.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function WelcomeUs.showRules()",
		fmt.Sprintf("UserId: %d, req (*types.Request): %v, res (*types.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Req == "GoReg" {
		srulBody(res, fm, dict.Dictionary[req.Language])
	} else {
		greetingsToUser(res, fm, dict.Dictionary[req.Language])
	}
}

// The body of welcomeToMainMenu
// Prepares all data
func wtmmBody(res *types.Response, fm *formatter.Formatter, dict map[string]string, id int, lang string) {
	producer.InterLogs("Start function WelcomeUs.wtmmBody()",
		fmt.Sprintf("UserId: %d, res (*types.Response): %v, fm (*formatter.Formatter): %v,id (int): %d, lang (string): %s", res.ChatID, res, fm, id, lang))
	res.Level = 3
	res.Act = "divarication"
	res.Status = true
	err := client.Update(id, lang)
	if err != nil {
		fm.Error(err)
	}
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
	fm.WriteString(dict["WelcomeToMainMenu"])
	producer.ActLogs("The user has completed registration", id)
}

// Redirect to Main Menu
// This is the first enter of Main Menu
func welcomeToMainMenu(req *types.Request, res *types.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function WelcomeUs.welcomeToMainMenu()",
		fmt.Sprintf("UserId: %d, req (*types.Request): %v, res (*types.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Req == "GoNext" {
		wtmmBody(res, fm, dict.Dictionary[req.Language], req.Id, req.Language)
	} else {
		srulBody(res, fm, dict.Dictionary[req.Language])
	}
}

// Directioner
func dir(req *types.Request, res *types.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function WelcomeUs.dir()",
		fmt.Sprintf("UserId: %d, req (*types.Request): %v, res (*types.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Level == START {
		if req.Status {
			producer.ActLogs("The user has started registration", res.ChatID)
			res.Status = false
		}
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
	producer.InterLogs("Start function WelcomeUs.WelcomeAct()", fmt.Sprintf("UserId: %d, req (*types.Request): %v, res (*types.Response): %v", req.Id, req, res))
	fm := new(formatter.Formatter)
	res.Level = req.Level
	res.Act = req.Act
	res.ChatID = req.Id
	res.Status = req.Status
	dir(req, res, fm)
	fm.ReadyKB()
	res.Keyboard = fm.Message.ReplyMarkup
	res.Message = fm.Message.Text
	if fm.Err != nil {
		res.Error = fmt.Sprint(fm.Err)
	}
}
