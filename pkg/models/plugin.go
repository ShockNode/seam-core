package models

// Plugin ...
type Plugin interface {
	Environment
	GetActions() ([]*Action, error)
}
