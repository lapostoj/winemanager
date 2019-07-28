package test

import (
	"time"

	"cloud.google.com/go/datastore"
	"github.com/lapostoj/winemanager/service/domain/model/cellar"
)

// ACellar returns a cellar instance to use in tests
func ACellar() cellar.Cellar {
	return cellar.Cellar{
		Name:         "Test Cellar",
		AccountID:    111,
		CreationTime: time.Now().UTC(),
	}
}

// ACellarWithID returns a cellar instance with ID to use in tests
func ACellarWithID() cellar.Cellar {
	return cellar.Cellar{
		Key:          datastore.IDKey("Cellar", 0, nil),
		Name:         "Test Cellar",
		AccountID:    111,
		CreationTime: time.Now().UTC(),
	}
}
