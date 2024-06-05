package tests

import "Game/apptype"

func mainReq() *apptype.Request {
	return &apptype.Request{
		Id:          1111,
		Level:       0,
		Req:         "",
		Language:    "ru",
		Act:         "game",
		LaunchPoint: 0,
		Limit:       7,
		Sport:       "",
		Date:        0,
		Time:        0,
		Seats:       0,
		Price:       0,
		Currency:    "",
		Link:        "",
		Address:     "",
		Status:      true,
	}
}

func sendHello() *apptype.Request {
	req := mainReq()
	req.Req = "HELOO!!!"
	req.Level = 0
	return req
}

func sendSport() *apptype.Request {
	req := mainReq()
	req.Req = "volleyball"
	req.Level = 2
	return req
}

func sendDate() *apptype.Request {
	req := mainReq()
	req.Req = "09-12-2024"
	req.Sport = "volleyball"
	req.Level = 3
	return req
}

func sendTime() *apptype.Request {
	req := mainReq()
	req.Req = "19-00"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Level = 4
	return req
}

func sendSeats() *apptype.Request {
	req := mainReq()
	req.Req = "15"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Level = 5
	return req
}

func sendPrice() *apptype.Request {
	req := mainReq()
	req.Req = "1000"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Level = 6
	return req
}

func sendCurrency() *apptype.Request {
	req := mainReq()
	req.Req = "RUB"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Level = 7
	return req
}

func sendLink() *apptype.Request {
	req := mainReq()
	req.Req = "https://www.google.com/maps?q=36.893445,30.709591"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	req.Level = 8
	return req
}

func sendAddress() *apptype.Request {
	req := mainReq()
	req.Req = "Игровая Площадка"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	req.Link = "https://www.google.com/maps?q=36.893445,30.709591"
	req.Level = 9
	return req
}

func sendSave() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Seats = 15
	req.Price = 1000
	req.Currency = "RUB"
	req.Link = "https://www.google.com/maps?q=36.893445,30.709591"
	req.Address = "Игровая Площадка"
	req.Level = 10
	return req
}
