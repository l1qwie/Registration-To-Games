package tests

import (
	"database/sql"
)

type Connection struct {
	con *sql.DB
}

func (c *Connection) UpdateAction() {
	_, err := c.con.Exec("UPDATE Users SET action = $1 WHERE userId = $2", "registration", 456)
	if err != nil {
		panic(err)
	}
}

func (c *Connection) UpdateLanguage() {
	_, err := c.con.Exec("UPDATE Users SET language = $1 WHERE userId = $2", "ru", 456)
	if err != nil {
		panic(err)
	}
}

func (c *Connection) UpdateLevel() {
	_, err := c.con.Exec("UPDATE Users SET level = $1 WHERE userId = $2", i, 456)
	if err != nil {
		panic(err)
	}

}

func (c *Connection) DeleteUser() {
	_, err := c.con.Exec("DELETE FROM Users WHERE userId = $1", 456)
	if err != nil {
		panic(err)
	}
}

func (c *Connection) CreateUser() {
	_, err := c.con.Exec("INSERT INTO Users (userId, action, language, level) VALUES ($1, $2, $3, $4)", 456, "registration", "ru", 0)
	if err != nil {
		panic(err)
	}
}
