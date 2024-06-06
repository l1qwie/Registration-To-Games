package tests

import (
	"Game/apptype"
)

func trash() *apptype.Request {
	req := mainReq()
	req.Req = "HELOO!!!"
	req.Level = 0
	return req
}
func trash1() *apptype.Request {
	req := mainReq()
	req.Req = "dAASDASW!!@#"
	req.Level = 0
	return req
}
func chtrash2() *apptype.Request {
	req := mainReq()
	req.Req = "я хочу питсы"
	req.Level = 1
	return req
}

func chtrash3() *apptype.Request {
	req := mainReq()
	req.Req = "я хочу питсы"
	req.Level = 1
	return req
}

func chtrash4() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "я хочу питсы"
	req.Level = 2
	return req
}

func chtrash5() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "я хочу питсы"
	req.Level = 2
	return req
}

func chtrash6() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.GameId = 6667
	req.Req = "я хочу питсы"
	req.Level = 3
	return req
}

func chtrash7() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.GameId = 6667
	req.Req = "я хочу питсы"
	req.Level = 3
	return req
}

func chtrash8() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.GameId = 6667
	req.Changeable = "sport"
	req.Req = "я хочу питсы"
	req.Level = 4
	return req
}

func chtrash9() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.GameId = 6667
	req.Changeable = "sport"
	req.Req = "я хочу питсы"
	req.Level = 4
	return req
}

func chtrash10() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.GameId = 6667
	req.Changeable = "sport"
	req.Sport = "football"
	req.Req = "я хочу питсы"
	req.Level = 5
	return req
}

func chtrash11() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.GameId = 6667
	req.Changeable = "sport"
	req.Sport = "football"
	req.Req = "я хочу питсы"
	req.Level = 5
	return req
}

func trash2() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 2
	return req
}
func trash3() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 2
	return req
}
func trash4() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 3
	req.Sport = "volleyball"
	return req
}
func trash5() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 3
	req.Sport = "volleyball"
	return req
}
func trash6() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 4
	req.Sport = "volleyball"
	req.Date = 20241209
	return req
}
func trash7() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 4
	req.Sport = "volleyball"
	req.Date = 20241209
	return req
}
func trash8() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 5
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	return req
}
func trash9() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 5
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	return req
}
func trash10() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 6
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	return req
}
func trash11() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 6
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	return req
}
func trash12() *apptype.Request {
	req := mainReq()
	req.Req = "я хочу питсы"
	req.Level = 7
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	return req
}
func trash13() *apptype.Request {
	req := mainReq()
	req.Req = "я хочу питсы"
	req.Level = 7
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	return req
}
func trash14() *apptype.Request {
	req := mainReq()
	req.Req = "я хочу питсы"
	req.Level = 8
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	return req
}
func trash15() *apptype.Request {
	req := mainReq()
	req.Req = "я хочу питсы"
	req.Level = 8
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	return req
}
func trash16() *apptype.Request {
	req := mainReq()
	req.Req = "я хочу питсы"
	req.Level = 9
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	req.Link = "https://www.google.com/maps?q=36.893445,30.709591"
	return req
}
func trash17() *apptype.Request {
	req := mainReq()
	req.Req = "я хочу питсы"
	req.Level = 9
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	req.Link = "https://www.google.com/maps?q=36.893445,30.709591"
	return req
}
func trash18() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 10
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	req.Link = "https://www.google.com/maps?q=36.893445,30.709591"
	req.Address = "Игровая Площадка"
	return req
}
func trash19() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "я хочу питсы"
	req.Level = 10
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	req.Link = "https://www.google.com/maps?q=36.893445,30.709591"
	req.Address = "Игровая Площадка"
	return req
}
