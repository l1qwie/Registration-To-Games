package tests

import (
	"Media/apptype"
	"Media/tests/functional"
	"encoding/json"
	"log"
)

var onetime, condtition bool

const wrongAnswers int = 2

// Update .. This stuct only for Update functions to update database
type Update struct {
	Act, Lang     string
	Level, UserId int
}

// TestStuct .. This stuct is for all tests. The main thought is you can use it anywhere
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
	if t.Trshcount < 2 {
		t.request = t.FuncTrsh[t.inficount]()
		t.inficount++
	}
	return t.Trshcount < 2
}

// The fucntion has to check what answer is and if it's worng answer - call a func and return true
// else - return false
// only if the main counter != 0
func (t *TestStuct) checkTheWorng() bool {
	if t.Wcounter < wrongAnswers {
		if t.TRcount != 0 {
			if onetime && t.TRcount == 1 {
				if condtition {
					functional.NochOnluUpAfewSecond(t.response)
				} else {
					functional.NoChOnluUpSecond(t.response)
				}
			} else {
				t.FuncRes[t.TRcount-1](t.response)
			}
		}
	}
	return (t.Wcounter < wrongAnswers) && (t.TRcount != 0)
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
	log.Printf(`req.Limit = %d`, t.request.Limit)
	log.Printf(`req.LaunchPoint = %d`, t.request.LaunchPoint)
	log.Printf(`req.Req = "%s"`, t.request.Req)
	log.Printf(`req.MediaDir = "%s"`, t.request.MediaDir)
	log.Printf(`req.GameId = %d`, t.request.GameId)
	log.Printf(`req.FileId = "%s"`, t.request.FileId)
	log.Printf(`req.TypeOffile = "%s"`, t.request.TypeOffile)
	log.Printf(`req.Media = "%s"`, t.request.Media)
	log.Print()
}

// Makes []bytes from json(object) and send it
func (t *TestStuct) doRequest() {
	t.logReq()
	jb, err := json.Marshal(t.request)
	if err != nil {
		panic(err)
	}
	t.response = callmedia(jb)
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
	log.Printf(`res.LaunchPoint = %d`, t.response.LaunchPoint)
	log.Printf(`res.Act = "%s"`, t.response.Act)
	log.Printf(`req.MediaDir = "%s"`, t.response.MediaDir)
	log.Printf(`req.GameId = %d`, t.response.GameId)
	log.Printf(`req.FileId = "%s"`, t.response.FileId)
	log.Printf(`req.TypeOffile = "%s"`, t.response.TypeOffile)
	log.Printf(`req.Media = "%s"`, t.response.Media)
	log.Printf(`res.ParseMode = "%s"`, t.response.ParseMode)
	log.Print()
}

// Accept bot answers.
// Call a func if it's not a wrong answer
func (t *TestStuct) acceptAnswers() {
	if !t.checkTheWorng() {
		t.logRes()
		t.FuncRes[t.TRcount](t.response)
	}
}

// DoTest is a body for all tests.
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
