package main

import (
	"User/api/server"
	"User/app"
	"User/apptype"

	"github.com/l1qwie/Fmtogram/fmtogram"
	"github.com/l1qwie/Fmtogram/types"
)

func main() {
	//defer grpc.DelEnv()
	apptype.Db = apptype.ConnectToDatabase()
	fmtogram.StartFunc = app.Receiving
	types.BotID = apptype.Id
	//client := grpc.AddCleint()
	//grpc.CreateSchRedis(client)
	//grpc.CreateEnv()

	go server.Start()
	fmtogram.Start()

	//defer grpc.DelSchRedis(client)
}

/*
func main() {
	grpc.SendMediaGroup()
}
*/
