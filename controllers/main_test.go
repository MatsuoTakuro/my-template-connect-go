package controllers_test

import (
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/controllers"
	"github.com/MatsuoTakuro/my-template-connect-go/services"
	_ "github.com/lib/pq"
)

var sCon *controllers.StoreController

func TestMain(m *testing.M) {
	ser := services.NewAppServiceMock()
	sCon = controllers.NewStoreController(ser)

	m.Run()
}
