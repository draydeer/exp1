package router

import (
	"envd/pkg/drivers"
	"envd/pkg/routes"
	"envd/pkg/routes/entire"
	"envd/pkg/routes/prefix"
	"envd/pkg/routes/regexp"
	"envd/pkg/routes/trimmed_prefix"
)

var noRoute = routes.RouteInstance{}
var noString = ""

type Router interface {
	AddEntireMatch(driver drivers.Driver, pattern string) Router
	AddPrefixMatch(driver drivers.Driver, pattern string) Router
	AddRegexpMatch(driver drivers.Driver, pattern string) Router
	AddTrimmedPrefixMatch(driver drivers.Driver, pattern string) Router
	Test(key string) (routes.Route, string, string, bool)
}

type RouterInstance struct {
	routes []routes.Route
}

func (router *RouterInstance) AddEntireMatch(driver drivers.Driver, entireMatch string) Router {
	router.routes = append(router.routes, entire.EntireRouteInstance{
		routes.RouteInstance{driver},
		entireMatch,
	})

	return router
}

func (router *RouterInstance) AddPrefixMatch(driver drivers.Driver, prefixMatch string) Router {
	router.routes = append(router.routes, prefix.PrefixRouteInstance{
		routes.RouteInstance{driver},
		prefixMatch,
	})

	return router
}

func (router *RouterInstance) AddRegexpMatch(driver drivers.Driver, regexpMatch string) Router {
	router.routes = append(router.routes, regexp.RegexpRouteInstance{
		routes.RouteInstance{driver},
		regexpMatch,
	})

	return router
}

func (router *RouterInstance) AddTrimmedPrefixMatch(driver drivers.Driver, trimmedPrefixMatch string) Router {
	router.routes = append(router.routes, trimmed_prefix.TrimmedPrefixRouteInstance{
		routes.RouteInstance{driver},
		trimmedPrefixMatch,
	})

	return router
}

func (router *RouterInstance) Test(key string) (routes.Route, string, string, bool) {
	for _, route := range router.routes {
		significantKey, prefix, isMatch := route.Test(key)

		if isMatch {
			return route, significantKey, prefix, true
		}
	}

	return noRoute, noString, noString, false
}

func NewRouter() RouterInstance {
	return RouterInstance{make([]routes.Route, 0)}
}
