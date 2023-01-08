package services

import (
	"context"

	"github.com/MatsuoTakuro/my-template-connect-go/apperrors"
	"github.com/MatsuoTakuro/my-template-connect-go/models"
	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
)

func (s *AppService) GetStoreListService(ctx context.Context, searchQuery string, companyCD int) ([]models.Store, error) {
	storeList, err := repositories.SelectStoreList(ctx, s.db, searchQuery, companyCD)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(storeList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return storeList, nil
}
