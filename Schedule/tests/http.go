package tests

import (
	"Schedule/apptype"
	"Schedule/tests/functional"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Call the server
func callSche(body []byte) *apptype.Response {
	result := new(apptype.Response)
	resp, err := http.Post("http://localhost:8083/Schedule", "application/json", bytes.NewBuffer(body))
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

// There is main logic
// these is the Head of tests
func Head() {
	apptype.Client = apptype.AddCleint()
	defer delSch()
	createSch()
	action()

}
