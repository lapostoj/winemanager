package wine

import (
	"errors"
	"strconv"
)

// Size defines the possible sizes for a wine.
//go:generate stringer -type=Size
type Size int

// Possible values for Size.
const (
	UNKNOWN        Size = 0
	HALFBOTTLE     Size = 37
	HALFLITER      Size = 50
	JURABOTTLE     Size = 62
	BOTTLE         Size = 75
	MAGNUM         Size = 150
	JEROBOAM       Size = 300
	METHUSELAH     Size = 600
	SALMANAZAR     Size = 1200
	BALTHAZAR      Size = 1800
	NABUCHADNEZZAR Size = 2400
)

// IntToSize returns the Size corresponding to the int provided.
func IntToSize(i int) Size {
	switch i {
	case 0:
		return UNKNOWN
	case 37:
		return HALFBOTTLE
	case 50:
		return HALFLITER
	case 62:
		return JURABOTTLE
	case 75:
		return BOTTLE
	case 150:
		return MAGNUM
	case 300:
		return JEROBOAM
	case 600:
		return METHUSELAH
	case 1200:
		return SALMANAZAR
	case 1800:
		return BALTHAZAR
	case 2400:
		return NABUCHADNEZZAR
	}
	panic(errors.New("Unknonw size: " + strconv.Itoa(i)))
}
