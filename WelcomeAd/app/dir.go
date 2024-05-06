package app

import (
	"Welcome/app/handler"
	"Welcome/apptype"
)

func Receiving(req *apptype.Request) *apptype.Response {
	resp := new(apptype.Response)
	handler.WelcomeAct(req, resp)
	return resp
}
