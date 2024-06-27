package handler

import (
	"Media/api/producer"
	"Media/app/dict"
	"Media/apptype"
	"fmt"
	"strconv"

	"github.com/l1qwie/Fmtogram/formatter"
)

const (
	START  int = 0
	LEVEL1 int = 1
	LEVEL2 int = 2
	LEVEL3 int = 3
	LEVEL4 int = 4
)

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	producer.InterLogs("Start function Media.setKb()",
		fmt.Sprintf("fm (*formatter.Formatter): %v, crd ([]int): %v, names ([]string): %v, data ([]string): %v", fm, crd, names, data))
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// Function directioner to MainMenu
func goToMainMenu(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, mes string) {
	producer.InterLogs("Start function Media.goToMainMenu()",
		fmt.Sprintf("res (*apptype.Response): %v, fm (*formatter.Formatter): %v, mes (string): %s", res, fm, mes))
	res.Level = 3
	res.Act = "divarication"
	res.Status = true
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
	fm.WriteString(mes)
}

// Makes from int date fromat: YYYYMMDD to string format: "DD-MM-YYYY"
func fromIntToStrDate(numberDate int) (date string) {
	var (
		year, month, day       int
		monthString, dayString string
	)
	producer.InterLogs("Start function Media.fromIntToStrDate()",
		fmt.Sprintf("numberDate (int): %d", numberDate))
	year = numberDate / 10000
	month = (numberDate - (year * 10000)) / 100
	day = (numberDate - (year * 10000) - (month * 100))
	if month <= 10 {
		monthString = fmt.Sprintf("%d%d", 0, month)
	}
	if day <= 10 {
		dayString = fmt.Sprintf("%d%d", 0, day)
	}
	if (monthString != "") && (dayString != "") {
		date = fmt.Sprintf("%s-%s-%d", dayString, monthString, year)
	} else if (monthString != "") && (dayString == "") {
		date = fmt.Sprintf("%d-%s-%d", day, monthString, year)
	} else if (monthString == "") && (dayString != "") {
		date = fmt.Sprintf("%s-%d-%d", dayString, month, year)
	} else {
		date = fmt.Sprintf("%d-%d-%d", day, month, year)
	}
	return date
}

// Makes from int time fromat: HHMM to string format: "HH-MM"
func fromIntToStrTime(numberTime int) (time string) {
	var (
		hour, minute             int
		hourString, minuteString string
	)
	producer.InterLogs("Start function Media.fromIntToStrTime()",
		fmt.Sprintf("numberTime (int): %d", numberTime))
	hour = numberTime / 100
	minute = (numberTime - (hour * 100))
	if hour < 10 {
		hourString = fmt.Sprintf("%d%d", 0, hour)
	}
	if minute < 10 {
		minuteString = fmt.Sprintf("%d%d", 0, minute)
	}
	if (hourString != "") && (minuteString != "") {
		time = fmt.Sprintf("%s:%s", hourString, minuteString)
	} else if (hourString != "") && (minuteString == "") {
		time = fmt.Sprintf("%s:%d", hourString, minute)
	} else if (hourString == "") && (minuteString != "") {
		time = fmt.Sprintf("%d:%s", hour, minuteString)
	} else if (hourString == "") && (minuteString == "") {
		time = fmt.Sprintf("%d:%d", hour, minute)
	}
	return time
}

// Checkes is there an int number or not
// If yes - returns true and the number
// Else - returns only false (and 0 in var num)
func intCheck(phrase string) (bool, int) {
	producer.InterLogs("Start function Media.intCheck()",
		fmt.Sprintf("phrase (string): %s", phrase))
	num, err := strconv.Atoi(phrase)
	return err == nil, num
}

// Checkes is it + or - to launch point and then changed if it's true
func increaseLaunchPoit(phrase string) (lp int) {
	producer.InterLogs("Start function Media.increaseLaunchPoit()",
		fmt.Sprintf("phrase (string): %s", phrase))
	if phrase == "next page" {
		lp = 7
	} else if phrase == "previous page" {
		lp = -7
	}
	return lp
}

// Only if you have games to upload and unload
func unAndUp(res *apptype.Response, fm *formatter.Formatter, dict map[string]string) {
	producer.InterLogs("Start function Media.unAndUp()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", res.ChatID, res, fm))
	res.Level = 1
	setKb(fm, []int{1, 1, 1}, []string{dict["unload"], dict["upload"], dict["MainMenu"]}, []string{"unload", "upload", "MainMenu"})
	fm.WriteString(dict["UpAndUn"])
}

