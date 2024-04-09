package app

import (
	"Schedule/app/handler"
	"Schedule/apptype"
	"log"
)

func Receiving(req *apptype.Request) *apptype.Response {
	log.Printf("req.Id = %d, req.Act = %s", req.Id, req.Act)
	res := new(apptype.Response)
	return handler.Schedule(req, res)
}
