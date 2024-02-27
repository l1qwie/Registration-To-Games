package preparationdata

import (
	"registrationtogames/fmtogram/types"
	"registrationtogames/tests/database"
	"registrationtogames/tests/registration"
	"registrationtogames/tests/routine"
	"registrationtogames/tests/schedule"
	"registrationtogames/tests/welcome"
)

const (
	userIdT1 int = 456
	userIdT2 int = 477
	userIdT3 int = 488
	gameId   int = 2
	wSTART   int = 0
	wLEVEL1  int = 1
	wLEVEL2  int = 2
	rSTART   int = 3
	rLEVEL1  int = 4
	rLEVEL2  int = 5
	rLEVEL3  int = 6
	rLEVEL4  int = 7
	schSTART int = 8
)

func PreparationToShowRules(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		welcome.JustTrash(responses)
	} else if counter == 1 {
		welcome.JustTrash2(responses)
	} else if counter == 2 {
		welcome.QueryForShowRules(responses)
	}
	routine.DeleteUser(userIdT1)
}

func WelcomeToMainMenu(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		welcome.JustTrash(responses)
	} else if counter == 1 {
		welcome.JustTrash2(responses)
	} else if counter == 2 {
		welcome.QueryForWelcomeToMainMenu(responses)
	}
	routine.DeleteUser(userIdT1)
}

func ChooseGame(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		registration.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForChooseGame(responses)
	}
	routine.DeleteUser(userIdT2)
	database.DeleteGame(gameId)
}

func ChooseSeats(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		registration.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForChooseSeats(responses)
	}
	routine.DeleteUser(userIdT2)
	database.DeleteGame(gameId)
}

func ChoosePayment(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		registration.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForChoosePayment(responses)
	}
	routine.DeleteUser(userIdT2)
	database.DeleteGame(gameId)
}

func BestWishes(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		registration.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForBestWishes(responses)
	}
	routine.DeleteUser(userIdT2)
	database.DeleteGame(gameId)
}

func WelcomeAct(counter, i int, responses chan *types.TelegramResponse) {
	if counter == wSTART {
		welcome.Start(responses)
		routine.DeleteUser(userIdT1)
	} else if counter == wLEVEL1 {
		PreparationToShowRules(i, responses)
	} else if counter == wLEVEL2 {
		WelcomeToMainMenu(i, responses)
	}
}

func SeeTheSchedule(counter int, responses chan *types.TelegramResponse) {
	if counter == schSTART {
		for i := 0; i < 4; i++ {
			if database.FoundGame(i) {
				database.DeleteGame(i)
			}
		}
		schedule.QueryForSeeSchedule(responses)
	}
}

func RegToGamesAct(counter, i int, responses chan *types.TelegramResponse) {
	if counter == rSTART {
		registration.QueryForPresentationScheduele(responses)
		routine.DeleteUser(userIdT2)
		database.DeleteGame(gameId)
	} else if counter == rLEVEL1 {
		ChooseGame(i, responses)
	} else if counter == rLEVEL2 {
		ChooseSeats(i, responses)
	} else if counter == rLEVEL3 {
		ChoosePayment(i, responses)
	} else if counter == rLEVEL4 {
		BestWishes(i, responses)
	}
}
