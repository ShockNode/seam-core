package models

import "time"

// Job ...
type Job interface {
	Seam
	ActionsAccessor
}

// JobRunner ...
type JobRunner interface {
	Seam
	Run(job Job) (*JobResult, error)
}

// JobResult ...
type JobResult struct {
	Metadata
	Status        Status
	DateTimeStart time.Time
	DateTimeEnd   time.Time
}
