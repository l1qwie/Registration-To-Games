package database

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/routine"
	"RegistrationToGames/fmtogram/types"
	"database/sql"

	_ "github.com/lib/pq"
)

func DeleteUser(userId int) {
	_, err := types.Db.Exec("DELETE FROM Users WHERE userId = $1", userId)
	if err != nil {
		panic(err)
	}
}

func UpdateLanguage(language string, userId int) {
	_, err := types.Db.Exec("UPDATE Users SET language = $1 WHERE userId = $2", language, userId)
	if err != nil {
		panic(err)
	}
}

func UpdateAction(action string, userId int) {
	_, err := types.Db.Exec("UPDATE Users SET action = $1 WHERE userId = $2", action, userId)
	if err != nil {
		panic(err)
	}
}

func UpdateLevel(level, userId int) {
	_, err := types.Db.Exec("UPDATE Users SET level = $1 WHERE userId = $2", level, userId)
	if err != nil {
		panic(err)
	}

}

func InsertGameId(gameId, userId int) {
	_, err := types.Db.Exec("UPDATE Users SET gameId = $1 WHERE userId = $2", gameId, userId)
	if err != nil {
		panic(err)
	}
}

func InsertSeats(seats, userId int) {
	_, err := types.Db.Exec("UPDATE Users SET seats = $1 WHERE userId = $2", seats, userId)
	if err != nil {
		panic(err)
	}
}

func InsertPayment(payment string, userId int) {
	_, err := types.Db.Exec("UPDATE Users SET payment = $1 WHERE userId = $2", payment, userId)
	if err != nil {
		panic(err)
	}
}

func CreateGame(gameId int) (err error) {
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
		VALUES ($1, 'volleyball', 20250212, 1200, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', 1)`, gameId)
	if err != nil {
		panic(err)
	}
	return err
}

func CreateNotFullMediaGame(gameId int) (err error) {
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES ($1, 'football', 20240212, 1200, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', -1)`, gameId)
	if err == nil {
		_, err = types.Db.Exec(`INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES ($1, 499, '!@#IOJSIOJE!@#**()!@#$*()SIOPE!@()#', 'photo', 0, 1)`, gameId)
	}
	return err
}

func CreateEmptyMediaGame(gameId int) (err error) {
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES ($1, 'volleyball', 20240212, 1200, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', -1)`, gameId)
	return err
}

func CreateSchedule() (err error) {
	var request string
	for i := 0; i < 4; i++ {
		if i == 0 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (0, 'volleyball', 20250212, 1200, 44, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'POUNDS', 1)`
		} else if i == 1 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (1, 'volleyball', 20260212, 1100, 34, 36.893445, 30.709591, 'Кладбище в Анталии', 10, 'POUNDS', 1)`
		} else if i == 2 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (2, 'football', 20250412, 1800, 14, 36.893445, 30.709591, 'Кладбище в Анталии', 1000, 'USD', 1)`
		} else if i == 3 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (3, 'volleyball', 20250202, 0800, 77, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'RUB', 1)`
		}
		_, err = types.Db.Exec(request)
		if err != nil {
			panic(err)
		}
	}
	return err
}

func FoundGame(gameId int) (detected bool) {
	var (
		rows    *sql.Rows
		err     error
		request string
		counter int
	)
	request = `SELECT COUNT(*) FROM Schedule WHERE gameId = $1`
	rows, err = types.Db.Query(request, gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&counter)
	if err != nil {
		panic(err)
	}
	if counter > 0 {
		detected = true
	}

	return detected
}

func DeleteGame(gameId int) {
	_, err := types.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
}

func DeleteMedia(gameId int) {
	_, err := types.Db.Exec("DELETE FROM MediaRepository WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
}

func AfterMainMenuCheckDb(userId int) {
	var (
		err  error
		user *bottypes.User
	)
	user = new(bottypes.User)
	user.Id = userId
	err = routine.DbRetrieveUser(user)
	if err != nil {
		panic(err)
	}
	user.ExMessageId, err = routine.SelectExMessageId(user.Id)
	if err != nil {
		panic(err)
	}
	if user.ExMessageId != 8888 {
		panic("user.ExMessageId != 8888")
	}
	if user.Id != 456 {
		panic("user.Id != 456")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Reg.GameId != 0 {
		panic("user.Reg.GameId != 0")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
	}
	if user.Level != 3 {
		panic("user.Level != 3")
	}
	if user.Reg.Seats != 0 {
		panic("user.Reg.Seats != 0")
	}
	if user.Reg.Payment != "" {
		panic("user.Reg.Payment != ``")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "" {
		panic("user.Media.Direction != ``")
	}
	if user.Media.Limit != 7 {
		panic("user.Media.Limit != 7")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
	if user.UserRec.Changeable != "" {
		panic("user.UserRec.Changeable != ``")
	}
	if user.UserRec.ActGame != "" {
		panic("user.UserRec.ActGame != ``")
	}
	if user.UserRec.WillChangeable != "" {
		panic("user.UserRec.WillChangeable != ``")
	}
	if user.UserRec.NewPay != "" {
		panic("user.UserRec.NewPay != ``")
	}
}
