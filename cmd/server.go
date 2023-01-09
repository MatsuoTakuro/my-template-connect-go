package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/MatsuoTakuro/my-template-connect-go/api"
	"github.com/MatsuoTakuro/my-template-connect-go/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(db *sql.DB, l net.Listener, cfg *config.Config) *Server {
	strPort := l.Addr().String()[strings.LastIndex(l.Addr().String(), ":")+1:]
	port, err := strconv.Atoi(strPort)
	if err != nil {
		log.Fatalf("failed to parseInt string port %s: %v", strPort, err)
	}

	if port == cfg.GrpcPort {
		gr := api.NewGrpcRouter(db)

		return &Server{
			srv: &http.Server{Handler: h2c.NewHandler(gr, &http2.Server{})},
			l:   l,
		}
	} else {
		hr := api.NewHttpRouter(db)

		return &Server{
			srv: &http.Server{Handler: hr},
			l:   l,
		}
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.srv.Serve(s.l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return eg.Wait()
}

func NewListner(port int) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(":%d", port))
}
