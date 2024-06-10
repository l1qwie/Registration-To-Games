package handler

import (
	"Game/apptype"
	"Game/fmtogram/formatter"
	"fmt"
)

type sportStruct struct{}
type dateStruct struct{}
type timeStruct struct{}
type seatsStruct struct{}
type priceStruct struct{}
type currencyStruct struct{}
type linkStruct struct{}

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

func (t *timeStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 4
	res.Changeable = "time"
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
	fm.WriteString(dict["writetime"])
}

func (s *seatsStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 4
	res.Changeable = "seats"
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
	fm.WriteString(dict["writeseats"])
}

func (p *priceStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 4
	res.Changeable = "price"
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
	fm.WriteString(dict["writeprice"])
}

func (c *currencyStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 4
	res.Changeable = "currency"
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
	fm.WriteString(dict["writecurrency"])
}

func (l *linkStruct) Offer(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 4
	res.Changeable = "link"
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
	fm.WriteString(dict["writelink"])
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

func (t *timeStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	time, ok := checkTime(val, fromIntToStrDate(selectDate(res.GameId, fm.Error)))
	if ok {
		res.Level = 5
		res.Time = time
		setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
		date, _, seats, price, sport, currency, link, address := selecAllGameInf(res.GameId, fm.Error)
		fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[sport]), "\n",
			fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(date)), "\n",
			fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)), "\n",
			fmt.Sprintf(dict["ShowSeats"], seats), "\n",
			fmt.Sprintf(dict["ShowTotalprice"], price, currency), "\n",
			fmt.Sprintf(dict["ShowLink"], link), "\n",
			fmt.Sprintf(dict["ShowAddress"], address),
			"\n\n\n", dict["woulduliketosave"]))
	} else {
		t.Offer(res, fm, dict)
	}
}

func (s *seatsStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	seats, ok := intCheck(val)
	if ok {
		res.Level = 5
		res.Seats = seats
		setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
		date, time, _, price, sport, currency, link, address := selecAllGameInf(res.GameId, fm.Error)
		fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[sport]), "\n",
			fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(date)), "\n",
			fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(time)), "\n",
			fmt.Sprintf(dict["ShowSeats"], res.Seats), "\n",
			fmt.Sprintf(dict["ShowTotalprice"], price, currency), "\n",
			fmt.Sprintf(dict["ShowLink"], link), "\n",
			fmt.Sprintf(dict["ShowAddress"], address),
			"\n\n\n", dict["woulduliketosave"]))
	} else {
		s.Offer(res, fm, dict)
	}
}

func (p *priceStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	price, ok := intCheck(val)
	if ok {
		res.Level = 5
		res.Price = price
		setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
		date, time, seats, _, sport, currency, link, address := selecAllGameInf(res.GameId, fm.Error)
		fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[sport]), "\n",
			fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(date)), "\n",
			fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(time)), "\n",
			fmt.Sprintf(dict["ShowSeats"], seats), "\n",
			fmt.Sprintf(dict["ShowTotalprice"], res.Price, currency), "\n",
			fmt.Sprintf(dict["ShowLink"], link), "\n",
			fmt.Sprintf(dict["ShowAddress"], address),
			"\n\n\n", dict["woulduliketosave"]))
	} else {
		p.Offer(res, fm, dict)
	}
}

func (c *currencyStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	res.Level = 5
	res.Currency = val
	setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
	date, time, seats, price, sport, _, link, address := selecAllGameInf(res.GameId, fm.Error)
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(date)), "\n",
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(time)), "\n",
		fmt.Sprintf(dict["ShowSeats"], seats), "\n",
		fmt.Sprintf(dict["ShowTotalprice"], price, res.Currency), "\n",
		fmt.Sprintf(dict["ShowLink"], link), "\n",
		fmt.Sprintf(dict["ShowAddress"], address),
		"\n\n\n", dict["woulduliketosave"]))
}

func (l *linkStruct) Accept(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	res.Level = 5
	res.Link = val
	setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
	date, time, seats, price, sport, currency, _, address := selecAllGameInf(res.GameId, fm.Error)
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(date)), "\n",
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(time)), "\n",
		fmt.Sprintf(dict["ShowSeats"], seats), "\n",
		fmt.Sprintf(dict["ShowTotalprice"], price, currency), "\n",
		fmt.Sprintf(dict["ShowLink"], res.Link), "\n",
		fmt.Sprintf(dict["ShowAddress"], address),
		"\n\n\n", dict["woulduliketosave"]))
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

func (t *timeStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	if val == "save" {
		res.Level = 2
		updateDBInt("time", res.Time, res.GameId, fm.Error)
		setKb(fm, []int{1, 1}, []string{dict["change"], dict["MainMenu"]}, []string{"change", "MainMenu"})
		fm.WriteString(dict["gameWasSaved"])
	} else {
		t.Accept(res, fm, dict, fromIntToStrTime(res.Time))
	}
}

func (s *seatsStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	if val == "save" {
		res.Level = 2
		updateDBInt("seats", res.Seats, res.GameId, fm.Error)
		setKb(fm, []int{1, 1}, []string{dict["change"], dict["MainMenu"]}, []string{"change", "MainMenu"})
		fm.WriteString(dict["gameWasSaved"])
	} else {
		s.Accept(res, fm, dict, fmt.Sprint(res.Seats))
	}
}

func (p *priceStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	if val == "save" {
		res.Level = 2
		updateDBInt("price", res.Price, res.GameId, fm.Error)
		setKb(fm, []int{1, 1}, []string{dict["change"], dict["MainMenu"]}, []string{"change", "MainMenu"})
		fm.WriteString(dict["gameWasSaved"])
	} else {
		p.Accept(res, fm, dict, fmt.Sprint(res.Price))
	}
}

func (c *currencyStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	if val == "save" {
		res.Level = 2
		updateDBStr("currency", res.Currency, res.GameId, fm.Error)
		setKb(fm, []int{1, 1}, []string{dict["change"], dict["MainMenu"]}, []string{"change", "MainMenu"})
		fm.WriteString(dict["gameWasSaved"])
	} else {
		c.Accept(res, fm, dict, res.Currency)
	}
}

func (l *linkStruct) Save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, val string) {
	if val == "save" {
		res.Level = 2
		updateDBStr("link", res.Link, res.GameId, fm.Error)
		setKb(fm, []int{1, 1}, []string{dict["change"], dict["MainMenu"]}, []string{"change", "MainMenu"})
		fm.WriteString(dict["gameWasSaved"])
	} else {
		l.Accept(res, fm, dict, res.Link)
	}
}

func initMap() map[string]apptype.Change {
	variables := make(map[string]apptype.Change)
	variables["sport"] = new(sportStruct)
	variables["date"] = new(dateStruct)
	variables["time"] = new(timeStruct)
	variables["seats"] = new(seatsStruct)
	variables["price"] = new(priceStruct)
	variables["currency"] = new(currencyStruct)
	variables["link"] = new(linkStruct)
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
