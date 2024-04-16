package tests

import (
	"Media/apptype"
	"Media/tests/functional"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Makes the requst and take a response
func callmedia(body []byte) *apptype.Response {
	result := new(apptype.Response)
	resp, err := http.Post("http://localhost:8085/Media", "application/json", bytes.NewBuffer(body))
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

// Tested the Unload One functional
func unOne() {
	defer DeleteUser()
	defer DeleteGame()
	defer DeleteMedia()
	CreateNotFullMediaGame()
	CreateUser()
	f := initlogs("unOne.log")
	ts := new(TestStuct)
	ts.Round = 3
	ts.Name = "Unload-One-Test"
	ts.FuncReq = []func() *apptype.Request{chMediaDir, chMUnload, unJustOne}
	ts.FuncRes = []func(*apptype.Response){functional.ChDirection, functional.ChMediaGameUn, functional.UnOne}
	ts.FuncTrsh = []func() *apptype.Request{common1, common2, common3, common4, unOnetr3, unOnetr4}
	ts.UpdtLevel = []int{0, 1, 3}
	ts.DoTest()
	defer f.Close()
}

// Tested the Unload a Few functional
func unAfew() {
	defer DeleteUser()
	defer DeleteSchedule()
	defer DeleteMediaSchedule()
	CreateMediaShedule()
	FillMediaSchedule()
	CreateUser()
	f := initlogs("unAfew.log")
	ts := new(TestStuct)
	ts.Round = 3
	ts.Name = "Unload-One-Test"
	ts.FuncReq = []func() *apptype.Request{chMediaDir, chMUnload, unJustAfew}
	ts.FuncRes = []func(*apptype.Response){functional.ChDirection, functional.ChMediaGamesUn, functional.UnAll}
	ts.FuncTrsh = []func() *apptype.Request{common1, common2, common3, common4, unOnetr3, unOnetr4}
	ts.UpdtLevel = []int{0, 1, 3}
	ts.DoTest()
	defer f.Close()
}

// Tested the Upload One functional
func upOne() {
	defer DeleteUser()
	defer DeleteGame()
	defer DeleteMedia()
	CreateEmptyMediaGame()
	CreateUser()
	f := initlogs("upOne.log")
	ts := new(TestStuct)
	ts.Round = 3
	ts.Name = "Upload-One-Test"
	ts.FuncReq = []func() *apptype.Request{chMediaDir, chMUpload, upJustOne}
	ts.FuncRes = []func(*apptype.Response){functional.NoChOnluUp, functional.WaitForMedia, functional.UpOne}
	ts.FuncTrsh = []func() *apptype.Request{upOnetr1, upOnetr2, upOnetr3, upOnetr4, upOnetr5, upOnetr6}
	ts.UpdtLevel = []int{0, 2, 3}
	onetime = true
	ts.DoTest()
	defer f.Close()
}

// Tested the Upload a Few functional
func upAfew() {
	defer DeleteUser()
	defer DeleteSchedule()
	defer DeleteMediaSchedule()
	CreateMediaShedule()
	CreateUser()
	f := initlogs("upAfew.log")
	ts := new(TestStuct)
	ts.Round = 3
	ts.Name = "Upload-a-few-Test"
	ts.FuncReq = []func() *apptype.Request{chMediaDir, chMUploadAfew, upJustAfew}
	ts.FuncRes = []func(*apptype.Response){functional.NoChOnluUpAfew, functional.WaitForMediaAfew, functional.UpAfew}
	ts.FuncTrsh = []func() *apptype.Request{upOnetr1, upOnetr2, upOnetr3, upOnetr4, upAfew1, upAfew2}
	ts.UpdtLevel = []int{0, 2, 3}
	onetime = true
	condtition = true
	ts.DoTest()
	defer f.Close()
}

// The list of testing functions
func testList() {
	unOne()
	unAfew()
	upOne()
	upAfew()
}

// The head of the directioner
// Only this function is imported
func Head() {
	apptype.Db = apptype.ConnectToDatabase(false)
	testList()

}
