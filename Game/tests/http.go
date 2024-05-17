package tests

import (
	"Game/apptype"
	"Game/tests/functional"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func callgame(body []byte) *apptype.Response {
	result := new(apptype.Response)
	resp, err := http.Post("http://localhost:8099/Game", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
	} else {
		json.NewDecoder(resp.Body).Decode(&result)
	}
	defer resp.Body.Close()
	return result
}

func createGame() {
	ts := new(TestStuct)
	ts.Round = 0
	ts.Name = "Create-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendSport, sendDate, sendTime, sendSeats, sendPrice, sendCurrency, sendLink, sendAddress, sendSave}
	ts.FuncRes = []func(*apptype.Response){functional.SelectSport, functional.SelectDate, functional.SelectTime,
		functional.SelectSeats, functional.SelectPrice, functional.SelectCurrency, functional.SelectLink,
		functional.SelectAddress, functional.SemiFinal, functional.Final}
	ts.FuncTrsh = []func() *apptype.Request{}
	ts.UpdtLevel = []int{}
	ts.DoTest()
}

func Start() {
	createGame()
	//changeGame()
	//deleteGame()
}
