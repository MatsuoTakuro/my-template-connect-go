package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/MatsuoTakuro/my-template-connect-go/api"
	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	httpPort = os.Getenv("HTTP1_PORT")
	grpcPort = os.Getenv("HTTP2_PORT")
)

func main() {

	db, cleanup, err := repositories.OpenDB(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	hr := api.NewHttpRouter(db)
	log.Println("http server start at port", httpPort)

	gr := api.NewGrpcRouter(db)
	log.Println("grpc server start at port", grpcPort)

	go func() {
		log.Fatal(http.ListenAndServe(":"+httpPort, hr))
	}()

	log.Fatal(http.ListenAndServe(":"+grpcPort, h2c.NewHandler(gr, &http2.Server{})))
}
