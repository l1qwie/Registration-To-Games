package grpc

import "Media/apptype"

func createGame(gameId int) {
	_, err := apptype.Db.Exec("INSERT INTO Schedule (gameId, sport, date, time, status) VALUES ($1, $2, $3, $4, $5)",
		gameId, "football", 20201212, 1200, 1)
	if err != nil {
		panic(err)
	}
}

func deleteGame(gameId int) {
	_, err := apptype.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", gameId)
	if err != nil {
		panic(err)
	}
}

func checkAddedGame() bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND status = $5",
		55, "football", 20201212, 1200, 1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangeGame() bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM Schedule WHERE gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND status = $5",
		55, "football", 20201212, 1200, -1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
