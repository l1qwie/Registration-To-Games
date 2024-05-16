package tests

import "Welcome/apptype"

// All defaul data makes here
func mainReq() *apptype.Request {
	return &apptype.Request{
		Id:       1999,
		Level:    0,
		Language: "ru",
		Act:      "registration",
		Req:      "",
		Status:   false,
	}
}

func hi() *apptype.Request {
	req := mainReq()
	req.Req = "PRIVET VSEM"
	req.Level = 0
	req.Status = true
	return req
}

func letsgo() *apptype.Request {
	req := mainReq()
	req.Req = "GoReg"
	req.Level = 1
	return req
}

func rulesAgree() *apptype.Request {
	req := mainReq()
	req.Req = "Start"
	req.Level = 2
	return req
}

// TRASH REQUESTS
func trash1() *apptype.Request {
	req := mainReq()
	req.Req = "StaASD:L<A>SM<D>ASD<>Mrt"
	req.Level = 0
	return req
}

func trash2() *apptype.Request {
	req := mainReq()
	req.Req = "1231254123oi0p1234890190"
	req.Level = 0
	return req
}

func trash3() *apptype.Request {
	req := mainReq()
	req.Req = "HELLLSDOP"
	req.Level = 1
	return req
}

func trash4() *apptype.Request {
	req := mainReq()
	req.Req = "KLAJDJKLASKJLDJKLASDKL"
	req.Level = 1
	return req
}

func trash5() *apptype.Request {
	req := mainReq()
	req.Req = "((((((((Start))))))))"
	req.Level = 2
	return req
}

func trash6() *apptype.Request {
	req := mainReq()
	req.Req = "||||||||||||||||||||||||||DDMJDJNMSMD"
	req.Level = 2
	return req
}
