package router

type Router interface {
	AddMatchByPrefix(pattern string) Router
	//AddMatchByRegexp(pattern string) Router
}

type RouterInstance struct {
	routes []Route
}

func (router RouterInstance) AddMatchByPrefix(pattern string) Router {
	append(router.routes, RouteMatchByPrefixInstance{pattern})

	return router
}

func NewRouter() Router {
	return RouterInstance{}
}
