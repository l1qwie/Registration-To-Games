package tests

import (
	"fmt"
	"log"
	rootmedia "registrationtogames/bot/media"
	rootregistration "registrationtogames/bot/registration"
	root "registrationtogames/bot/routine"
	"registrationtogames/fmtogram"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/fmtogram/types"
	"registrationtogames/tests/database"
	"registrationtogames/tests/media"
	preparationdata "registrationtogames/tests/preparationData"
	"registrationtogames/tests/registration"
	"registrationtogames/tests/routine"
	"registrationtogames/tests/schedule"
	"registrationtogames/tests/welcome"
)

const (
	userIdT1     int    = 456
	userIdT2     int    = 477
	userIdT3     int    = 488
	userIdT4     int    = 499
	gameId       int    = 2
	pastgameId   int    = 10
	seats        int    = 2
	payment      string = "card"
	Welcome      int    = 3
	RegToGames   int    = 8
	ShowSchedule int    = 8
	Media        int    = 9
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
	mStart       int    = 9
	mLEVEL1      int    = 10
	mLEVEL2      int    = 11
	mLEVEL3      int    = 12
	mLEVEL4      int    = 13
	START        int    = 0
	LEVEL1       int    = 1
	LEVEL2       int    = 2
	LEVEL3       int    = 3
	LEVEL4       int    = 4
	LEVEL5       int    = 5
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
	if !rootregistration.FindAGame(gameId) {
		err := database.CreateGame(gameId)
		if err != nil {
			panic(err)
		}
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

func preparationDatabaseForMediaUnload(counter int) {
	var err error
	if !rootmedia.FindMediaGame(pastgameId) {
		err = database.CreateNotFullMediaGame(pastgameId)
		if err != nil {
			panic(err)
		}
	}
	if !root.Find(userIdT4) {
		err = root.CreateUser(userIdT4, "ru")
		if err != nil {
			panic(err)
		}
	}
	if counter == START {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(0, userIdT4)
	} else if counter == LEVEL1 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(1, userIdT4)
	} else if counter == LEVEL2 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(2, userIdT4)
	}
}

func preparationDatabaseForMediaUpload(counter int) {
	var err error
	if !rootmedia.FindMediaGame(pastgameId) {
		err = database.CreateEmptyMediaGame(pastgameId)
		if err != nil {
			panic(err)
		}
	}
	if !root.Find(userIdT4) {
		err = root.CreateUser(userIdT4, "ru")
		if err != nil {
			panic(err)
		}
	}
	if counter == START {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(0, userIdT4)
	} else if counter == LEVEL1 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(1, userIdT4)
	} else if counter == LEVEL2 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(2, userIdT4)
	} else if counter == LEVEL3 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(3, userIdT4)
	} else if counter == LEVEL4 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(4, userIdT4)
	}
}

func PreparationDatabase(counter int) {
	if counter < Welcome {
		PreparationDatabaseForWelcome(counter)
	} else if counter >= Welcome && counter < RegToGames {
		PreparationDatabaseForRegToGames(counter)
	} else if counter >= ShowSchedule && counter < Media {
		PreparationDatabaseForSchedule(counter)
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

func acceptanceOfResOfUnloadMedia(output *formatter.Formatter, counter, i int) {
	if counter == START {
		media.ChooseDirection(output)
	} else if counter == LEVEL1 {
		if i < wrongAnswers {
			media.ChooseDirection(output)
		} else {
			media.ChooseMediaGameUnload(output)
		}
	} else if counter == LEVEL2 {
		if i < wrongAnswers {
			media.ChooseMediaGameUnload(output)
		} else {
			media.Unload(output)
		}
	}
}

func acceptanceOfResOfUploadMedia(output *formatter.Formatter, counter, i int) {
	if counter == START {
		media.ChooseDirection(output)
	} else if counter == LEVEL1 {
		if i < wrongAnswers {
			media.ChooseDirection(output)
		} else {
			media.ChooseMediaGameUpload(output)
		}
	} else if counter == LEVEL2 {
		if i < wrongAnswers {
			media.ChooseMediaGameUpload(output)
		} else {
			media.WaitingYourMedia(output)
		}
	} else if counter == LEVEL3 {
		if i < wrongAnswers {
			media.WaitingYourMedia(output)
		} else {
			media.Upload(output)
		}
	} else if counter == LEVEL4 {
		if i < wrongAnswers {
			media.Upload(output)
		} else {
			media.CheckUpload(output)
		}
	}
}

func Unload(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.MediaUnload(i, j, response, output)
			preparationDatabaseForMediaUnload(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			acceptanceOfResOfUnloadMedia(r, i, j)
		}
		fmt.Printf("UploadTest %d has been complete", i)
	}
	fmt.Println("Whole UnloadTest has been complete")
}

func Upload(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.MediaUpload(i, j, response, output)
			preparationDatabaseForMediaUpload(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			acceptanceOfResOfUploadMedia(r, i, j)
		}
		fmt.Printf("UploadTest %d has been complete", i)
	}
	fmt.Println("Whole UploadTest has been complete")
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
	} else if counter >= ShowSchedule && counter < Media {
		AcceptanceOfResOfSchedule(output, counter)
	}
}
