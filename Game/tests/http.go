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
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			panic(err)
		}
	}
	defer resp.Body.Close()
	return result
}

func createGame() {
	defer deleteGame()
	ts := new(TestStuct)
	ts.Round = 10
	ts.Name = "Create-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendSport, sendDate, sendTime, sendSeats, sendPrice, sendCurrency, sendLink, sendAddress, sendSave}
	ts.FuncRes = []func(*apptype.Response){functional.SelectSport, functional.SelectDate, functional.SelectTime,
		functional.SelectSeats, functional.SelectPrice, functional.SelectCurrency, functional.SelectLink,
		functional.SelectAddress, functional.SemiFinal, functional.Final}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, trash2, trash3, trash4, trash5, trash6, trash7, trash8, trash9, trash10, trash11, trash18, trash19}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ts.DoTest()
}

func changeSport() {
	defer deleteGame()
	ts := new(TestStuct)
	ts.Round = 0
	ts.Name = "ChangeSport-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendGame, choseSport, chsendSport, sendChSave}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable, functional.ChSportSemiFinal, functional.ChFinal}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5}
	ts.UpdtLevel = []int{0, 1, 2, 3}
	ts.DoTest()
}

func changeGame() {
	changeSport()
	//changeDate()
	//changeTime()
	//changeSeats()
	//changePrice()
	//changeCurency()
	//changeLink()
	//changeAddress()
}

func Start() {
	createGame()
	//changeGame()
	//deleteGame()
}
