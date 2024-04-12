package tests

import (
	"Settings/apptype"
	"Settings/tests/functional"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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

// Tests change language functional
func language() {
	defer DeleteUser()
	CreateUser()
	ts := new(TestStuct)
	ts.Round = 2
	ts.Name = "ChangeLanguageTest"
	ts.FuncReq = []func() *apptype.Request{chOpt, chLang}
	ts.FuncRes = []func(*apptype.Response){functional.ChOptOnlyLang, functional.ChooseLang}
	ts.FuncTrsh = []func(int) *apptype.Request{trash, trash2}
	ts.UpdtLevel = []int{0, 2}
	onetime = true
	ts.DoTest()
}

// The list of testing functions
func testList() {
	language()
	//delGame()
	//chSeats()
	//chPayment()
}

// The head of the directioner
// Only this function is imported
func Head() {
	apptype.Db = apptype.ConnectToDatabase(false)
	testList()

}
