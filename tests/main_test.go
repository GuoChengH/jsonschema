package tests

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Start the server and ensure it's only started once
	server := startTestServer()

	// Run the tests
	code := m.Run()

	// Shut down the server
	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}

	os.Exit(code)
}
