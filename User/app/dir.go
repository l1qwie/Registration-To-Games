package app

import (
	"User/app/handler"
	"User/apptype"
	"User/fmtogram/formatter"
	"User/fmtogram/helper"
	"User/fmtogram/types"
)

func Receiving(tr *types.TelegramResponse, mes *types.MessageResponse) *formatter.Formatter {
	user := new(apptype.User)
	fm := new(formatter.Formatter)
	user.Request = helper.ReturnText(tr)
	user.Id = helper.ReturnChatId(tr)
	user.Language = helper.ReturnLanguage(tr)
	user.Photo, _ = helper.ReturnPhotoResp(mes)
	user.Video, _ = helper.ReturnVideoResp(mes)
	user.MediaF, _ = helper.ReturnMediaResp(mes)
	user.Media.MediaF, _ = helper.ReturnMediaReq(tr)
	user.Media.Photo, _ = helper.ReturnPhotoReq(tr)
	user.Media.Video, _ = helper.ReturnVideoReq(tr)
	user.ExMessageId, _ = helper.ReturnMessageId(mes)
	handler.DispatcherPhrase(user, fm)
	return fm
}
