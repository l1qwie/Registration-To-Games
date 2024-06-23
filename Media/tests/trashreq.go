package tests

import "Media/apptype"

// This module about trash functions
// I need it because I must test any kind of requests

func common1() *apptype.Request {
	req := mainReq()
	req.Req = "A:KLDALK:SD:LKASD:LA"
	req.Level = 0
	return req
}

func common2() *apptype.Request {
	req := mainReq()
	req.Req = "PAMAGITIIII"
	req.Level = 0
	return req
}

func common3() *apptype.Request {
	req := mainReq()
	req.Req = "PARASYATA HRU HRU HRU"
	req.Level = 1
	return req
}

func common4() *apptype.Request {
	req := mainReq()
	req.Req = "NU DAITE UJE zagruzit' hot' shto-to"
	req.Level = 1
	return req
}

func unOnetr3() *apptype.Request {
	req := mainReq()
	req.Req = "nu i cho?"
	req.Level = 3
	req.MediaDir = "unload"
	return req
}

func unOnetr4() *apptype.Request {
	req := mainReq()
	req.Req = "nu i cho teper'?"
	req.Level = 3
	req.MediaDir = "unload"
	return req
}

func upOnetr1() *apptype.Request {
	req := mainReq()
	req.Req = "NU DAITE UJE zagruzit' hot' shto-to"
	req.Level = 0
	//req.MediaDir = "upload"
	return req
}

func upOnetr2() *apptype.Request {
	req := mainReq()
	req.Req = "nu i cho?"
	req.Level = 0
	//req.MediaDir = "upload"
	return req
}

func upOnetr3() *apptype.Request {
	req := mainReq()
	req.Req = "PAMAGITIIII"
	req.Level = 2
	req.MediaDir = "upload"
	return req
}

func upOnetr4() *apptype.Request {
	req := mainReq()
	req.Req = "PARASYATA HRU HRU HRU"
	req.Level = 2
	req.MediaDir = "upload"
	return req
}

func upAfew1() *apptype.Request {
	req := mainReq()
	req.Req = "Dai mne MEDIA"
	req.Level = 3
	req.MediaDir = "upload"
	req.GameId = 1
	return req
}

func upAfew2() *apptype.Request {
	req := mainReq()
	req.Req = "NU DAITE UJE zagruzit' hot' shto-to"
	req.Level = 3
	req.MediaDir = "upload"
	req.GameId = 1
	return req
}

func upOnetr5() *apptype.Request {
	req := mainReq()
	req.Req = "Dai mne MEDIA"
	req.Level = 3
	req.MediaDir = "upload"
	req.GameId = 10
	return req
}

func upOnetr6() *apptype.Request {
	req := mainReq()
	req.Req = "NU DAITE UJE zagruzit' hot' shto-to"
	req.Level = 3
	req.MediaDir = "upload"
	req.GameId = 10
	return req
}
