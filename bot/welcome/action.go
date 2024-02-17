package welcome

import "registrationtogames/bot/bottypes"

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

func (user *User) GreetingsToUser() {

}
