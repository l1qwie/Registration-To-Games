package media

import (
	"database/sql"
	"registrationtogames/bot/forall"
	"registrationtogames/fmtogram/types"
)

func FindMediaGame(gameId int) (detected bool) {
	var (
		rows    *sql.Rows
		counter int
		err     error
	)
	rows, err = types.Db.Query("SELECT COUNT(*) FROM Schedule WHERE gameId = $1 AND status = -1", gameId)
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

func findAnyMGames() (detected bool) {
	var (
		rows    *sql.Rows
		counter int
		err     error
	)
	rows, err = types.Db.Query("SELECT COUNT(*) FROM Schedule WHERE status = -1")
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

func findEveryGames() (unload, upload int) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = types.Db.Query(`SELECT
								COUNT(DISTINCT s.gameId) AS passedGamesCount,
								COUNT(DISTINCT m.gameId) AS uniqueMediaGamesCount
							FROM
								schedule s
							LEFT JOIN
								mediarepository m ON s.gameId = m.gameId
							WHERE
								s.status = -1
								AND m.gameId IS NULL`)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&upload, unload)
	return unload, upload
}

func queryUnload(limit, lm int) (rows *sql.Rows) {
	rows, err := types.Db.Query(`SELECT Schedule.gameId, sport, date, time 
							FROM Schedule JOIN MediaRepository ON 
							Schedule.gameId = MediaRepository.gameId AND 
							Schedule.status = -1 AND 
							MediaRepository.status = 1
							ORDER BY Schedule DESC LIMIT $1 OFFSET $2`, limit, lp)
	if err != nil {
		panic(err)
	}
	return rows
}

func queryUpload(limit, lm int) (rows *sql.Rows) {
	rows, err := types.Db.Query(`SELECT gameId, sport, date, time 
							FROM Schedule JOIN MediaRepository ON 
							(MediaRepository.gameId = Schedule.gameId AND MediaRepository.status = 1 AND Schedule.status = -1) OR
							(MediaRepository.gameId = NULL AND Schedule.status = -1)
							ORDER BY Schedule DESC LIMIT $1 OFFSET $2`, limit, lp)
	if err != nil {
		panic(err)
	}
	return rows
}

func selectGamesInf(limit, lp int, vector bool) (schedule []*forall.Game) {
	var (
		rows          *sql.Rows
		err           error
		date, time, i int
	)
	if vector {
		rows = queryUnload(limit, lp)
	} else {
		rows = queryUpload(limit, lp)
	}
	for rows.Next() {
		err = rows.Scan(&schedule[i].Id, &schedule[i].Sport, &date, &time)
		if err != nil {
			panic(err)
		}
		schedule[i].Date = forall.FromIntToStrDate(date)
		schedule[i].Time = forall.FromIntToStrTime(time)
		i++
	}

	return schedule
}
