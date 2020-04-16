package models

// Metadata ...
type Metadata struct {
	ID          string
	Name        string
	Description string
}

// MetadataAccessor ...
type MetadataAccessor interface {
	GetMetaData() (*Metadata, error)
}
