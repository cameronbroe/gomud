package pkg

import (
	"regexp"
)

type GomudRoute struct {
	route   *regexp.Regexp
	handler func()
}

type GomudRouter struct {
	routes []GomudRoute
}

func NewRouter() *GomudRouter {
	return new(GomudRouter)
}

func (r *GomudRouter) AddRoute(route *regexp.Regexp, handler func()) *GomudRouter {
	routeObj := GomudRoute{
		route:   route,
		handler: handler,
	}
	r.routes = append(r.routes, routeObj)
	return r
}

func (r *GomudRouter) Match(message string) {
	for _, route := range r.routes {
		if route.route.MatchString(message) {
			route.handler()
			return
		}
	}
}
