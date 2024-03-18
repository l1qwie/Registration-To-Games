package texecuter

import (
	"RegistrationToGames/fmtogram"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests/database"
	"fmt"
)

const (
	userIdT1     int  = 456
	userIdT2     int  = 477
	userIdT3     int  = 488
	userIdT4     int  = 499
	userIdT5     int  = 899
	gameId       int  = 2
	wrongAnswers int  = 2
	pastgameId   int  = 10
	START        int  = 0
	LEVEL1       int  = 1
	LEVEL2       int  = 2
	LEVEL3       int  = 3
	LEVEL4       int  = 4
	LEVEL5       int  = 5
	wSTART       int  = 0
	wLEVEL1      int  = 1
	wLEVEL2      int  = 2
	rSTART       int  = 3
	rLEVEL1      int  = 4
	rLEVEL2      int  = 5
	rLEVEL3      int  = 6
	rLEVEL4      int  = 7
	schSTART     int  = 8
	mStart       int  = 9
	mLEVEL1      int  = 10
	mLEVEL2      int  = 11
	mLEVEL3      int  = 12
	Unload       bool = true
	Upload       bool = false
)

// This stuct only for Update functions to update database
type Update struct {
	Act, Lang     string
	Level, UserId int
}

// This stuct is for all tests. The main thought is you can use it anywhere
type TestStuct struct {
	TRcount, Trshcount, MRcount, Round, Wcounter int
	ArrFuncTr, ArrFuncTrash                      []func() *types.TelegramResponse
	ArrFuncMr                                    []func() *types.MessageResponse
	ArrFuncAss                                   []func(*formatter.Formatter)
	ArrFuncChDB                                  []func(int)
	Query                                        chan *types.TelegramResponse
	Response                                     chan *types.MessageResponse
	Fmt                                          chan *formatter.Formatter
	Update                                       []Update
	Name                                         string
}

// Check trash-query, call a func and return true
// else - return false
func (t *TestStuct) checkTheTrash() bool {
	if len(t.ArrFuncTrash) > t.Trshcount {
		t.Query <- t.ArrFuncTrash[t.Trshcount]()
	}
	return t.Trshcount < len(t.ArrFuncTrash)
}

// The fucntion has to check what answer is and if it's worng answer - call a func and return true
// else - return false
// only if the main counter != 0
func (t *TestStuct) checkTheWorng() bool {
	if t.Wcounter < wrongAnswers {
		if t.TRcount != 0 {
			t.ArrFuncAss[t.TRcount-1](<-t.Fmt)
			t.ArrFuncChDB[t.TRcount-1](t.Update[t.TRcount].UserId)
		}
	}
	return (t.Wcounter < wrongAnswers) && (t.TRcount != 0)
}

// The head of making a query from 'user' to bot
// check the trash and then if all is OK call a func
// also create a response to previos query from user
func (t *TestStuct) theHead() {
	if !t.checkTheTrash() {
		t.Query <- t.ArrFuncTr[t.TRcount]()
	}
	t.Response <- t.ArrFuncMr[t.MRcount]()
}

// Preparation the datbase. Just the routine
func (t *TestStuct) prepDatabase() {
	if t.TRcount < len(t.Update) {
		database.UpdateAction(t.Update[t.TRcount].Act, t.Update[t.TRcount].UserId)
		database.UpdateLanguage(t.Update[t.TRcount].Lang, t.Update[t.TRcount].UserId)
		database.UpdateLevel(t.Update[t.TRcount].Level, t.Update[t.TRcount].UserId)
	}
}

// Accept bot answers
// Call a func if it's not a wrong answer
func (t *TestStuct) acceptAnswers() {
	if !t.checkTheWorng() {
		t.ArrFuncAss[t.TRcount](<-t.Fmt)
		t.ArrFuncChDB[t.TRcount](t.Update[t.TRcount].UserId)
	}
}

// The body for all tests
func (t *TestStuct) DoTest() {
	for t.TRcount < t.Round {
		t.Trshcount = 0
		t.Wcounter = 0
		for t.Trshcount < 3 {
			t.theHead()
			t.prepDatabase()
			fmtogram.Worker(t.Query, t.Response, t.Fmt)
			t.acceptAnswers()
			t.Trshcount++
			t.Wcounter++
		}
		fmt.Printf("%s %d has been complete\n", t.Name, t.TRcount)
		t.TRcount++
		if t.MRcount+1 < len(t.ArrFuncMr) {
			t.MRcount++
		}
	}
}
