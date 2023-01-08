package testutils

import (
	"database/sql"
	"fmt"
	"testing"
)

var (
	dbHost     = "127.0.0.1"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbPassword = "sa"
	dbName     = "template-db"
	dbConn     = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
)

func OpenDBForTest(t *testing.T) *sql.DB {
	t.Helper()

	testDB, err := sql.Open("postgres", dbConn)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = testDB.Close() })

	return testDB
}
