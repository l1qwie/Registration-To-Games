package handler

import (
	"Game/apptype"
	"Game/fmtogram/formatter"
)

func deleteTheGame(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, gameid int) {
	res.Level = 1
	res.GameId = gameid
	deleteGameDB(res.GameId, fm.Error)
	fm.WriteString(dict["gameWasDeleted"])
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["createGame"], dict["changeGame"], dict["deleteGame"], dict["MainMenu"]}, []string{"create", "change", "delete", "MainMenu"})
}
