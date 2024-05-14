package app

import (
	"User/api/producer"
	"User/app/handler"
	"User/apptype"
	"User/fmtogram/formatter"
	"User/fmtogram/helper"
	"User/fmtogram/types"
	"fmt"
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
	producer.InterLogs("Start function User.Receiving()", fmt.Sprintf("UserId: %d, tr (*types.TelegramResponse): %v, mes (*types.MessageResponse): %v", user.Id, tr, mes))
	handler.DispatcherPhrase(user, fm)
	return fm
}
