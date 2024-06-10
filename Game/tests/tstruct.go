package tests

import (
	"Game/apptype"
	"encoding/json"
	"log"
)

const wrongAnswers int = 2

// Update stuct only for Update functions to update database
type Update struct {
	Act, Lang     string
	Level, UserId int
}

var onetime bool
var twotimes int
var unverifiable bool

// TestStuct stuct is for all tests. The main thought is you can use it anywhere
type TestStuct struct {
	TRcount, Trshcount, Round, Wcounter, inficount int
	FuncReq                                        []func() *apptype.Request
	FuncRes                                        []func(*apptype.Response)
	FuncTrsh                                       []func() *apptype.Request
	request                                        *apptype.Request
	response                                       *apptype.Response
	UpdtLevel                                      []int
	Name                                           string
	Logf                                           string
}

// Check trash-query, call a func and return true
// else - return false
func (t *TestStuct) checkTheTrash() bool {
	var condition bool
	if t.Trshcount < 2 && (t.TRcount < 6 || t.TRcount > 8) {
		if !unverifiable || unverifiable && t.TRcount < 3 {
			condition = true
			t.request = t.FuncTrsh[t.inficount]()
			t.inficount++
		}
	}
	return condition
}

// The fucntion has to check what answer is and if it's worng answer - call a func and return true
// else - return false
// only if the main counter != 0
func (t *TestStuct) checkTheWorng() bool {
	var condition bool
	if t.Wcounter < wrongAnswers && (t.TRcount < 6 || t.TRcount > 8) {
		if !unverifiable || unverifiable && t.TRcount < 3 {
			if t.TRcount != 0 {
				condition = true
				t.FuncRes[t.TRcount-1](t.response)
			}
		}
	}
	return condition
}

// The head of making a query from 'user' to bot
// check the trash and then if all is OK call a func
// also create a response to previos query from user
func (t *TestStuct) theHead() {
	if !t.checkTheTrash() {
		t.request = t.FuncReq[t.TRcount]()
	}
}

// Logs data from Request
func (t *TestStuct) logReq() {
	log.Print("Request:")
	log.Printf(`req.Id = %d`, t.request.Id)
	log.Printf(`req.Level = %d`, t.request.Level)
	log.Printf(`req.Language = "%s"`, t.request.Language)
	log.Printf(`req.Act = "%s"`, t.request.Act)
	log.Printf(`req.Direction = "%s"`, t.request.Direction)
	log.Printf(`req.Changeable = "%s"`, t.request.Changeable)
	log.Printf(`req.GameId = %d`, t.request.GameId)
	log.Printf(`req.Limit = %d`, t.request.Limit)
	log.Printf(`req.LaunchPoint = %d`, t.request.LaunchPoint)
	log.Printf(`req.Req = "%s"`, t.request.Req)
	log.Printf(`req.Sport = "%s"`, t.request.Sport)
	log.Printf(`req.Date = %d`, t.request.Date)
	log.Printf(`req.Time = "%d"`, t.request.Time)
	log.Printf(`req.Seats = "%d"`, t.request.Seats)
	log.Printf(`req.Price = "%d"`, t.request.Price)
	log.Printf(`req.Currency = "%s"`, t.request.Currency)
	log.Printf(`req.Link = "%s"`, t.request.Link)
	log.Printf(`req.Address = "%s"`, t.request.Address)
	log.Printf(`req.Status = "%v"`, t.request.Status)
	log.Print()
}

// Makes []bytes from json(object) and send it
func (t *TestStuct) doRequest() {
	t.logReq()
	jb, err := json.Marshal(t.request)
	if err != nil {
		panic(err)
	}
	t.response = callgame(jb)
	if t.response.Error != "" {
		panic(t.response.Error)
	}
}

// Logs data from Response
func (t *TestStuct) logRes() {
	log.Print("Response:")
	log.Printf(`res.Message = "%s"`, t.response.Message)
	log.Printf(`res.Keyboard = "%s"`, t.response.Keyboard)
	log.Printf(`res.ChatId = %d`, t.response.ChatID)
	log.Printf(`res.level = %d`, t.response.Level)
	log.Printf(`res.Direction = "%s"`, t.response.Direction)
	log.Printf(`res.Changeable = "%s"`, t.response.Changeable)
	log.Printf(`res.GameId = %d`, t.response.GameId)
	log.Printf(`res.LaunchPoint = %d`, t.response.LaunchPoint)
	log.Printf(`res.Act = "%s"`, t.response.Act)
	log.Printf(`req.Sport = "%s"`, t.response.Sport)
	log.Printf(`req.Date = %d`, t.response.Date)
	log.Printf(`req.Time = "%d"`, t.response.Time)
	log.Printf(`req.Seats = "%d"`, t.response.Seats)
	log.Printf(`req.Price = "%d"`, t.response.Price)
	log.Printf(`req.Currency = "%s"`, t.response.Currency)
	log.Printf(`req.Link = "%s"`, t.response.Link)
	log.Printf(`req.Address = "%s"`, t.response.Address)
	log.Printf(`req.Error = "%s"`, t.response.Error)
	log.Printf(`req.Status = %v`, t.response.Status)
	log.Print()
}

// Accept bot answers
// Call a func if it's not a wrong answer
func (t *TestStuct) acceptAnswers() {
	t.logRes()
	if !t.checkTheWorng() {
		t.FuncRes[t.TRcount](t.response)
	}
}

// DoTest is a body for all tests
// Only this function could be imported
func (t *TestStuct) DoTest() {
	for t.TRcount < t.Round {
		t.Trshcount = 0
		t.Wcounter = 0
		for t.Trshcount < 3 {
			t.theHead()
			//t.prepDatabase()
			t.doRequest()
			t.acceptAnswers()
			t.Trshcount++
			t.Wcounter++
		}
		log.Printf("%s %d has been complete\n", t.Name, t.TRcount)
		t.TRcount++
	}
}
