package app

import (
	"Welcome/api/producer"
	"Welcome/app/handler"
	"Welcome/apptype"
	"fmt"
)

func Receiving(req *apptype.Request) *apptype.Response {
	producer.InterLogs("Start function WelcomeUs.Receiving()", fmt.Sprintf("UserId: %d, req (*types.Request): %v", req.Id, req))
	resp := new(apptype.Response)
	handler.WelcomeAct(req, resp)
	return resp
}
