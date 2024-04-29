package handler

import (
	"User/app/dict"
	"User/apptype"
	"User/fmtogram/formatter"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	START   int = 0
	LEVEL1  int = 1
	LEVEL2  int = 2
	LEVEL3  int = 3
	OPTIONS int = 3
	LEVEL4  int = 4
	LEVEL5  int = 5
	LEVEL6  int = 6
	LEVEL7  int = 7
	LEVEL8  int = 8
)

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

func retrieveUser(user *apptype.User, fm *formatter.Formatter) {
	if find(user.Id, fm.Error) {
		dbRetrieveUser(user, fm.Error)
	} else {
		createUser(user.Id, user.Language, fm.Error)
		user.Act = "registration"
		user.Media.Limit = 7
	}
}

func retainUser(user *apptype.User, fm *formatter.Formatter) {
	dbRetainUser(user, fm.Error)
}

func edit(user *apptype.User, fm *formatter.Formatter) (exMessageId int) {
	if user.ExMessageId == 0 {
		exMessageId = SelectExMessageId(user.Id, fm.Error)
	} else {
		exMessageId = user.ExMessageId
		updateExMessageId(exMessageId, user.Id, fm.Error)
	}
	if user.Photo == "" && user.Video == "" {
		fm.WriteEditMesId(exMessageId)
	} else {
		fm.WriteDeleteMesId(exMessageId)
	}
	return exMessageId
}

func mainMenu(user *apptype.User, fm *formatter.Formatter, dict map[string]string) {
	user.Act = "divarication"
	user.Level = OPTIONS
	setKb(fm, []int{1, 1, 1, 1}, []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}, []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"})
	fm.WriteString(dict["MainMenu"])
	fm.WriteChatId(user.Id)
}

func send(req, res any, path, name string) error {
	jb, err := json.Marshal(req)
	if err == nil {
		resp, err := http.Post(fmt.Sprintf("http://localhost:%s/%s", path, name), "application/json", bytes.NewBuffer(jb))
		if err != nil {
			log.Print("Ошибка при выполнении запроса:", err)
		} else {
			json.NewDecoder(resp.Body).Decode(res)
		}
		defer resp.Body.Close()
	}
	return err
}

func wreq(user *apptype.User, req *apptype.WelcomeReq) {
	req.Id = user.Id
	req.Level = user.Level
	req.Language = user.Language
	req.Req = user.Request
	req.Act = user.Act
}

func welcomeAct(user *apptype.User, fm *formatter.Formatter) {
	req := new(apptype.WelcomeReq)
	wreq(user, req)
	res := new(apptype.WelcomeRes)
	err := send(req, res, "8081", "Welcome")
	if err == nil {
		if res.Error == "" {
			fm.WriteString(res.Message)
			fm.WriteChatId(res.ChatID)
			fm.Message.ReplyMarkup = res.Keyboard
			user.Act = res.Act
			user.Level = res.Level
		} else {
			fm.Error(fmt.Errorf(res.Error))
		}
	} else {
		fm.Error(err)
	}
}

func rreq(user *apptype.User, req *apptype.RegistrationReq) {
	req.Id = user.Id
	req.Level = user.Level
	req.Language = user.Language
	req.Act = user.Act
	req.Limit = user.Media.Limit
	req.LaunchPoint = user.LaunchPoint
	req.Req = user.Request
	req.GameId = user.Reg.GameId
	req.Seats = user.Reg.Seats
	req.Payment = user.Reg.Payment
}

func registrationAct(user *apptype.User, fm *formatter.Formatter) {
	req := new(apptype.RegistrationReq)
	rreq(user, req)
	res := new(apptype.RegistrationRes)
	err := send(req, res, "8094", "Registration")
	if err == nil {
		if res.Error == "" {
			fm.WriteString(res.Message)
			fm.WriteChatId(res.ChatID)
			fm.WriteParseMode(res.ParseMode)
			fm.Message.ReplyMarkup = res.Keyboard
			user.Act = res.Act
			user.Level = res.Level
			user.LaunchPoint = res.LaunchPoint
			user.Reg.GameId = res.GameId
			user.Reg.Seats = res.Seats
			user.Reg.Payment = res.Payment
		} else {
			fm.Error(fmt.Errorf(res.Error))
		}
	} else {
		fm.Error(err)
	}
}

func schreq(user *apptype.User, req *apptype.ScheduleReq) {
	req.Id = user.Id
	req.Act = user.Act
	req.Language = user.Language
}

func scheduleAct(user *apptype.User, fm *formatter.Formatter) {
	req := new(apptype.ScheduleReq)
	schreq(user, req)
	res := new(apptype.ScheduleRes)
	err := send(req, res, "8083", "Schedule")
	if err == nil {
		if res.Error == "" {
			fm.WriteString(res.Message)
			fm.WriteChatId(res.ChatID)
			fm.WriteParseMode(res.ParseMode)
			fm.Message.ReplyMarkup = res.Keyboard
			user.Act = res.Act
			user.Level = res.Level
		} else {
			fm.Error(fmt.Errorf(res.Error))
		}
	} else {
		fm.Error(err)
	}
}

