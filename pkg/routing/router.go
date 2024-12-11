// A very basic, version 1 of a router/mux. This router is not meant to be fit
// for production or any purpose. It is a starting point to build the shape of
// the framework.
package routing

import (
	"net/http"
	"regexp"
)

type Router struct {
	routes []*Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Delete(path string) *Route {
	return r.Method(http.MethodDelete, path)
}

func (r *Router) Get(path string) *Route {
	return r.Method(http.MethodGet, path)
}

func (r *Router) Patch(path string) *Route {
	return r.Method(http.MethodPatch, path)
}

func (r *Router) Post(path string) *Route {
	return r.Method(http.MethodPost, path)
}

func (r *Router) Method(method, path string) *Route {
	route := &Route{
		Method: method,
		Path:   path_to_regex(path),
	}

	r.routes = append(r.routes, route)
	return route
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, route := range r.routes {
		if route.Match(request) {
			route.Handler(writer, request)
			return
		}
	}

	http.NotFound(writer, request)
}

var keyMatcher = regexp.MustCompile(`:(\w+)`)

func path_to_regex(path string) *regexp.Regexp {
	return regexp.MustCompile(keyMatcher.ReplaceAllString(path, `(?<$1>\w+)`))
}