// The first function of the act
func chooseDirection(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Media.chooseDirection()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", res.ChatID, res, fm))
	if findAnyMGames(fm.Error) {
		ungames, upgames := findEveryGames(fm.Error)
		if (ungames != 0) && (upgames != 0) {
			unAndUp(res, fm, dict.Dictionary[req.Language])
		} else if (ungames != 0) && (upgames == 0) {
			prepareToUn(res, fm, dict.Dictionary[req.Language], req.Limit, dict.Dictionary[req.Language]["UnloadGames"])
		} else if (ungames == 0) && (upgames != 0) {
			prepareToUp(res, fm, dict.Dictionary[req.Language], req.Limit, dict.Dictionary[req.Language]["UploadGames"])
		}
	} else {
		goToMainMenu(res, fm, dict.Dictionary[req.Language], fmt.Sprint(dict.Dictionary[req.Language]["NoMGames"], dict.Dictionary[req.Language]["MainMenu"]))
	}
}

// Makes a response body (with the schedule in the keyboard)
func schKb(fm *formatter.Formatter, sch []*apptype.Game, dict map[string]string, text string) {
	var (
		crd         []int
		names, data []string
	)
	producer.InterLogs("Start function Media.schKb()",
		fmt.Sprintf("fm (*formatter.Formatter): %v, sch ([]*apptype.Game): %v, text (string): %s", fm, sch, text))
	for i := 0; sch[i] != nil && i < len(sch); i++ {
		names = append(names, fmt.Sprintf("%s %s %s", dict[sch[i].Sport], sch[i].Date, sch[i].Time))
		data = append(data, fmt.Sprint(sch[i].Id))
		crd = append(crd, 1)
	}
	names = append(names, dict["MainMenu"])
	data = append(data, "MainMenu")
	crd = append(crd, 1)
	setKb(fm, crd, names, data)
	fm.WriteString(fmt.Sprint(text, dict["ChooseAnyGame"]))
}

// Does a few changes and redirect to response builder
func prepareToUn(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, limit int, text string) {
	producer.InterLogs("Start function Media.prepareToUn()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, limit (int): %d, text (string): %s", res.ChatID, res, fm, limit, text))
	res.Level = 3
	res.MediaDir = "unload"
	sch := selectGamesInf(limit, res.LaunchPoint, true, fm.Error)
	schKb(fm, sch, dict, text)
}

// Does a few changes and redirect to response builder
func prepareToUp(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, limit int, text string) {
	producer.InterLogs("Start function Media.prepareToUp()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, limit (int): %d, text (string): %s", res.ChatID, res, fm, limit, text))
	res.Level = 2
	res.MediaDir = "upload"
	sch := selectGamesInf(limit, res.LaunchPoint, false, fm.Error)
	schKb(fm, sch, dict, text)
}

// Redirects the request depending on the input data
func chooseMediaGame(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Media.chooseMediaGame()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Req == "unload" {
		prepareToUn(res, fm, dict.Dictionary[req.Language], req.Limit, "")
	} else if req.Req == "upload" {
		prepareToUp(res, fm, dict.Dictionary[req.Language], req.Limit, "")
	} else {
		unAndUp(res, fm, dict.Dictionary[req.Language])
	}
}

// Makes a body for a response
func waitBody(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, gameId int) {
	producer.InterLogs("Start function Media.waitBody()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, gameId (int): %d", res.ChatID, res, fm, gameId))
	space, ok := howMuchSpace(gameId, fm.Error)
	if ok {
		res.Level = 3
		res.GameId = gameId
		setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
		fm.WriteString(fmt.Sprintf(dict["WaitForYourFiles"], space))
		//Delete message should be here
	} else {
		goToMainMenu(res, fm, dict, dict["NotEnoughSpace"])
	}
}

// Checks a few moments and redirects to a function that make response-body
func waitingYourMedia(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Media.waitingYourMedia()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	det, num := intCheck(req.Req)
	if det {
		if findMediaGame(num, fm.Error) {
			waitBody(res, fm, dict.Dictionary[req.Language], num)
		} else {
			prepareToUp(res, fm, dict.Dictionary[req.Language], req.Limit, "")
		}
	} else {
		res.LaunchPoint = increaseLaunchPoit(req.Req)
		prepareToUp(res, fm, dict.Dictionary[req.Language], req.Limit, "")
	}
}

// Finds out what media files are there and chooses the path to transfer them to the variable
func whichWillbeShowed(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, gameId int) {
	producer.InterLogs("Start function Media.whichWillbeShowed()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, gameId (int): %d", res.ChatID, res, fm, gameId))
	res.Level = 4
	res.GameId = gameId
	res.Status = true
	quan := selectQuantity(gameId, fm.Error)
	if quan > 1 {
		fm.AddMapOfMedia(selectArrOrMedia(gameId, quan, fm.Error))
	} else {
		fid, ty := selectOneMedia(gameId, fm.Error)
		if ty == "photo" {
			fm.AddPhotoFromTG(fid)
		} else if ty == "video" {
			fm.AddVideoFromTG(fid)
		}
	}
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
	fm.WriteString(dict["HereYouAre"])
	producer.ActLogs("User has completed media action", res.ChatID)
}

