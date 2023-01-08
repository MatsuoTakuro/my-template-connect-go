package services

import (
	"context"

	"github.com/MatsuoTakuro/my-template-connect-go/models"
)

type StoreServicer interface {
	GetStoreListService(ctx context.Context, searchQuery string, companyCD int) ([]models.Store, error)
}
type GreetServicer interface {
	GreetService()
}
