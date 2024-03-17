package settings

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/formatter"
)

func ChooseOptions(user *bottypes.User, fm *formatter.Formatter) {
	dict := dictionary.Dictionary[user.Language]
	if FindUserGames(user.Id) {
		//create a schedule of user games
	} else {
		//only change language
		user.Level = 2
		user.UserRec.WillChangeable = "language"
		fm.WriteString(dict["NoGamesChangeLang"])
		forall.SetTheKeyboard(fm, []int{1, 1, 1, 1}, []string{dict["en"], dict["ru"], dict["tur"], dict["MainMenu"]}, []string{"en", "ru", "tur", "MainMenu"})
		fm.WriteChatId(user.Id)
	}
}

func ChangeLanguge(user *bottypes.User, fm *formatter.Formatter) {
	if user.Request == "en" || user.Request == "ru" || user.Request == "tur" {
		user.Language = updateLanguage(user.Id, user.Request)
		user.Level = 3
		user.Act = "divarication"
		dict := dictionary.Dictionary[user.Language]
		forall.GoToMainMenu(user, fm, dict["Lanchanged"])
		fm.WriteChatId(user.Id)
	} else {
		ChooseOptions(user, fm)
	}
}

func WhatWay(user *bottypes.User, fm *formatter.Formatter) {
	if user.UserRec.WillChangeable == "language" {
		ChangeLanguge(user, fm)
	} else {
		//
	}
}
