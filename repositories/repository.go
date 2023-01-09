package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MatsuoTakuro/my-template-connect-go/config"
	_ "github.com/lib/pq"
)

func OpenDB(ctx context.Context, cfg *config.Config) (*sql.DB, func(), error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	))
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, err
	}
	return db, func() { _ = db.Close() }, nil
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type Queryer interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

var (
	_ Queryer = (*sql.DB)(nil)
	_ Queryer = (*sql.Tx)(nil)
	_ Execer  = (*sql.DB)(nil)
	_ Execer  = (*sql.Tx)(nil)
)
