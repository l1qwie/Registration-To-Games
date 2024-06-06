package handler

import (
	"Game/apptype"
	"Game/fmtogram/formatter"
)

type sportStruct struct {
	sport string
}

type dateStruct struct {
	date int
}

func (s *sportStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	res.Level = 4
	res.Changeable = "sport"
	setKb(fm, []int{1, 1, 1}, []string{dict["volleyball"], dict["football"], dict["MainMenu"]}, []string{"volleyball", "football", "MainMenu"})
	fm.WriteString(dict["writesport"])
}

func (s *dateStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
}

func (s *sportStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {

}

func (s *dateStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {

}

func (s *sportStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {

}

func (s *dateStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {

}

func initMap() map[string]apptype.Change {
	variables := make(map[string]apptype.Change)
	variables["sport"] = new(sportStruct)
	variables["date"] = new(dateStruct)
	return variables
}

func changeCLient(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, key, val string) {
	variables := initMap()
	if value, ok := variables[key]; ok {
		if res.Level == LEVEL3 {
			value.Offer(res, fm, dict, val)
		} else if res.Level == LEVEL4 {
			value.Accept(res, fm, dict, val)
		} else if res.Level == LEVEL5 {
			value.Save(res, fm, dict, val)
		}
	}
}

func showOptions(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, gameid int) {
	res.Level = 3
	res.GameId = gameid
	setKb(fm, []int{1, 1, 1, 1, 1, 1, 1, 1},
		[]string{dict["sport"], dict["date"], dict["time"], dict["seats"], dict["price"], dict["currency"], dict["link"], dict["nameaddress"]},
		[]string{"sport", "date", "time", "seats", "price", "currency", "link", "address"})
	fm.WriteString(dict["chooseChangable"])
}
