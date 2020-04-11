package models

// Status ...
type Status string

//
const (
	Ready      Status = "ready"
	InProgress Status = "progress"
	Done       Status = "done"
	Failed     Status = "failed"
)
