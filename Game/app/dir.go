package app

import (
	"Game/app/handler"
	"Game/apptype"
)

func Receiving(req *apptype.Request) *apptype.Response {
	res := new(apptype.Response)
	handler.GameAction(req, res)
	return res
}
