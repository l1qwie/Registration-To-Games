package functional

import "Media/apptype"

// Checks of there is the media
func checkUploadedMedia() bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND userId = $2 and fileId = $3 AND type = $4 AND counter = $5",
		10, 499, "!@#UIO!@#IOJJKLASEDKLKL#IO!JKLASJKL13419", "photo", 1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

// Checks the first media in the table "MediaRepository"
func thefirst() bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND userId = $2 and fileId = $3 AND type = $4 AND counter = $5",
		1, 499, "!@#IOJSJE!@#**()!@#$*()SIOPE!@()#", "photo", 3).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

// Checks the second media in the table "MediaRepository"
func thesecond() bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND userId = $2 and fileId = $3 AND type = $4 AND counter = $5",
		1, 499, "!@#IOJSIOJE!@#**()!@HFF#$*()SIOPE$#@$!@()#", "photo", 3).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

// Checks the third media in the table "MediaRepository"
func thethird() bool {
	var count int
	err := apptype.Db.QueryRow("SELECT COUNT(*) FROM MediaRepository WHERE gameId = $1 AND userId = $2 and fileId = $3 AND type = $4 AND counter = $5",
		1, 499, "!@#IOJSIOJE!$@!$!@#@#**()!@#$*()SIOPE!@()#", "photo", 3).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
