package wine

// Type defines the possible types for a wine.
//go:generate stringer -type=Type
type Type int

// Possible values for Type.
const (
	STILL Type = iota
	SPARKLING
)
