package services

import (
	"context"

	"github.com/MatsuoTakuro/my-template-connect-go/models"
	"github.com/MatsuoTakuro/my-template-connect-go/models/testdata"
)

type AppserviceMock struct{}

func NewAppServiceMock() *AppserviceMock {
	return &AppserviceMock{}
}

func (s *AppserviceMock) GetStoreListService(ctx context.Context, searchQuery string, companyCD int) ([]models.Store, error) {
	return testdata.Stores, nil
}
