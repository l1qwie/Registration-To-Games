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

// Trash func
func trash(lvl int) *apptype.Request {
	req := mainReq()
	req.Req = "()#!JKASKLDKLAJSLKJSALKD"
	req.Level = lvl
	req.IsChanged = "language"
	return req
}

// Trash func 2
func trash2(lvl int) *apptype.Request {
	req := mainReq()
	req.Req = "PRIVET GUYS!"
	req.Level = lvl
	req.IsChanged = "language"
	return req
}

// Makes req.Req for the Change Option part
func chOpt() *apptype.Request {
	req := mainReq()
	req.Req = "YA HOCHU CHEGO-NIBUD'"
	req.Level = 0
	return req
}

// Make req.Req for the Change Language part
func chLang() *apptype.Request {
	req := mainReq()
	req.Req = "en"
	req.Level = 2
	req.IsChanged = "language"
	return req
}
