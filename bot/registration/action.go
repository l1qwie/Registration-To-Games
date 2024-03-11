package registration

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/bot/schedule"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"fmt"
	"strconv"
)

func SeatsAreFull(user *bottypes.User, fm *formatter.Formatter) {
	var (
		coordinates    []int
		kbName, kbData []string
		dict           map[string]string
	)
	dict = dictionary.Dictionary[user.Language]
	user.Level = 5
	kbName = []string{dict["go-ahead"], dict["no-deleteAll"]}
	kbData = []string{"Go-ahead", "MainMenu"}
	coordinates = []int{1, 1}
	forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
	fm.WriteString(fmt.Sprintf(dict["SeatsAreFull"], dict["Review"]))
	fm.WriteChatId(user.Id)
}

func PresentationScheduele(user *bottypes.User, fm *formatter.Formatter) {
	var (
		sch            []*forall.Game
		coordinates    []int
		kbName, kbData []string
		dict           map[string]string
	)
	dict = dictionary.Dictionary[user.Language]
	if schedule.FindGame() {
		user.Level = 1
		sch = selectTheSchedule(user.Media.Limit, user.LaunchPoint, user.Language)
		for i := 0; sch[i] != nil && i < len(sch); i++ {
			coordinates = append(coordinates, 1)
			kbName = append(kbName, fmt.Sprintf("%s %s %s %s %d", sch[i].Sport, sch[i].Date, sch[i].Time, dict["freeSpace"], sch[i].Seats))
			kbData = append(kbData, strconv.Itoa(sch[i].Id))
		}
		coordinates = append(coordinates, 2)
		coordinates = append(coordinates, 1)
		forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
		fm.WriteInlineButtonCmd(dict["previous"], "previous page")
		fm.WriteInlineButtonCmd(dict["next"], "next page")
		fm.WriteInlineButtonCmd(dict["MainMenu"], "MainMenu")
		fm.WriteString(dict["ChooseAnyGame"])
		fm.WriteChatId(user.Id)
	} else {
		forall.GoToMainMenu(user, fm, (fmt.Sprint(dict["NoGames"], dict["MainMenu"])))
	}
}

func ChooseGame(user *bottypes.User, fm *formatter.Formatter, seatsStatus bool) {
	var (
		detected             bool
		gameId, price, space int
		currency             string
		coordinates          []int
		kbName, kbData       []string
		dict                 map[string]string
	)
	dict = dictionary.Dictionary[user.Language]
	detected, gameId = forall.IntCheck(user.Request)
	if detected {
		if FindAGame(gameId) {
			user.Reg.GameId = gameId
			user.Level = 2
			price, space, currency = selectThePrice(gameId)
			for i := 0; i < space && i < 3; i++ {
				coordinates = append(coordinates, 1)
				kbName = append(kbName, dict[strconv.Itoa(i+1)])
				kbData = append(kbData, strconv.Itoa(i+1))
			}
			forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
			if seatsStatus {
				fm.WriteString(fmt.Sprintf(dict["ChooseSeats"], price, currency))
			} else {
				fm.WriteString(fmt.Sprintf(dict["NoMoreSeats"], dict["ChooseSeats"], price, currency))
			}
			fm.WriteChatId(user.Id)
		} else {
			user.LaunchPoint += forall.IncreaseLaunchPoit(user.Request)
			PresentationScheduele(user, fm)
		}
	} else {
		user.LaunchPoint += forall.IncreaseLaunchPoit(user.Request)
		PresentationScheduele(user, fm)
	}
}

func ChooseSeats(user *bottypes.User, fm *formatter.Formatter) {
	var (
		detected       bool
		dict           map[string]string
		seats          int
		coordinates    []int
		kbName, kbData []string
	)
	dict = dictionary.Dictionary[user.Language]
	detected, seats = forall.IntCheck(user.Request)
	if detected {
		if howManyIsLeft(user.Reg.GameId, seats) {
			user.Reg.Seats = seats
			user.Level = 3
			coordinates = []int{1, 1, 1}
			kbName = []string{dict["payByCard"], dict["payByCash"], dict["MainMenu"]}
			kbData = []string{"card", "cash", "MainMenu"}
			forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
			fm.WriteString(dict["ChoosePaymethod"])
			fm.WriteChatId(user.Id)
		} else {
			user.Request = strconv.Itoa(user.Reg.GameId)
			ChooseGame(user, fm, false)
		}
	} else {
		user.Request = strconv.Itoa(user.Reg.GameId)
		ChooseGame(user, fm, true)
	}

}

func ChoosePayment(user *bottypes.User, fm *formatter.Formatter) {
	var (
		price, cost int
		currency    string
		dict        map[string]string
	)
	dict = dictionary.Dictionary[user.Language]
	if user.Request == "card" || user.Request == "cash" {
		if howManyIsLeft(user.Reg.GameId, user.Reg.Seats) {
			user.Reg.Payment = user.Request
			if user.Request == "card" {
				user.Level = 4
				price, _, currency = selectThePrice(user.Reg.GameId)
				cost = price * user.Reg.Seats
				fm.SetIkbdDim([]int{1, 1, 1})
				fm.WriteInlineButtonUrl(dict["pay"], "https://www.papara.com/personal/qr?karekod=7502100102120204082903122989563302730612230919141815530394954120000000000006114081020219164116304DDE3")
				fm.WriteInlineButtonCmd(dict["GoNext"], "Next")
				fm.WriteInlineButtonCmd(dict["MainMenu"], "MainMenu")
				fm.AddPhotoFromStorage("qr.jpg")
				fm.WriteString(fmt.Sprintf(dict["WaitForYourMoney"], cost, currency))
				fm.WriteChatId(user.Id)
				fm.WriteDeleteMesId(user.ExMessageId)
			} else {
				BestWishes(user, fm)
			}
		} else {
			SeatsAreFull(user, fm)
		}
	} else {
		user.Request = strconv.Itoa(user.Reg.Seats)
		ChooseSeats(user, fm)
	}
}

func BestWishes(user *bottypes.User, fm *formatter.Formatter) {
	var (
		details        *forall.Game
		coordinates    []int
		kbName, kbData []string
		dict           map[string]string
		cost           int
	)
	dict = dictionary.Dictionary[user.Language]
	if user.Request == "cash" || user.Request == "Next" {
		if howManyIsLeft(user.Reg.GameId, user.Reg.Seats) {
			completeRegistration(user.Id, user.Reg.GameId, user.Reg.Seats, user.Reg.Payment)
			user.Level = 3
			user.Act = "divarication"
			details = new(forall.Game)
			details = selectDetailOfGame(user.Reg.GameId, user.Language)
			cost = details.Price * user.Reg.Seats
			kbName = []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}
			kbData = []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}
			coordinates = []int{1, 1, 1, 1}
			forall.SetTheKeyboard(fm, coordinates, kbName, kbData)
			res := fmt.Sprintf(dict["RegistrationCompleted"], details.Sport, details.Date, details.Time, user.Reg.Seats, user.Reg.Payment, cost, details.Currency, details.Address, details.Lattitude, details.Longitude)
			fm.WriteString(res)
			fm.WriteParseMode(types.HTML)
			fm.WriteChatId(user.Id)
		} else {
			SeatsAreFull(user, fm)
		}
	} else {
		user.Request = user.Reg.Payment
		ChoosePayment(user, fm)
	}
}
