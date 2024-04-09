package handler

import (
	"Schedule/app/dict"
	"Schedule/apptype"
	"Schedule/fmtogram/formatter"
	"Schedule/fmtogram/types"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// make the schedule
func makeSchedule(client *redis.Client, fm *formatter.Formatter, dict map[string]string, l int) {
	var mes string
	lsg := selecGames(client, fm.Error, l)
	for i := 0; i < len(lsg); i++ {
		mes += fmt.Sprintf(dict["Schedule"], i+1, dict[lsg[i].Sport], lsg[i].Date, lsg[i].Time, lsg[i].Seats, lsg[i].Price, lsg[i].Currency)
	}
	fm.WriteString(mes)
	fm.SetIkbdDim([]int{1})
	fm.WriteInlineButtonCmd(dict["MainMenu"], "MainMenu")
	fm.WriteParseMode(types.HTML)
}

// GO to Main Menu
func goToMain(res *apptype.Response, fm *formatter.Formatter, tresp, lang string) {
	res.Level = 3
	res.Act = "divarication"
	d := dict.Dictionary[lang]
	names := []string{d["first"], d["second"], d["third"], d["fourth"]}
	data := []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}
	crd := []int{1, 1, 1, 1}
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
	fm.WriteString(tresp)
}

// The name says itself
func mainLogic(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, cl *redis.Client) {
	l := findAnyGame(cl, fm.Error)
	if l > 0 {
		makeSchedule(cl, fm, dict.Dictionary[req.Language], l)
	} else {
		goToMain(res, fm, dict.Dictionary[req.Language]["NoGames"], req.Language)
	}
	fm.WriteChatId(req.Id)
}

// Checks are there any games and then redirect to making function
func Schedule(req *apptype.Request, res *apptype.Response) *apptype.Response {
	res.Act = req.Act
	fm := new(formatter.Formatter)
	cl, err := addClient()
	if err == nil {
		mainLogic(req, res, fm, cl)
	} else {
		fm.Error(err)
	}
	fm.ReadyKB()
	res.Message = fm.Message.Text
	res.Keyboard = fm.Message.ReplyMarkup
	res.ChatId = fm.Message.ChatID
	res.ParseMode = fm.Message.ParseMode
	return res
}
