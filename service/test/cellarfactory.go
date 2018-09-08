package test

import (
	"time"

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
