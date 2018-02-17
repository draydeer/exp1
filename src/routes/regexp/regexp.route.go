package regexp

import (
	"envd/src/routes"
)

var noString = ""

type RegexpRouteInstance struct {
	routes.RouteInstance
	Regexp string
}

func (route RegexpRouteInstance) Test(key string) (string, string, bool) {
	return noString, noString, false
}
