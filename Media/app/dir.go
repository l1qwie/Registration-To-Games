package app

import (
	"Media/api/producer"
	"Media/app/handler"
	"Media/apptype"
	"fmt"
)

func Receiving(req *apptype.Request) *apptype.Response {
	producer.InterLogs("Start function Media.Receiving()", fmt.Sprintf("UserId: %d, req (*apptype.Request): %v", req.Id, req))
	res := new(apptype.Response)
	handler.MediaAct(req, res)
	return res
}
