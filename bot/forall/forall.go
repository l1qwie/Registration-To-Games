package forall

import (
	"fmt"
	"registrationtogames/fmtogram/formatter"
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
		monthString = fmt.Sprint(0, month)
	}
	if day <= 10 {
		dayString = fmt.Sprint(0, day)
	}
	if monthString != "" && dayString != "" {
		date = fmt.Sprint(dayString, monthString, year)
	} else if monthString != "" && dayString == "" {
		date = fmt.Sprint(day, monthString, year)
	} else if monthString == "" && dayString != "" {
		date = fmt.Sprint(dayString, month, year)
	} else {
		date = fmt.Sprint(day, month, year)
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
		hourString = fmt.Sprint(0, hour)
	}
	if minute <= 10 {
		minuteString = fmt.Sprint(0, minute)
	}

	if hourString != "" && minuteString != "" {
		time = fmt.Sprint(hourString, minuteString)
	} else if hourString != "" && minuteString == "" {
		time = fmt.Sprint(hourString, minute)
	} else if hourString == "" && minuteString != "" {
		time = fmt.Sprint(hour, minuteString)
	} else if hourString == "" && minuteString == "" {
		time = fmt.Sprint(hour, minute)
	}

	return time
}

func IntCheck(numS string) (detected bool, num int) {
	var err error
	num, err = strconv.Atoi(numS)
	if err != nil {
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
