package main

import (
	"User/fmtogram"
	"User/fmtogram/types"
)

func main() {
	types.Db = types.ConnectToDatabase(false)
	fmtogram.Start()
}
