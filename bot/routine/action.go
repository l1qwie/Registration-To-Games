package routine

import (
	"fmt"
	"log"
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/welcome"
	"registrationtogames/fmtogram/formatter"
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

type User struct {
	Id          int
	Request     string
	Language    string
	LaunchPoint int
	Act         string
	Level       int
	Reg         bottypes.RegToGames
	Media       bottypes.Media
	UserRec     bottypes.UserRec
}

func (user *User) retrieveUser() {
	var err error
	if user.find() {
		err = user.dbRetrieveUser()
	} else {
		err = user.createUser()
	}

	if err != nil {
		log.Fatal(err)
	}
}

func (user *User) retainUser() {
	err := user.dbRetainUser()
	if err != nil {
		log.Fatal(err)
	}
}

func (user *User) Welcome() {
	var welcome *welcome.User
	welcome = user
	if user.Level == START {
		welcome.GreetingsToUser()
	}
}

func (user *User) DispatcherPhrase(fm *formatter.Formatter) {
	user.retrieveUser()
	fmt.Println("level =", user.Level, "phrase =", user.Request, "action =", user.Act)
	if user.Act == "registration" {
		user.Welcome()
	}
	user.retainUser()
}
