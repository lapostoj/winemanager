package test

import (
	"github.com/lapostoj/winemanager/service/domain/model/wine"
)

// AWine returns a wine instance to use in tests
func AWine() wine.Wine {
	return wine.Wine{
		Name:        "Test Wine",
		Designation: "Test Designation",
		Growth:      "Test Growth",
		Country:     "FR",
		Region:      "Bourgogne",
		Producer:    "Test Producer",
		Color:       wine.RED,
		Type:        wine.SEC,
	}
}
