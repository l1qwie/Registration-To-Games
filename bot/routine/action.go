package routine

import (
	"fmt"
	"log"
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/dictionary"
	"registrationtogames/bot/forall"
	"registrationtogames/bot/registration"
	"registrationtogames/bot/welcome"
	"registrationtogames/fmtogram/formatter"
	"runtime"
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
		_, file, line, _ := runtime.Caller(0)
		log.Fatalf("Error at %s:%d: %v", file, line, err)
	}
}

func retainUser(user *bottypes.User) {
	err := DbRetainUser(user)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatalf("Error at %s:%d: %v", file, line, err)
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
		user.Level = 0
		user.Act = "reg to games"
		RegToGames(user, fm)
	} else {
		MainMenu(user, fm)
	}
}

func DispatcherPhrase(user *bottypes.User, fm *formatter.Formatter) {
	retrieveUser(user)
	fmt.Println("level =", user.Level, "phrase =", user.Request, "action =", user.Act)
	if user.Request == "MainMenu" {
		MainMenu(user, fm)
	} else if user.Act == "registration" {
		Welcome(user, fm)
	} else if user.Act == "reg to games" {
		RegToGames(user, fm)
	} else if user.Act == "divarication" {
		Options(user, fm)
	}
	retainUser(user)
}
