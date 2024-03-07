package routine

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/routine"
	"RegistrationToGames/fmtogram/types"

	_ "github.com/lib/pq"
)

func DeleteUser(id int) {
	_, err := types.Db.Exec("DELETE FROM Users WHERE userId = $1", id)
	if err != nil {
		panic(err)
	}
}

func TestRetrevenUser() {
	var (
		user *bottypes.User
		err  error
	)
	user = new(bottypes.User)
	user.Id = 8223
	user.Language = "ru"
	user.Act = "registration"

	defer func() {
		if routine.Find(user.Id) {
			panic("User detected")
		}
	}()
	defer DeleteUser(user.Id)

	err = routine.CreateUser(user.Id, user.Language)
	if err != nil {
		panic(err)
	}
	err = routine.DbRetrieveUser(user)
	if err != nil {
		panic(err)
	}

	if user.Id != 8223 {
		panic("user.Id != 8223")
	} else if user.Language != "ru" {
		panic("user.Language != `ru`")
		//Change it in the future!!
		//} else if gameId != 0 {
		//panic("gameId != 0")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "registration" {
		panic("user.Act != `registration`")
	}
	if user.Level != 0 {
		panic("user.Level != 0")
	}
	if user.Reg.Seats != 0 {
		panic("user.Reg.Seats != 0")
	}
	if user.Reg.Payment != "" {
		panic("user.Reg.Payment != ``")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "" {
		panic("user.Media.Direction != ``")
	}
	if user.Media.Limit != 7 {
		panic("user.Media.Limit != 7")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
	if user.UserRec.Changeable != "" {
		panic("user.UserRec.Changeable != ``")
	}
	if user.UserRec.ActGame != "" {
		panic("user.UserRec.ActGame != ``")
	}
	if user.UserRec.WillChangeable != "" {
		panic("user.UserRec.WillChangeable != ``")
	}
	if user.UserRec.NewPay != "" {
		panic("user.UserRec.NewPay != ``")
	}

}

func TestRetainUser() {
	var (
		user *bottypes.User
		err  error
	)
	user = new(bottypes.User)
	user.Id = 8217
	user.Language = "en"
	user.Act = "registration"
	user.LaunchPoint = 18
	user.UserRec.Changeable = "seats"

	defer func() {
		if routine.Find(user.Id) {
			panic("User detected")
		}
	}()
	defer DeleteUser(user.Id)

	err = routine.CreateUser(user.Id, user.Language)
	if err != nil {
		panic(err)
	}
	err = routine.DbRetainUser(user)
	if err != nil {
		panic(err)
	}
	err = routine.DbRetrieveUser(user)
	if err != nil {
		panic(err)
	}
	if user.Id != 8217 {
		panic("user.Id != 8223")
	} else if user.Language != "en" {
		panic("user.Language != `en`")
		//Change it in the future!!
		//} else if gameId != 0 {
		//panic("gameId != 0")
	}
	if user.LaunchPoint != 18 {
		panic("user.LaunchPoint != 18")
	}
	if user.Act != "registration" {
		panic("user.Act != `registration`")
	}
	if user.Level != 0 {
		panic("user.Level != 0")
	}
	if user.Reg.Seats != 0 {
		panic("user.Reg.Seats != 0")
	}
	if user.Reg.Payment != "" {
		panic("user.Reg.Payment != ``")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "" {
		panic("user.Media.Direction != ``")
	}
	if user.Media.Limit != 7 {
		panic("user.Media.Limit != 7")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
	if user.UserRec.Changeable != "seats" {
		panic("user.UserRec.Changeable != `seats`")
	}
	if user.UserRec.ActGame != "" {
		panic("user.UserRec.ActGame != ``")
	}
	if user.UserRec.WillChangeable != "" {
		panic("user.UserRec.WillChangeable != ``")
	}
	if user.UserRec.NewPay != "" {
		panic("user.UserRec.NewPay != ``")
	}
}
