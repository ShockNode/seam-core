package models

// Action ...
type Action interface {
	Environment
	GetMetadata() Metadata
	Act(build *Build) error
}
