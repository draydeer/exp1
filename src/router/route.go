package router

import (
	"envd/src/drivers"
	"strings"
)

type Route interface {
	GetDriver() drivers.Driver
	Test(key string) (string, bool)
}

type RouteInstance struct {
	driver drivers.Driver
}

func (route RouteInstance) GetDriver() drivers.Driver {
	return route.driver
}

func (route RouteInstance) Test(key string) (string, bool) {
	return noString, false
}

// Route match by prefix

type RouteMatchByPrefixInstance struct {
	RouteInstance
	prefix string
}

func (route RouteMatchByPrefixInstance) Test(key string) (string, bool) {
	if strings.HasPrefix(key, route.prefix) {
		return key, true
	}

	return noString, false
}

// Route match by regexp

type RouteMatchByRegexpInstance struct {
	RouteInstance
	regexp string
}

func (route RouteMatchByRegexpInstance) Test(key string) (string, bool) {
	return noString, false
}

// Route match by trimmed prefix

type RouteMatchByTrimmedPrefixInstance struct {
	RouteInstance
	prefix string
}

func (route RouteMatchByTrimmedPrefixInstance) Test(key string) (string, bool) {
	if strings.HasPrefix(key, route.prefix) {
		return strings.TrimPrefix(key, route.prefix), true
	}

	return noString, false
}
