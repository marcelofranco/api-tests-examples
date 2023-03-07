package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcelofranco/api-tests-examples/contract-test/provider/data"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var testApp Config

func Test_GetStudentClasses(t *testing.T) {
	repo := data.NewPosgrestTestRepository(nil)
	testApp.Repo = repo

	expected := `
	[
		{
			"id": 1,
			"discipline": "Math",
			"day": "Tuesday",
			"hour": "10:00"
		},
		{
			"id": 2,
			"discipline": "Geography",
			"day": "Wednesday",
			"hour": "11:00"
		}
	]
`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/classes/1" {
			t.Errorf("Expected to request '/classes/1', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expected))
	}))
	defer server.Close()

	testApp.Client = NewClient(server.URL)

	w := httptest.NewRecorder()

	ctx := GetTestGinContext(w)

	//configure path params
	params := []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	MockJsonGet(ctx, params)
	testApp.GetStudentClasses(ctx)

	if w.Code != http.StatusOK {
		t.Errorf("expected http.StatusAccepted but got %d", w.Code)
	}
}

// mock gin context
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

// mock getrequest
func MockJsonGet(c *gin.Context, params gin.Params) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	// set path params
	c.Params = params
}
