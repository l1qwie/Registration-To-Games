package tests

import (
	"Registration/apptype"
	"Registration/tests/functional"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var (
	i       = 0
	j       = 0
	request = []string{"I don't want to type anything", "2", "2", "card", "Next"}
	trash   = []string{"O, I want to registrait myself to game!", "!#!@#QWESASED#!#!#$%#3123e1"}
)

// Makes the requst and take a response
func callreg(body []byte) *apptype.Response {
	result := new(apptype.Response)
	resp, err := http.Post("http://localhost:8094/Registration", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
	} else {
		json.NewDecoder(resp.Body).Decode(&result)
	}
	defer resp.Body.Close()
	return result
}

// Preparing a request
func handleReq() *apptype.Request {
	req := new(apptype.Request)
	req.Id = 477
	req.Act = "reg to games"
	req.Level = i
	req.Language = "ru"
	req.Limit = 7
	if i == 0 && j == 0 {
		req.Status = true
	} else {
		req.Status = false
	}
	if j < 2 {
		req.Req = trash[j]
	} else {
		req.Req = request[i]
	}
	if i == 2 {
		req.GameId = 2
	}
	if i == 3 {
		req.GameId = 2
		req.Seats = 2
	}
	if i == 4 {
		req.GameId = 2
		req.Seats = 2
		req.Payment = "card"
	}
	return req
}

// All main actions happen here
func action() {
	for i < 5 {
		j = 0
		for j < 3 {
			req := handleReq()
			jsonBytes, err := json.Marshal(req)
			if err != nil {
				panic(err)
			}
			res := callreg(jsonBytes)
			if res.Error == "" {
				functional.Dir(res, i, j)
				log.Printf("Secondary test %d was completed", j)
			} else {
				log.Print(res.Error)
			}
			j++
		}
		i++
	}
}

// The head of the directioner
// Only this function is imported
func Head() {
	apptype.Db = apptype.ConnectToDatabase()
	defer DeleteGameWithUser()
	defer DeleteGame()
	CreateGame()
	action()
}
