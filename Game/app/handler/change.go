package handler

import (
	"Game/apptype"
	"Game/fmtogram/formatter"
	"fmt"
)

type sportStruct struct{}
type dateStruct struct{}

func (s *sportStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 4
	res.Changeable = "sport"
	setKb(fm, []int{1, 1, 1}, []string{dict["volleyball"], dict["football"], dict["MainMenu"]}, []string{"volleyball", "football", "MainMenu"})
	fm.WriteString(dict["writesport"])
}

func (d *dateStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 4
	res.Changeable = "date"
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
	fm.WriteString(dict["writedate"])
}

func (s *sportStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	if val == "volleyball" || val == "football" {
		res.Level = 5
		res.Sport = val
		setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
		date, time, seats, price, _, currency, link, address := selecAllGameInf(res.GameId, fm.Error)
		fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n",
			fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(date)), "\n",
			fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(time)), "\n",
			fmt.Sprintf(dict["ShowSeats"], seats), "\n",
			fmt.Sprintf(dict["ShowTotalprice"], price, currency), "\n",
			fmt.Sprintf(dict["ShowLink"], link), "\n",
			fmt.Sprintf(dict["ShowAddress"], address),
			"\n\n\n", dict["woulduliketosave"]))
	} else {
		s.Offer(res, fm, dict)
	}
}

func (d *dateStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	date, ok := checkDate(val)
	if ok {
		res.Level = 5
		res.Date = date
		setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
		_, time, seats, price, sport, currency, link, address := selecAllGameInf(res.GameId, fm.Error)
		fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[sport]), "\n",
			fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)), "\n",
			fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(time)), "\n",
			fmt.Sprintf(dict["ShowSeats"], seats), "\n",
			fmt.Sprintf(dict["ShowTotalprice"], price, currency), "\n",
			fmt.Sprintf(dict["ShowLink"], link), "\n",
			fmt.Sprintf(dict["ShowAddress"], address),
			"\n\n\n", dict["woulduliketosave"]))
	} else {
		d.Offer(res, fm, dict)
	}
}

func (s *sportStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	if val == "save" {
		res.Level = 2
		updateDBStr("sport", res.Sport, res.GameId, fm.Error)
		setKb(fm, []int{1, 1}, []string{dict["change"], dict["MainMenu"]}, []string{"change", "MainMenu"})
		fm.WriteString(dict["gameWasSaved"])
	} else {
		s.Accept(res, fm, dict, res.Sport)
	}
}

func (d *dateStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	if val == "save" {
		res.Level = 2
		updateDBInt("date", res.Date, res.GameId, fm.Error)
		setKb(fm, []int{1, 1}, []string{dict["change"], dict["MainMenu"]}, []string{"change", "MainMenu"})
		fm.WriteString(dict["gameWasSaved"])
	} else {
		d.Accept(res, fm, dict, fromIntToStrDate(res.Date))
	}
}

func initMap() map[string]apptype.Change {
	variables := make(map[string]apptype.Change)
	variables["sport"] = new(sportStruct)
	variables["date"] = new(dateStruct)
	return variables
}

func changeClient(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, key, val string) {
	variables := initMap()
	if value, ok := variables[key]; ok {
		if res.Level == LEVEL3 {
			value.Offer(res, fm, dict)
		} else if res.Level == LEVEL4 {
			value.Accept(res, fm, dict, val)
		} else if res.Level == LEVEL5 {
			value.Save(res, fm, dict, val)
		}
	} else {
		showOptions(res, fm, dict, res.GameId)
	}

}

func showOptions(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, gameid int) {
	res.Level = 3
	res.GameId = gameid
	setKb(fm, []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]string{dict["sport"], dict["date"], dict["time"], dict["seats"], dict["price"], dict["currency"], dict["link"], dict["nameaddress"], dict["MainMenu"]},
		[]string{"sport", "date", "time", "seats", "price", "currency", "link", "address", "MainMenu"})
	fm.WriteString(dict["chooseChangable"])
}
