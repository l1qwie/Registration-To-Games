package app

import (
	"User/app/handler"
	"User/apptype"
	"User/fmtogram/formatter"
	"User/fmtogram/helper"
	"User/fmtogram/types"
	"log"
)

func Receiving(tr *types.TelegramResponse, mes *types.MessageResponse) *formatter.Formatter {
	var (
		fm   *formatter.Formatter
		user *apptype.User
		err  error
	)
	user = new(apptype.User)
	fm = new(formatter.Formatter)
	user.Request = helper.ReturnText(tr)
	user.Id = helper.ReturnChatId(tr)
	user.Language = helper.ReturnLanguage(tr)
	user.Photo, err = helper.ReturnPhotoResp(mes)
	user.Video, err = helper.ReturnVideoResp(mes)
	user.MediaF, err = helper.ReturnMediaResp(mes)
	user.Media.MediaF, err = helper.ReturnMediaReq(tr)
	user.Media.Photo, err = helper.ReturnPhotoReq(tr)
	user.Media.Video, err = helper.ReturnVideoReq(tr)
	user.ExMessageId, err = helper.ReturnMessageId(mes)
	log.Printf("You have an error: %s", err)
	handler.DispatcherPhrase(user, fm)
	return fm
}
