package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantHttpPort := 8080
	t.Setenv("PORT", fmt.Sprint(wantHttpPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	if got.HttpPort != wantHttpPort {
		t.Errorf("want %d, but %d", wantHttpPort, got.HttpPort)
	}
}
