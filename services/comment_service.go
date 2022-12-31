package services

import (
	"github.com/MatsuoTakuro/my-template-connect-go/apperrors"
	"github.com/MatsuoTakuro/my-template-connect-go/models"
	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
)

// PostCommentHandlerで使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
