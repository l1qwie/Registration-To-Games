package app

import (
	"Schedule/api/producer"
	"Schedule/app/handler"
	"Schedule/apptype"
	"fmt"
)

func Receiving(req *apptype.Request) *apptype.Response {
	producer.InterLogs("Start function Schedule.Receiving()", fmt.Sprintf("UserId: %d, req (*apptype.Request): %v", req.Id, req))
	res := new(apptype.Response)
	return handler.Schedule(req, res)
}
