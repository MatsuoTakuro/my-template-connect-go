package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/MatsuoTakuro/my-template-connect-go/repositories"
)

var (
	httpPort = os.Getenv("HTTP1_PORT")
	grpcPort = os.Getenv("HTTP2_PORT")
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	db, cleanup, err := repositories.OpenDB(context.Background())
	if err != nil {
		return err
	}
	defer cleanup()

	hl, err := net.Listen("tcp", fmt.Sprintf(":%s", httpPort))
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", httpPort, err)
	}
	hu := fmt.Sprintf("http://%s", hl.Addr().String())
	log.Printf("start with: %v", hu)
	hs := NewServer(db, hl)

	gl, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", grpcPort, err)
	}
	gu := fmt.Sprintf("grpc://%s", gl.Addr().String())
	log.Printf("start with: %v", gu)
	gs := NewServer(db, gl)

	go func() error {
		return hs.Run(ctx)
	}()

	return gs.Run(ctx)
}
