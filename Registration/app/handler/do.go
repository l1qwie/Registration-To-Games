package handler

import (
	"Registration/api/client"
	"Registration/api/producer"
	"Registration/app/dict"
	"Registration/apptype"
	"fmt"
	"strconv"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	START  int    = 0
	LEVEL1 int    = 1
	LEVEL2 int    = 2
	LEVEL3 int    = 3
	LEVEL4 int    = 4
	LEVEL5 int    = 5
	prev   string = "previous page"
	next   string = "next page"
)

type btn struct {
	crd   []int
	names []string
	data  []string
}

// Chechs the data can become integer or no
func intCheck(numS string) (bool, int) {
	producer.InterLogs("Start function Registration.intCheck()", fmt.Sprintf("numS (string): %s", numS))
	num, err := strconv.Atoi(numS)
	return err == nil, num
}

// The name of the function speaks for itself
// Returns date (string) format looks like DD-MM-YYYY
func fromIntToStrDate(numberDate int) (date string) {
	producer.InterLogs("Start function Registration.fromIntToStrDate()", fmt.Sprintf("numberDate (int): %d", numberDate))
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
	if monthString != "" && dayString != "" {
		date = fmt.Sprintf("%s-%s-%d", dayString, monthString, year)
	} else if monthString != "" && dayString == "" {
		date = fmt.Sprintf("%d-%s-%d", day, monthString, year)
	} else if monthString == "" && dayString != "" {
		date = fmt.Sprintf("%s-%d-%d", dayString, month, year)
	} else {
		date = fmt.Sprintf("%d-%d-%d", day, month, year)
	}
	return date
}

// The name of the function speaks for itself
// Returns time (string) format looks like HH:MM
func fromIntToStrTime(numberTime int) (time string) {
	producer.InterLogs("Start function Registration.fromIntToStrTime()", fmt.Sprintf("numberTime (int): %d", numberTime))
	var (
		hour, minute             int
		hourString, minuteString string
	)
	hour = numberTime / 100
	minute = (numberTime - (hour * 100))
	if hour < 10 {
		hourString = fmt.Sprintf("%d%d", 0, hour)
	}
	if minute < 10 {
		minuteString = fmt.Sprintf("%d%d", 0, minute)
	}

	if hourString != "" && minuteString != "" {
		time = fmt.Sprintf("%s:%s", hourString, minuteString)
	} else if hourString != "" && minuteString == "" {
		time = fmt.Sprintf("%s:%d", hourString, minute)
	} else if hourString == "" && minuteString != "" {
		time = fmt.Sprintf("%d:%s", hour, minuteString)
	} else if hourString == "" && minuteString == "" {
		time = fmt.Sprintf("%d:%d", hour, minute)
	}
	return time
}

// Makes previos and/or next page buttons
func (b *btn) prevNext(req *apptype.Request, length int) {
	producer.InterLogs("Start function Registration.prevNext()", fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, length (int): %d", req.Id, req, length))
	if (length > req.Limit) && (req.LaunchPoint == 0) && (length != 0) {
		b.crd = append(b.crd, 1)
		b.names = append(b.names, ">>")
		b.data = append(b.data, next)
	} else if (length > req.Limit) && (req.Limit <= (req.Limit + req.LaunchPoint)) && (length != 0) {
		b.crd = append(b.crd, 1)
		b.names = append(b.names, "<<")
		b.data = append(b.data, prev)
	} else if (length > req.Limit) && (length > (req.Limit + req.LaunchPoint)) && (length != 0) {
		b.crd = append(b.crd, 2)
		b.names = append(b.names, "<<")
		b.names = append(b.names, ">>")
		b.data = append(b.data, prev)
		b.data = append(b.data, next)
	}

}

// Make only one data for one button
// There is the MainMenu button
func (b *btn) makeMainButt(dict map[string]string) {
	producer.InterLogs("Start function Registration.makeMainButt()", "mo enter data")
	b.crd = append(b.crd, 1)
	b.names = append(b.names, dict["MainMenu"])
	b.data = append(b.data, "MainMenu")
}

// Makes the schedule with selected games
func (b *btn) selGames(req *apptype.Request, f func(error), dict map[string]string) {
	producer.InterLogs("Start function Registration.selGames()", fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, f (func(error)): %T", req.Id, req, f))
	sch := selectTheSchedule(req.Limit, req.LaunchPoint, req.Language, f)
	for i := 0; sch[i] != nil && i < len(sch); i++ {
		b.crd = append(b.crd, 1)
		b.names = append(b.names, fmt.Sprintf("%s %s %s %s %d", sch[i].Sport, sch[i].Date, sch[i].Time, dict["freeSpace"], sch[i].Seats))
		b.data = append(b.data, strconv.Itoa(sch[i].Id))
	}
}

// Main action about making schedule
func makeSch(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Registration.makeSch()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	b := new(btn)
	res.Level = 1
	b.selGames(req, fm.Error, dict.Dictionary[req.Language])
	b.prevNext(req, len(b.crd))
	b.makeMainButt(dict.Dictionary[req.Language])
	setKb(fm, b.crd, b.names, b.data)
	fm.WriteString(dict.Dictionary[req.Language]["ChooseAnyGame"])
}

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	producer.InterLogs("Start function Registration.setKb()",
		fmt.Sprintf("fm (*formatter.Formatter): %v, crd ([]int): %v, names ([]string): %v, data ([]string): %v", fm, crd, names, data))
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// GO to Main Menu
func goToMain(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, tresp string) {
	producer.InterLogs("Start function Registration.goToMain()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, tresp (string): %s", req.Id, req, res, fm, tresp))
	res.Level = 3
	res.Act = "divarication"
	d := dict.Dictionary[req.Language]
	names := []string{d["first"], d["second"], d["third"], d["fourth"]}
	data := []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}
	crd := []int{1, 1, 1, 1}
	setKb(fm, crd, names, data)
	fm.WriteString(tresp)
}

