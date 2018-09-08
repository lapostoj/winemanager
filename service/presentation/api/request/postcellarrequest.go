package request

import (
	"time"

	"github.com/lapostoj/winemanager/service/domain/model/cellar"
)

// PostCellarRequest defines the request that should be sent when posting a cellar
type PostCellarRequest struct {
	Name      string `json:"name"`
	AccountID int    `json:"accountID"`
}

// NewCellar transforms a PostCellarRequest into a Cellar
func (request PostCellarRequest) NewCellar() *cellar.Cellar {
	return &cellar.Cellar{
		Name:         request.Name,
		AccountID:    request.AccountID,
		CreationTime: time.Now().UTC(),
	}
}
