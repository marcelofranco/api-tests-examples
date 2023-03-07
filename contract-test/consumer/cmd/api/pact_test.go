package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestClientPact_GetUser(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "Students",
		Provider: "Classes",
		Host:     "localhost",
	}
	defer pact.Teardown()

	var test = func() (err error) {
		u := fmt.Sprintf("http://localhost:%d/classes/1", pact.Server.Port)
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return
		}

		// NOTE: by default, request bodies are expected to be sent with a Content-Type
		// of application/json. If you don't explicitly set the content-type, you
		// will get a mismatch during Verification.
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer 1234")

		_, err = http.DefaultClient.Do(req)
		return
	}

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given("Class of student exist").
		UponReceiving("A request to get classes from student").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.String("/classes/1"),
			Body:   nil,
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Body:   dsl.Match(&[]Class{}),
		})

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}
