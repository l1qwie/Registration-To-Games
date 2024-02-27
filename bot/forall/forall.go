package forall

import (
	"fmt"
	"registrationtogames/fmtogram/formatter"
	"strconv"
)

type Game struct {
	Id        int
	Sport     string
	Date      string
	Time      string
	Seats     int
	Price     int
	Currency  string
	Address   string
	Lattitude float32
	Longitude float32
}

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

	if hour <= 10 {
		hourString = fmt.Sprintf("%d%d", 0, hour)
	}
	if minute <= 10 {
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

func IntCheck(numS string) (detected bool, num int) {
	var err error
	num, err = strconv.Atoi(numS)
	if err == nil {
		detected = true
	}

	return detected, num
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
