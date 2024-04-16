package handler

import (
	"Media/apptype"
	"Media/fmtogram/types"
	"database/sql"
)

// Finds any game to upload/unload media
func findAnyMGames(f func(error)) bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE status = -1").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

// Finds every game that the user could upload or unload.
// Returns two int values
func findEveryGames(f func(error)) (int, int) {
	var un, up int
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
		panic(err)
	}
	return un, up
}

// Make a query for unload games
func queryUnload(limit, lp int, f func(error)) (rows *sql.Rows) {
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
func queryUpload(limit, lp int, f func(error)) (rows *sql.Rows) {
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
func selectGamesInf(limit, lp int, vector bool, f func(error)) []*apptype.Game {
	var (
		rows          *sql.Rows
		err           error
		date, time, i int
	)
	schedule := make([]*apptype.Game, limit)
	if vector {
		rows = queryUnload(limit, lp, f)
	} else {
		rows = queryUpload(limit, lp, f)
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
func findMediaGame(gameId int, f func(error)) bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE gameId = $1 AND status = -1", gameId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

// Selects the quantity of the game
func selectQuantity(gameId int, f func(error)) int {
	var res int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1", gameId).Scan(&res)
	if err != nil {
		f(err)
	}

	return res
}

// Select max 20 mediafiles from the game
func selectArrOrMedia(gameId int, f func(error)) []types.Media {
	var (
		rows   *sql.Rows
		err    error
		id, ty string
		i      int
	)
	fIds := make([]types.Media, 20)
	rows, err = apptype.Db.Query("SELECT fileId, type FROM MediaRepository WHERE gameId = $1", gameId)
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
func selectOneMedia(gameId int, f func(error)) (string, string) {
	var fId, ty string
	err := apptype.Db.QueryRow("SELECT fileId, type FROM MediaRepository WHERE gameId = $1", gameId).Scan(&fId, &ty)
	if err != nil {
		f(err)
	}
	return fId, ty
}

// Finds how mutch space we have for the game
func howMuchSpace(gameId int, f func(error)) (int, bool) {
	var res int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId  = $1", gameId).Scan(&res)
	if err != nil {
		f(err)
	}
	return 20 - res, (20 - res) > 0
}

// Saves the mediafile to the game
func insertOneNewMedia(fileId, ty string, gameId, userId int, f func(error)) {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 1", gameId).Scan(&count)
	if err != nil {
		f(err)
	}
	_, err = apptype.Db.Exec("INSERT INTO MediaRepository (gameId, fileId, type, counter, userId, status) VALUES ($1, $2, $3, $4, $5, 1)", gameId, fileId, ty, count+1, userId)
	if err != nil {
		f(err)
	}
}

// Saves a lot of media (array media) to the game
func insertAfewNewMedia(media []types.Media, gameId, userId int, f func(error)) {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 1", gameId).Scan(&count)
	if err != nil {
		f(err)
	}
	count += len(media)
	for i := range media {
		_, err = apptype.Db.Exec("INSERT INTO MediaRepository (gameId, fileId, type, counter, userId, status) VALUES ($1, $2, $3, $4, $5, 1)", gameId, media[i].Media, media[i].Type, count, userId)
		if err != nil {
			f(err)
		}
	}
}
