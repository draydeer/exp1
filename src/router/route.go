package router

type Route interface {

}

type RouteMatchByPrefixInstance struct {
	driver string
	prefix string
}

type RouteMatchByRegexpInstance struct {
	driver string
	regexp string
}
