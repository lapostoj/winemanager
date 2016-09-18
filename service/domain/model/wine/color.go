package wine

// Color defines the possible colors for a wine.
//go:generate stringer -type=Color
type Color int

// Possible values for Color.
const (
	RED Color = iota
	WHITE
	ROSE
)
