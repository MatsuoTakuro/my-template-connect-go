package services

import "github.com/MatsuoTakuro/my-template-connect-go/models"

// /article関連を引き受けるサービス
type ArticleServicer interface {
	GetArticleListService(page int) ([]models.Article, error)
}
type StoreServicer interface {
	GetStoreListService(searchQuery string, companyCD int) ([]models.Store, error)
}
