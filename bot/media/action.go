package media

import (
	"fmt"
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/dictionary"
	"registrationtogames/bot/forall"
	"registrationtogames/fmtogram/formatter"
)

func createScheduleKeyboard(user *bottypes.User, fm *formatter.Formatter, schedule []*forall.Game) {
	var (
		dict           map[string]string
		kbName, kbData []string
		coordinates    []int
	)
	dict = dictionary.Dictionary[user.Language]
	for i := 0; i < len(schedule); i++ {
		kbName = []string{fmt.Sprintf("%s %s %s", dict[schedule[i].Sport], schedule[i].Date, schedule[i].Time)}
		kbData = []string{fmt.Sprint(schedule[i].Id)}
		coordinates = append(coordinates, 1)
	}
	forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
	fm.WriteString(dict["ChooseGame"])
	fm.WriteChatId(user.Id)
}

func ChooseDirection(user *bottypes.User, fm *formatter.Formatter) {
	var (
		dict                     map[string]string
		unloadgames, uploadgames int
		kbName, kbData           []string
		coordinates              []int
	)
	dict = dictionary.Dictionary[user.Language]
	if findAnyMGames() {
		user.Level = 1
		unloadgames, uploadgames = findEveryGames()
		if unloadgames != 0 && uploadgames != 0 {
			kbName = []string{dict["unload"], dict["upload"]}
			kbData = []string{"unload", "upload"}
			coordinates = []int{1, 1}
			fm.WriteString(dict["UpAndUn"])
		} else if unloadgames != 0 && uploadgames == 0 {
			kbName = []string{dict["unload"]}
			kbData = []string{"unload"}
			coordinates = []int{1}
			fm.WriteString(dict["UnloadGames"])
		} else if unloadgames == 0 && uploadgames != 0 {
			kbName = []string{dict["upload"]}
			kbData = []string{"upload"}
			coordinates = []int{1}
			fm.WriteString(dict["UploadGames"])
		}
		forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
		fm.WriteChatId(user.Id)
	} else {
		forall.GoToMainMenu(user, fm, dict["WeDon'tHaveAnyGames"])
	}
}

func ChooseMediaGame(user *bottypes.User, fm *formatter.Formatter) {
	if user.Request == "unload" {
		user.Media.Direction = user.Request
		user.Level = 3
		schedule := selectGamesInf(user.Media.Limit, user.LaunchPoint, true)
		createScheduleKeyboard(user, fm, schedule)
	} else if user.Request == "upload" {
		user.Media.Direction = user.Request
		user.Level = 2
		schedule := selectGamesInf(user.Media.Limit, user.LaunchPoint, false)
		createScheduleKeyboard(user, fm, schedule)
	} else {
		ChooseDirection(user, fm)
	}
}

func WaitingYourMedia(user *bottypes.User, fm *formatter.Formatter) {
	var (
		detected bool
		gameId   int
		dict     map[string]string
	)
	dict = dictionary.Dictionary[user.Language]
	detected, gameId = forall.IntCheck(user.Request)
	if detected {
		if FindMediaGame(gameId) {
			user.Level = 3
			user.Media.DelGameId = gameId
			forall.SetTheKeyboard(fm, []int{1, 1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
			fm.WriteString(dict["WaitForYourFiles"])
			fm.WriteChatId(user.Id)
			fm.WriteDeleteMesId(user.ExMessageId)
		} else {
			user.Request = user.Media.Direction
			ChooseMediaGame(user, fm)
		}
	} else {
		forall.CheckPages(user.Request, user.LaunchPoint)
		user.Request = user.Media.Direction
		ChooseMediaGame(user, fm)
	}
}

func unload(user *bottypes.User, fm *formatter.Formatter) {
	var (
		detected bool
		gameId   int
		dict     map[string]string
	)
	dict = dictionary.Dictionary[user.Language]
	detected, gameId = forall.IntCheck(user.Request)
	if detected {
		if FindMediaGame(gameId) {
			user.Level = 4
			user.Media.DelGameId = gameId
			forall.SetTheKeyboard(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
			fm.WriteString(dict["HereYouAre"])
			fm.WriteChatId(user.Id)
		} else {
			user.Request = user.Media.Direction
			ChooseMediaGame(user, fm)
		}
	} else {
		forall.CheckPages(user.Request, user.LaunchPoint)
		user.Request = user.Media.Direction
		ChooseMediaGame(user, fm)
	}
}

func upload(user *bottypes.User, fm *formatter.Formatter) {
	if 
}

func UnloadAndUnload(user *bottypes.User, fm *formatter.Formatter) {
	if user.Media.Direction == "unload" {
		unload(user, fm)
	} else if user.Media.Direction == "upload" {
		upload(user, fm)
	}
}

func CheckUpload(user *bottypes.User, fm *formatter.Formatter) {

}
