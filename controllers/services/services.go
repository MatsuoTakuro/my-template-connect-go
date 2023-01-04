package services

import "github.com/MatsuoTakuro/my-template-connect-go/models"

type StoreServicer interface {
	GetStoreListService(searchQuery string, companyCD int) ([]models.Store, error)
}
