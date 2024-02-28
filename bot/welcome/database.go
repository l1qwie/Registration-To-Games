package welcome

import (
	"registrationtogames/fmtogram/types"

	_ "github.com/lib/pq"
)

func CompletionOfRegistration(userId int) (err error) {
	_, err = types.Db.Exec("UPDATE Users SET setup_reg = 'completed' WHERE userId = $1", userId)
	return err
}
