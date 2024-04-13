package tests

import "Settings/apptype"

// All defaul data makes here
func mainReq() *apptype.Request {
	return &apptype.Request{
		Id:          899,
		Level:       0,
		Language:    "ru",
		Act:         "settings",
		Limit:       7,
		LaunchPoint: 0,
		Req:         "",
		IsChanged:   "",
		GameId:      0,
		Connection:  apptype.Db,
	}
}

// Makes data for the Change Option part
func chOpt() *apptype.Request {
	req := mainReq()
	req.Req = "YA HOCHU CHEGO-NIBUD'"
	req.Level = 0
	return req
}

// Make data and for the Change Language part
func chLang() *apptype.Request {
	req := mainReq()
	req.Req = "en"
	req.Level = 2
	req.IsChanged = "language"
	return req
}

// Make data and for the Choose the Way part
func chRec() *apptype.Request {
	req := mainReq()
	req.Level = 1
	req.Req = "records"
	return req
}

// Make data and for the Choose a Game part
func delGameId() *apptype.Request {
	req := mainReq()
	req.Level = 2
	req.IsChanged = "records"
	req.Req = "1"
	return req
}

// Make data and for the Choose a Changeable Option part
func del() *apptype.Request {
	req := mainReq()
	req.Level = 3
	req.IsChanged = "records"
	req.GameId = 1
	req.Req = "del"
	return req
}
