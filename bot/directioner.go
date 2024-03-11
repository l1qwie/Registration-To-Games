package bot

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/routine"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/helper"
	"RegistrationToGames/fmtogram/types"
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
	user.ExMessageId, _ = helper.ReturnMessageId(mes)
	user.PhotosFileId, _ = helper.ReturnPhotosFileId(mes)
	user.VideosFileId, _ = helper.ReturnVideosFileId(mes)

	routine.DispatcherPhrase(user, fm)

	return fm
}
