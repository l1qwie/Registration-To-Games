package preparationdata

import (
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests/database"
	"RegistrationToGames/tests/media"
	"RegistrationToGames/tests/registration"
	"RegistrationToGames/tests/routine"
	"RegistrationToGames/tests/schedule"
	"RegistrationToGames/tests/settings"
	"RegistrationToGames/tests/welcome"
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

type TestStuct struct {
	TRcount, Trshcount, MRcount int
	ArrFuncTr                   []func() *types.TelegramResponse
	ArrFuncMr                   []func() *types.MessageResponse
	ArrFuncTrash                []func() *types.TelegramResponse
	Query                       chan *types.TelegramResponse
	Response                    chan *types.MessageResponse
}

func (t *TestStuct) checkTheTrash() bool {
	if len(t.ArrFuncTrash) > t.Trshcount {
		t.Query <- t.ArrFuncTrash[t.Trshcount]()
	}
	return t.Trshcount < len(t.ArrFuncTrash)
}

func (t *TestStuct) TheHead() {
	if !t.checkTheTrash() {
		t.Query <- t.ArrFuncTr[t.TRcount]()
	}
	t.Response <- t.ArrFuncMr[t.MRcount]()
}

func preparationToShowRules(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		welcome.JustTrash(responses, output)
	} else if counter == 1 {
		welcome.JustTrash2(responses, output)
	} else if counter == 2 {
		welcome.QueryForShowRules(responses, output)
	}
	routine.DeleteUser(userIdT1)
}

func welcomeToMainMenu(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		welcome.JustTrash(responses, output)
	} else if counter == 1 {
		welcome.JustTrash2(responses, output)
	} else if counter == 2 {
		welcome.QueryForWelcomeToMainMenu(responses, output)
	}
	routine.DeleteUser(userIdT1)
}

func chooseGame(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		registration.JustTrash(responses, output)
	} else if counter == 1 {
		registration.JustTrash2(responses, output)
	} else if counter == 2 {
		registration.QueryForChooseGame(responses, output)
	}
	routine.DeleteUser(userIdT2)
	database.DeleteGame(gameId)
}

func chooseSeats(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		registration.JustTrash(responses, output)
	} else if counter == 1 {
		registration.JustTrash2(responses, output)
	} else if counter == 2 {
		registration.QueryForChooseSeats(responses, output)
	}
	routine.DeleteUser(userIdT2)
	database.DeleteGame(gameId)
}

func choosePayment(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		registration.JustTrash(responses, output)
	} else if counter == 1 {
		registration.JustTrash2(responses, output)
	} else if counter == 2 {
		registration.QueryForChoosePayment(responses, output)
	}
	routine.DeleteUser(userIdT2)
	database.DeleteGame(gameId)
}

func bestWishes(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		registration.JustTrash(responses, output)
	} else if counter == 1 {
		registration.JustTrash2(responses, output)
	} else if counter == 2 {
		registration.QueryForBestWishes(responses, output)
	}
	routine.DeleteUser(userIdT2)
	database.DeleteGame(gameId)
}

/*
func chooseMediaGame(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse, direction bool) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		if direction {
			media.QueryForChooseMediaGameUnload(responses, output)
		} else {
			media.QueryForChooseMediaGameUpload(responses, output)
		}
	}
}

func waitingYourMediaOne(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForWaitingYourMediaOne(responses, output)
	}
}

func waitingYourMediaAfew(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForWaitingYourMediaAfew(responses, output)
	}
}

func unloadOne(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForUnloadone(responses, output)

	}
}

func unloadAfew(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForUnloadAfew(responses, output)
	}
}
func uploadOne(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForUploadOne(responses, output)
	}
}

func uploadAfew(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForUploadAfew(responses, output)
	}
}

/*
func checkTheTrash(c int, arrtrash []func() *types.TelegramResponse, resp chan<- *types.TelegramResponse) bool {
	if len(arrtrash) > c {
		resp <- arrtrash[c]()
	}
	return c < len(arrtrash)

		if c == 0 {
			if act == "set" {
				resp <- settings.Trash()
			} else if act == "media" {
				resp <- media.Trash()
			}
		} else if c == 1 {
			if act == "set" {
				resp <- settings.Trash2()
			} else if act == "media" {
				resp <- media.Trash2()
			}
		}
		return c < wrongAnswers
*/

