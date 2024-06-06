package handler

import (
	"Game/apptype"
	"Game/fmtogram/formatter"
	"fmt"
	"log"
)

// startCreate starts the create action and prepares a message (text, keyboard) to an admin to choose a sport
func startCreate(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 2
	res.Direction = "create"
	fm.WriteString(dict["writesport"])
	setKb(fm, []int{1, 1, 1}, []string{dict["volleyball"], dict["football"], dict["MainMenu"]}, []string{"volleyball", "football", "MainMenu"})
}

// writeDate prepares a message (text, keybaord) to an admin to write a date of the new game
func writeDate(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, sport string) {
	res.Level = 3
	res.Sport = sport
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n\n\n", dict["writedate"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

// isItSport decides if in req.Req is a right request and if yes then it redirects to another function
func isItSport(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	if req.Req == "volleyball" || req.Req == "football" {
		writeDate(res, fm, dict, req.Req)
	} else {
		startCreate(res, fm, dict)
	}
}

// isItDate checks if it's a real date by calling other functions and then redirect to a correct way
func isItDate(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	date, success := checkDate(req.Req)
	if success {
		writeTime(res, fm, dict, date)
	} else {
		writeDate(res, fm, dict, req.Sport)
	}
}

// writeSeats prepares a message (text, keybaord) to an admin to write seats of the new game
func writeSeats(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, timeint int) {
	res.Level = 5
	res.Time = timeint
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)), "\n",
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)),
		"\n\n\n", dict["writeseats"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

// isItTime checks if it's a real time by calling other functions and then redirect to a correct way
func isItTime(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	timeint, success := checkTime(req.Req, fromIntToStrDate(req.Date))
	log.Print(timeint, success, "PAMAGITI")
	if success {
		writeSeats(res, fm, dict, timeint)
	} else {
		writeTime(res, fm, dict, req.Date)
	}
}

// writePrice prepares a message (text, keybaord) to an admin to write a price of the new game
func writePrice(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, seats int) {
	res.Level = 6
	res.Seats = seats
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)), "\n",
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)), "\n",
		fmt.Sprintf(dict["ShowSeats"], seats),
		"\n\n\n", dict["writeprice"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

// isItSeats checks if it's a real int value by calling other functions and then redirect to a correct way
func isItSeats(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	seats, success := intCheck(req.Req)
	if success {
		writePrice(res, fm, dict, seats)
	} else {
		writeSeats(res, fm, dict, req.Time)
	}
}

// writeCurrency prepares a message (text, keybaord) to an admin to write a currency of the new game
func writeCurrency(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, price int) {
	res.Level = 7
	res.Price = price
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)), "\n",
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)), "\n",
		fmt.Sprintf(dict["ShowSeats"], res.Seats),
		"\n\n\n", dict["writecurrency"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

// isItPrice checks if it's a real int value by calling other functions and then redirect to a correct way
func isItPrice(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	price, success := intCheck(req.Req)
	if success {
		writeCurrency(res, fm, dict, price)
	} else {
		writePrice(res, fm, dict, req.Seats)
	}
}

// writeLink prepares a message (text, keybaord) to an admin to write a link of the location of the new game
func writeLink(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, currency string) {
	res.Level = 8
	res.Currency = currency
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)), "\n",
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)), "\n",
		fmt.Sprintf(dict["ShowSeats"], res.Seats), "\n",
		fmt.Sprintf(dict["ShowTotalprice"], res.Price, currency),
		"\n\n\n", dict["writelink"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

// writeAddress prepares a message (text, keybaord) to an admin to write a name address of the new game
func writeAddress(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, link string) {
	res.Level = 9
	res.Link = link
	log.Print(fmt.Sprintf(dict["ShowLink"], link), dict["ShowLink"], link)
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)), "\n",
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)), "\n",
		fmt.Sprintf(dict["ShowSeats"], res.Seats), "\n",
		fmt.Sprintf(dict["ShowTotalprice"], res.Price, res.Currency), "\n",
		fmt.Sprintf(dict["ShowLink"], link),
		"\n\n\n", dict["writeaddress"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func checkAllGameInf(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, address string) {
	res.Level = 10
	res.Address = address
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)), "\n",
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)), "\n",
		fmt.Sprintf(dict["ShowSeats"], res.Seats), "\n",
		fmt.Sprintf(dict["ShowTotalprice"], res.Price, res.Currency), "\n",
		fmt.Sprintf(dict["ShowLink"], res.Link), "\n",
		fmt.Sprintf(dict["ShowAddress"], address),
		"\n\n\n", dict["clarification"]))
	setKb(fm, []int{1, 1}, []string{dict["save"], dict["MainMenu"]}, []string{"save", "MainMenu"})
}

func save(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	saveToDatabase(res, fm.Error)
	res.Act = "divarication"
	fm.WriteString(dict["gameWasSaved"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func saveTheGame(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	if req.Req == "save" {
		save(res, fm, dict)
	} else {
		checkAllGameInf(res, fm, dict, res.Address)
	}
}
