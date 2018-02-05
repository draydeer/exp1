package router

type Router interface {
	AddMatchByPrefix(pattern string) Router
	//AddMatchByRegexp(pattern string) Router
}

type RouterInstance struct {
	routes []Route
}

func (router RouterInstance) AddMatchByPrefix(prefix string) Router {
	router.routes = append(router.routes, RouteMatchByPrefixInstance{prefix})

	return router
}

func NewRouter() Router {
	return RouterInstance{}
}
