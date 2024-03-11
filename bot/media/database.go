package media

import (
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/types"
	"database/sql"
)

func FindMediaGame(gameId int) bool {
	var (
		rows    *sql.Rows
		counter int
		err     error
		dec     bool
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
		dec = true
	}
	return dec
}

func findAnyMGames() bool {
	var (
		rows    *sql.Rows
		counter int
		err     error
		d       bool
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
		d = true
	}
	return d
}

func howMuchSpace(gameId int) (int, bool) {
	var (
		rows *sql.Rows
		err  error
		res  int
		ok   bool
	)
	rows, err = types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId  = $1", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&res)
	if err != nil {
		panic(err)
	}
	if (20 - res) > 0 {
		ok = true
	}
	return 20 - res, ok
}

func selectQuantity(gameId int) int {
	var (
		rows *sql.Rows
		err  error
		res  int
	)
	rows, err = types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&res)
	if err != nil {
		panic(err)
	}
	return res
}

func selectArrOrMedia(gameId int) map[string]string {
	var (
		rows   *sql.Rows
		err    error
		id, ty string
	)
	fIds := make(map[string]string, 20)
	rows, err = types.Db.Query("SELECT fileId, type FROM MediaRepository WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&id, &ty)
		if err != nil {
			panic(err)
		}
		fIds[id] = ty
	}
	return fIds
}

func insertOneNewMedia(fileId, ty string, gameId, userId int) {
	var (
		//rows *sql.Rows
		//err  error
		cc int
	)
	rows, err := types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 0", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&cc)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec("INSERT INTO MediaRepository (gameId, fileId, type, counter, userId) VALUES ($1, $2, $3, $4, $5)", gameId, fileId, ty, cc+1, userId)
	if err != nil {
		panic(err)
	}
}

func selectOneMedia(gameId int) (string, string) {
	var (
		rows    *sql.Rows
		err     error
		fId, ty string
	)
	rows, err = types.Db.Query("SELECT fileId, type FROM MediaRepository WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&fId, &ty)
	if err != nil {
		panic(err)
	}
	return fId, ty
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
								s.status = -1`)
	if err != nil {
		panic(err)
	}
	rows.Next()
	err = rows.Scan(&upload, &unload)
	if err != nil {
		panic(err)
	}
	return unload, upload
}

func queryUnload(limit, lp int) (rows *sql.Rows) {
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

func queryUpload(limit, lp int) (rows *sql.Rows) {
	rows, err := types.Db.Query(`SELECT gameId, sport, date, time 
								FROM Schedule WHERE status = -1
								ORDER BY Schedule DESC LIMIT $1 OFFSET $2`, limit, lp)
	if err != nil {
		panic(err)
	}
	return rows
}

func selectGamesInf(limit, lp int, vector bool) []*forall.Game {
	var (
		rows          *sql.Rows
		err           error
		date, time, i int
	)
	schedule := make([]*forall.Game, limit)
	if vector {
		rows = queryUnload(limit, lp)
	} else {
		rows = queryUpload(limit, lp)
	}
	for rows.Next() {
		schedule[i] = &forall.Game{}
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
