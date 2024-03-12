package media

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/formatter"
	"fmt"
	"strconv"
)

func createScheduleKeyboard(user *bottypes.User, fm *formatter.Formatter, schedule []*forall.Game, pretext string) {
	var (
		dict           map[string]string
		kbName, kbData []string
		coordinates    []int
	)
	dict = dictionary.Dictionary[user.Language]
	for i := 0; schedule[i] != nil && i < len(schedule); i++ {
		kbName = []string{fmt.Sprintf("%s %s %s", dict[schedule[i].Sport], schedule[i].Date, schedule[i].Time)}
		kbData = []string{fmt.Sprint(schedule[i].Id)}
		coordinates = append(coordinates, 1)
	}
	forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
	fm.WriteString(fmt.Sprint(pretext, dict["ChooseAnyGame"]))
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
			forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
			fm.WriteChatId(user.Id)
		} else if unloadgames != 0 && uploadgames == 0 {
			user.Request = "unload"
			ChooseMediaGame(user, fm, dict["UnloadGames"])
		} else if unloadgames == 0 && uploadgames != 0 {
			user.Request = "upload"
			ChooseMediaGame(user, fm, dict["UploadGames"])
		}
	} else {
		forall.GoToMainMenu(user, fm, dict["WeDon'tHaveAnyGames"])
	}
}

func ChooseMediaGame(user *bottypes.User, fm *formatter.Formatter, pretext string) {
	if user.Request == "unload" {
		user.Media.Direction = user.Request
		user.Level = 3
		schedule := selectGamesInf(user.Media.Limit, user.LaunchPoint, true)
		createScheduleKeyboard(user, fm, schedule, pretext)
	} else if user.Request == "upload" {
		user.Media.Direction = user.Request
		user.Level = 2
		schedule := selectGamesInf(user.Media.Limit, user.LaunchPoint, false)
		createScheduleKeyboard(user, fm, schedule, pretext)
	} else {
		ChooseDirection(user, fm)
	}
}

func WaitingYourMedia(user *bottypes.User, fm *formatter.Formatter) {
	var (
		detected      bool
		gameId, space int
		dict          map[string]string
		ok            bool
	)
	dict = dictionary.Dictionary[user.Language]
	detected, gameId = forall.IntCheck(user.Request)
	if detected {
		if FindMediaGame(gameId) {
			space, ok = howMuchSpace(gameId)
			if ok {
				user.Level = 3
				user.Media.DelGameId = gameId
				forall.SetTheKeyboard(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
				fm.WriteString(fmt.Sprintf(dict["WaitForYourFiles"], space))
				fm.WriteChatId(user.Id)
				fm.WriteDeleteMesId(user.ExMessageId)
			} else {
				forall.SetTheKeyboard(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
				fm.WriteString(dict["NotEnoughSpace"])
				fm.WriteChatId(user.Id)
			}
		} else {
			user.Request = user.Media.Direction
			ChooseMediaGame(user, fm, "")
		}
	} else {
		user.LaunchPoint = forall.CheckPages(user.Request, user.LaunchPoint)
		user.Request = user.Media.Direction
		ChooseMediaGame(user, fm, "")
	}
}

func unload(user *bottypes.User, fm *formatter.Formatter) {
	var (
		detected     bool
		gameId, quan int
		dict         map[string]string
		fid, ty      string
	)
	dict = dictionary.Dictionary[user.Language]
	detected, gameId = forall.IntCheck(user.Request)
	if detected {
		if FindMediaGame(gameId) {
			quan = selectQuantity(gameId)
			if quan > 1 {
				fm.AddMapOfMedia(selectArrOrMedia(gameId)) //this func doesn't have functional
			} else {
				fid, ty = selectOneMedia(gameId)
				if ty == "photo" {
					fm.AddPhotoFromTG(fid)
				} else if ty == "video" {
					fm.AddVideoFromTG(fid)
				}
			}
			user.Level = 4
			user.Media.DelGameId = gameId
			forall.SetTheKeyboard(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
			fm.WriteString(dict["HereYouAre"])
			fm.WriteChatId(user.Id)
		} else {
			user.Request = user.Media.Direction
			ChooseMediaGame(user, fm, "")
		}
	} else {
		user.LaunchPoint = forall.CheckPages(user.Request, user.LaunchPoint)
		user.Request = user.Media.Direction
		ChooseMediaGame(user, fm, "")
	}
}

func upload(user *bottypes.User, fm *formatter.Formatter) {
	var (
		space int
		ok    bool
		dict  map[string]string
	)
	dict = dictionary.Dictionary[user.Language]
	if len(user.Media.Photo) != 0 || len(user.Media.Video) != 0 {
		space, ok = howMuchSpace(user.Media.DelGameId)
		if ok && space > user.Media.Counter {
			user.Level = 4
			fmt.Println(user.Media.Counter)
			if user.Media.Counter > 1 {
				//
			} else {
				if len(user.Media.Photo) != 0 {
					insertOneNewMedia(user.Media.Photo[0], "photo", user.Media.DelGameId, user.Id)
				} else {
					insertOneNewMedia(user.Media.Video[0], "video", user.Media.DelGameId, user.Id)
				}
			}
			forall.SetTheKeyboard(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
			fm.WriteString(fmt.Sprint(dict["Succesful"]))
			fm.WriteChatId(user.Id)
		} else {
			user.Request = strconv.Itoa(user.Media.DelGameId)
			WaitingYourMedia(user, fm)
		}
	} else {
		user.Request = strconv.Itoa(user.Media.DelGameId)
		WaitingYourMedia(user, fm)
	}
}

func UnloadAndUnload(user *bottypes.User, fm *formatter.Formatter) {
	if user.Media.Direction == "unload" {
		unload(user, fm)
	} else if user.Media.Direction == "upload" {
		upload(user, fm)
	}
}

func JustToMenu(user *bottypes.User, fm *formatter.Formatter) {

}
