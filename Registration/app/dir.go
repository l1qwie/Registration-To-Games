package app

import (
	"Registraion/app/handler"
	apptype "Registraion/apptype"
	"log"
)

func Receiving(req *apptype.Request) *apptype.Response {
	log.Printf("req.Level = %d, req.Req = %s, req.Id = %d, req.Act = %s", req.Level, req.Req, req.Id, req.Act)
	res := new(apptype.Response)
	handler.RegistrationAct(req, res)
	return res
}
