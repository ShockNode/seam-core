package models

import "time"

//Job ...
type Job struct {
	Metadata
	Environment
	Status        Status
	Actions       []*Action
	DateTimeStart time.Time
	DateTimeEnd   time.Time
}
