package function

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHandle ensures that Handle executes without error, returns 2000 and sets
// and sets a content-type.
func TestHandle_GET(t *testing.T) {
	var (
		w   = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "http://example.com/test", nil)
		res *http.Response
		err error
	)

	// Invoke the Handler via a standard Go http.Handler
	func(w http.ResponseWriter, req *http.Request) {
		Handle(context.Background(), w, req)
	}(w, req)
	res = w.Result()
	defer res.Body.Close()

	// Assert postconditions
	if err != nil {
		t.Fatalf("unepected error in Handle: %v", err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("unexpected response code: %v", res.StatusCode)
	}
}

func TestHandle_POST(t *testing.T) {
	var (
		w   = httptest.NewRecorder()
		req = httptest.NewRequest("POST", "http://example.com/test", strings.NewReader(`{"text":"test"}`))
		res *http.Response
		err error
	)

	// Invoke the Handler via a standard Go http.Handler
	func(w http.ResponseWriter, req *http.Request) {
		Handle(context.Background(), w, req)
	}(w, req)
	res = w.Result()
	defer res.Body.Close()

	// Assert postconditions
	if err != nil {
		t.Fatalf("unepected error in Handle: %v", err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("unexpected response code: %v", res.StatusCode)
	}
}
