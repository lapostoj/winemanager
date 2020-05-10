package wine

import (
	"errors"
	"log"
)

// Type defines the possible types for a wine.
type Type string

// Possible values for Type.
// Need to be translated but it is quite specific vocabulary.
const (
	LIQUOREUX      Type = "LIQUOREUX"
	SEC                 = "SEC"
	VINDOUXNATUREL      = "VINDOUXNATUREL"
	EFFERVESCENT        = "EFFERVESCENT"
	DOUX                = "DOUX"
	MOELLEUX            = "MOELLEUX"
	TRANQUILLE          = "TRANQUILLE"
	AUTRE               = "AUTRE"
)

func (t Type) String() string {
	switch t {
	case LIQUOREUX:
		return "LIQUOREUX"
	case SEC:
		return "SEC"
	case VINDOUXNATUREL:
		return "VINDOUXNATUREL"
	case EFFERVESCENT:
		return "EFFERVESCENT"
	case DOUX:
		return "DOUX"
	case MOELLEUX:
		return "MOELLEUX"
	case TRANQUILLE:
		return "TRANQUILLE"
	case AUTRE:
		return "AUTRE"
	}

	panic(errors.New("Unknown type"))
}

// StringToType returns the Type corresponsing to the string provided.
// The strings considered here (French and first letter cap) were for a specfic use case.
func StringToType(s string) Type {
	switch s {
	case "Liquoreux":
		return LIQUOREUX
	case "Sec":
		return SEC
	case "Vin Doux Naturel":
		return VINDOUXNATUREL
	case "Effervescent":
		return EFFERVESCENT
	case "Doux":
		return DOUX
	case "Moelleux":
		return MOELLEUX
	case "Tranquille":
		return TRANQUILLE
	}

	log.Printf("Unknown wine type string %s. Defaulting to AUTRE", s)
	return AUTRE
}
