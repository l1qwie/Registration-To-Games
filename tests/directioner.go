package tests

import (
	rootmedia "RegistrationToGames/bot/media"
	rootregistration "RegistrationToGames/bot/registration"
	root "RegistrationToGames/bot/routine"
	"RegistrationToGames/fmtogram"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests/database"
	"RegistrationToGames/tests/media"
	preparationdata "RegistrationToGames/tests/preparationData"
	"RegistrationToGames/tests/registration"
	"RegistrationToGames/tests/routine"
	"RegistrationToGames/tests/schedule"
	"RegistrationToGames/tests/settings"
	"RegistrationToGames/tests/welcome"
	"fmt"
)

const (
	userIdT1     int    = 456
	userIdT2     int    = 477
	userIdT3     int    = 488
	userIdT4     int    = 499
	userIdT5     int    = 899
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

func PreparationDatabaseForSchedule() {
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
	database.UpdateLanguage("ru", userIdT3)
	database.UpdateAction("see schedule", userIdT3)
	database.UpdateLevel(0, userIdT3)
}

func prparDbForWelcome(counter int) {
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

func prparDbForRegToGames(counter int) {
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
	if counter == START {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(0, userIdT2)
	} else if counter == LEVEL1 {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(1, userIdT2)
	} else if counter == LEVEL2 {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(2, userIdT2)
		database.InsertGameId(gameId, userIdT2)
	} else if counter == LEVEL3 {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(3, userIdT2)
		database.InsertGameId(gameId, userIdT2)
		database.InsertSeats(seats, userIdT2)
	} else if counter == LEVEL4 {
		database.UpdateLanguage("ru", userIdT2)
		database.UpdateAction("reg to games", userIdT2)
		database.UpdateLevel(4, userIdT2)
		database.InsertGameId(gameId, userIdT2)
		database.InsertSeats(seats, userIdT2)
		database.InsertPayment(payment, userIdT2)
	}
}

func UnloadAnoOpt(counter int) {
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
		database.UpdateLevel(3, userIdT4)
	}
}

func UploadAnoOpt(counter int) {
	if !root.Find(userIdT4) {
		err := root.CreateUser(userIdT4, "ru")
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
		database.UpdateLevel(2, userIdT4)
	} else if counter == LEVEL2 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(3, userIdT4)
	} else if counter == LEVEL3 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(4, userIdT4)
	} else if counter == LEVEL4 {
		database.UpdateLanguage("ru", userIdT4)
		database.UpdateAction("photos and videos", userIdT4)
		database.UpdateLevel(5, userIdT4)
	}
}

func prparDbForUnloadAfew(counter int) {
	if !root.Find(userIdT4) {
		err := root.CreateUser(userIdT4, "ru")
		if err != nil {
			panic(err)
		}
	}
	// only one option
	UnloadAnoOpt(counter)
}

func prparDbForUnloadOne(counter int) {
	if !rootmedia.FindMediaGame(pastgameId) {
		err := database.CreateNotFullMediaGame(pastgameId)
		if err != nil {
			panic(err)
		}
	}
	if !root.Find(userIdT4) {
		err := root.CreateUser(userIdT4, "ru")
		if err != nil {
			panic(err)
		}
	}
	// only one option
	UnloadAnoOpt(counter)
}

func prparDbForUploadOne(counter int) {
	if !rootmedia.FindMediaGame(pastgameId) {
		err := database.CreateEmptyMediaGame(pastgameId)
		if err != nil {
			panic(err)
		}
	}
	UploadAnoOpt(counter)
}

func prparDbForUploadAfew(counter int) {
	UploadAnoOpt(counter)
}

func prparDbForChangeLang(counter int) {
	if !root.Find(userIdT5) {
		err := root.CreateUser(userIdT5, "ru")
		if err != nil {
			panic(err)
		}
	}
	if counter == START {
		database.UpdateLanguage("ru", userIdT5)
		database.UpdateAction("settings", userIdT5)
		database.UpdateLevel(0, userIdT5)
	} else if counter == LEVEL1 {
		database.UpdateLanguage("ru", userIdT5)
		database.UpdateAction("settings", userIdT5)
		database.UpdateLevel(2, userIdT5)
	}
}

func prparDbForDeleteGame(counter int) {
	if !root.Find(userIdT5) {
		err := root.CreateUser(userIdT5, "ru")
		if err != nil {
			panic(err)
		}
	}
	if counter == START {
		database.UpdateLanguage("ru", userIdT5)
		database.UpdateAction("settings", userIdT5)
		database.UpdateLevel(0, userIdT5)
	} else if counter == LEVEL1 {
		database.UpdateLanguage("ru", userIdT5)
		database.UpdateAction("settings", userIdT5)
		database.UpdateLevel(1, userIdT5)
	}
}

func AcceptanceOfResOfSchedule(output *formatter.Formatter) {
	schedule.ShowTheSchedule(output)
	database.AfterShowTheSchedule(userIdT3)
	for i := 0; i < 4; i++ {
		database.DeleteGame(i)
	}
}