// Checks some information and if everything is allright
// Directs and compiles schedules
func presentationScheduele(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Registration.presentationScheduele()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	res.Level = 3
	if findGame(fm.Error) {
		makeSch(req, res, fm)
	} else {
		goToMain(req, res, fm, (fmt.Sprint(dict.Dictionary[req.Language]["NoGames"], dict.Dictionary[req.Language]["MainMenu"])))
	}
	fm.WriteChatId(req.Id)
}

// Makes some data to a question like "how many seats would you want?"
func makeDataForChSeats(res *apptype.Response, fm *formatter.Formatter, gameId int, dict map[string]string, sestat bool) {
	producer.InterLogs("Start function Registration.makeDataForChSeats()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, gameId (int): %d, seatat (bool): %v", res.ChatID, res, fm, gameId, sestat))
	var (
		crd         []int
		names, data []string
	)
	res.Level = 2
	res.GameId = gameId
	price, space, currency := selectThePrice(res.GameId, fm.Error)
	for i := 0; i < space && i < 3; i++ {
		crd = append(crd, 1)
		names = append(names, dict[strconv.Itoa(i+1)])
		data = append(data, strconv.Itoa(i+1))
	}
	setKb(fm, crd, names, data)
	if sestat {
		fm.WriteString(fmt.Sprintf(dict["ChooseSeats"], price, currency))
	} else {
		fm.WriteString(fmt.Sprintf(dict["NoMoreSeats"], dict["ChooseSeats"], price, currency))
	}
}

// Has to check information about launch poit
func checkLP(mes string) int {
	producer.InterLogs("Start function Registration.checkLP()",
		fmt.Sprintf("mes (string): %s", mes))
	lp := 0
	if mes == next {
		lp = 7
	} else if mes == prev {
		lp = -7
	}
	return lp
}

// There is a function which can work with gameId and offers to choose seats
// Directs the main logic of this task
func chooseGame(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, sstat bool) {
	producer.InterLogs("Start function Registration.chooseGame()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, seatat (bool): %v", req.Id, req, res, fm, sstat))
	det, gid := intCheck(req.Req)
	if det {
		if findThatGame(gid, fm.Error) {
			makeDataForChSeats(res, fm, gid, dict.Dictionary[req.Language], sstat)
		} else {
			makeSch(req, res, fm)
		}
	} else {
		res.LaunchPoint = checkLP(req.Req)
		makeSch(req, res, fm)
	}
	fm.WriteChatId(req.Id)
}

// Makes data to a question like "how will you pay?"
func makeDataForChPay(res *apptype.Response, fm *formatter.Formatter, seats int, dict map[string]string) {
	producer.InterLogs("Start function Registration.makeDataForChPay()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, seats (int): %d", res.ChatID, res, fm, seats))
	res.Seats = seats
	res.Level = 3
	setKb(fm, []int{1, 1, 1}, []string{dict["payByCard"], dict["payByCash"], dict["MainMenu"]}, []string{"card", "cash", "MainMenu"})
	fm.WriteString(dict["ChoosePaymethod"])
}

// There is a function which can work with gameId and num of seats and offers to choose paymethod
// Directs the main logic of this task
func chooseSeats(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Registration.chooseSeats()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	det, seats := intCheck(req.Req)
	if det {
		if howManyIsLeft(req.GameId, seats, fm.Error) {
			makeDataForChPay(res, fm, seats, dict.Dictionary[req.Language])
		} else {
			makeDataForChSeats(res, fm, req.GameId, dict.Dictionary[req.Language], false)
		}
	} else {
		makeDataForChSeats(res, fm, req.GameId, dict.Dictionary[req.Language], true)
	}
	fm.WriteChatId(req.Id)
}

