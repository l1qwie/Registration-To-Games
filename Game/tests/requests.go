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

func sendDiretionChange() *apptype.Request {
	req := mainReq()
	req.Req = "change"
	req.Level = 1
	return req
}

func sendDiretionDel() *apptype.Request {
	req := mainReq()
	req.Req = "delete"
	req.Level = 1
	return req
}

func choseSport() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "sport"
	req.GameId = 6667
	req.Level = 3
	return req
}

func choseDate() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "date"
	req.GameId = 6667
	req.Level = 3
	return req
}

func choseTime() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "time"
	req.GameId = 6667
	req.Level = 3
	return req
}

func choseSeats() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "seats"
	req.GameId = 6667
	req.Level = 3
	return req
}

func chosePrice() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "price"
	req.GameId = 6667
	req.Level = 3
	return req
}

func choseCurrency() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "currency"
	req.GameId = 6667
	req.Level = 3
	return req
}

func choseLink() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "link"
	req.GameId = 6667
	req.Level = 3
	return req
}

func choseAddress() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "address"
	req.GameId = 6667
	req.Level = 3
	return req
}

func chsendSport() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "football"
	req.GameId = 6667
	req.Changeable = "sport"
	req.Level = 4
	return req
}

func chsendDate() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "12-12-2024"
	req.GameId = 6667
	req.Changeable = "date"
	req.Level = 4
	return req
}

func chsendTime() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "20:00"
	req.GameId = 6667
	req.Changeable = "time"
	req.Level = 4
	return req
}

func chsendSeats() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "99"
	req.GameId = 6667
	req.Changeable = "seats"
	req.Level = 4
	return req
}

func chsendPrice() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "199"
	req.GameId = 6667
	req.Changeable = "price"
	req.Level = 4
	return req
}

func chsendCurrency() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "USDT"
	req.GameId = 6667
	req.Changeable = "currency"
	req.Level = 4
	return req
}

func chsendLink() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "https://www.google.com/maps/place/31%C2%B051'47.5%22N+34%C2%B051'50.8%22E/@31.863181,34.8626321,17"
	req.GameId = 6667
	req.Changeable = "link"
	req.Level = 4
	return req
}

func chsendAddress() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "Кудыкина Гора"
	req.GameId = 6667
	req.Changeable = "address"
	req.Level = 4
	return req
}

func sendGame() *apptype.Request {
	req := mainReq()
	req.Direction = "change"
	req.Req = "6667"
	req.Level = 2
	return req
}

func sendDelGame() *apptype.Request {
	req := mainReq()
	req.Direction = "delete"
	req.Req = "6667"
	req.Level = 2
	return req
}

func sendSport() *apptype.Request {
	req := mainReq()
	req.Req = "volleyball"
	req.Direction = "create"
	req.Level = 2
	return req
}

func sendDate() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "09-12-2024"
	req.Sport = "volleyball"
	req.Level = 3
	return req
}

func sendTime() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "19-00"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Level = 4
	return req
}

func sendSeats() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
	req.Req = "15"
	req.Sport = "volleyball"
	req.Date = 20241209
	req.Time = 1900
	req.Level = 5
	return req
}

func sendPrice() *apptype.Request {
	req := mainReq()
	req.Direction = "create"
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
	req.Direction = "create"
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
	req.Direction = "create"
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
	req.Direction = "create"
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
	req.Direction = "create"
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

func sendChSaveSport() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Direction = "change"
	req.Changeable = "sport"
	req.GameId = 6667
	req.Sport = "football"
	req.Level = 5
	return req
}
func sendChSaveDate() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Direction = "change"
	req.Changeable = "date"
	req.GameId = 6667
	req.Date = 20241212
	req.Level = 5
	return req
}

func sendChSaveTime() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Direction = "change"
	req.Changeable = "time"
	req.GameId = 6667
	req.Time = 2000
	req.Level = 5
	return req
}

func sendChSaveSeats() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Direction = "change"
	req.Changeable = "seats"
	req.GameId = 6667
	req.Seats = 99
	req.Level = 5
	return req
}

func sendChSavePrice() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Direction = "change"
	req.Changeable = "price"
	req.GameId = 6667
	req.Price = 199
	req.Level = 5
	return req
}

func sendChSaveCurrency() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Direction = "change"
	req.Changeable = "currency"
	req.GameId = 6667
	req.Currency = "USDT"
	req.Level = 5
	return req
}

func sendChSaveLink() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Direction = "change"
	req.Changeable = "link"
	req.GameId = 6667
	req.Link = "https://www.google.com/maps/place/31%C2%B051'47.5%22N+34%C2%B051'50.8%22E/@31.863181,34.8626321,17"
	req.Level = 5
	return req
}

func sendChSaveAddress() *apptype.Request {
	req := mainReq()
	req.Req = "save"
	req.Direction = "change"
	req.Changeable = "address"
	req.GameId = 6667
	req.Address = "Кудыкина Гора"
	req.Level = 5
	return req
}