// Unload function, checks a little and redicrect to bodybuilder
func unload(res *apptype.Response, fm *formatter.Formatter, dict map[string]string, phrase string, limit int) {
	producer.InterLogs("Start function Media.unload()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v, phrase (string): %s, limit (int): %d", res.ChatID, res, fm, phrase, limit))
	det, num := intCheck(phrase)
	if det {
		if findMediaGame(num, fm.Error) {
			whichWillbeShowed(res, fm, dict, num)
		} else {
			prepareToUn(res, fm, dict, limit, "")
		}
	} else {
		res.LaunchPoint = increaseLaunchPoit(phrase)
		prepareToUn(res, fm, dict, limit, "")
	}
}

// Saves the media, checks which way would it be
func saveMedia(res *apptype.Response, req *apptype.Request, fm *formatter.Formatter, dict map[string]string) {
	var mes string
	producer.InterLogs("Start function Media.saveMedia()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, req (*apptype.Request): %v, fm (*formatter.Formatter): %v", req.Id, res, req, fm))
	space, ok := howMuchSpace(res.GameId, fm.Error)
	if ok && (space > req.MediaCounter) {
		res.Level = 4
		res.Status = true
		if req.MediaCounter > 1 {
			ok = insertAfewNewMedia(req.Media, req.GameId, req.Id, fm.Error)
		} else {
			ok = insertOneNewMedia(req.FileId, req.TypeOffile, req.GameId, req.Id, fm.Error)
		}
		setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
		if ok {
			mes = dict["Succesful"]
		} else {
			mes = dict["Oops!"]
		}
		fm.WriteString(mes)
		producer.ActLogs("User has completed media action", res.ChatID)
	} else {
		waitBody(res, fm, dict, res.GameId)
	}
}

// Checks if there is enough data and redirect to function which save the media
func upload(res *apptype.Response, req *apptype.Request, fm *formatter.Formatter, dict map[string]string) {
	producer.InterLogs("Start function Media.upload()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, req (*apptype.Request): %v, fm (*formatter.Formatter): %v", req.Id, res, req, fm))
	if ((req.FileId != "") && (req.TypeOffile != "")) || (req.Media != nil) {
		saveMedia(res, req, fm, dict)
	} else {
		waitBody(res, fm, dict, res.GameId)
	}
}

// Directs to unload function or upload
func unloadAndUpload(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Media.unloadAndUnload()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.MediaDir == "unload" {
		unload(res, fm, dict.Dictionary[req.Language], req.Req, req.Limit)
	} else if req.MediaDir == "upload" {
		upload(res, req, fm, dict.Dictionary[req.Language])
	}
}

// Not the main directioner, but also.
func dir(req *apptype.Request, res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Media.MediaAct()",
		fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", req.Id, req, res, fm))
	if req.Level == START {
		if req.Status {
			producer.ActLogs("The user has started media action", req.Id)
			res.Status = false
		}
		chooseDirection(req, res, fm)
	} else if req.Level == LEVEL1 {
		chooseMediaGame(req, res, fm)
	} else if req.Level == LEVEL2 {
		waitingYourMedia(req, res, fm)
	} else if req.Level == LEVEL3 {
		unloadAndUpload(req, res, fm)
	} else if req.Level == LEVEL4 {
		goToMainMenu(res, fm, dict.Dictionary[req.Language], dict.Dictionary[req.Language]["MainMenu"])
	}
}

// Transfers data changed during the program to the response variable
func media(res *apptype.Response, fm *formatter.Formatter) {
	producer.InterLogs("Start function Media.unloadAndUnload()",
		fmt.Sprintf("UserId: %d, res (*apptype.Response): %v, fm (*formatter.Formatter): %v", res.ChatID, res, fm))
	if res.MediaDir == "unload" {
		if fm.Message.Photo != "" {
			res.FileId = fm.Message.Photo
			res.TypeOffile = "photo"
		} else if fm.Message.Video != "" {
			res.FileId = fm.Message.Video
			res.TypeOffile = "video"
		} else if len(fm.Message.InputMedia) != 0 {
			res.Media = fm.Message.InputMedia
		}
	}
}

// The directioner.
// Only this func is imported
func MediaAct(req *apptype.Request, res *apptype.Response) {
	producer.InterLogs("Start function Media.MediaAct()", fmt.Sprintf("UserId: %d, req (*apptype.Request): %v, res (*apptype.Response): %v", req.Id, req, res))
	fm := new(formatter.Formatter)
	apptype.Db = apptype.ConnectToDatabase()
	res.Level = req.Level
	res.LaunchPoint = req.LaunchPoint
	res.Act = req.Act
	res.MediaDir = req.MediaDir
	res.GameId = req.GameId
	res.Status = req.Status
	dir(req, res, fm)
	fm.ReadyKB()
	res.Keyboard = fm.Message.ReplyMarkup
	res.Message = fm.Message.Text
	res.ChatID = req.Id
	res.ParseMode = fm.Message.ParseMode
	media(res, fm)
	if fm.Err != nil {
		res.Error = fmt.Sprint(fm.Err)
	}
}
