package app

import (
	"Welcome/api/producer"
	"Welcome/app/handler"
	"Welcome/types"
	"fmt"
)

func Receiving(req *types.Request) *types.Response {
	producer.InterLogs("Start function WelcomeUs.Receiving()", fmt.Sprintf("UserId: %d, req (*types.Request): %v", req.Id, req))
	resp := new(types.Response)
	handler.WelcomeAct(req, resp)
	return resp
}
