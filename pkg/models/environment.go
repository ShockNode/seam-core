package models

// Environment ...
type Environment interface {
	GetEnvironmentVariables() map[string]interface{}
}
