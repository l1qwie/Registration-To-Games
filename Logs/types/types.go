package types

import "time"

type Called struct {
	Timestamp time.Time `json:"timestamp"`
	LogId     int       `json:"id"`
	Message   string    `json:"message"`
	Data      string    `json:"data"`
}
