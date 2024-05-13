package database

import (
	"Logs/types"

	_ "github.com/lib/pq"
)

func Add(msg *types.ClientAct) {
	_, err := types.Db.Exec("INSERT INTO Actions (id, userId, act, eventTime) VALUES (nextval('action_id_seq'), $1, $2, $3)", msg.UserId, msg.Action, msg.Com.Timestamp)
	if err != nil {
		panic(err)
	}
}
