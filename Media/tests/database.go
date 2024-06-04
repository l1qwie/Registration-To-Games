package tests

import (
	"Media/apptype"
)

// UpdateAction updates column "action" in table "Users" in the database
func UpdateAction() {
	_, err := apptype.Db.Exec("UPDATE Users SET action = $1 WHERE userId = $2", "reg to games", 899)
	if err != nil {
		panic(err)
	}
}

// UpdateLanguage updates column "language" in the table "Users" in the database
func UpdateLanguage() {
	_, err := apptype.Db.Exec("UPDATE Users SET language = $1 WHERE userId = $2", "ru", 899)
	if err != nil {
		panic(err)
	}
}

// UpdateLevel updates column "level" in the table "Users" in the database
func UpdateLevel(level int) {
	_, err := apptype.Db.Exec("UPDATE Users SET level = $1 WHERE userId = $2", level, 899)
	if err != nil {
		panic(err)
	}

}

// DeleteUser deletes a user from the table "Users"
func DeleteUser() {
	_, err := apptype.Db.Exec("DELETE FROM Users WHERE userId = $1", 899)
	if err != nil {
		panic(err)
	}
}

// CreateUser creates a user in the table "Users"
func CreateUser() {
	_, err := apptype.Db.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", 899, "reg to games", "ru", 0)
	if err != nil {
		panic(err)
	}
}

// DeleteGame deletes a Game from the table "Schedule"
func DeleteGame() {
	_, err := apptype.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", 10)
	if err != nil {
		panic(err)
	}
}

// DeleteMedia deletes media from the table "MediaRepository"
func DeleteMedia() {
	_, err := apptype.Db.Exec("DELETE FROM MediaRepository WHERE gameId = $1", 10)
	if err != nil {
		panic(err)
	}
}

// CreateNotFullMediaGame creates a game and a mediaGame with one fileId
func CreateNotFullMediaGame() {
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, status) 
	VALUES ($1, 'football', 20240212, 1200, -1)`, 10)
	if err != nil {
		panic(err)
	}
	_, err = apptype.Db.Exec(`INSERT INTO MediaRepository (gameId, userId, fileId, type, counter, status)
			VALUES ($1, 499, '!@#IOJSIOJE!@#**()!@#$*()SIOPE!@()#', 'photo', 0, 1)`, 10)
	if err != nil {
		panic(err)
	}
}

// CreateEmptyMediaGame creates empty media game (just a game in the table "Schedule" with status -1)
func CreateEmptyMediaGame() {
	_, err := apptype.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, status) 
	VALUES ($1, 'volleyball', 20240212, 1200, -1)`, 10)
	if err != nil {
		panic(err)
	}
}

// DeleteSchedule deletes a schedule from the table "Schedule"
func DeleteSchedule() {
	_, err := apptype.Db.Exec("DELETE FROM Schedule WHERE gameId = 0 OR gameId = 1 OR gameId = 2 OR gameId = 3 OR gameId = 4")
	if err != nil {
		panic(err)
	}
}

// DeleteMediaSchedule deletes a media inf from the table "MediaRepository"
func DeleteMediaSchedule() {
	_, err := apptype.Db.Exec("DELETE FROM MediaRepository WHERE gameId = 1 OR gameId = 2 OR gameId = 3 OR gameId = 4")
	if err != nil {
		panic(err)
	}
}

// CreateMediaShedule creates a media schedule on the table "MediaRepository"
func CreateMediaShedule() {
	var request string
	for i := 1; i < 5; i++ {
		if i == 1 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, status) 
				VALUES (1, 'volleyball', 20260212, 1100, -1)`
		} else if i == 2 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, status) 
				VALUES (2, 'football', 20250412, 1800, -1)`
		} else if i == 3 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, status) 
				VALUES (3, 'volleyball', 20250202, 0800, -1)`
		} else if i == 4 {
			request = `INSERT INTO Schedule (gameId, sport, date, time, status) 
				VALUES (4, 'volleyball', 20250202, 0800, -1)`
		}
		_, err := apptype.Db.Exec(request)
		if err != nil {
			panic(err)
		}
	}
}

// FillMediaSchedule fills the media
func FillMediaSchedule() {
	var request string
	for i := 0; i < 8; i++ {
		if i == 0 {
			request = `INSERT INTO MediaRepository (id, gameId, userId, fileId, type, counter, status)
			VALUES (nextval('mediarepository_id_seq'), 1, 499, '!@#IOJSJE!@#**()!@#$*()SIOPE!@()#', 'photo', 3, 1)`
		} else if i == 1 {
			request = `INSERT INTO MediaRepository (id, gameId, userId, fileId, type, counter, status)
			VALUES (nextval('mediarepository_id_seq'), 2, 499, '!@#IOJSJE!@#**()!@#$*()SIOPE!@(#!@*(IOOI)', 'photo', 2, 1)`
		} else if i == 2 {
			request = `INSERT INTO MediaRepository (id, gameId, userId, fileId, type, counter, status)
			VALUES (nextval('mediarepository_id_seq'), 3, 499, '!@#IOJSIOJASDGE!@#**()!@#$*()SIOPE!@()#', 'photo', 2, 1)`
		} else if i == 3 {
			request = `INSERT INTO MediaRepository (id, gameId, userId, fileId, type, counter, status)
			VALUES (nextval('mediarepository_id_seq'), 4, 499, '!@#IOJSIOJESL:DFK**()!@#$*()SIOPE!@()#', 'photo', 1, 1)`
		} else if i == 4 {
			request = `INSERT INTO MediaRepository (id, gameId, userId, fileId, type, counter, status)
			VALUES (nextval('mediarepository_id_seq'), 1, 499, '!@#IOJSIOJE!@#**()!@HFF#$*()SIOPE$#@$!@()#', 'photo', 3, 1)`
		} else if i == 5 {
			request = `INSERT INTO MediaRepository (id, gameId, userId, fileId, type, counter, status)
			VALUES (nextval('mediarepository_id_seq'), 2, 499, '!@#IOJSIOJE!@#**()!()()998890-@#$*()SIOPE!@()#', 'photo', 2, 1)`
		} else if i == 6 {
			request = `INSERT INTO MediaRepository (id, gameId, userId, fileId, type, counter, status)
			VALUES (nextval('mediarepository_id_seq'), 3, 499, '!@#IOJSIOJE!@#**ASDG()!@#$*()SIOPE!@()#', 'photo', 2, 1)`
		} else if i == 7 {
			request = `INSERT INTO MediaRepository (id, gameId, userId, fileId, type, counter, status)
			VALUES (nextval('mediarepository_id_seq'), 1, 499, '!@#IOJSIOJE!$@!$!@#@#**()!@#$*()SIOPE!@()#', 'photo', 3, 1)`
		}
		_, err := apptype.Db.Exec(request)
		if err != nil {
			panic(err)
		}
	}
}
