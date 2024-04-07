package app

import (
	"Welcome/app/handler"
	"Welcome/types"
)

func Receiving(req *types.Request) *types.Response {
	resp := new(types.Response)
	handler.WelcomeAct(req, resp)
	return resp
}
