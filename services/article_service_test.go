package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/services"

	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.AppService

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func TestMain(m *testing.M) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aSer = services.NewAppService(db)

	m.Run()
}
