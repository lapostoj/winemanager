package wine

import "errors"

// Color defines the possible colors for a wine.
type Color string

// Possible values for Color.
const (
	RED   Color = "RED"
	WHITE       = "WHITE"
	ROSE        = "ROSE"
)

func (color Color) String() string {
	switch color {
	case RED:
		return "RED"
	case WHITE:
		return "WHITE"
	case ROSE:
		return "ROSE"
	}
	panic(errors.New("Unknown color"))
}

// StringToColor returns the Color corresponsing to the string provided.
// The strings considered here (French and first letter cap) were for a specfic use case.
func StringToColor(s string) Color {
	switch s {
	case "Blanc":
		return WHITE
	case "Rosé":
		return ROSE
	case "Rouge":
		return RED
	}
	panic(errors.New("Unknown color: " + s))
}
