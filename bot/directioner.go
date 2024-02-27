package bot

import (
	"registrationtogames/bot/bottypes"
	"registrationtogames/bot/routine"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/fmtogram/helper"
	"registrationtogames/fmtogram/types"
)

func Receiving(tr *types.TelegramResponse, mes *types.MessageResponse) *formatter.Formatter {
	var (
		fm   *formatter.Formatter
		user *bottypes.User
	)
	user = new(bottypes.User)
	fm = new(formatter.Formatter)
	user.Request = helper.ReturnText(tr)
	user.Id = helper.ReturnChatId(tr)
	user.Language = helper.ReturnLanguage(tr)
	user.ExMessageId, _ = helper.MessageId(mes)

	routine.DispatcherPhrase(user, fm)

	return fm
}
