package forall

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/fmtogram/formatter"
	"fmt"
	"strconv"
)

func FromIntToStrDate(numberDate int) (date string) {
	var (
		year, month, day       int
		monthString, dayString string
	)

	year = numberDate / 10000
	month = (numberDate - (year * 10000)) / 100
	day = (numberDate - (year * 10000) - (month * 100))

	if month <= 10 {
		monthString = fmt.Sprintf("%d%d", 0, month)
	}
	if day <= 10 {
		dayString = fmt.Sprintf("%d%d", 0, day)
	}
	if monthString != "" && dayString != "" {
		date = fmt.Sprintf("%s-%s-%d", dayString, monthString, year)
	} else if monthString != "" && dayString == "" {
		date = fmt.Sprintf("%d-%s-%d", day, monthString, year)
	} else if monthString == "" && dayString != "" {
		date = fmt.Sprintf("%s-%d-%d", dayString, month, year)
	} else {
		date = fmt.Sprintf("%d-%d-%d", day, month, year)
	}

	return date
}

func FromIntToStrTime(numberTime int) (time string) {
	var (
		hour, minute             int
		hourString, minuteString string
	)
	hour = numberTime / 100
	minute = (numberTime - (hour * 100))
	if hour < 10 {
		hourString = fmt.Sprintf("%d%d", 0, hour)
	}
	if minute < 10 {
		minuteString = fmt.Sprintf("%d%d", 0, minute)
	}

	if hourString != "" && minuteString != "" {
		time = fmt.Sprintf("%s:%s", hourString, minuteString)
	} else if hourString != "" && minuteString == "" {
		time = fmt.Sprintf("%s:%d", hourString, minute)
	} else if hourString == "" && minuteString != "" {
		time = fmt.Sprintf("%d:%s", hour, minuteString)
	} else if hourString == "" && minuteString == "" {
		time = fmt.Sprintf("%d:%d", hour, minute)
	}
	return time
}

func IntCheck(numS string) (bool, int) {
	num, err := strconv.Atoi(numS)
	return err == nil, num
}

func SetTheKeyboard(fm *formatter.Formatter, coordinates []int, kbName, kbData []string) {
	fm.SetIkbdDim(coordinates)
	for i := 0; i < len(coordinates) && i < len(kbName); i++ {
		fm.WriteInlineButtonCmd(kbName[i], kbData[i])
	}
}

func IncreaseLaunchPoit(phrase string) (launchPoint int) {
	if phrase == "next page" {
		launchPoint = 7
	} else if phrase == "previous page" {
		launchPoint = -7
	}

	return launchPoint
}

func GoToMainMenu(user *bottypes.User, fm *formatter.Formatter, textresponse string) {
	var (
		dict           map[string]string
		kbName, kbData []string
		coordinates    []int
	)
	user.Level = 3
	user.Act = "divarication"
	dict = dictionary.Dictionary[user.Language]
	kbName = []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}
	kbData = []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}
	coordinates = []int{1, 1, 1, 1}
	SetTheKeyboard(fm, coordinates, kbName, kbData)
	fm.WriteString(textresponse)
	fm.WriteChatId(user.Id)
}

func CheckPages(req string, lp int) int {
	if req == "next page" {
		lp = lp + 7
	} else if req == "previuos page" {
		lp = lp - 7
	}

	return lp
}

func Seats(hm int, fm *formatter.Formatter) {
	crd := make([]int, hm)
	names := make([]string, hm)
	datas := make([]string, hm)
	for i := range crd {
		crd[i] = 1
		names[i] = fmt.Sprint(i + 1)
		datas[i] = fmt.Sprint(i + 1)
	}
	SetTheKeyboard(fm, crd, names, datas)
}
