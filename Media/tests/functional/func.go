package functional

import (
	"Media/apptype"
	"Media/fmtogram/types"
	"fmt"
)

// [{"text":"","callback_data":"","url":""}]
// `{"inline_keyboard":[[{"text":"","callback_data":"","url":""}],[{"text":"","callback_data":"","url":""}]]}`
/*
	ch.mes = ""
	ch.kb = ``
	ch.lvl = 0
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = ""
	ch.gameid = 0
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.maintest()
*/

var ch = new(check)

type check struct {
	mes        string
	kb         string
	lvl        int
	lp         int
	act        string
	mediaDir   string
	gameid     int
	fileId     string
	typeoffile string
	media      []types.Media
	prmode     string
	status     bool
	res        *apptype.Response
}

// Compares media arrays
func (ch *check) compareMedia() {
	if len(ch.res.Media) != len(ch.media) {
		fmt.Println(ch.res.Media)
		fmt.Println(ch.media)
		panic("len(ch.res.Media) != len(ch.media)")
	}
	for i, val := range ch.res.Media {
		if val.Media != ch.media[i].Media {
			panic(fmt.Sprintf(`ch.res.Media[%d].Fileid != %s because %s`, i, ch.media[i].Media, fmt.Sprintf("ch.res.Media[%d].Fileid == %s", i, val.Media)))
		}
		if val.Type != ch.media[i].Type {
			panic(fmt.Sprintf(`ch.res.Media[%d].TypeOfFile != %s because %s`, i, ch.media[i].Type, fmt.Sprintf("ch.res.Media[%d].TypeOfFile == %s", i, val.Type)))
		}
	}
}

// Main tester
// Takes data from struct ckeck and ckecks it between Response and what I want
func (ch *check) maintest() {
	if ch.res.ChatID != 499 {
		panic((fmt.Sprintf(`ch.res.ChatID != 499 because %s`, fmt.Sprintf("ch.res.ChatID == %d", ch.res.ChatID))))
	}
	if ch.res.Message != ch.mes {
		panic(fmt.Sprintf(`ch.res.Message != "%s" because %s`, ch.mes, fmt.Sprintf(`ch.res.Message == "%s"`, ch.res.Message)))
	}
	if ch.res.Keyboard != ch.kb {
		panic(fmt.Sprintf(`ch.res.Keyboard != "%s" because %s`, ch.kb, fmt.Sprintf(`ch.res.Keyboard == "%s"`, ch.res.Keyboard)))
	}
	if ch.res.Level != ch.lvl {
		panic(fmt.Sprintf(`ch.res.Level != %d because %s`, ch.lvl, fmt.Sprintf("ch.res.level == %d", ch.res.Level)))
	}
	if ch.res.LaunchPoint != ch.lp {
		panic(fmt.Sprintf(`ch.res.LaunchPoint != %d because %s`, ch.lp, fmt.Sprintf("ch.res.LaunchPoint == %d", ch.res.LaunchPoint)))
	}
	if ch.res.Act != ch.act {
		panic(fmt.Sprintf(`ch.res.Act != "%s" because %s`, ch.act, fmt.Sprintf(`ch.res.Act == "%s"`, ch.res.Act)))
	}
	if ch.res.MediaDir != ch.mediaDir {
		panic(fmt.Sprintf(`ch.res.MediaDir != "%s" because %s`, ch.mediaDir, fmt.Sprintf(`ch.res.MediaDir == "%s"`, ch.res.MediaDir)))
	}
	if ch.res.GameId != ch.gameid {
		panic(fmt.Sprintf(`ch.res.GameId != %d because %s`, ch.gameid, fmt.Sprintf("ch.res.GameId == %d", ch.res.GameId)))
	}
	if ch.res.FileId != ch.fileId {
		panic(fmt.Sprintf(`ch.res.FileId != %s because %s`, ch.fileId, fmt.Sprintf("ch.res.FileId == %s", ch.res.FileId)))
	}
	if ch.res.TypeOffile != ch.typeoffile {
		panic(fmt.Sprintf(`ch.res.TypeOffile != %s because %s`, ch.typeoffile, fmt.Sprintf("ch.res.TypeOffile == %s", ch.res.TypeOffile)))
	}
	ch.compareMedia()
	if ch.res.ParseMode != ch.prmode {
		panic(fmt.Sprintf(`ch.res.ParseMode != %s because %s`, ch.prmode, fmt.Sprintf(`ch.res.ParseMode == "%s"`, ch.res.ParseMode)))
	}
	if ch.res.Status != ch.status {
		panic(fmt.Sprintf(`ch.res.Status != %v because %v`, ch.status, fmt.Sprintf(`ch.res.Status == "%v"`, ch.res.Status)))
	}
}

