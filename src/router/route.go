package router

type Route interface {

}

type RouteMatchByPrefixInstance struct {
	pattern string
}

type RouteMatchByRegexpInstance struct {
	pattern string
}