func accOfResOfWelcome(output *formatter.Formatter, counter, i int) {
	if counter == wSTART {
		welcome.TestGreetingsToUser(output)
		database.AfterGreetingsToUser(userIdT1)
	} else if counter == wLEVEL1 {
		if i < wrongAnswers {
			welcome.TestGreetingsToUser(output)
			database.AfterGreetingsToUser(userIdT1)
		} else {
			welcome.TestShowRules(output)
			database.AfterShowRules(userIdT1)
		}
	} else if counter == wLEVEL2 {
		if i < wrongAnswers {
			welcome.TestShowRules(output)
			database.AfterShowRules(userIdT1)
		} else {
			welcome.TestWelcomeToMainMenu(output)
			database.AfterWelcomeMainMenu(userIdT1)
		}
	}
}

func accOfResOfRegToGames(output *formatter.Formatter, counter, i int) {
	if counter == START {
		registration.PresentationScheduele(output)
		database.AfterPresentationScheduele(userIdT2)
	} else if counter == LEVEL1 {
		if i < wrongAnswers {
			registration.PresentationScheduele(output)
			database.AfterPresentationScheduele(userIdT2)
		} else {
			registration.ChooseGame(output)
			database.AfterChooseGameR(userIdT2)
		}
	} else if counter == LEVEL2 {
		if i < wrongAnswers {
			registration.ChooseGame(output)
			database.AfterChooseGameR(userIdT2)
		} else {
			registration.ChooseSeats(output)
			database.AfterAfterChooseSeats(userIdT2)
		}
	} else if counter == LEVEL3 {
		if i < wrongAnswers {
			registration.ChooseSeats(output)
			database.AfterAfterChooseSeats(userIdT2)
		} else {
			registration.ChoosePayment(output)
			database.AfterChoosePayment(userIdT2)
		}
	} else if counter == LEVEL4 {
		if i < wrongAnswers {
			registration.ChoosePayment(output)
			database.AfterChoosePayment(userIdT2)
		} else {
			registration.BestWishes(output)
			database.AfterBestWishes(userIdT2)
		}
	}
}

func accOfResOfUnloadAfew(output *formatter.Formatter, counter, i int) {
	if counter == START {
		media.ChooseDirection(output)
		database.AfterChooseDirection(userIdT4)
	} else if counter == LEVEL1 {
		if i < wrongAnswers {
			media.ChooseDirection(output)
			database.AfterChooseDirection(userIdT4)
		} else {
			media.ChooseMediaGamesUnload(output)
			database.AfterChooseMediaGameUnload(userIdT4)
		}
	} else if counter == LEVEL2 {
		if i < wrongAnswers {
			media.ChooseMediaGamesUnload(output)
			database.AfterChooseMediaGameUnload(userIdT4)
		} else {
			media.Unloadall(output)
			database.AfterUnloadAlot(userIdT4)
		}
	}
}
func accOfResOfUnloadOne(output *formatter.Formatter, counter, i int) {
	if counter == START {
		media.ChooseDirection(output)
		database.AfterChooseDirection(userIdT4)
	} else if counter == LEVEL1 {
		if i < wrongAnswers {
			media.ChooseDirection(output)
			database.AfterChooseDirection(userIdT4)
		} else {
			media.ChooseMediaGameUnload(output)
			database.AfterChooseMediaGameUnload(userIdT4)
		}
	} else if counter == LEVEL2 {
		if i < wrongAnswers {
			media.ChooseMediaGameUnload(output)
			database.AfterChooseMediaGameUnload(userIdT4)
		} else {
			media.Unloadone(output)
			database.AfterUnloadone(userIdT4)
		}
	}
}

func accOfResOfUploadOne(output *formatter.Formatter, counter, i int) {
	if counter == START {
		media.NoChoiseOnlyUploadOne(output)
		//database.AfterNoChoiseOnlyUpload(userIdT4)
	} else if counter == LEVEL1 {
		if i < wrongAnswers {
			media.ChooseGameOne(output)
			database.AfterChooseGame(userIdT4)
		} else {
			media.WaitingYourMedia(output)
			database.AfterWaitingYourMediaOne(userIdT4)
		}
	} else if counter == LEVEL2 {
		if i < wrongAnswers {
			media.WaitingYourMedia(output)
			database.AfterChooseMediaGameUploadOne(userIdT4)
		} else {
			media.Upload(output)
			database.AfterUploadOne(userIdT4)
		}
	}
}

func accOfResOfUploadAfew(output *formatter.Formatter, counter, i int) {
	if counter == START {
		media.NoChoiseOnlyUploadAfew(output)
	} else if counter == LEVEL1 {
		if i < wrongAnswers {
			media.ChooseGameAfew(output)
			database.AfterChooseGame(userIdT4)
		} else {
			media.WaitingYourMedia(output)
			database.AfterWaitingYourMediaAfew(userIdT4)
		}
	} else if counter == LEVEL2 {
		if i < wrongAnswers {
			media.WaitingYourMedia(output)
			database.AfterChooseMediaGameUploadAfew(userIdT4)
		} else {
			media.Upload(output)
			database.AfterUploadAfew(userIdT4)
		}
	}
}

