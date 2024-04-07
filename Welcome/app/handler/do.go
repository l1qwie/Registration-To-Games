package handler

import (
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

func greetingsToUser(req *types.Request, fm *formatter.Formatter) {
	var (
		kbName, kbData []string
		coordinates    []int
	)
	req.Level = 1
	kbName = []string{dict.Dictionary[req.Language]["reg"]}
	kbData = []string{"GoReg"}
	coordinates = []int{1}

	fm.SetIkbdDim(coordinates)
	for i := 0; i < len(kbName); i++ {
		fm.WriteInlineButtonCmd(kbName[i], kbData[i])
	}
	fm.WriteString(dict.Dictionary[req.Language]["WelcomeToBot"])
	fm.WriteChatId(req.Id)
}

func showRules(req *types.Request, fm *formatter.Formatter) {
	var (
		kbName, kbData []string
		coordinates    []int
	)
	if req.Req == "GoReg" {
		req.Level = 2
		kbName = []string{dict.Dictionary[req.Language]["allright"]}
		kbData = []string{"GoNext"}
		coordinates = []int{1}

		fm.SetIkbdDim(coordinates)
		for i := 0; i < len(kbName); i++ {
			fm.WriteInlineButtonCmd(kbName[i], kbData[i])
		}
		fm.WriteString(dict.Dictionary[req.Language]["BotRules"])
		fm.WriteChatId(req.Id)
	} else {
		greetingsToUser(req, fm)
	}
}

func welcomeToMainMenu(req *types.Request, fm *formatter.Formatter) {
	var (
		kbName, kbData []string
		coordinates    []int
		d              map[string]string
	)
	if req.Req == "GoNext" {
		req.Level = 3
		req.Act = "divarication"
		d = dict.Dictionary[req.Language]
		kbName = []string{d["first"], d["second"], d["third"], d["fourth"]}
		kbData = []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}
		coordinates = []int{1, 1, 1, 1}
		fm.SetIkbdDim(coordinates)
		for i := 0; i < len(kbName); i++ {
			fm.WriteInlineButtonCmd(kbName[i], kbData[i])
		}
		fm.WriteString(d["WelcomeToMainMenu"])
		fm.WriteChatId(req.Id)
		fmt.Println(fm.Message.ReplyMarkup, "!#@$!")
	} else {
		req.Req = "GoReg"
		showRules(req, fm)
	}
}

func dir(req *types.Request, fm *formatter.Formatter) {
	if req.Level == START {
		greetingsToUser(req, fm)
	} else if req.Level == LEVEL1 {
		showRules(req, fm)
	} else if req.Level == LEVEL2 {
		welcomeToMainMenu(req, fm)
	}
}

func WelcomeAct(req *types.Request, resp *types.Response) {
	fm := new(formatter.Formatter)
	dir(req, fm)
	fm.ReadyKB()
	resp.Keyboard = fm.Message.ReplyMarkup
	resp.Message = fm.Message.Text
	resp.ChatID = fm.Message.ChatID
}
