package testdata

import "github.com/MatsuoTakuro/my-template-connect-go/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) GetStoreListService(searchQuery string, companyCD int) ([]models.Store, error) {
	return storeTestData, nil
}
