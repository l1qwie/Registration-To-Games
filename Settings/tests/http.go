package tests

import (
	"Settings/apptype"
	"Settings/tests/functional"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Makes the requst and take a response
func callreg(body []byte) *apptype.Response {
	result := new(apptype.Response)
	resp, err := http.Post("http://localhost:8089/Settings", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
	} else {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&result)
	}
	return result
}

// Initialization of logs
func initlogs(file string) *os.File {
	os.Remove(file)
	lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(lf)
	return lf
}

// Tests Change the Language functional
func language() {
	defer DeleteUser()
	CreateUser()
	//f := initlogs("applang.log")
	ts := new(TestStuct)
	ts.Round = 2
	ts.Name = "ChangeLanguageTest"
	ts.FuncReq = []func() *apptype.Request{chOpt, chLang}
	ts.FuncRes = []func(*apptype.Response){functional.ChOptOnlyLang, functional.ChooseLang}
	ts.FuncTrsh = []func() *apptype.Request{commontrash, commontrash2, langtr2, langtr3}
	ts.UpdtLevel = []int{0, 2}
	onetime = true
	ts.DoTest()
	onetime = false
	//defer f.Close()
}

// Tests Delete a Game functional
func delGame() {
	defer DeleteUser()
	defer DeleteUserSchedule()
	defer DeleteSchedule()
	CreateScheduleForUser()
	CreateUserScehdule()
	CreateUser()
	//f := initlogs("appdelGame.log")
	ts := new(TestStuct)
	ts.Round = 4
	ts.Name = "DeleteGameTest"
	ts.FuncReq = []func() *apptype.Request{chOpt, chRec, delGameId, del}
	ts.FuncRes = []func(*apptype.Response){functional.ChTwoOpt, functional.ChGame, functional.Del, functional.DelGame}
	ts.FuncTrsh = []func() *apptype.Request{commontrash, commontrash2, dGametr2, dGametr3, dGametr4, dGametr5, dGametr6, dGametr7}
	ts.UpdtLevel = []int{0, 1, 2, 3}
	ts.DoTest()
	//defer f.Close()
}

// Tests Change the Seats functional
func chSeats() {
	defer DeleteUser()
	defer DeleteUserSchedule()
	defer DeleteSchedule()
	CreateScheduleForUser()
	CreateUserScehdule()
	CreateUser()
	//f := initlogs("appchSeats.log")
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeSeatsTest"
	ts.FuncReq = []func() *apptype.Request{chOpt, chRec, delGameIdCh, change, seats, numOfSeats}
	ts.FuncRes = []func(*apptype.Response){functional.ChTwoOpt, functional.ChGame, functional.Ch, functional.ChThwWay, functional.WrtSeats, functional.ChangeSeats}
	ts.FuncTrsh = []func() *apptype.Request{commontrash, commontrash2, chSeatstr2, chSeatstr3, chSeatstr4, chSeatstr5, chSeatstr6, chSeatstr7, chSeatstr8, chSeatstr9, chSeatstr10, chSeatstr11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
	//defer f.Close()
}

// Tests Change the Paymethod functional
func chPayment() {
	defer DeleteUser()
	defer DeleteUserSchedule()
	defer DeleteSchedule()
	CreateScheduleForUser()
	CreateUserScehdule()
	CreateUser()
	//f := initlogs("appchPay.log")
	ts := new(TestStuct)
	ts.Round = 5
	ts.Name = "ChangePaymentTest"
	ts.FuncReq = []func() *apptype.Request{chOpt, chRec, delGameIdCh, change, payment}
	ts.FuncRes = []func(*apptype.Response){functional.ChTwoOpt, functional.ChGame, functional.Ch, functional.ChThwWay, functional.PaybyCash}
	ts.FuncTrsh = []func() *apptype.Request{commontrash, commontrash2, chSeatstr2, chSeatstr3, chSeatstr4, chSeatstr5, chSeatstr6, chSeatstr7, chPay1, chPay}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4}
	ts.DoTest()
	//defer f.Close()
}

// The list of testing functions
func testList() {
	language()
	delGame()
	chSeats()
	chPayment()
}

// The head of the directioner
// Only this function is imported
func Head() {
	apptype.Db = apptype.ConnectToDatabase(true)
	testList()
}
