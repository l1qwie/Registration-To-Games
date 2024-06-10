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
	day = numberDate - (year * 10000) - (month * 100)
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

// showGames prepares a message (text and keyboard) to an admin to choose a game
func showGames(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, direct string) {
	res.Level = 2
	res.Direction = direct
	dates, times, gameIds := selectDateTime(fm.Error)
	crd := make([]int, len(dates)+1)
	names := make([]string, len(dates)+1)
	data := make([]string, len(dates)+1)
	for i, date := range dates {
		crd[i] = 1
		names[i] = fmt.Sprintf("%s %s", date, times[i])
		data[i] = fmt.Sprint(gameIds[i])
	}
	crd[len(crd)-1] = 1
	names[len(names)-1] = dict["MainMenu"]
	data[len(data)-1] = "MainMenu"
	setKb(fm, crd, names, data)
	fm.WriteString(dict["chooseGame"])
}

// showGamesOrCreate decides in which way an admin will go next
func showGamesOrCreate(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	if req.Req == "create" || req.Req == "change" || req.Req == "delete" {
		if req.Req != "create" {
			showGames(res, fm, dict, req.Req)
		} else {
			startCreate(res, fm, dict)
		}
	} else {
		chooseGameDir(req, res, fm, dict)
	}
}

// createOrElse decides what req.Direction is equal and then it redirects to another function
func createOrElse(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	num, ok := intCheck(req.Req)
	if req.Direction == "create" {
		isItSport(req, res, fm, dict)
	} else {
		if ok {
			if findGame(num, fm.Error) {
				if req.Direction == "change" {
					showOptions(res, fm, dict, num)
				} else {
					deleteTheGame(res, fm, dict, num)
				}
			} else {
				showGames(res, fm, dict, res.Direction)
			}
		} else {
			showGames(res, fm, dict, res.Direction)
		}
	}
}

// validDate check if the date is a valid date
func validDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// realDateTime is a universual function to find correct number to a date or a time
func realDateTime(input, pattern string) (int, int, int, bool) {
	var components []string
	match, _ := regexp.MatchString(pattern, input)
	if match {
		components = regexp.MustCompile(`\d+`).FindAllString(regexp.MustCompile(pattern).FindString(input), -1)
	}
	if len(components) == 3 {
		day, _ := strconv.Atoi(components[0])
		month, _ := strconv.Atoi(components[1])
		year, _ := strconv.Atoi(components[2])
		return day, month, year, match
	} else if len(components) == 2 {
		minutes, _ := strconv.Atoi(components[0])
		hours, _ := strconv.Atoi(components[1])
		return minutes, hours, 0, match
	}
	return 0, 0, 0, match
}

// checkDate checks the input data and tries to find a real date in it.
// Uses pattern to do it
func checkDate(input string) (date int, success bool) {
	day, month, year, _ := realDateTime(input, datepattern)
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
	fm.WriteString(fmt.Sprint(fmt.Sprintf(dict["ShowSport"], dict[res.Sport]), "\n",
		fmt.Sprintf(dict["ShowDate"], fromIntToStrDate(res.Date)),
		"\n\n\n", dict["writetime"]))
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

// createOrChange is a function just for knowing which way should be choosed changeClient() or isItDate()
func createOrChange(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	if req.Direction == "create" {
		isItDate(req, res, fm, dict)
	} else {
		changeClient(res, fm, dict, req.Req, "")
	}
}

// chClientOrTime is a function just for knowing which way should be choosed changeClient() or isItTime()
func chClientOrTime(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	if req.Direction == "create" {
		isItTime(req, res, fm, dict)
	} else {
		changeClient(res, fm, dict, req.Changeable, req.Req)
	}
}

// chClientOrSeats is a function just for knowing which way should be choosed changeClient() or isItSeats()
func chClientOrSeats(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	if req.Direction == "create" {
		isItSeats(req, res, fm, dict)
	} else {
		changeClient(res, fm, dict, req.Changeable, req.Req)
	}
}

// CheckTime checks the input data and tries to find a real time in it.
// Uses pattern to do it
func checkTime(input, dateInput string) (timeint int, success bool) {
	date, err := time.Parse("02-01-2006", dateInput)
	if err == nil {
		hour, minute, _, match := realDateTime(input, timepattern)
		if match {
			if hour >= 0 && hour < 24 && minute >= 0 && minute < 60 {
				inputDateTime := time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, date.Location())
				if inputDateTime.After(time.Now()) {
					timeint = (hour * 100) + minute
					success = true
				}
			}
		}
	}
	return timeint, success
}

// intCheck checks whether the phrase has some potential to become an int value
func intCheck(phrase string) (int, bool) {
	num, err := strconv.Atoi(phrase)
	return num, err == nil
}

func dir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	if req.Level == START {
		chooseGameDir(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL1 {
		showGamesOrCreate(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL2 {
		createOrElse(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL3 {
		createOrChange(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL4 {
		chClientOrTime(req, res, fm, dict.Dictionary[req.Language])
	} else if req.Level == LEVEL5 {
		chClientOrSeats(req, res, fm, dict.Dictionary[req.Language])
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
	res.Direction = req.Direction
	res.Changeable = req.Changeable
	res.GameId = req.GameId
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
	//apptype.Db = apptype.ConnectToDatabase()
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
