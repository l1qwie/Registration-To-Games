package welcome

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/fmtogram/formatter"
)

func GreetingsToUser(user *bottypes.User, fm *formatter.Formatter) {
	var (
		kbName, kbData []string
		coordinates    []int
	)
	user.Level = 1
	kbName = []string{dictionary.Dictionary[user.Language]["reg"]}
	kbData = []string{"GoReg"}
	coordinates = []int{1}

	fm.SetIkbdDim(coordinates)
	for i := 0; i < len(kbName); i++ {
		fm.WriteInlineButtonCmd(kbName[i], kbData[i])
	}
	fm.WriteString(dictionary.Dictionary[user.Language]["WelcomeToBot"])
	fm.WriteChatId(user.Id)
}

func ShowRules(user *bottypes.User, fm *formatter.Formatter) {
	var (
		kbName, kbData []string
		coordinates    []int
	)
	if user.Request == "GoReg" {
		user.Level = 2
		kbName = []string{dictionary.Dictionary[user.Language]["allright"]}
		kbData = []string{"GoNext"}
		coordinates = []int{1}

		fm.SetIkbdDim(coordinates)
		for i := 0; i < len(kbName); i++ {
			fm.WriteInlineButtonCmd(kbName[i], kbData[i])
		}
		fm.WriteString(dictionary.Dictionary[user.Language]["BotRules"])
		fm.WriteChatId(user.Id)
	} else {
		GreetingsToUser(user, fm)
	}
}

func WelcomeToMainMenu(user *bottypes.User, fm *formatter.Formatter) {
	var (
		kbName, kbData []string
		coordinates    []int
		err            error
		dict           map[string]string
	)
	if user.Request == "GoNext" {
		//err = CompletionOfRegistration(user.Id)
		if err == nil {
			user.Level = 3
			user.Act = "divarication"
			dict = dictionary.Dictionary[user.Language]
			kbName = []string{dict["first"], dict["second"], dict["third"], dict["fourth"]}
			kbData = []string{"Looking Schedule", "Reg to games", "Photo&Video", "My records"}
			coordinates = []int{1, 1, 1, 1}
			fm.SetIkbdDim(coordinates)
			for i := 0; i < len(kbName); i++ {
				fm.WriteInlineButtonCmd(kbName[i], kbData[i])
			}
			fm.WriteString(dict["WelcomeToMainMenu"])
			fm.WriteChatId(user.Id)
		} else {
			panic(err)
		}
	} else {
		user.Request = "GoReg"
		ShowRules(user, fm)
	}
}
