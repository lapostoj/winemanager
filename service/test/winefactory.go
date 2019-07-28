package test

import (
	"cloud.google.com/go/datastore"
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

// AWineWithID returns a wine instance with ID to use in tests
func AWineWithID() wine.Wine {
	return wine.Wine{
		Key:         datastore.IDKey("Wine", 0, nil),
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
