package trimmed_prefix

import (
	"strings"
	
	"envd/pkg/routes"
)

var noString = ""

type TrimmedPrefixRouteInstance struct {
	routes.RouteInstance
	TrimmedPrefix string
}

func (route TrimmedPrefixRouteInstance) Test(key string) (string, string, bool) {
	if strings.HasPrefix(key, route.TrimmedPrefix) {
		return strings.TrimPrefix(key, route.TrimmedPrefix), route.TrimmedPrefix, true
	}

	return noString, noString, false
}