// The data after the first function of the act
func ChDirection(res *apptype.Response) {
	ch.mes = "Вы хотите загрузить или посмотреть?"
	ch.kb = `{"inline_keyboard":[[{"text":"Посмотреть","callback_data":"unload","url":""}],[{"text":"Загрузить","callback_data":"upload","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 1
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = ""
	ch.gameid = 0
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// The data after "chooseMediaGame" unload (only one game)
func ChMediaGameUn(res *apptype.Response) {
	ch.mes = "Выберите интересующую вас игру"
	ch.kb = `{"inline_keyboard":[[{"text":"Футбол 12-02-2024 12:00","callback_data":"10","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 3
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "unload"
	ch.gameid = 0
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// The data after "unloadAndUnload" unload (only one game)
func UnOne(res *apptype.Response) {
	ch.mes = "Вот все медиа по это игре"
	ch.fileId = "!@#IOJSIOJE!@#**()!@#$*()SIOPE!@()#"
	ch.typeoffile = "photo"
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 4
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "unload"
	ch.gameid = 10
	ch.prmode = ""
	ch.res = res
	ch.status = true
	ch.maintest()
}

// The data after the first function of the act,
// but without games to unload. Only with one game
func NoChOnluUp(res *apptype.Response) {
	ch.mes = "Игр с уже загруженными медиафайлами еще нет, но вы можете стать первым!\nВыберите интересующую вас игру"
	ch.kb = `{"inline_keyboard":[[{"text":"Волейбол 12-02-2024 12:00","callback_data":"10","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 2
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "upload"
	ch.gameid = 0
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// The data after "chooseMediaGame" upload (only one game)
func NoChOnluUpSecond(res *apptype.Response) {
	ch.mes = "Выберите интересующую вас игру"
	ch.kb = `{"inline_keyboard":[[{"text":"Волейбол 12-02-2024 12:00","callback_data":"10","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 2
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "upload"
	ch.gameid = 0
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// The data after "chooseMediaGame" upload (a few games). Second time
func NochOnluUpAfewSecond(res *apptype.Response) {
	ch.mes = "Выберите интересующую вас игру"
	ch.kb = `{"inline_keyboard":[[{"text":"Волейбол 02-02-2025 08:00","callback_data":"4","url":""}],[{"text":"Волейбол 02-02-2025 08:00","callback_data":"3","url":""}],[{"text":"Футбол 12-04-2025 18:00","callback_data":"2","url":""}],[{"text":"Волейбол 12-02-2026 11:00","callback_data":"1","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 2
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "upload"
	ch.gameid = 0
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// The data after waitingYourMedia upload one game
func WaitForMedia(res *apptype.Response) {
	ch.mes = "Прекрасно! Вы выбрали игру. У нее осталось 20 свободных мест для медиафайлов (максимум 20). Присылайте ваших фотографий или видео, которые были сделаны на этой игре. За один раз можно прислать не более 10 фалов. Кончено же, желательно, чтобы это число не привышало свободных мест"
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 3
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "upload"
	ch.gameid = 10
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// The data after waitingYourMedia upload a few games
func WaitForMediaAfew(res *apptype.Response) {
	ch.mes = "Прекрасно! Вы выбрали игру. У нее осталось 20 свободных мест для медиафайлов (максимум 20). Присылайте ваших фотографий или видео, которые были сделаны на этой игре. За один раз можно прислать не более 10 фалов. Кончено же, желательно, чтобы это число не привышало свободных мест"
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 3
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "upload"
	ch.gameid = 1
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// The data after "unloadAndUnload". Only woth upload one game
func UpOne(res *apptype.Response) {
	ch.mes = "Все файлы успешно загружены"
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 4
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "upload"
	ch.gameid = 10
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = true
	ch.maintest()
	if !checkUploadedMedia() {
		panic("The media wasn't uploaded")
	}
}

// The data after "chooseMediaGame" unload (a few games)
func ChMediaGamesUn(res *apptype.Response) {
	ch.mes = "Выберите интересующую вас игру"
	ch.kb = `{"inline_keyboard":[[{"text":"Волейбол 02-02-2025 08:00","callback_data":"4","url":""}],[{"text":"Волейбол 02-02-2025 08:00","callback_data":"3","url":""}],[{"text":"Футбол 12-04-2025 18:00","callback_data":"2","url":""}],[{"text":"Волейбол 12-02-2026 11:00","callback_data":"1","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 3
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "unload"
	ch.gameid = 0
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// The data after "unloadAndUnload". Only woth upload a few games
func UnAll(res *apptype.Response) {
	ch.mes = "Вот все медиа по это игре"
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 4
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "unload"
	ch.gameid = 1
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = make([]types.Media, 3)
	ch.media[0] = types.Media{Type: "photo", Media: "!@#IOJSJE!@#**()!@#$*()SIOPE!@()#"}
	ch.media[1] = types.Media{Type: "photo", Media: "!@#IOJSIOJE!@#**()!@HFF#$*()SIOPE$#@$!@()#"}
	ch.media[2] = types.Media{Type: "photo", Media: "!@#IOJSIOJE!$@!$!@#@#**()!@#$*()SIOPE!@()#"}
	ch.res = res
	ch.status = true
	ch.maintest()
}

// The data after the first function of the act,
// but without games to upload. Only with one game
func NoChOnluUpAfew(res *apptype.Response) {
	ch.mes = "Игр с уже загруженными медиафайлами еще нет, но вы можете стать первым!\nВыберите интересующую вас игру"
	ch.kb = `{"inline_keyboard":[[{"text":"Волейбол 02-02-2025 08:00","callback_data":"4","url":""}],[{"text":"Волейбол 02-02-2025 08:00","callback_data":"3","url":""}],[{"text":"Футбол 12-04-2025 18:00","callback_data":"2","url":""}],[{"text":"Волейбол 12-02-2026 11:00","callback_data":"1","url":""}],[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 2
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "upload"
	ch.gameid = 0
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = false
	ch.maintest()
}

// Checks all uploaded files
func checkUploadAfewMedia() {
	if !thefirst() {
		panic("The first media file wasn't saved")
	} else if !thesecond() {
		panic("The second media file wasn't saved")
	} else if !thethird() {
		panic("The third media file wasn't saved")
	}
}

// The data after "unloadAndUnload". Only woth upload a few games
func UpAfew(res *apptype.Response) {
	ch.mes = "Все файлы успешно загружены"
	ch.kb = `{"inline_keyboard":[[{"text":"Главное Меню","callback_data":"MainMenu","url":""}]]}`
	ch.lvl = 4
	ch.lp = 0
	ch.act = "photos and videos"
	ch.mediaDir = "upload"
	ch.gameid = 1
	ch.prmode = ""
	ch.fileId = ""
	ch.typeoffile = ""
	ch.media = []types.Media{}
	ch.res = res
	ch.status = true
	ch.maintest()
	checkUploadAfewMedia()
}
