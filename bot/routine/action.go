package routine

import (
	"fmt"
	"log"
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/registration"
	"registrationtogames/bot/welcome"
	"registrationtogames/fmtogram/formatter"
	"runtime"
	"strconv"
)

const (
	START   = 0
	LEVEL1  = 1
	LEVEL2  = 2
	LEVEL3  = 3
	OPTIONS = 3
	LEVEL4  = 4
	LEVEL5  = 5
	LEVEL6  = 6
	LEVEL7  = 7
	LEVEL8  = 8
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

func DispatcherPhrase(user *bottypes.User, fm *formatter.Formatter) {
	retrieveUser(user)
	fmt.Println("level =", user.Level, "phrase =", user.Request, "action =", user.Act)
	if user.Act == "registration" {
		Welcome(user, fm)
	} else if user.Act == "reg to games" {
		RegToGames(user, fm)
	}
	retainUser(user)
}
