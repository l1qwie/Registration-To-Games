package main

import (
	"User/api/server"
	"User/fmtogram"
	"User/fmtogram/types"
)

func main() {
	//defer grpc.DelEnv()
	types.Db = types.ConnectToDatabase(true)
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
