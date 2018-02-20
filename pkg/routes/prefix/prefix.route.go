package prefix

import (
	"strings"
	
	"envd/pkg/routes"
)

var noString = ""

type PrefixRouteInstance struct {
	routes.RouteInstance
	Prefix string
}

func (route PrefixRouteInstance) Test(key string) (string, string, bool) {
	if strings.HasPrefix(key, route.Prefix) {
		return key, "", true
	}

	return noString, noString, false
}
