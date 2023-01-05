package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MatsuoTakuro/my-template-connect-go/api"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	httpPort   = os.Getenv("HTTP1_PORT")
	grpcPort   = os.Getenv("HTTP2_PORT")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	hr := api.NewHttpRouter(db)
	log.Println("http server start at port", httpPort)

	gr := api.NewGrpcRouter(db)
	log.Println("grpc server start at port", grpcPort)

	go func() {
		log.Fatal(http.ListenAndServe(":"+httpPort, hr))
	}()

	log.Fatal(http.ListenAndServe(":"+grpcPort, h2c.NewHandler(gr, &http2.Server{})))
}
