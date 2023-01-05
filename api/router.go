package api

import (
	"database/sql"
	"net/http"

	"github.com/MatsuoTakuro/my-template-connect-go/api/middlewares"
	"github.com/MatsuoTakuro/my-template-connect-go/controllers"
	"github.com/MatsuoTakuro/my-template-connect-go/services"
	"github.com/gorilla/mux"
)

func NewHttpRouter(db *sql.DB) *mux.Router {
	ser := services.NewAppService(db)
	sCon := controllers.NewStoreController(ser)

	hr := mux.NewRouter()

	hr.HandleFunc("/hello", sCon.HelloHandler).Methods(http.MethodGet)

	hr.HandleFunc("/store", sCon.StoreListHandler).Methods(http.MethodGet)

	hr.Use(middlewares.JsonResponseMiddleware)
	hr.Use(middlewares.LoggingMiddleware)

	return hr
}

func NewGrpcRouter(db *sql.DB) *http.ServeMux {
	ser := services.NewAppService(db)
	gCon := controllers.NewGreetController(ser)

	gr := http.NewServeMux()

	greetPath, greetHandler := gCon.GreetHandler()
	gr.Handle(greetPath, middlewares.LoggingMiddleware(greetHandler))

	return gr
}
