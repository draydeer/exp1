package router

import "envd/src/drivers"

var noRoute = RouteInstance{}

type Router interface {
	AddMatchByPrefix(driver drivers.Driver, pattern string) Router
	AddMatchByRegexp(driver drivers.Driver, pattern string) Router
	AddMatchByTrimmedPrefix(driver drivers.Driver, pattern string) Router
	Test(key string) (Route, string, bool)
}

type RouterInstance struct {
	routes []Route
}

func (router *RouterInstance) AddMatchByPrefix(driver drivers.Driver, prefix string) Router {
	router.routes = append(router.routes, RouteMatchByPrefixInstance{
		RouteInstance{driver},
		prefix,
	})

	return router
}

func (router *RouterInstance) AddMatchByRegexp(driver drivers.Driver, regexp string) Router {
	router.routes = append(router.routes, RouteMatchByRegexpInstance{
		RouteInstance{driver},
		regexp,
		})

	return router
}

func (router *RouterInstance) AddMatchByTrimmedPrefix(driver drivers.Driver, prefix string) Router {
	router.routes = append(router.routes, RouteMatchByTrimmedPrefixInstance{
		RouteInstance{driver},
		prefix,
		})

	return router
}

func (router *RouterInstance) Test(key string) (Route, string, bool) {
	for _, route := range router.routes {
		var significantPart, isMatch = route.Test(key)

		if isMatch {
			return route, significantPart, true
		}
	}

	return noRoute, noString, false
}

func NewRouter() RouterInstance {
	return RouterInstance{make([]Route, 0)}
}
