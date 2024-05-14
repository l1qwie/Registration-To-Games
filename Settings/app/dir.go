package app

import (
	"Settings/api/producer"
	"Settings/app/handler"
	"Settings/apptype"
	"fmt"
)

func Receiving(req *apptype.Request) *apptype.Response {
	producer.InterLogs("Start function Settings.Receiving()", fmt.Sprintf("UserId: %d, req (*apptype.Request): %v", req.Id, req))
	res := new(apptype.Response)
	handler.SettingsAct(req, res)
	return res
}
