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

// Makes data the Change Option part
func chOpt() *apptype.Request {
	req := mainReq()
	req.Req = "YA HOCHU CHEGO-NIBUD'"
	req.Level = 0
	return req
}

// Make data for the Change Language part
func chLang() *apptype.Request {
	req := mainReq()
	req.Req = "en"
	req.Level = 2
	req.IsChanged = "language"
	return req
}

// Make data for the Choose the Way part
func chRec() *apptype.Request {
	req := mainReq()
	req.Level = 1
	req.Req = "records"
	return req
}

// Make data for the Choose a Game part
func delGameId() *apptype.Request {
	req := mainReq()
	req.Level = 2
	req.IsChanged = "records"
	req.Req = "1"
	return req
}

func delGameIdCh() *apptype.Request {
	req := mainReq()
	req.Level = 2
	req.IsChanged = "records"
	req.Req = "2" // if you want to change
	return req
}

// Make data for the Choose a Changeable Option part
func del() *apptype.Request {
	req := mainReq()
	req.Level = 3
	req.IsChanged = "records"
	req.GameId = 1
	req.Req = "del"
	return req
}

// Make data for the Choose an ChangeOption part
func change() *apptype.Request {
	req := mainReq()
	req.Level = 3
	req.Req = "change"
	req.GameId = 2
	req.IsChanged = "records"
	return req
}

// Make data for the Choose an change option of the game
func seats() *apptype.Request {
	req := mainReq()
	req.Level = 4
	req.Req = "myseats"
	req.GameId = 2
	req.IsChanged = "records"
	return req
}

// Make data to send num of seats
func numOfSeats() *apptype.Request {
	req := mainReq()
	req.Level = 5
	req.Req = "6"
	req.GameId = 2
	req.IsChanged = "myseats"
	return req
}

// Make data to send paymethod
func payment() *apptype.Request {
	req := mainReq()
	req.Level = 4
	req.Req = "payment"
	req.GameId = 2
	req.IsChanged = "records"
	return req
}