func accOfResOfChangeLanguage(output *formatter.Formatter, counter, i int) {
	if counter == START {
		settings.ChooseOptionOnlyLang(output)
		database.AfterChooseOptionOnlyLang(userIdT5)
	} else if counter == LEVEL1 {
		if i < wrongAnswers {
			settings.ChooseOptionOnlyLang(output)
			database.AfterChooseOptionOnlyLang(userIdT5)
		} else {
			settings.ChooseLanguage(output)
			database.AfterChooseLanguage(userIdT5)
		}
	}
}

func accOfResOfDeleteGame(output *formatter.Formatter, counter, i int) {

}

func UnloadOne(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	defer database.DeleteUser(userIdT4)
	defer database.DeleteGame(pastgameId)
	defer database.DeleteMedia(pastgameId)
	//arrfunc := []func() *types.TelegramResponse{media.ChDir, media.ChMUnload, media.UnlOne}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.MediaUnloadOne(i, j, response, output)
			prparDbForUnloadOne(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			accOfResOfUnloadOne(r, i, j)
		}
		fmt.Printf("UploadTest %d has been complete\n", i)
	}
}

func UploadOne(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	defer database.DeleteUser(userIdT4)
	defer database.DeleteGame(pastgameId)
	defer database.DeleteMedia(pastgameId)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.MediaUploadOne(i, j, response, output)
			prparDbForUploadOne(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			accOfResOfUploadOne(r, i, j)
		}
		fmt.Printf("UploadTest %d has been complete\n", i)
	}
}

func UnloadAfew(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	defer database.DeleteUser(userIdT4)
	defer database.DeleteSchedule()
	defer database.DeleteMediaSchedule()
	database.CreateMediaShedule()
	database.FillMEdiaSchedule()
	arrfunc := []func() *types.TelegramResponse{media.ChDir, media.ChMUnload, media.UnlAfew}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.MediaUnloadAfew(i, j, response, output, arrfunc)
			prparDbForUnloadAfew(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			accOfResOfUnloadAfew(r, i, j)
		}
		fmt.Printf("UploadTest %d has been complete\n", i)
	}
	fmt.Println("Whole UnloadTest has been complete")
}

func UploadAfew(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	defer database.DeleteUser(userIdT4)
	defer database.DeleteSchedule()
	defer database.DeleteMediaSchedule()
	database.CreateMediaShedule()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.MediaUploadAfew(i, j, response, output)
			prparDbForUploadAfew(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			accOfResOfUploadAfew(r, i, j)
		}
		fmt.Printf("UploadTest %d has been complete\n", i)
	}
	fmt.Println("Whole UnloadTest has been complete")
}

func CheckTheRoutine() {
	routine.TestRetrevenUser()
	routine.TestRetainUser()
}

func JustWelcome(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	defer database.DeleteUser(userIdT1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.WelcomeAct(i, j, response, output)
			prparDbForWelcome(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			accOfResOfWelcome(r, i, j)
		}
		fmt.Printf("WelcomeTest %d has been complete\n", i)
	}
}

func RegToOneGame(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	defer database.DeleteUser(userIdT2)
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.RegToGamesAct(i, j, response, output)
			prparDbForRegToGames(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			accOfResOfRegToGames(r, i, j)
		}
		fmt.Printf("RegToGamesTest %d has been complete\n", i)
	}
}

func ChangeLanguage(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	defer database.DeleteUser(userIdT5)
	ts := new(preparationdata.TestStuct)
	ts.Query = make(chan *types.TelegramResponse, 1)
	ts.Response = make(chan *types.MessageResponse, 1)
	ts.ArrFuncTr = []func() *types.TelegramResponse{settings.ChOpt, settings.ChLang}
	ts.ArrFuncMr = []func() *types.MessageResponse{settings.MesResp}
	ts.ArrFuncTrash = []func() *types.TelegramResponse{settings.Trash, settings.Trash2}
	for ts.TRcount < 2 {
		ts.Trshcount = 0
		for ts.Trshcount < 3 {
			ts.TheHead()
			//preparationdata.SetActChanLang(i, j, response, output, arrfunc)
			prparDbForChangeLang(ts.TRcount)
			fmtogram.Worker(ts.Query, ts.Response, request)
			r := <-request
			accOfResOfChangeLanguage(r, ts.TRcount, ts.Trshcount)
			ts.Trshcount++
		}
		fmt.Printf("ChangeLanguageTest %d has been complete\n", ts.TRcount)
		ts.TRcount++
	}
}

func DeleteGame(response chan *types.TelegramResponse, request chan *formatter.Formatter, output chan *types.MessageResponse) {
	defer database.DeleteUser(userIdT5)
	defer database.DeleteSchedule()
	defer database.DeleteUserGames(userIdT5)
	database.CreateUserSchedule(userIdT5)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			preparationdata.SetActDeleteGame(i, j, response, output)
			prparDbForDeleteGame(i)
			fmtogram.Worker(response, output, request)
			r := <-request
			accOfResOfDeleteGame(r, i, j)
		}
		fmt.Printf("DeleteGameTest %d has been complete\n", i)
	}
}
