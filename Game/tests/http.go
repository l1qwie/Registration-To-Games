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
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			panic(err)
		}
	}
	return result
}

func createGame() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	defer deleteTGame()
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
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeSport-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseSport, chsendSport, sendChSaveSport}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChSport, functional.ChSemiFinalSport, functional.ChFinalSport}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrash8, chtrash9, chtrash10, chtrash11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeDate() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeDate-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseDate, chsendDate, sendChSaveDate}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChDate, functional.ChSemiFinalDate, functional.ChFinalDate}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashdate8, chtrashdate9, chtrashdate10, chtrashdate11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeTime() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeTime-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseTime, chsendTime, sendChSaveTime}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChTime, functional.ChSemiFinalTime, functional.ChFinalTime}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashtime8, chtrashtime9, chtrashtime10, chtrashtime11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeSeats() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeSeats-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseSeats, chsendSeats, sendChSaveSeats}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChSeats, functional.ChSemiFinalSeats, functional.ChFinalSeats}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashseats8, chtrashseats9, chtrashtseats10, chtrashseats11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changePrice() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangePrice-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, chosePrice, chsendPrice, sendChSavePrice}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChPrice, functional.ChSemiFinalPrice, functional.ChFinalPrice}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashprice8, chtrashprice9, chtrashtprice10, chtrashprice11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeCurency() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	unverifiable = true
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeCurrency-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseCurrency, chsendCurrency, sendChSaveCurrency}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChCurrency, functional.ChSemiFinalCurrency, functional.ChFinalCurrency}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashtcurrency10, chtrashcurrency11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeLink() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	unverifiable = true
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeLink-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseLink, chsendLink, sendChSaveLink}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChLink, functional.ChSemiFinalLink, functional.ChFinalLink}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashtlink10, chtrashlink11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeAddress() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	unverifiable = true
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeAddress-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseAddress, chsendAddress, sendChSaveAddress}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChAddress, functional.ChSemiFinalAddress, functional.ChFinalAddress}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashaddress10, chtrashaddress11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func deleteGame() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	unverifiable = false
	ts := new(TestStuct)
	ts.Round = 3
	ts.Name = "Delete-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionDel, sendDelGame}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseDelGame, functional.FinalDel}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrashdel4, chtrashdel5}
	ts.UpdtLevel = []int{0, 1, 2}
	ts.DoTest()
}

func changeGame() {
	changeSport()
	changeDate()
	changeTime()
	changeSeats()
	changePrice()
	changeCurency()
	changeLink()
	changeAddress()
}

func Start() {
	createGame()
	changeGame()
	deleteGame()
}
