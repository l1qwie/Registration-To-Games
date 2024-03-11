package preparationdata

import (
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests/database"
	"RegistrationToGames/tests/media"
	"RegistrationToGames/tests/registration"
	"RegistrationToGames/tests/routine"
	"RegistrationToGames/tests/schedule"
	"RegistrationToGames/tests/welcome"
)

const (
	userIdT1   int  = 456
	userIdT2   int  = 477
	userIdT3   int  = 488
	userIdT4   int  = 499
	gameId     int  = 2
	pastgameId int  = 10
	START      int  = 0
	LEVEL1     int  = 1
	LEVEL2     int  = 2
	LEVEL3     int  = 3
	LEVEL4     int  = 4
	LEVEL5     int  = 5
	wSTART     int  = 0
	wLEVEL1    int  = 1
	wLEVEL2    int  = 2
	rSTART     int  = 3
	rLEVEL1    int  = 4
	rLEVEL2    int  = 5
	rLEVEL3    int  = 6
	rLEVEL4    int  = 7
	schSTART   int  = 8
	mStart     int  = 9
	mLEVEL1    int  = 10
	mLEVEL2    int  = 11
	mLEVEL3    int  = 12
	Unload     bool = true
	Upload     bool = false
)

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

func waitingYourMedia(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForWaitingYourMedia(responses, output)
	}
}

func unload(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForUnload(responses, output)
	}
}

func upload(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == 0 {
		media.JustTrash(responses, output)
	} else if counter == 1 {
		media.JustTrash2(responses, output)
	} else if counter == 2 {
		media.QueryForUpload(responses, output)
	}
}

func WelcomeAct(counter, i int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == wSTART {
		welcome.Start(responses, output)
		routine.DeleteUser(userIdT1)
	} else if counter == wLEVEL1 {
		preparationToShowRules(i, responses, output)
	} else if counter == wLEVEL2 {
		welcomeToMainMenu(i, responses, output)
	}
}

func SeeTheSchedule(counter int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == schSTART {
		for i := 0; i < 4; i++ {
			if database.FoundGame(i) {
				database.DeleteGame(i)
			}
		}
		schedule.QueryForSeeSchedule(responses, output)
	}
}

func RegToGamesAct(counter, i int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == rSTART {
		registration.QueryForPresentationScheduele(responses, output)
		routine.DeleteUser(userIdT2)
		database.DeleteGame(gameId)
	} else if counter == rLEVEL1 {
		chooseGame(i, responses, output)
	} else if counter == rLEVEL2 {
		chooseSeats(i, responses, output)
	} else if counter == rLEVEL3 {
		choosePayment(i, responses, output)
	} else if counter == rLEVEL4 {
		bestWishes(i, responses, output)
	}
}

func MediaUnload(counter, i int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == START {
		media.QueryForChooseDirection(responses, output)
		routine.DeleteUser(userIdT2)
		database.DeleteMedia(pastgameId)
		database.DeleteGame(pastgameId)
	} else if counter == LEVEL1 {
		chooseMediaGame(i, responses, output, Unload)
	} else if counter == LEVEL2 {
		unload(i, responses, output)
	}
}

func MediaUpload(counter, i int, responses chan *types.TelegramResponse, output chan *types.MessageResponse) {
	if counter == START {
		media.QueryForChooseDirection(responses, output)
		routine.DeleteUser(userIdT4)
		database.DeleteMedia(pastgameId)
		database.DeleteGame(pastgameId)
	} else if counter == LEVEL1 {
		//chooseMediaGame(i, responses, output, Upload)
		waitingYourMedia(i, responses, output)
	} else if counter == LEVEL2 {
		//waitingYourMedia(i, responses, output)
		upload(i, responses, output)
		//} else if counter == LEVEL3 {
		//upload(i, responses, output)
		//}
	}
}
