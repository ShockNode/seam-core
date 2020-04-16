package models

// EnvironmentAccessor ...
type EnvironmentAccessor interface {
	GetEnvironmentVariables() map[string]interface{}
}
