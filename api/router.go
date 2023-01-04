package api

import (
	"database/sql"
	"net/http"

	"github.com/MatsuoTakuro/my-template-connect-go/api/middlewares"
	"github.com/MatsuoTakuro/my-template-connect-go/controllers"
	"github.com/MatsuoTakuro/my-template-connect-go/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewAppService(db)
	sCon := controllers.NewStoreController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/hello", sCon.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/store", sCon.StoreListHandler).Methods(http.MethodGet)

	r.Use(middlewares.JsonResponseMiddleware)
	r.Use(middlewares.LoggingMiddleware)

	return r
}
