package handler

import (
	"Game/api/client"
	"Game/apptype"

	"github.com/l1qwie/Fmtogram/formatter"
)

func deleteTheGame(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, gameid int) {
	res.Level = 1
	res.GameId = gameid
	deleteGameDB(res.GameId, fm.Error)
	upd := new(client.Upd)
	//date, time, seats, price int, sport, currency, link, address string
	upd.Date, upd.Time, upd.Seats, upd.Price, upd.Sport, upd.Currency, upd.Link, upd.Address = selecAllGameInf(res.GameId, fm.Error)
	fm.WriteString(dict["gameWasDeleted"])
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["createGame"], dict["changeGame"], dict["deleteGame"], dict["MainMenu"]}, []string{"create", "change", "delete", "MainMenu"})
}
