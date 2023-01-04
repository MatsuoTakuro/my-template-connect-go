package controllers_test

import (
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/controllers"
	"github.com/MatsuoTakuro/my-template-connect-go/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var sCon *controllers.StoreController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	sCon = controllers.NewStoreController(ser)

	m.Run()
}
