package app

import (
	"Welcome/api/producer"
	"Welcome/app/handler"
	"Welcome/apptype"
	"fmt"
)

func Receiving(req *apptype.Request) *apptype.Response {
	producer.InterLogs("Start function Receiving()", fmt.Sprintf("AdminId: %d, req (*apptype.Request): %v", req.Id, req))
	resp := new(apptype.Response)
	handler.WelcomeAct(req, resp)
	return resp
}
