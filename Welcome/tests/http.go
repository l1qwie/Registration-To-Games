package tests

import (
	"Welcome/tests/functional"
	"Welcome/types"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var (
	i       = 0
	j       = 0
	request = []string{"/start", "GoReg", "GoNext"}
	trash   = []string{"Hello World ma booooooooy", "Owww, nooo, you are the best human in the world!"}
)

// Makes the requst and take a response
func callwelcome(body []byte) *types.Response {
	result := new(types.Response)
	resp, err := http.Post("http://localhost:8081/Welcome", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
	} else {
		json.NewDecoder(resp.Body).Decode(&result)
	}
	defer resp.Body.Close()
	return result
}

// Preparing the Database
func handleDB(con *Connection) {
	con.UpdateAction()
	con.UpdateLevel()
	con.UpdateLevel()
}

// Preparing a request
func handleReq(con *Connection) *types.Request {
	req := new(types.Request)
	req.Id = 456
	req.Act = "registration"
	req.Level = i
	req.Language = "ru"
	req.Connection = con.con
	if j < 2 {
		req.Req = trash[j]
	} else {
		req.Req = request[i]
	}
	return req
}

// All main actions happen here
func action(con *Connection) {
	for i < 3 {
		j = 0
		for j < 3 {
			handleDB(con)
			req := handleReq(con)
			jsonBytes, err := json.Marshal(req)
			if err != nil {
				panic(err)
			}
			res := callwelcome(jsonBytes)
			if res.Error == "" {
				functional.Dir(i, j, res)
				log.Printf("Secondary test %d was completed", j)
				j++
			} else {
				panic(res.Error)
			}
		}
		log.Printf("Main test %d was completed", i)
		i++
	}
}

// The head of the directioner
// Only this function is imported
func Head() {
	con := new(Connection)
	con.con = types.ConnectToDatabase(false)
	defer con.DeleteUser()
	con.CreateUser()
	action(con)

}
