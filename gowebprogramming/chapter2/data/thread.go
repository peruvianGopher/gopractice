package data

import "time"

type Thread struct {
	Id int
	UUid string
	Topic string
	UserId int
	CreatedAt time.Time
}
