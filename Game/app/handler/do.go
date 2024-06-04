package handler

import (
	"Game/app/dict"
	"Game/apptype"
	"Game/fmtogram/formatter"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	START       = 0
	LEVEL1      = 1
	LEVEL2      = 2
	LEVEL3      = 3
	LEVEL4      = 4
	LEVEL5      = 5
	LEVEL6      = 6
	LEVEL7      = 7
	LEVEL8      = 8
	LEVEL9      = 9
	LEVEL10     = 10
	datepattern = `\d{1,2}[^0-9]+\d{1,2}[^0-9]+\d{4}`
	timepattern = `\d{1,2}[^0-9]+\d{1,2}`
)

// fromIntToStrDate make A REAL date but in integer to a date in string in a format DDMMYYYY
func fromIntToStrDate(numberDate int) (date string) {
	var (
		year, month, day       int
		monthString, dayString string
	)
	year = numberDate / 10000
	month = (numberDate - (year * 10000)) / 100
	day = (numberDate - (year * 10000) - (month * 100))
	if month <= 10 {
		monthString = fmt.Sprintf("%d%d", 0, month)
	}
	if day <= 10 {
		dayString = fmt.Sprintf("%d%d", 0, day)
	}
	if (monthString != "") && (dayString != "") {
		date = fmt.Sprintf("%s-%s-%d", dayString, monthString, year)
	} else if (monthString != "") && (dayString == "") {
		date = fmt.Sprintf("%d-%s-%d", day, monthString, year)
	} else if (monthString == "") && (dayString != "") {
		date = fmt.Sprintf("%s-%d-%d", dayString, month, year)
	} else {
		date = fmt.Sprintf("%d-%d-%d", day, month, year)
	}
	return date
}

// fromIntToStrTime makes A REAL time from an int fromat: HHMM to string a format: "HH-MM"
func fromIntToStrTime(numberTime int) (time string) {
	var (
		hour, minute             int
		hourString, minuteString string
	)
	hour = numberTime / 100
	minute = numberTime - (hour * 100)
	if hour < 10 {
		hourString = fmt.Sprintf("%d%d", 0, hour)
	}
	if minute < 10 {
		minuteString = fmt.Sprintf("%d%d", 0, minute)
	}
	if (hourString != "") && (minuteString != "") {
		time = fmt.Sprintf("%s:%s", hourString, minuteString)
	} else if (hourString != "") && (minuteString == "") {
		time = fmt.Sprintf("%s:%d", hourString, minute)
	} else if (hourString == "") && (minuteString != "") {
		time = fmt.Sprintf("%d:%s", hour, minuteString)
	} else if (hourString == "") && (minuteString == "") {
		time = fmt.Sprintf("%d:%d", hour, minute)
	}
	return time
}

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// startCreate starts the create action and prepares a message (text, keyboard) to an admin to choose a sport
func startCreate(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	res.Level = 2
	fm.WriteString(dict["writesport"])
	setKb(fm, []int{1, 1, 1}, []string{dict["volleyball"], dict["football"], dict["MainMenu"]}, []string{"volleyball", "football", "MainMenu"})
}

// chooseGameDir prepares a message (text and keyboard) to an admin to choose Game direction
func chooseGameDir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	if isThereAnyGame(fm.Error) {
		res.Level = 1
		fm.WriteString(dict["gameDirections"])
		setKb(fm, []int{1, 1, 1, 1}, []string{dict["createGame"], dict["changeGame"], dict["deleteGame"], dict["MainMenu"]}, []string{"create", "change", "delete", "MainMenu"})
	} else {
		startCreate(res, fm, dict)
	}
}

// showGamesOrCreate decides in which way an admin will go next
func showGamesOrCreate(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	if req.Req != "create" {
		//showGames()
	} else {
		startCreate(res, fm, dict)
	}
}

// writeDate prepares a message (text, keybaord) to an admin to write a date of the new game
func writeDate(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, sport string) {
	res.Level = 3
	res.Sport = sport
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[sport]), "\n\n\n", dict["writedate"]))
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

// validDate check if the date is a valid date
func validDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// realDateTime is a universual function to find correct number to a date or a time
func realDateTime(input, pattern string) (int, int, int) {
	var components []string
	match, _ := regexp.MatchString(pattern, input)
	if match {
		components = regexp.MustCompile(`\d+`).FindAllString(regexp.MustCompile(pattern).FindString(input), -1)
	}
	if len(components) == 3 {
		day, _ := strconv.Atoi(components[0])
		month, _ := strconv.Atoi(components[1])
		year, _ := strconv.Atoi(components[2])
		return day, month, year
	} else if len(components) == 2 {
		minutes, _ := strconv.Atoi(components[0])
		hours, _ := strconv.Atoi(components[1])
		return minutes, hours, 0
	}
	return 0, 0, 0
}

