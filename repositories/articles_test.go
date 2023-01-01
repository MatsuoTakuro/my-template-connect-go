package repositories_test

import (
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
	"github.com/MatsuoTakuro/my-template-connect-go/repositories/testdata"
	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}
