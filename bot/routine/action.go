package routine

import (
	"fmt"
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/dictionary"
	"registrationtogames/bot/forall"
	"registrationtogames/bot/registration"
	"registrationtogames/bot/schedule"
	"registrationtogames/bot/welcome"
	"registrationtogames/fmtogram/formatter"
	"strconv"
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

func retrieveUser(user *bottypes.User) {
	var err error
	if Find(user.Id) {
		err = DbRetrieveUser(user)
	} else {
		err = CreateUser(user.Id, user.Language)
	}

	if err != nil {
		panic(err)
	}
}

func retainUser(user *bottypes.User) {
	err := DbRetainUser(user)
	if err != nil {
		panic(err)
	}
}

func Welcome(user *bottypes.User, fm *formatter.Formatter) {
	if user.Level == START {
		welcome.GreetingsToUser(user, fm)
	} else if user.Level == LEVEL1 {
		welcome.ShowRules(user, fm)
	} else if user.Level == LEVEL2 {
		welcome.WelcomeToMainMenu(user, fm)
	}
}

func RegToGames(user *bottypes.User, fm *formatter.Formatter) {
	if user.Level == START {
		registration.PresentationScheduele(user, fm)
	} else if user.Level == LEVEL1 {
		registration.ChooseGame(user, fm, true)
	} else if user.Level == LEVEL2 {
		registration.ChooseSeats(user, fm)
	} else if user.Level == LEVEL3 {
		registration.ChoosePayment(user, fm)
	} else if user.Level == LEVEL4 {
		registration.BestWishes(user, fm)
	} else if user.Level == LEVEL5 {
		user.Request = strconv.Itoa(user.Reg.GameId)
		registration.ChooseGame(user, fm, true)
	}
}

func MainMenu(user *bottypes.User, fm *formatter.Formatter) {
	var (
		kbName, kbData []string
		coordinates    []int
		dict           map[string]string
	)
	dict = dictionary.Dictionary[user.Language]
	kbName = []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}
	kbData = []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}
	coordinates = []int{1, 1, 1, 1}
	forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
	fm.WriteString(dict["Can'tUnderstend"])
	fm.WriteChatId(user.Id)
}

func Options(user *bottypes.User, fm *formatter.Formatter) {
	if user.Request == "Reg to games" {
		user.Level = START
		user.Act = "reg to games"
		RegToGames(user, fm)
	} else if user.Request == "Looking Schedule" {
		user.Level = START
		user.Act = "see schedule"
		Schedule(user, fm)
	} else {
		user.Act = "divarication"
		user.Level = OPTIONS
		MainMenu(user, fm)
	}
}

func Schedule(user *bottypes.User, fm *formatter.Formatter) {
	if user.Level == START {
		schedule.ShowTheSchedule(user, fm)
	}
}

func Edit(user *bottypes.User, fm *formatter.Formatter) (exMessageId int) {
	var err error
	if user.ExMessageId == 0 {
		fmt.Println("SELECT")
		exMessageId, err = SelectExMessageId(user.Id)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("UPDATE")
		exMessageId = user.ExMessageId
		err = updateExMessageId(exMessageId, user.Id)
		if err != nil {
			panic(err)
		}
	}
	if user.PhotoFileId == "" {
		fmt.Println("Edit", exMessageId)
		fm.WriteEditMesId(exMessageId)
	} else {
		fmt.Println("Delete", exMessageId)
		fm.WriteDeleteMesId(exMessageId)
	}
	return exMessageId
}

func DispatcherPhrase(user *bottypes.User, fm *formatter.Formatter) {
	retrieveUser(user)
	//fm.WriteDeleteMesId(user.ExMessageId)
	//fm.WriteChatId(user.Id)
	user.ExMessageId = Edit(user, fm)
	fmt.Println("level =", user.Level, fmt.Sprintf(`phrase = "%s"`, user.Request), fmt.Sprintf(`action = "%s"`, user.Act), "user.Id =", user.Id)
	if user.Request == "MainMenu" {
		user.Act = "divarication"
		user.Level = OPTIONS
		//fm.WriteDeleteMesId(user.ExMessageId)
		//fm.AddPhotoFromStorage("qr.jpg")
		MainMenu(user, fm)
	} else if user.Act == "registration" {
		Welcome(user, fm)
	} else if user.Act == "reg to games" {
		RegToGames(user, fm)
	} else if user.Act == "divarication" {
		Options(user, fm)
	} else if user.Act == "see schedule" {
		Schedule(user, fm)
	}
	retainUser(user)
}
