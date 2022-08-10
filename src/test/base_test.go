package test

import (
	"go-testing/server"
	"os"
	"testing"
)

const (
	baseURL = "http://localhost:8080"
)

func TestMain(m *testing.M) {
	go server.StartServer()
	os.Exit(m.Run())
}
