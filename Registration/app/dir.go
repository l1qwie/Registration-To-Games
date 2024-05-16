package app

import (
	"Registration/api/producer"
	"Registration/app/handler"
	"Registration/apptype"
	"fmt"
)

func Receiving(req *apptype.Request) *apptype.Response {
	producer.InterLogs("Start function Registration.Receiving()", fmt.Sprintf("UserId: %d, req (*apptype.Request): %v", req.Id, req))
	res := new(apptype.Response)
	handler.RegistrationAct(req, res)
	return res
}
