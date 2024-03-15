package media

import (
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/types"
	"database/sql"
)

func FindMediaGame(gameId int) bool {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM Schedule WHERE gameId = $1 AND status = -1", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	counter := 0
	err = rows.Scan(&counter)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return counter > 0
}

func findAnyMGames() bool {
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
	defer rows.Close()
	return counter > 0
}

func howMuchSpace(gameId int) (int, bool) {
	var (
		rows *sql.Rows
		err  error
		res  int
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
	defer rows.Close()
	return 20 - res, (20 - res) > 0
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
	defer rows.Close()
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
	defer rows.Close()
	return fIds
}

func insertAfewNewMeda(media map[string]string, gameId, userId int) {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 0", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	cc := 0
	err = rows.Scan(&cc)
	if err != nil {
		panic(err)
	}
	cc += len(media)
	for key := range media {
		_, err = types.Db.Exec("INSERT INTO MediaRepository (gameId, fileId, type, counter, userId) VALUES ($1, $2, $3, $4, $5)", gameId, key, media[key], cc, userId)
		if err != nil {
			panic(err)
		}
	}
	defer rows.Close()
}

func insertOneNewMedia(fileId, ty string, gameId, userId int) {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 0", gameId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	cc := 0
	err = rows.Scan(&cc)
	if err != nil {
		panic(err)
	}
	_, err = types.Db.Exec("INSERT INTO MediaRepository (gameId, fileId, type, counter, userId) VALUES ($1, $2, $3, $4, $5)", gameId, fileId, ty, cc+1, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
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
	defer rows.Close()
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
	defer rows.Close()
	return unload, upload
}

func queryUnload(limit, lp int) (rows *sql.Rows) {
	rows, err := types.Db.Query(`
		SELECT DISTINCT Schedule.gameId, sport, date, time 
		FROM Schedule 
		JOIN MediaRepository ON Schedule.gameId = MediaRepository.gameId 
		WHERE Schedule.status = -1 AND MediaRepository.status = 1
		ORDER BY Schedule.gameId DESC LIMIT $1 OFFSET $2`, limit, lp)
	if err != nil {
		panic(err)
	}
	return rows
}

func queryUpload(limit, lp int) (rows *sql.Rows) {
	rows, err := types.Db.Query(`SELECT DISTINCT(gameId), sport, date, time 
								FROM Schedule 
								WHERE status = -1
								ORDER BY gameId DESC LIMIT $1 OFFSET $2`, limit, lp)
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
	defer rows.Close()
	return schedule
}
