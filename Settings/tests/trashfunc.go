package tests

// This module about trash functions
// I need it because I must test any kind of requests
import "Settings/apptype"

func commontrash() *apptype.Request {
	req := mainReq()
	req.Req = "()#!JKASKLDKLAJSLKJSALKD"
	req.Level = 0
	return req
}

func commontrash2() *apptype.Request {
	req := mainReq()
	req.Req = ":ASDKL::KLDKOL(((((9999)))))"
	req.Level = 0
	return req
}

func langtr2() *apptype.Request {
	req := mainReq()
	req.Req = "Zachem tak mnogo musornih funkcii?"
	req.Level = 2
	req.IsChanged = "language"
	return req
}

func langtr3() *apptype.Request {
	req := mainReq()
	req.Req = "Ya prosto hochu pomen'yat' yazik"
	req.Level = 2
	req.IsChanged = "language"
	return req
}

func dGametr2() *apptype.Request {
	req := mainReq()
	req.Req = "Ya prosto hochu pomen'yat' yazik"
	req.Level = 1
	return req
}

func dGametr3() *apptype.Request {
	req := mainReq()
	req.Req = "Zachem tak mnogo musornih funkcii?"
	req.Level = 1
	return req
}

func dGametr4() *apptype.Request {
	req := mainReq()
	req.Req = "QQ pupsiki?"
	req.Level = 2
	req.IsChanged = "records"
	return req
}

func dGametr5() *apptype.Request {
	req := mainReq()
	req.Req = "L:ASDL:K:AASDASDASDASD"
	req.Level = 2
	req.IsChanged = "records"
	return req
}

func dGametr6() *apptype.Request {
	req := mainReq()
	req.Level = 3
	req.Req = "NU PACHIMUUUUUUUUUUU"
	req.IsChanged = "records"
	req.GameId = 1
	return req
}

func dGametr7() *apptype.Request {
	req := mainReq()
	req.Level = 3
	req.Req = "Cho za TREASHSHSHSH"
	req.IsChanged = "records"
	req.GameId = 1
	return req
}

func chSeatstr2() *apptype.Request {
	req := mainReq()
	req.Level = 1
	req.Req = "KL:ASDKL:ASDL:KASDL:ASD"
	return req
}

func chSeatstr3() *apptype.Request {
	req := mainReq()
	req.Level = 1
	req.Req = "PAMAGITIIII"
	return req
}

func chSeatstr4() *apptype.Request {
	req := mainReq()
	req.Level = 2
	req.Req = "Cho za TREASHSHSHSH"
	req.IsChanged = "records"
	return req
}

func chSeatstr5() *apptype.Request {
	req := mainReq()
	req.Level = 2
	req.Req = "NU PACHIMUUUUUUUUUUU"
	req.IsChanged = "records"
	return req
}

func chSeatstr6() *apptype.Request {
	req := mainReq()
	req.Level = 3
	req.Req = "Zachem tak mnogo musornih funkcii?"
	req.IsChanged = "records"
	req.GameId = 2
	return req
}

func chSeatstr7() *apptype.Request {
	req := mainReq()
	req.Level = 3
	req.Req = ":ASDKL::KLDKOL(((((9999)))))"
	req.IsChanged = "records"
	req.GameId = 2
	return req
}

func chSeatstr8() *apptype.Request {
	req := mainReq()
	req.Level = 4
	req.Req = ":A(((9999)))))"
	req.IsChanged = "records"
	req.GameId = 2
	return req
}

func chPay1() *apptype.Request {
	req := mainReq()
	req.Level = 4
	req.Req = ":A(((9999)))))"
	req.IsChanged = "records"
	req.GameId = 2
	return req
}

func chSeatstr9() *apptype.Request {
	req := mainReq()
	req.Level = 4
	req.Req = "AKSDKASDKASDKASKDKDASKDDASKDAKSDAKSDAKSDKAS"
	req.IsChanged = "records"
	req.GameId = 2
	return req
}

func chPay() *apptype.Request {
	req := mainReq()
	req.Level = 4
	req.Req = "ASDLKMFSKLJDFJKL"
	req.IsChanged = "records"
	req.GameId = 2
	return req
}

func chSeatstr10() *apptype.Request {
	req := mainReq()
	req.Level = 5
	req.Req = "Cho za TREASHSHSHSH"
	req.IsChanged = "myseats"
	req.GameId = 2
	return req
}

func chSeatstr11() *apptype.Request {
	req := mainReq()
	req.Level = 5
	req.Req = "Ya prosto hochu pomen'yat' yazik"
	req.IsChanged = "myseats"
	req.GameId = 2
	return req
}