// There is a function which can make data to send (only if req.Req was "card")
func makeDataToPayByCard(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	producer.InterLogs("Start function Registration.makeDataToPayByCard()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	res.Level = 4
	price, _, currency := selectThePrice(req.GameId, fm.Error)
	cost := price * req.Seats
	fm.SetIkbdDim([]int{1, 1, 1})
	fm.WriteInlineButtonUrl(dict["pay"], "https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3")
	fm.WriteInlineButtonCmd(dict["GoNext"], "Next")
	fm.WriteInlineButtonCmd(dict["MainMenu"], "MainMenu")
	fm.AddPhotoFromStorage("qr.jpg")
	fm.WriteString(fmt.Sprintf(dict["WaitForYourMoney"], cost, currency))
	fm.WriteDeleteMesId(req.ExMesId)
}

// There is a function which make data if we don't have more seats to the game
func seatsAreFull(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	producer.InterLogs("Start function Registration.seatsAreFull()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", res.ChatID, res, fm))
	res.Level = 5
	setKb(fm, []int{1, 1}, []string{dict["go-ahead"], dict["no-deleteAll"]}, []string{"Go-ahead", "MainMenu"})
	fm.WriteString(fmt.Sprintf(dict["SeatsAreFull"], dict["Review"]))
}

// There is a function which can work with gameId, num of seats, and payment
// Offers to pay, or directs to next function that wishs you good luck
// Directs the main logic of this task
func choosePayment(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Registration.choosePayment()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if (req.Req == "card") || (req.Req == "cash") {
		if howManyIsLeft(req.GameId, req.Seats, fm.Error) {
			res.Payment = req.Req
			if req.Req == "card" {
				makeDataToPayByCard(req, res, fm, dict.Dictionary[req.Language])
			} else {
				bestWishes(req, res, fm)
			}
		} else {
			seatsAreFull(res, fm, dict.Dictionary[req.Language])
		}
	} else {
		makeDataForChPay(res, fm, req.Seats, dict.Dictionary[req.Language])
	}
	fm.WriteChatId(req.Id)
}

// Makes body for wish
func wishUGoodLuck(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	producer.InterLogs("Start function Registration.wishUGoodLuck()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	res.Level = 3
	res.Act = "divarication"
	res.Status = true
	details := selectDetailOfGame(req.GameId, req.Language, fm.Error)
	cost := details.Price * req.Seats
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
	result := fmt.Sprintf(dict["RegistrationCompleted"], details.Sport, details.Date, details.Time, req.Seats, req.Payment, cost, details.Currency, details.Address, details.Lattitude, details.Longitude)
	fm.WriteString(result)
	fm.WriteParseMode(types.HTML)
}

// Prepares data to update function (the enter to gRPC)
func upd(req *apptype.Request, f func(error)) {
	producer.InterLogs("Start function Registration.upd()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, f (func(error)): %T", req.Id, req, f))
	u, err := fill(req.GameId, req.Id)
	u.Action = "change"
	u.ActionGWU = "new"
	err = client.Updates(u, err)
	if err != nil {
		f(err)
	}
}

// There is a function which can work with gameId, num of seats, and payment
// The main diffrent between this function and previois (choosePayment)
// Is this function wichs you good luck and directs to MainMenu
func bestWishes(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Registration.bestWishes()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Req == "cash" || req.Req == "Next" {
		if howManyIsLeft(req.GameId, req.Seats, fm.Error) {
			ok := completeRegistration(req.Id, req.GameId, req.Seats, req.Payment, fm.Error)
			if ok {
				wishUGoodLuck(req, res, fm, dict.Dictionary[req.Language])
				upd(req, fm.Error)
			} else {
				goToMain(req, res, fm, fmt.Sprint(dict.Dictionary[req.Language]["Oops!"], "\n\n", dict.Dictionary[req.Language]["MainMenu"]))
			}
			producer.ActLogs("The user has completed registration (to game) action", req.Id)
		} else {
			seatsAreFull(res, fm, dict.Dictionary[req.Language])
		}
	} else {
		makeDataToPayByCard(req, res, fm, dict.Dictionary[req.Language])
	}
	fm.WriteChatId(req.Id)

}

// Directioner
func dir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Registration.dir()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Level == START {
		if req.Status {
			producer.ActLogs("The user has started registration (to game) action", req.Id)
		}
		presentationScheduele(req, res, fm)
	} else if req.Level == LEVEL1 {
		chooseGame(req, res, fm, true)
	} else if req.Level == LEVEL2 {
		chooseSeats(req, res, fm)
	} else if req.Level == LEVEL3 {
		choosePayment(req, res, fm)
	} else if req.Level == LEVEL4 {
		bestWishes(req, res, fm)
	} else if req.Level == LEVEL5 {
		req.Req = strconv.Itoa(req.GameId)
		chooseGame(req, res, fm, true)
	}
}

// The main
// The head
// The directioner
func RegistrationAct(req *apptype.Request, res *apptype.Response) {
	producer.InterLogs("Start function Registration.RegistrationAct()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v", req.Id, req, res))
	fm := new(formatter.Formatter)
	apptype.Db = apptype.ConnectToDatabase()
	res.Level = req.Level
	res.LaunchPoint = req.LaunchPoint
	res.GameId = req.GameId
	res.Seats = req.Seats
	res.Payment = req.Payment
	res.Act = req.Act
	res.ChatID = req.Id
	dir(req, res, fm)
	fm.ReadyKB()
	res.Keyboard = fm.Message.ReplyMarkup
	res.Message = fm.Message.Text
	res.ParseMode = fm.Message.ParseMode
	if fm.Err != nil {
		res.Error = fmt.Sprint(fm.Err)
	}
}
