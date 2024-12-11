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

func (r *Route) Match(path string) bool {
	return r.Path.MatchString(path)
}

func (r *Route) With(h http.HandlerFunc) *Route {
	r.Handler = h
	return r
}
