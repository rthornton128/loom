package routing

import (
	"net/http"
	"regexp"
)

type Route struct {
	Handler http.HandlerFunc
	Method  string
	Path    *regexp.Regexp
}

func (r *Route) Match(request *http.Request) bool {
	return r.Method == request.Method && r.Path.MatchString(request.URL.Path)
}

func (r *Route) With(h http.HandlerFunc) *Route {
	r.Handler = h
	return r
}
