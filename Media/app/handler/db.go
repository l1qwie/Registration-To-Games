package handler

import (
	"Media/api/producer"
	"Media/apptype"
	"database/sql"
	"fmt"

	"github.com/l1qwie/Fmtogram/types"
)

type Conn struct {
	Db *sql.DB
}

// Finds any game to upload/unload media
func (c *Conn) findAnyMGames(f func(error)) bool {
	var count int
	producer.InterLogs("Start function Media.findAnyMGames()",
		fmt.Sprintf("f (func(error)): %T", f))
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE status = -1").Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

// Finds every game that the user could upload or unload.
// Returns two int values
func (c *Conn) findEveryGames(f func(error)) (int, int) {
	var un, up int
	producer.InterLogs("Start function Media.findEveryGames()",
		fmt.Sprintf("f (func(error)): %T", f))
	err := apptype.Db.QueryRow(`SELECT
								COUNT(DISTINCT s.gameId) AS passedGamesCount,
								COUNT(DISTINCT m.gameId) AS uniqueMediaGamesCount
							FROM
								schedule s
							LEFT JOIN
								mediarepository m ON s.gameId = m.gameId
							WHERE
								s.status = -1`).Scan(&up, &un)
	if err != nil {
		f(err)
	}
	return un, up
}

// Make a query for unload games
func (c *Conn) queryUnload(limit, lp int, f func(error)) (rows *sql.Rows) {
	producer.InterLogs("Start function Media.queryUnload()",
		fmt.Sprintf("limit (int): %d, lp (int): %d, f (func(error)): %T", limit, lp, f))
	rows, err := apptype.Db.Query(`
		SELECT DISTINCT(Schedule.gameId), sport, date, time 
		FROM Schedule 
		JOIN MediaRepository ON Schedule.gameId = MediaRepository.gameId 
		WHERE Schedule.status = -1 AND MediaRepository.status = 1
		ORDER BY Schedule.gameId DESC LIMIT $1 OFFSET $2`, limit, lp)
	if err != nil {
		f(err)
	}
	return rows
}

// Make a query for upload games
func (c *Conn) queryUpload(limit, lp int, f func(error)) (rows *sql.Rows) {
	producer.InterLogs("Start function Media.queryUpload()",
		fmt.Sprintf("limit (int): %d, lp (int): %d, f (func(error)): %T", limit, lp, f))
	rows, err := apptype.Db.Query(`SELECT DISTINCT(gameId), sport, date, time 
								FROM Schedule 
								WHERE status = -1
								ORDER BY gameId DESC LIMIT $1 OFFSET $2`, limit, lp)
	if err != nil {
		f(err)
	}
	return rows
}

// Selects all games inf depens on upload or unload option
func (c *Conn) selectGamesInf(limit, lp int, vector bool, f func(error)) []*apptype.Game {
	var (
		rows          *sql.Rows
		err           error
		date, time, i int
	)
	producer.InterLogs("Start function Media.selectGamesInf()",
		fmt.Sprintf("limit (int): %d, lp (int): %d, vector (bool): %v, f (func(error)): %T", limit, lp, vector, f))
	schedule := make([]*apptype.Game, limit)
	if vector {
		rows = c.queryUnload(limit, lp, f)
	} else {
		rows = c.queryUpload(limit, lp, f)
	}
	for rows.Next() {
		schedule[i] = &apptype.Game{}
		err = rows.Scan(&schedule[i].Id, &schedule[i].Sport, &date, &time)
		if err != nil {
			f(err)
		}
		schedule[i].Date = fromIntToStrDate(date)
		schedule[i].Time = fromIntToStrTime(time)
		i++
	}
	defer rows.Close()
	return schedule
}

// Finds the media game
func (c *Conn) findMediaGame(gameId int, f func(error)) bool {
	var count int
	producer.InterLogs("Start function Media.findMediaGame()",
		fmt.Sprintf("gameId (int): %d, f (func(error)): %T", gameId, f))
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE gameId = $1 AND status = -1", gameId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

// Selects the quantity of the game
func (c *Conn) selectQuantity(gameId int, f func(error)) int {
	var res int
	producer.InterLogs("Start function Media.selectQuantity()",
		fmt.Sprintf("gameId (int): %d, f (func(error)): %T", gameId, f))
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1", gameId).Scan(&res)
	if err != nil {
		f(err)
	}

	return res
}

// Select max 20 mediafiles from the game
func (c *Conn) selectArrOrMedia(gameId, quan int, f func(error)) []types.Media {
	var (
		id, ty string
		i      int
	)
	producer.InterLogs("Start function Media.selectArrOrMedia()",
		fmt.Sprintf("gameId (int): %d, quan (int): %d, f (func(error)): %T", gameId, quan, f))
	fIds := make([]types.Media, quan)
	rows, err := apptype.Db.Query("SELECT fileId, type FROM MediaRepository WHERE gameId = $1", gameId)
	if err != nil {
		f(err)
	}
	for rows.Next() {
		err = rows.Scan(&id, &ty)
		if err != nil {
			f(err)
		}
		fIds[i].Type = ty
		fIds[i].Media = id
		i++
	}
	defer rows.Close()
	return fIds
}

// Selects only one mediafile from the game
func (c *Conn) selectOneMedia(gameId int, f func(error)) (string, string) {
	var fId, ty string
	producer.InterLogs("Start function Media.selectOneMedia()",
		fmt.Sprintf("gameId (int): %d, f (func(error)): %T", gameId, f))
	err := apptype.Db.QueryRow("SELECT fileId, type FROM MediaRepository WHERE gameId = $1", gameId).Scan(&fId, &ty)
	if err != nil {
		f(err)
	}
	return fId, ty
}

// Finds how mutch space we have for the game
func (c *Conn) howMuchSpace(gameId int, f func(error)) (int, bool) {
	var res int
	producer.InterLogs("Start function Media.howMuchSpace()",
		fmt.Sprintf("gameId (int): %d, f (func(error)): %T", gameId, f))
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId  = $1", gameId).Scan(&res)
	if err != nil {
		f(err)
	}
	return 20 - res, (20 - res) > 0
}

func (c *Conn) endOfTransaction(err error, f func(error)) {
	var res string
	if err != nil {
		res = "ROLLBACK"
	} else {
		res = "COMMIT"
	}
	_, err = apptype.Db.Exec(res)
	if err != nil {
		f(err)
	}
}

// Saves the mediafile to the game
func (c *Conn) insertOneNewMedia(fileId, ty string, gameId, userId int, f func(error)) bool {
	var count int
	producer.InterLogs("Start function Media.insertOneNewMedia()",
		fmt.Sprintf("fileId (string): %s, ty (string): %s, gameId (int): %d, userId (int): %d, f (func(error)): %T", fileId, ty, gameId, userId, f))
	_, err := apptype.Db.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	if err == nil {
		err = apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 1", gameId).Scan(&count)
	}
	if err == nil {
		_, err = apptype.Db.Exec(`
			INSERT INTO MediaRepository 
			(id, gameId, fileId, type, counter, userId, status)
			VALUES
			(nextval('mediarepository_id_seq'), $1, $2, $3, $4, $5, 1)`, gameId, fileId, ty, count+1, userId)
	}
	if err == nil {
		_, err = apptype.Db.Exec("UPDATE MediaRepository SET counter = $1 WHERE gameId = $2", count+1, gameId)
	}
	defer c.endOfTransaction(err, f)
	return err == nil
}

// Saves a lot of media (array media) to the game
func (c *Conn) insertAfewNewMedia(media []types.Media, gameId, userId int, f func(error)) bool {
	var count int
	producer.InterLogs("Start function Media.insertAfewNewMedia()",
		fmt.Sprintf("media ([]types.Media): %v, gameId (int): %d, userId (int): %d, f (func(error)): %T", media, gameId, userId, f))
	_, err := apptype.Db.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	if err == nil {
		err = apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 1", gameId).Scan(&count)
	}
	if err == nil {
		count += len(media)
		for i := 0; i < len(media) && err == nil; i++ {
			_, err = apptype.Db.Exec(`
				INSERT INTO MediaRepository 
				(id, gameId, fileId, type, counter, userId, status)
				VALUES 
				(nextval('mediarepository_id_seq'), $1, $2, $3, $4, $5, 1)`, gameId, media[i].Media, media[i].Type, count, userId)
		}
	}
	if err == nil {
		_, err = apptype.Db.Exec("UPDATE MediaRepository SET counter = $1 WHERE gameId = $2", count, gameId)
		if err != nil {
			f(err)
		}
	}
	defer c.endOfTransaction(err, f)
	return err == nil
}

func (c *Conn) UpdateTheSchedule(date, time, status int, g *apptype.Game, act string) error {
	var request string
	producer.InterLogs("Start function Media.UpdateTheSchedule()",
		fmt.Sprintf("date (int): %d, time (int): %d, status (int): %d, g (*apptype.Game): %v, act (string): %s", date, time, status, g, act))
	if act == "new" {
		request = "INSERT INTO Schedule (gameId, sport, date, time, status) VALUES ($1, $2, $3, $4, $5)"
	} else if act == "change" {
		request = "UPDATE Schedule SET gameId = $1, sport = $2, date = $3, time = $4, status = $5 WHERE gameId = $1"
	}
	_, err := apptype.Db.Exec(request, g.Id, g.Sport, date, time, status)
	return err
}
