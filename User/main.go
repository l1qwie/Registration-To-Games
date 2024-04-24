package main

import (
	"User/api/server"
	"User/fmtogram"
	"User/fmtogram/types"
)

func main() {
	types.Db = types.ConnectToDatabase(false)
	go server.Start()
	fmtogram.Start()
}
