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
	i = 0
	j = 0
)

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

func handleDB(con *Connection) {
	con.UpdateAction()
	con.UpdateLevel()
	con.UpdateLevel()
}

func handleReq(con *Connection) *types.Request {
	req := new(types.Request)
	req.Id = 0
	req.Act = "registration"
	req.Level = i
	req.Language = "ru"
	req.Connection = con.con
	if i == 0 {
		if j == 0 {
			req.Req = "Hello World ma booooooooy"
		} else if j == 1 {
			req.Req = "Owww, nooo, you are the best human in the world!"
		} else if j == 2 {
			req.Req = "/start"
		}
	} else if i == 1 {
		if j == 0 {
			req.Req = "Hello World ma booooooooy"
		} else if j == 1 {
			req.Req = "Owww, nooo, you are the best human in the world!"
		} else if j == 2 {
			req.Req = "GoReg"
		}
	} else if i == 2 {
		if j == 0 {
			req.Req = "Hello World ma booooooooy"
		} else if j == 1 {
			req.Req = "Owww, nooo, you are the best human in the world!"
		} else if j == 2 {
			req.Req = "GoNext"
		}
	}
	return req
}

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
			functional.Dir(i, j, res)
			log.Printf("Secondary test %d was completed", j)
			j++
		}
		log.Printf("Main test %d was completed", i)
		i++
	}
}

func Head() {
	con := new(Connection)
	con.con = types.ConnectToDatabase(false)
	defer con.DeleteUser()
	con.CreateUser()
	action(con)

}
