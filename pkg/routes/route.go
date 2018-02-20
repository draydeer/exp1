package routes

import (
	"envd/pkg/drivers"
)

var noString = ""

type Route interface {
	GetDriver() drivers.Driver
	Test(key string) (string, string, bool)
}

type RouteInstance struct {
	Driver drivers.Driver
}

func (route RouteInstance) GetDriver() drivers.Driver {
	return route.Driver
}

func (route RouteInstance) Test(key string) (string, string, bool) {
	return noString, noString, false
}
