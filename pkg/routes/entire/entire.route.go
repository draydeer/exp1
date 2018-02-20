package entire

import "envd/pkg/routes"

var noString = ""

type EntireRouteInstance struct {
	routes.RouteInstance
	Entire string
}

func (route EntireRouteInstance) Test(key string) (string, string, bool) {
	if key == route.Entire {
		return key, "", true
	}

	return noString, noString, false
}
