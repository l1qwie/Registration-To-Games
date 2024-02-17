package bot

import (
	"registrationtogames/bot/routine"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/fmtogram/helper"
	"registrationtogames/fmtogram/types"
)

func Receiving(tr *types.TelegramResponse) *formatter.Formatter {
	var (
		fm   formatter.Formatter
		user routine.User
	)
	user.Request = helper.ReturnText(tr)
	user.Id = helper.ReturnChatId(tr)
	user.Language = helper.ReturnLanguage(tr)

	user.DispatcherPhrase(&fm)

	return &fm
}