// checkDate checks the input data and tries to find a real date in it.
// Uses pattern to do it
func checkDate(input string) (date int, success bool) {
	day, month, year := realDateTime(input, datepattern)
	if validDate(fmt.Sprintf("%04d-%02d-%02d", year, month, day)) {
		inputDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		if inputDate.After(time.Now()) || inputDate.Equal(time.Now()) {
			date = (day * 1) + (month * 100) + (year * 10000)
			success = true
		}
	}
	return date, success
}

// writeTime prepares a message (text, keybaord) to an admin to write a time of the new game
func writeTime(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, date int) {
	res.Level = 4
	res.Date = date
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]),
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)),
		"\n\n\n", dict["writetime"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
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

// checkTime checks the input data and tries to find a real time in it.
// Uses pattern to do it
func checkTime(input string) (timeint int, success bool) {
	hour, minute, _ := realDateTime(input, timepattern)
	if hour >= 0 && hour < 24 && minute >= 0 && minute < 60 {
		inputDateTime := time.Now().Add(time.Hour*time.Duration(hour) + time.Minute*time.Duration(minute))
		if inputDateTime.After(time.Now()) {
			timeint = (hour * 100) + minute
			success = true
		}
	}
	return timeint, success
}

// writeSeats prepares a message (text, keybaord) to an admin to write seats of the new game
func writeSeats(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, timeint int) {
	res.Level = 5
	res.Time = timeint
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]),
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)),
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)),
		"\n\n\n", dict["writeseats"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

// isItTime checks if it's a real time by calling other functions and then redirect to a correct way
func isItTime(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	timeint, success := checkTime(req.Req)
	if success {
		writeSeats(res, fm, dict, timeint)
	} else {
		writeTime(res, fm, dict, req.Date)
	}
}

// intCheck checks whether the phrase has some potential to become an int value
func intCheck(phrase string) (int, bool) {
	num, err := strconv.Atoi(phrase)
	return num, err == nil
}

// writePrice prepares a message (text, keybaord) to an admin to write a price of the new game
func writePrice(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, seats int) {
	res.Level = 6
	res.Seats = seats
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]),
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)),
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)),
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
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]),
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)),
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)),
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
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]),
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)),
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)),
		fmt.Sprintf(dict["ShowSeats"], res.Seats),
		fmt.Sprintf(dict["ShowTotalprice"], res.Price, currency),
		"\n\n\n", dict["writelink"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

// writeAddress prepares a message (text, keybaord) to an admin to write a name address of the new game
func writeAddress(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, link string) {
	res.Level = 9
	res.Link = link
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]),
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)),
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)),
		fmt.Sprintf(dict["ShowSeats"], res.Seats),
		fmt.Sprintf(dict["ShowTotalprice"], res.Price, res.Currency),
		fmt.Sprintf(dict["Showlink"], link),
		"\n\n\n", dict["writeaddress"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func checkAllGameInf(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, address string) {
	res.Level = 10
	res.Address = address
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]),
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)),
		fmt.Sprintf(dict["ShowTime"], fromIntToStrTime(res.Time)),
		fmt.Sprintf(dict["ShowSeats"], res.Seats),
		fmt.Sprintf(dict["ShowTotalprice"], res.Price, res.Currency),
		fmt.Sprintf(dict["Showlink"], res.Link),
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

func dir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Level == START {
		chooseGameDir(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL1 {
		showGamesOrCreate(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL2 {
		isItSport(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL3 {
		isItDate(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL4 {
		isItTime(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL5 {
		isItSeats(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL6 {
		isItPrice(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL7 {
		writeLink(res, fm, dict.Dictionary[req.Language], req.Req)
	} else if req.Level == LEVEL8 {
		writeAddress(res, fm, dict.Dictionary[req.Language], req.Req)
	} else if req.Level == LEVEL9 {
		checkAllGameInf(res, fm, dict.Dictionary[req.Language], req.Req)
	} else if req.Level == LEVEL10 {
		saveTheGame(req, res, fm, dict.Dictionary[req.Language])
	}
}

func fromReqToRes(req *apptype.Request, res *apptype.Response) {
	res.ChatID = req.Id
	res.Level = req.Level
	res.Act = req.Act
	res.LaunchPoint = req.LaunchPoint
	res.Sport = req.Sport
	res.Date = req.Date
	res.Time = req.Time
	res.Seats = req.Seats
	res.Price = req.Price
	res.Currency = req.Currency
	res.Link = req.Link
	res.Address = req.Address
	res.Status = req.Status
}

// GameAction is only one function which is imported from this package.
// It makes a few values like formatter.Formatter and it opens connection to database
func GameAction(req *apptype.Request, res *apptype.Response) {
	fm := new(formatter.Formatter)
	apptype.Db = apptype.ConnectToDatabase()
	fromReqToRes(req, res)
	dir(req, res, fm)
	fm.ReadyKB()
	res.Keyboard = fm.Message.ReplyMarkup
	res.Message = fm.Message.Text
	//res.ParseMode = fm.Message.ParseMode
	if fm.Err != nil {
		res.Error = fmt.Sprint(fm.Err)
	}
}
