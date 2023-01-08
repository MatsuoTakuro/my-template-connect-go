package services_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/services"

	_ "github.com/lib/pq"
)

var aSer *services.AppService

var (
	dbHost     = "127.0.0.1"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbPassword = "sa"
	dbName     = "template-db"
	dbConn     = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
)

func TestMain(m *testing.M) {
	testDB, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal(err)
	}

	aSer = services.NewAppService(testDB)

	m.Run()
}

func BenchmarkGetStoreListService(b *testing.B) {
	ctx := context.Background()

	searchQuery := "田"
	companyCD := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetStoreListService(ctx, searchQuery, companyCD)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
