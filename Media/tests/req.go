package tests

import (
	"Media/apptype"

	"github.com/l1qwie/Fmtogram/types"
)

// All defaul data makes here
func mainReq() *apptype.Request {
	return &apptype.Request{
		Id:          499,
		Level:       0,
		Language:    "ru",
		Act:         "photos and videos",
		Limit:       7,
		LaunchPoint: 0,
		Req:         "",
		GameId:      0,
		Status:      false,
	}
}

// Make data for a request for the first level
func chMediaDir() *apptype.Request {
	req := mainReq()
	req.Req = "!OK!"
	req.Level = 0
	req.Status = true
	return req
}

// Make data for a request to choose Unload
func chMUnload() *apptype.Request {
	req := mainReq()
	req.Req = "unload"
	req.Level = 1
	return req
}

// Make data for a request to unload just a one picture (fileId)
func unJustOne() *apptype.Request {
	req := mainReq()
	req.Req = "10"
	req.Level = 3
	req.MediaDir = "unload"
	return req
}

func chMUpload() *apptype.Request {
	req := mainReq()
	req.Req = "10"
	req.Level = 2
	req.MediaDir = "upload"
	return req
}

func upJustOne() *apptype.Request {
	req := mainReq()
	req.Req = "TAKE MY PHOTO"
	req.Level = 3
	req.GameId = 10
	req.MediaDir = "upload"
	req.FileId = "!@#UIO!@#IOJJKLASEDKLKL#IO!JKLASJKL13419"
	req.TypeOffile = "photo"
	req.MediaCounter = 1
	return req
}

func unJustAfew() *apptype.Request {
	req := mainReq()
	req.Req = "1"
	req.MediaDir = "unload"
	req.Level = 3
	return req
}

func chMUploadAfew() *apptype.Request {
	req := mainReq()
	req.Req = "1"
	req.MediaDir = "upload"
	req.Level = 2
	return req
}

func upJustAfew() *apptype.Request {
	req := mainReq()
	req.Req = "take please my full media"
	req.Level = 3
	req.MediaDir = "upload"
	req.Media = []types.Media{{Type: "photo", Media: "!@#IOJSJE!@#**()!@#$*()SIOPE!@()#"},
		{Type: "photo", Media: "!@#IOJSIOJE!@#**()!@HFF#$*()SIOPE$#@$!@()#"},
		{Type: "photo", Media: "!@#IOJSIOJE!$@!$!@#@#**()!@#$*()SIOPE!@()#"}}
	req.MediaCounter = 3
	req.GameId = 1
	return req
}
