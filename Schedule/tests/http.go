package tests

import (
	"Schedule/apptype"
	"Schedule/tests/functional"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Call the server
func callSche(body []byte) *apptype.Response {
	result := new(apptype.Response)
	resp, err := http.Post("http://localhost:8079/Schedule", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
	} else {
		json.NewDecoder(resp.Body).Decode(&result)
	}
	defer resp.Body.Close()
	return result
}

// Preparation request
func handleReq() *apptype.Request {
	req := new(apptype.Request)
	req.Id = 488
	req.Act = "see schedule"
	req.Language = "ru"
	return req
}

// All ations are planed here
func action() {
	req := handleReq()
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	res := callSche(jsonBytes)
	functional.IsItSch(res)
	log.Printf("The test was completed")
}

// Inits logs
func initlogs(file string) *os.File {
	os.Remove(file)
	lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(lf)
	return lf
}

// There is main logic
// these is the Head of tests
func Head() {
	f := initlogs("app.log")
	apptype.Client = apptype.AddCleint()
	defer delSch()
	createSch()
	action()
	defer f.Close()

}
