package handler

import (
	"Schedule/api/producer"
	"Schedule/app/dict"
	"Schedule/apptype"

	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

// make the schedule
func makeSchedule(client *redis.Client, fm *formatter.Formatter, dict map[string]string, l int) {
	producer.InterLogs("Start function Schedule.makeSchedule()",
		fmt.Sprintf("client (*redis.Client): %v, fm (*formatter.Formatter): %v, l (int): %d", client, fm, l))
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
	producer.InterLogs("Start function Schedule.goToMain()",
		fmt.Sprintf("res (*apptype.Response): %v, fm (*formatter.Formatter): %v, tresp (string): %s, lang (string): %s", res, fm, tresp, lang))
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
	producer.InterLogs("Start function Schedule.mainLogic()",
		fmt.Sprintf("req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, cl (*redis.Client): %v", req, res, fm, cl))
	l := findAnyGame(cl, fm.Error)
	if l > 0 {
		makeSchedule(cl, fm, dict.Dictionary[req.Language], l)
	} else {
		goToMain(res, fm, dict.Dictionary[req.Language]["NoGames"], req.Language)
	}
}

// Updates the schedule
func UpdateTheSchedule(g *apptype.Game) error {
	producer.InterLogs("Start function Schedule.UpdateTheSchedule()", fmt.Sprintf("g (*apptype.Game): %v", g))
	cl, err := addClient()
	if g.Action == "new" {
		err = newGame(cl, g)
	} else if g.Action == "del" {
		err = delGame(cl, g.Id)
	} else if g.Action == "change" {
		err = changeGame(cl, g)
	}
	return err
}

// Checks are there any games and then redirect to making function
func Schedule(req *apptype.Request, res *apptype.Response) *apptype.Response {
	producer.InterLogs("Start function Schedule.Schedule()", fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v", req.Id, req, res))
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
	res.ChatId = req.Id
	res.ParseMode = fm.Message.ParseMode
	if fm.Err != nil {
		res.Error = fmt.Sprint(fm.Err)
	}
	producer.ActLogs("User started watching in the schedule", req.Id)
	return res
}
