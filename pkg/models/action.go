package models

// Action ...
type Action interface {
	Environment
	GetMetadata() Metadata
	Act(job *Job) error
}
