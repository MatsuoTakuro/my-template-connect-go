package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MatsuoTakuro/my-template-connect-go/config"
	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	db, cleanup, err := repositories.OpenDB(context.Background(), cfg)
	if err != nil {
		return err
	}
	defer cleanup()

	httpLner, err := NewListner(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.HttpPort, err)
	}
	httpSrv := NewServer(db, httpLner, cfg)

	grpcLner, err := NewListner(cfg.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.GrpcPort, err)
	}
	grpcSrv := NewServer(db, grpcLner, cfg)

	go func() error {
		log.Printf("start with: %v", fmt.Sprintf("http://%s", httpSrv.l.Addr().String()))
		return httpSrv.Run(ctx)
	}()

	log.Printf("start with: %v", fmt.Sprintf("grpc://%s", grpcSrv.l.Addr().String()))
	return grpcSrv.Run(ctx)
}
