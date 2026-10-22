package types

import "time"

type URLs struct {
	Id          string
	OriginalUrl string
	Alias       string
	CreatedAt   time.Time
	ExpireAt    time.Time
}
