package database

import (
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

func CreateUser(userId int) {
	_, err := types.Db.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", userId, "registration", "ru", 0)
	if err != nil {
		panic(err)
	}
}

func DeleteSchedule() {
	_, err := types.Db.Exec("DELETE FROM Schedule WHERE gameId = 0 OR gameId = 1 OR gameId = 2 OR gameId = 3 OR gameId = 4")
	if err != nil {
		panic(err)
	}
}

func DeleteMediaSchedule() {
	_, err := types.Db.Exec("DELETE FROM MediaRepository WHERE gameId = 1 OR gameId = 2 OR gameId = 3 OR gameId = 4")
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

func CreateUserSchedule(userId int) {
	_, err := types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (1, 'volleyball', 20250212, 1200, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'POUNDS', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec(`INSERT INTO GamesWithSchedule (userId, gameId, seats, payment, statuspayment, status)
		VALUES ($1, 1, 3, "card", 1, 0)`, userId)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (2, 'volleyball', 20251212, 1100, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec(`INSERT INTO GamesWithSchedule (userId, gameId, seats, payment, statuspayment, status)
		VALUES ($1, 2, 6, "card", 0, 0)`, userId)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (3, 'volleyball', 20250612, 1000, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'RUB', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec(`INSERT INTO GamesWithSchedule (userId, gameId, seats, payment, statuspayment, status)
		VALUES ($1, 3, 7, "cash", 0, 0)`, userId)
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

func FillMEdiaSchedule() {
	var request string
	for i := 0; i < 8; i++ {
		if i == 0 {
			request = `INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES (1, 499, '!@#IOJSJE!@#**()!@#$*()SIOPE!@()#', 'photo', 3, 1)`
		} else if i == 1 {
			request = `INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES (2, 499, '!@#IOJSJE!@#**()!@#$*()SIOPE!@(#!@*(IOOI)', 'photo', 2, 1)`
		} else if i == 2 {
			request = `INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES (3, 499, '!@#IOJSIOJASDGE!@#**()!@#$*()SIOPE!@()#', 'photo', 2, 1)`
		} else if i == 3 {
			request = `INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES (4, 499, '!@#IOJSIOJESL:DFK**()!@#$*()SIOPE!@()#', 'photo', 1, 1)`
		} else if i == 4 {
			request = `INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES (1, 499, '!@#IOJSIOJE!@#**()!@HFF#$*()SIOPE$#@$!@()#', 'photo', 3, 1)`
		} else if i == 5 {
			request = `INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES (2, 499, '!@#IOJSIOJE!@#**()!()()998890-@#$*()SIOPE!@()#', 'photo', 2, 1)`
		} else if i == 6 {
			request = `INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES (3, 499, '!@#IOJSIOJE!@#**ASDG()!@#$*()SIOPE!@()#', 'photo', 2, 1)`
		} else if i == 7 {
			request = `INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES (1, 499, '!@#IOJSIOJE!$@!$!@#@#**()!@#$*()SIOPE!@()#', 'photo', 3, 1)`
		}
		_, err := types.Db.Exec(request)
		if err != nil {
			panic(err)
		}
	}
}

func CreateMediaShedule() {
	var request string
	for i := 1; i < 5; i++ {
		if i == 1 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (1, 'volleyball', 20260212, 1100, 34, 36.893445, 30.709591, 'Кладбище в Анталии', 10, 'POUNDS', -1)`
		} else if i == 2 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (2, 'football', 20250412, 1800, 14, 36.893445, 30.709591, 'Кладбище в Анталии', 1000, 'USD', -1)`
		} else if i == 3 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (3, 'volleyball', 20250202, 0800, 77, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'RUB', -1)`
		} else if i == 4 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
				VALUES (4, 'volleyball', 20250202, 0800, 77, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'RUB', -1)`
		}
		_, err := types.Db.Exec(request)
		if err != nil {
			panic(err)
		}
	}
}

func findMediaSchedule() bool {
	var (
		rows *sql.Rows
		err  error
	)
	querys := []string{
		"SELECT COUNT(*) FROM MediaRepository WHERE gameId = 1 AND fileId = '!@#UIO!@#IOJJKLASEDKLKL#IO!JKLASJKL13419' AND type = 'photo' AND counter = 4",
		"SELECT COUNT(*) FROM MediaRepository WHERE gameId = 1 AND fileId = '!@#UIO!@#IOJJKLIO!JKLASJKL13419' AND type = 'photo' AND counter = 4",
		"SELECT COUNT(*) FROM MediaRepository WHERE gameId = 1 AND fileId = 'IJ!#JJKLASERJKLIOPEIO*()%*()IOPSDKL:ASDOPK#I!#~!@31313' AND type = 'photo' AND counter = 4",
		"SELECT COUNT(*) FROM MediaRepository WHERE gameId = 1 AND fileId = 'H!UIO@#HUI!@HASJKLDIOJKL#*()!@()_IOASDEUIO%()_)_' AND type = 'photo' AND counter = 4",
	}
	counter := 1
	for i := 0; i < 3 && counter > 0; i++ {
		rows, err = types.Db.Query(querys[i])
		if err != nil {
			panic(err)
		}
		rows.Next()
		err = rows.Scan(&counter)
		if err != nil {
			panic(err)
		}
	}
	defer rows.Close()
	return counter > 0
}

func checkGamesWithUsersTable(userId, gameId, seats int, payment string) bool {
	var (
		rows    *sql.Rows
		err     error
		request string
		counter int
	)
	request = `SELECT COUNT(*) FROM GamesWithUsers WHERE userId = $1 AND gameId = $2 AND seats = $3 AND payment = $4` //AND statuspayment = 0 AND status = 0
	rows, err = types.Db.Query(request, userId, gameId, seats, payment)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&counter)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return counter > 0
}

func CreateEmptyMediaGame(gameId int) (err error) {
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES ($1, 'volleyball', 20240212, 1200, 55, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', -1)`, gameId)
	return err
}

func CreateUserScehdule(userId int) {
	_, err := types.Db.Exec("INSERT INTO GamesWithUsers (userId, gameId, seats, payment, statuspayment, status) VALUES ($1, 0, 3, 'card', 1, 1)", userId)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec("INSERT INTO GamesWithUsers (userId, gameId, seats, payment, statuspayment, status) VALUES ($1, 1, 6, 'card', 0, 1)", userId)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec("INSERT INTO GamesWithUsers (userId, gameId, seats, payment, statuspayment, status) VALUES ($1, 2, 7, 'cash', 0, 1)", userId)
	if err != nil {
		panic(err)
	}
}

func CreateScheduleForUser() {
	_, err := types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (0, 'volleyball', 20250212, 1200, 44, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'POUNDS', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (1, 'volleyball', 20251212, 1100, 44, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'USD', 1)`)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
	VALUES (2, 'volleyball', 20250612, 1000, 44, 36.893445, 30.709591, 'Кладбище в Анталии', 100, 'RUB', 1)`)
	if err != nil {
		panic(err)
	}
}

func DeleteUserSchedule(userId int) {
	_, err := types.Db.Exec("DELETE FROM GamesWithUsers WHERE userId = $1", userId)
	if err != nil {
		panic(err)
	}
}

func checkDelGame(gameId, userId int) bool {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM GamesWithUsers WHERE gameId = $1 AND userId = $2 AND status = -1", gameId, userId)
	if err != nil {
		panic(err)
	}
	cc := 0
	rows.Next()
	err = rows.Scan(&cc)
	if err != nil {
		panic(err)
	}
	return cc > 0
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

func DeleteGame(gameId int) {
	_, err := types.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
}

func DeleteGameWithUser(gameId, userId int) {
	_, err := types.Db.Exec("DELETE FROM GamesWithUsers WHERE gameId = $2 AND userId = $1", gameId, userId)
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
