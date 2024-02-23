package tests

import (
	"log"
	root "registrationtogames/bot/routine"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/tests/database"
	"registrationtogames/tests/registration"
	"registrationtogames/tests/routine"
	"registrationtogames/tests/welcome"
)

const (
	userId       int    = 456
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
	wrongAnswers int    = 2
)

func PreparationDatabaseForWelcome(counter int) {
	if counter == wSTART {
		database.UpdateLanguage("ru", userId)
		database.UpdateAction("registration", userId)
		database.UpdateLevel(0, userId)
	} else if counter == wLEVEL1 {
		database.UpdateLanguage("ru", userId)
		database.UpdateAction("registration", userId)
		database.UpdateLevel(1, userId)
	} else if counter == wLEVEL2 {
		database.UpdateLanguage("ru", userId)
		database.UpdateAction("registration", userId)
		database.UpdateLevel(2, userId)
	}
}

func PreparationDatabaseForRegToGames(counter int) {
	err := database.CreateGame()
	if err != nil {
		panic(err)
	}
	if counter == rSTART {
		database.UpdateLanguage("ru", userId)
		database.UpdateAction("reg to games", userId)
		database.UpdateLevel(0, userId)
	} else if counter == rLEVEL1 {
		database.UpdateLanguage("ru", userId)
		database.UpdateAction("reg to games", userId)
		database.UpdateLevel(1, userId)
	} else if counter == rLEVEL2 {
		database.UpdateLanguage("ru", userId)
		database.UpdateAction("reg to games", userId)
		database.UpdateLevel(2, userId)
		database.InsertGameId(gameId, userId)
	} else if counter == rLEVEL3 {
		database.UpdateLanguage("ru", userId)
		database.UpdateAction("reg to games", userId)
		database.UpdateLevel(3, userId)
		database.InsertGameId(gameId, userId)
		database.InsertSeats(seats, userId)
	} else if counter == rLEVEL4 {
		database.UpdateLanguage("ru", userId)
		database.UpdateAction("reg to games", userId)
		database.UpdateLevel(4, userId)
		database.InsertGameId(gameId, userId)
		database.InsertSeats(seats, userId)
		database.InsertPayment(payment, userId)
	}
}

func PreparationDatabase(counter int) {
	if !root.Find(userId) {
		err := root.CreateUser(userId, "ru")
		if err != nil {
			panic(err)
		}
	}
	if counter < Welcome {
		PreparationDatabaseForWelcome(counter)
	} else if counter >= Welcome && counter < RegToGames {
		PreparationDatabaseForRegToGames(counter)
	}
}

func AcceptanceOfResOfWelcome(output *formatter.Formatter, counter, i int) {
	if counter == wSTART {
		welcome.TestGreetingsToUser(output)
		database.AfterGreetingsToUserCheckDb(userId)
	} else if counter == wLEVEL1 {
		if i < wrongAnswers {
			welcome.TestGreetingsToUser(output)
			database.AfterGreetingsToUserCheckDb(userId)
		} else {
			welcome.TestShowRules(output)
			database.AfterShowRulesCheckDb(userId)
		}
	} else if counter == wLEVEL2 {
		if i < wrongAnswers {
			welcome.TestShowRules(output)
			database.AfterShowRulesCheckDb(userId)
		} else {
			welcome.TestWelcomeToMainMenu(output)
			database.AfterMainMenuCheckDb(userId)
		}
	}
}

func AcceptanceOfResOfRegToGames(output *formatter.Formatter, counter, i int) {
	if counter == rSTART {
		registration.PresentationScheduele(output)
		database.AfterPresentationSchedueleCheckDb(userId)
	} else if counter == rLEVEL1 {
		if i < wrongAnswers {
			registration.PresentationScheduele(output)
			database.AfterPresentationSchedueleCheckDb(userId)
		} else {
			registration.ChooseGame(output)
			database.AfterChooseGameCheckDb(userId)
		}
	} else if counter == rLEVEL2 {
		if i < wrongAnswers {
			registration.ChooseGame(output)
			database.AfterChooseGameCheckDb(userId)
		} else {
			registration.ChooseSeats(output)
			database.AfterChooseSeatsCheckDb(userId)
		}
	} else if counter == rLEVEL3 {
		if i < wrongAnswers {
			registration.ChooseSeats(output)
			database.AfterChooseSeatsCheckDb(userId)
		} else {
			registration.ChoosePayment(output)
			database.AfterChoosePaymentCheckDb(userId)
		}
	} else if counter == rLEVEL4 {
		if i < wrongAnswers {
			registration.ChoosePayment(output)
			database.AfterChoosePaymentCheckDb(userId)
		} else {
			registration.BestWishes(output)
			database.AfterBestWishes(userId)
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
