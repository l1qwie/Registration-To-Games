package welcome

import (
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/dictionary"
	"registrationtogames/fmtogram/formatter"
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
