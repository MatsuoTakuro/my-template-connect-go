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
	sCon := controllers.NewStoreController(ser)

	gr := http.NewServeMux()

	// grpcurl -plaintext -v -proto ./proto/greet/v1/greet.proto -d '{"name": "test" }' localhost:9090 greet.v1.GreetService/Greet
	greetPath, greetHandler := gCon.GreetHandler()
	gr.Handle(greetPath, middlewares.LoggingMiddleware(greetHandler))

	// grpcurl -plaintext -v -proto ./proto/templateconnectgo/v1/store.proto  -d '{"search_query": "田", "company_cd": 1}' localhost:9090 templateconnectgo.v1.StoreService/ListStores
	storeListPath, storeListHandler := sCon.GrpcStoreListHandler()
	gr.Handle(storeListPath, middlewares.LoggingMiddleware(storeListHandler))

	return gr
}
