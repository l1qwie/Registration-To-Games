package routine

import (
	"fmt"
	"log"
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/welcome"
	"registrationtogames/fmtogram/formatter"
	"runtime"
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
	if find(user) {
		err = dbRetrieveUser(user)
	} else {
		err = createUser(user)
	}

	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatalf("Error at %s:%d: %v", file, line, err)
	}
}

func retainUser(user *bottypes.User) {
	err := dbRetainUser(user)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatalf("Error at %s:%d: %v", file, line, err)
	}
}

func Welcome(user *bottypes.User, fm *formatter.Formatter) {
	if user.Level == START {
		welcome.GreetingsToUser(user, fm)
	}
}

func DispatcherPhrase(user *bottypes.User, fm *formatter.Formatter) {
	retrieveUser(user)
	fmt.Println("level =", user.Level, "phrase =", user.Request, "action =", user.Act)
	if user.Act == "registration" {
		Welcome(user, fm)
	}
	retainUser(user)
}
