package models

import "time"

//Build ...
type Build struct {
	Metadata
	Environment
	Status        Status
	Actions       []*Action
	DateTimeStart time.Time
	DateTimeEnd   time.Time
}
