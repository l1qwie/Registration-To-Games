package tests

import (
	"log"
	root "registrationtogames/bot/routine"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/tests/database"
	"registrationtogames/tests/registration"
	"registrationtogames/tests/routine"
	"registrationtogames/tests/schedule"
	"registrationtogames/tests/welcome"
)

const (
	userIdT1     int    = 456
	userIdT2     int    = 477
	userIdT3     int    = 488
	gameId       int    = 2
	seats        int    = 2
	payment      string = "card"
	Welcome      int    = 3
	RegToGames   int    = 8
	wSTART       int    = 0
	wLEVEL1      int    = 1
	wLEVEL2      int    = 2
	rSTART       int    = 3
	rLEVEL1      int    = 4
	rLEVEL2      int    = 5
	rLEVEL3      int    = 6
	rLEVEL4      int    = 7
	schSTART     int    = 8
	wrongAnswers int    = 2
)

func PreparationDatabaseForSchedule(counter int) {
	if !root.Find(userIdT3) {
		err := root.CreateUser(userIdT3, "ru")
		if err != nil {
			panic(err)
		}
	}
	err := database.CreateSchedule()
	if err != nil {
		panic(err)
	}
	if counter == schSTART {
		database.UpdateLanguage("ru", userIdT3)
		database.UpdateAction("see schedule", userIdT3)
		database.UpdateLevel(0, userIdT3)
	}
}

func PreparationDatabaseForWelcome(counter int) {
	if !root.Find(userIdT1) {
		err := root.CreateUser(userIdT1, "ru")
		if err != nil {
			panic(err)
		}
	}
	if counter == wSTART {
		database.UpdateLanguage("ru", userIdT1)
		database.UpdateAction("registration", userIdT1)
		database.UpdateLevel(0, userIdT1)
	} else if counter == wLEVEL1 {
		database.UpdateLanguage("ru", userIdT1)
		database.UpdateAction("registration", userIdT1)
		database.UpdateLevel(1, userIdT1)
	} else if counter == wLEVEL2 {
		database.UpdateLanguage("ru", userIdT1)
		database.UpdateAction("registration", userIdT1)
		database.UpdateLevel(2, userIdT1)
	}
}

func PreparationDatabaseForRegToGames(counter int) {
	err := database.CreateGame()
	if err != nil {
		panic(err)
	}
	if !root.Find(userIdT2) {
		err := root.CreateUser(userIdT2, "ru")
		if err != nil {
			panic(err)
		}
	}
	if counter == rSTART {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(0, userIdT2)
	} else if counter == rLEVEL1 {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(1, userIdT2)
	} else if counter == rLEVEL2 {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(2, userIdT2)
		database.InsertGameId(gameId, userIdT2)
	} else if counter == rLEVEL3 {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(3, userIdT2)
		database.InsertGameId(gameId, userIdT2)
		database.InsertSeats(seats, userIdT2)
	} else if counter == rLEVEL4 {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(4, userIdT2)
		database.InsertGameId(gameId, userIdT2)
		database.InsertSeats(seats, userIdT2)
		database.InsertPayment(payment, userIdT2)
	}
}

func PreparationDatabase(counter int) {
	if counter < Welcome {
		PreparationDatabaseForWelcome(counter)
	} else if counter >= Welcome && counter < RegToGames {
		PreparationDatabaseForRegToGames(counter)
	}
}

func AcceptanceOfResOfSchedule(output *formatter.Formatter, counter int) {
	if counter == schSTART {
		schedule.ShowTheSchedule(output)
		database.AfterShowTheSchedule(userIdT3)
	}
	for i := 0; i < 4; i++ {
		database.DeleteGame(i)
	}
}

func AcceptanceOfResOfWelcome(output *formatter.Formatter, counter, i int) {
	if counter == wSTART {
		welcome.TestGreetingsToUser(output)
		database.AfterGreetingsToUserCheckDb(userIdT1)
	} else if counter == wLEVEL1 {
		if i < wrongAnswers {
			welcome.TestGreetingsToUser(output)
			database.AfterGreetingsToUserCheckDb(userIdT1)
		} else {
			welcome.TestShowRules(output)
			database.AfterShowRulesCheckDb(userIdT1)
		}
	} else if counter == wLEVEL2 {
		if i < wrongAnswers {
			welcome.TestShowRules(output)
			database.AfterShowRulesCheckDb(userIdT1)
		} else {
			welcome.TestWelcomeToMainMenu(output)
			database.AfterMainMenuCheckDb(userIdT1)
		}
	}
}

func AcceptanceOfResOfRegToGames(output *formatter.Formatter, counter, i int) {
	if counter == rSTART {
		registration.PresentationScheduele(output)
		database.AfterPresentationSchedueleCheckDb(userIdT2)
	} else if counter == rLEVEL1 {
		if i < wrongAnswers {
			registration.PresentationScheduele(output)
			database.AfterPresentationSchedueleCheckDb(userIdT2)
		} else {
			registration.ChooseGame(output)
			database.AfterChooseGameCheckDb(userIdT2)
		}
	} else if counter == rLEVEL2 {
		if i < wrongAnswers {
			registration.ChooseGame(output)
			database.AfterChooseGameCheckDb(userIdT2)
		} else {
			registration.ChooseSeats(output)
			database.AfterChooseSeatsCheckDb(userIdT2)
		}
	} else if counter == rLEVEL3 {
		if i < wrongAnswers {
			registration.ChooseSeats(output)
			database.AfterChooseSeatsCheckDb(userIdT2)
		} else {
			registration.ChoosePayment(output)
			database.AfterChoosePaymentCheckDb(userIdT2)
		}
	} else if counter == rLEVEL4 {
		if i < wrongAnswers {
			registration.ChoosePayment(output)
			database.AfterChoosePaymentCheckDb(userIdT2)
		} else {
			registration.BestWishes(output)
			database.AfterBestWishes(userIdT2)
		}
	}
}

func CheckTheRoutine() {
	routine.TestRetrevenUser()
	routine.TestRetainUser()
}

func AcceptanceOfResults(output *formatter.Formatter, counter, i int) {
	if output == nil {
		log.Fatal()
	}
	//CheckTheRoutine()

	if counter < Welcome {
		AcceptanceOfResOfWelcome(output, counter, i)
	} else if counter >= Welcome && counter < RegToGames {
		AcceptanceOfResOfRegToGames(output, counter, i)
	}
}
