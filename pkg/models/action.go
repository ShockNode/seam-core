package models

// Action ...
type Action interface {
	Seam
	Act(job Job) error
}

// ActionsAccessor ...
type ActionsAccessor interface {
	GetActions() ([]Action, error)
}
