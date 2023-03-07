package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var pactDir = "./pacts"

func TestProvider(t *testing.T) {

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Provider: "Classes",
	}
	DBConnect()
	// Start provider API in the background
	go HandleRequests(":8085")

	// Verify the Provider using the locally saved Pact Files
	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8085",
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/students-classes.json", pactDir))},
	})
}
