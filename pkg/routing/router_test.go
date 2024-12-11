package routing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rthornton128/loom/pkg/routing"
)

func TestRouter(t *testing.T) {
	r := routing.NewRouter()

	route := r.Delete("/")
	if route.Method != http.MethodDelete {
		t.Error("route should have set DELETE method")
	}

	route = r.Get("/")
	if route.Method != http.MethodGet {
		t.Error("route should have set GET method")
	}

	route = r.Patch("/")
	if route.Method != http.MethodPatch {
		t.Error("route should have set PATCH method")
	}

	route = r.Post("/")
	if route.Method != http.MethodPost {
		t.Error("route should have set POST method")
	}

	route = r.Get("/:id/b")
	request := httptest.NewRequest(http.MethodGet, "/a/b", nil)
	if !route.Match(request) {
		t.Error("expected route to match wildcard")
	}
}