func mreq(user *apptype.User, req *apptype.MediaReq) {
	if user.Media.Photo != "" {
		req.FileId = user.Media.Photo
		req.TypeOffile = "photo"
		req.MediaCounter = 1
	} else if user.Media.Video != "" {
		req.FileId = user.Media.Video
		req.TypeOffile = "video"
		req.MediaCounter = 1
	} else if len(user.Media.MediaF) > 0 {
		req.Media = user.Media.MediaF
		req.MediaCounter = len(user.Media.MediaF)
	}
	req.Id = user.Id
	req.Level = user.Level
	req.Language = user.Language
	req.Act = user.Act
	req.Limit = user.Media.Limit
	req.LaunchPoint = user.LaunchPoint
	req.Req = user.Request
	req.MediaDir = user.Media.Direction
	req.GameId = user.Media.DelGameId
}

func mediaAct(user *apptype.User, fm *formatter.Formatter) {
	req := new(apptype.MediaReq)
	mreq(user, req)
	res := new(apptype.MediaRes)
	err := send(req, res, "8085", "Media")
	if err == nil {
		if res.Error == "" {
			fm.WriteString(res.Message)
			fm.WriteChatId(res.ChatID)
			fm.WriteParseMode(res.ParseMode)
			if res.TypeOffile == "photo" {
				fm.WriteDeleteMesId(user.ExMessageId)
				fm.AddPhotoFromTG(res.FileId)
			} else if res.TypeOffile == "video" {
				fm.WriteDeleteMesId(user.ExMessageId)
				fm.AddVideoFromTG(res.FileId)
			} else if len(res.Media) > 0 {
				fm.WriteDeleteMesId(user.ExMessageId)
				fm.AddMapOfMedia(res.Media)
			}
			log.Print(fm.Message.InputMedia)
			fm.Message.ReplyMarkup = res.Keyboard
			user.Level = res.Level
			user.LaunchPoint = res.LaunchPoint
			user.Act = res.Act
			user.Media.Direction = res.MediaDir
			user.Media.DelGameId = res.GameId
		} else {
			fm.Error(fmt.Errorf(res.Error))
			panic(res.Error)
		}
	} else {
		fm.Error(err)
	}
}

func setreq(user *apptype.User, req *apptype.SettingsReq) {
	req.Id = user.Id
	req.Level = user.Level
	req.Language = user.Language
	req.Act = user.Act
	req.Limit = user.Media.Limit
	req.LaunchPoint = user.LaunchPoint
	req.Req = user.Request
	req.IsChanged = user.UserRec.Changeable
	req.GameId = user.UserRec.GameId
}

func settingsAct(user *apptype.User, fm *formatter.Formatter) {
	req := new(apptype.SettingsReq)
	setreq(user, req)
	res := new(apptype.SettingsRes)
	err := send(req, res, "8089", "Settings")
	if err == nil {
		if res.Error == "" {
			fm.WriteString(res.Message)
			fm.WriteChatId(res.ChatID)
			fm.WriteParseMode(res.ParseMode)
			fm.Message.ReplyMarkup = res.Keyboard
			user.Level = res.Level
			user.LaunchPoint = res.LaunchPoint
			user.Act = res.Act
			user.UserRec.Changeable = res.IsChanged
			user.Language = res.Language
			user.UserRec.GameId = res.GameId
		} else {
			fm.Error(fmt.Errorf(res.Error))
		}
	} else {
		fm.Error(err)
	}
}

func opt(user *apptype.User, fm *formatter.Formatter) {
	if user.Request == "Reg to games" {
		user.Level = START
		user.Act = "reg to games"
		registrationAct(user, fm)
	} else if user.Request == "Looking Schedule" {
		user.Level = START
		user.Act = "see schedule"
		scheduleAct(user, fm)
	} else if user.Request == "Photo&Video" {
		user.Level = START
		user.Act = "photos and videos"
		mediaAct(user, fm)
	} else if user.Request == "My records" {
		user.Level = START
		user.Act = "settings"
		settingsAct(user, fm)
	} else {
		mainMenu(user, fm, dict.Dictionary[user.Language])
	}
}

func DispatcherPhrase(user *apptype.User, fm *formatter.Formatter) {
	retrieveUser(user, fm)
	user.ExMessageId = edit(user, fm)
	log.Printf(`DispatcherPhrase basic data:				
			user.Level = %d, user.Request = %s, user.Act = %s user.Id = %d`, user.Level, user.Request, user.Act, user.Id)
	if user.Request == "MainMenu" {
		mainMenu(user, fm, dict.Dictionary[user.Language])
	} else if user.Act == "registration" {
		welcomeAct(user, fm)
	} else if user.Act == "reg to games" {
		registrationAct(user, fm)
	} else if user.Act == "divarication" {
		opt(user, fm)
	} else if user.Act == "see schedule" {
		scheduleAct(user, fm)
	} else if user.Act == "photos and videos" {
		mediaAct(user, fm)
	} else if user.Act == "settings" {
		settingsAct(user, fm)
	}
	retainUser(user, fm)

}
