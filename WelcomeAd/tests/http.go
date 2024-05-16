package tests

import (
	"Welcome/apptype"
	"Welcome/tests/functional"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func callreg(body []byte) *apptype.Response {
	result := new(apptype.Response)
	resp, err := http.Post("http://localhost:8100/Welcome", "application/json", bytes.NewBuffer(body))
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

func welcomeTest() {
	//f := initlogs("welcome.log")
	ts := new(TestStuct)
	ts.Round = 3
	ts.Name = "WelcomeTest"
	ts.FuncReq = []func() *apptype.Request{hi, letsgo, rulesAgree}
	ts.FuncRes = []func(*apptype.Response){functional.SayHello, functional.SendTheRules, functional.SendMainMenu}
	ts.FuncTrsh = []func() *apptype.Request{trash1, trash2, trash3, trash4, trash5, trash6}
	ts.UpdtLevel = []int{0, 1, 2}
	ts.DoTest()
	//defer f.Close()
}

// The head of the directioner
// Only this function is imported
func Start() {
	welcomeTest()
}
