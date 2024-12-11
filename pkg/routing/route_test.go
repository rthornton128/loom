package routing_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/rthornton128/loom/pkg/routing"
)

func TestRouteMatch(t *testing.T) {
	path := "/a/b"
	route := &routing.Route{
		Method: http.MethodGet,
		Path:   regexp.MustCompile(path),
	}

	request := httptest.NewRequest(http.MethodGet, "/a/b", nil)
	if !route.Match(request) {
		t.Error("expected path to match route")
	}

	request = httptest.NewRequest(http.MethodGet, "/z", nil)
	if route.Match(request) {
		t.Error("expected path not to match route")
	}
}

func TestRouteWith(t *testing.T) {
	route := &routing.Route{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
	result := route.With(handler)

	if result != route {
		t.Error("expected With to return self")
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	route.Handler(w, r)
	if w.Code != http.StatusNoContent {
		t.Error("expected With to have set Handler to passed function")
	}
}
