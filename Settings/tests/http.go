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
	resp, err := http.Post("http://localhost:8084/Settings", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
	} else {
		json.NewDecoder(resp.Body).Decode(&result)
	}
	defer resp.Body.Close()
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
	f := initlogs("applang.log")
	ts := new(TestStuct)
	ts.Round = 2
	ts.Name = "ChangeLanguageTest"
	ts.FuncReq = []func() *apptype.Request{chOpt, chLang}
	ts.FuncRes = []func(*apptype.Response){functional.ChOptOnlyLang, functional.ChooseLang}
	ts.FuncTrsh = []func() *apptype.Request{commontrash, commontrash2, langtr2, langtr3}
	ts.UpdtLevel = []int{0, 2}
	onetime = true
	ts.DoTest()
	defer f.Close()
}

// Tests Delete a Game functional
func delGame() {
	defer DeleteUser()
	defer DeleteUserSchedule()
	defer DeleteSchedule()
	CreateScheduleForUser()
	CreateUserScehdule()
	CreateUser()
	f := initlogs("appdelGame.log")
	ts := new(TestStuct)
	ts.Round = 4
	ts.Name = "DeleteGameTest"
	ts.FuncReq = []func() *apptype.Request{chOpt, chRec, delGameId, del}
	ts.FuncRes = []func(*apptype.Response){functional.ChTwoOpt, functional.ChGame, functional.ChOrDel, functional.DelGame}
	ts.FuncTrsh = []func() *apptype.Request{commontrash, commontrash2, dGametr2, dGametr3, dGametr4, dGametr5, dGametr6, dGametr7}
	ts.UpdtLevel = []int{0, 1, 2, 3}
	ts.DoTest()
	defer f.Close()
}

// The list of testing functions
func testList() {
	//language()
	delGame()
	//chSeats()
	//chPayment()
}

// The head of the directioner
// Only this function is imported
func Head() {
	apptype.Db = apptype.ConnectToDatabase(false)
	testList()

}