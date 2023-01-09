package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/MatsuoTakuro/my-template-connect-go/config"
	"github.com/MatsuoTakuro/my-template-connect-go/testutils"
	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		// start http server
		s := NewServer(testutils.OpenDBForTest(t), l, cfg)
		return s.Run(ctx)
	})
	// GET HelloHandler (on http)
	path := "hello"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), path)
	t.Logf("try request to %q", url)
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body) // read response message
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// verify the server's termination behavior
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}

	// verify the return value
	want := `{"message": "Hello, world! by store"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
