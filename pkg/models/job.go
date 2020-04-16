package models

// Job ...
type Job interface {
	Seam
	ActionsAccessor
}

// JobRunner ...
type JobRunner interface {
	Seam
	Build(job Job) error
}
