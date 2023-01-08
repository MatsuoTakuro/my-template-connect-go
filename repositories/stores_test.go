package repositories_test

import (
	"context"
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/models"
	"github.com/MatsuoTakuro/my-template-connect-go/models/testdata"
	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
	"github.com/MatsuoTakuro/my-template-connect-go/testutils"
	_ "github.com/lib/pq"
)

func TestSelectStoreList(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	tx, err := testutils.OpenDBForTest(t).BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = tx.Rollback() })

	err = testutils.PrepareStores(ctx, t, tx)
	if err != nil {
		t.Fatal(err)
	}

	want := []models.Store{
		testdata.Stores[0],
		testdata.Stores[6],
	}
	wantedNum := len(want)
	got, err := repositories.SelectStoreList(ctx, tx, "田", 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != wantedNum {
		t.Errorf("want %d but got %d stores\n", wantedNum, num)
	}
}
