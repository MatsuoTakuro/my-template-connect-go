package services

import (
	"github.com/MatsuoTakuro/my-template-connect-go/apperrors"
	"github.com/MatsuoTakuro/my-template-connect-go/models"
	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
)

// ArticleListHandlerで使うことを想定したサービス
// 指定pageの記事一覧を返却
func (s *AppService) GetStoreListService(searchQuery string, companyCD int) ([]models.Store, error) {
	storeList, err := repositories.SelectStoreList(s.db, searchQuery, companyCD)
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
