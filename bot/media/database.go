package media

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"database/sql"
)

func FindMediaGame(gameId int, fm *formatter.Formatter) bool {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM Schedule WHERE gameId = $1 AND status = -1", gameId)
	if err != nil {
		fm.Error(err)
	}
	rows.Next()
	counter := 0
	err = rows.Scan(&counter)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return counter > 0
}

func findAnyMGames(fm *formatter.Formatter) bool {
	var (
		rows    *sql.Rows
		counter int
		err     error
	)
	rows, err = types.Db.Query("SELECT COUNT(*) FROM Schedule WHERE status = -1")
	if err != nil {
		fm.Error(err)
	}
	rows.Next()
	err = rows.Scan(&counter)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return counter > 0
}

func howMuchSpace(gameId int, fm *formatter.Formatter) (int, bool) {
	var (
		rows *sql.Rows
		err  error
		res  int
	)
	rows, err = types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId  = $1", gameId)
	if err != nil {
		fm.Error(err)
	}
	rows.Next()
	err = rows.Scan(&res)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return 20 - res, (20 - res) > 0
}

func selectQuantity(gameId int, fm *formatter.Formatter) int {
	var (
		rows *sql.Rows
		err  error
		res  int
	)
	rows, err = types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1", gameId)
	if err != nil {
		fm.Error(err)
	}
	rows.Next()
	err = rows.Scan(&res)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return res
}

func selectArrOrMedia(gameId int, fm *formatter.Formatter) []types.Media {
	var (
		rows   *sql.Rows
		err    error
		id, ty string
		i      int
	)
	fIds := make([]types.Media, 20)
	rows, err = types.Db.Query("SELECT fileId, type FROM MediaRepository WHERE gameId = $1", gameId)
	if err != nil {
		fm.Error(err)
	}
	for rows.Next() {
		err = rows.Scan(&id, &ty)
		if err != nil {
			fm.Error(err)
		}
		fIds[i].Type = ty
		fIds[i].Media = id
		i++
	}
	defer rows.Close()
	return fIds
}

func insertAfewNewMeda(media map[string]string, gameId, userId int, fm *formatter.Formatter) {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 1", gameId)
	if err != nil {
		fm.Error(err)
	}
	rows.Next()
	cc := 0
	err = rows.Scan(&cc)
	if err != nil {
		fm.Error(err)
	}
	cc += len(media)
	for key := range media {
		_, err = types.Db.Exec("INSERT INTO MediaRepository (gameId, fileId, type, counter, userId, status) VALUES ($1, $2, $3, $4, $5, 1)", gameId, key, media[key], cc, userId)
		if err != nil {
			fm.Error(err)
		}
	}
	defer rows.Close()
}

func insertOneNewMedia(fileId, ty string, gameId, userId int, fm *formatter.Formatter) {
	rows, err := types.Db.Query("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND status = 1", gameId)
	if err != nil {
		fm.Error(err)
	}
	rows.Next()
	cc := 0
	err = rows.Scan(&cc)
	if err != nil {
		fm.Error(err)
	}
	_, err = types.Db.Exec("INSERT INTO MediaRepository (gameId, fileId, type, counter, userId, status) VALUES ($1, $2, $3, $4, $5, 1)", gameId, fileId, ty, cc+1, userId)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
}

func selectOneMedia(gameId int, fm *formatter.Formatter) (string, string) {
	var (
		rows    *sql.Rows
		err     error
		fId, ty string
	)
	rows, err = types.Db.Query("SELECT fileId, type FROM MediaRepository WHERE gameId = $1", gameId)
	if err != nil {
		fm.Error(err)
	}
	rows.Next()
	err = rows.Scan(&fId, &ty)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return fId, ty
}

func findEveryGames(fm *formatter.Formatter) (unload, upload int) {
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
		fm.Error(err)
	}
	rows.Next()
	err = rows.Scan(&upload, &unload)
	if err != nil {
		fm.Error(err)
	}
	defer rows.Close()
	return unload, upload
}

func queryUnload(limit, lp int, fm *formatter.Formatter) (rows *sql.Rows) {
	rows, err := types.Db.Query(`
		SELECT DISTINCT(Schedule.gameId), sport, date, time 
		FROM Schedule 
		JOIN MediaRepository ON Schedule.gameId = MediaRepository.gameId 
		WHERE Schedule.status = -1 AND MediaRepository.status = 1
		ORDER BY Schedule.gameId DESC LIMIT $1 OFFSET $2`, limit, lp)
	if err != nil {
		fm.Error(err)
	}
	return rows
}

func queryUpload(limit, lp int, fm *formatter.Formatter) (rows *sql.Rows) {
	rows, err := types.Db.Query(`SELECT DISTINCT(gameId), sport, date, time 
								FROM Schedule 
								WHERE status = -1
								ORDER BY gameId DESC LIMIT $1 OFFSET $2`, limit, lp)
	if err != nil {
		fm.Error(err)
	}
	return rows
}

func selectGamesInf(limit, lp int, vector bool, fm *formatter.Formatter) []*bottypes.Game {
	var (
		rows          *sql.Rows
		err           error
		date, time, i int
	)
	schedule := make([]*bottypes.Game, limit)
	if vector {
		rows = queryUnload(limit, lp, fm)
	} else {
		rows = queryUpload(limit, lp, fm)
	}
	for rows.Next() {
		schedule[i] = &bottypes.Game{}
		err = rows.Scan(&schedule[i].Id, &schedule[i].Sport, &date, &time)
		if err != nil {
			fm.Error(err)
		}
		schedule[i].Date = forall.FromIntToStrDate(date)
		schedule[i].Time = forall.FromIntToStrTime(time)
		i++
	}
	defer rows.Close()
	return schedule
}