func WelcomeAct(counter, i int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == START {
		welcome.Start(responses, output)
		routine.DeleteUser(userIdT1)
	} else if counter == LEVEL1 {
		preparationToShowRules(i, responses, output)
	} else if counter == LEVEL2 {
		welcomeToMainMenu(i, responses, output)
	}
}

func SeeTheSchedule(responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	for i := 0; i < 4; i++ {
		if database.FoundGame(i) {
			database.DeleteGame(i)
		}
	}
	schedule.QueryForSeeSchedule(responses, output)
}

func RegToGamesAct(counter, i int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == START {
		registration.QueryForPresentationScheduele(responses, output)
		routine.DeleteUser(userIdT2)
		database.DeleteGame(gameId)
	} else if counter == LEVEL1 {
		chooseGame(i, responses, output)
	} else if counter == LEVEL2 {
		chooseSeats(i, responses, output)
	} else if counter == LEVEL3 {
		choosePayment(i, responses, output)
	} else if counter == LEVEL4 {
		bestWishes(i, responses, output)
	}
}

func MediaUnloadOne(counter, i int, resp chan *types.TelegramResponse, op chan *types.MessageResponse) {
	//if !checkTheTrash(i, "media", resp) {
	if counter == START {
		resp <- media.ChDir()
	} else if counter == LEVEL1 {
		resp <- media.ChMUnload()
	} else if counter == LEVEL2 {
		resp <- media.UnlOne()
	}
	//}
	op <- settings.MesResp()
}

func MediaUnloadAfew(counter, i int, resp chan *types.TelegramResponse, op chan *types.MessageResponse, arrfunc []func() *types.TelegramResponse) {
	//if !checkTheTrash(i, "media", resp) {
	resp <- arrfunc[counter]()
	//}
	op <- media.MesResp()
}

func SetActChanLang(counter, i int, resp chan *types.TelegramResponse, op chan *types.MessageResponse, arrfunc []func() *types.TelegramResponse) {
	//if !checkTheTrash(i, "set", resp) {
	resp <- arrfunc[counter]()
	//}
	op <- settings.MesResp()
}

func SetActDeleteGame(counter, i int, resp chan *types.TelegramResponse, op chan *types.MessageResponse) {
	//if !checkTheTrash(i, "set", resp) {
	if counter == START {
		resp <- settings.ChOpt()
	} else if counter == LEVEL1 {
		resp <- settings.ChGame()
	} else if counter == LEVEL2 {
		resp <- settings.ChDelGame()
	}
	//}
	op <- settings.MesResp()
}

func MediaUploadOne(counter, i int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == START {
		//media.QueryForChooseDirection(responses, output)
		routine.DeleteUser(userIdT4)
		database.DeleteMedia(pastgameId)
		database.DeleteGame(pastgameId)
	} else if counter == LEVEL1 {
		//waitingYourMediaOne(i, responses, output)
	} else if counter == LEVEL2 {
		//uploadOne(i, responses, output)
	}
}

func MediaUploadAfew(counter, i int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == START {
		//	media.QueryForChooseDirection(responses, output)
		routine.DeleteUser(userIdT4)
		database.DeleteMedia(pastgameId)
		database.DeleteGame(pastgameId)
	} else if counter == LEVEL1 {
		//waitingYourMediaAfew(i, responses, output)
	} else if counter == LEVEL2 {
		//uploadAfew(i, responses, output)
	}
}
