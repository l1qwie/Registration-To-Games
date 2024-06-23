package database

import (
	"Logs/apptype"

	_ "github.com/lib/pq"
)

func Add(msg *apptype.ClientAct) {
	_, err := apptype.Db.Exec("INSERT INTO Actions (id, userId, message, act, eventTime) VALUES (nextval('action_id_seq'), $1, $2, $3, $4)", msg.UserId, msg.Message, msg.Action, msg.Timestamp)
	if err != nil {
		panic(err)
	}
}
