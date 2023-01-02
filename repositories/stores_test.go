package repositories_test

import (
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
	"github.com/MatsuoTakuro/my-template-connect-go/repositories/testdata"
	_ "github.com/go-sql-driver/mysql"
)

func TestSelectStoreList(t *testing.T) {
	expectedNum := len(testdata.StoreTestData)
	got, err := repositories.SelectStoreList(testDB, "田", 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d stores: %v\n", expectedNum, num, got)
	}
}
