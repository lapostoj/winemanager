package wine

// Size defines the possible sizes for a wine.
//go:generate stringer -type=Size
type Size int

// Possible values for Size.
const (
	BOTTLE         Size = 1
	MAGNUM         Size = 2
	JEROBOAM       Size = 4
	METHUSELAH     Size = 8
	SALMANAZAR     Size = 12
	BALTHAZAR      Size = 16
	NABUCHADNEZZAR Size = 20
)
