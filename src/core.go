package src

import (
	"envd/src/drivers"
	"envd/src/router"
	"envd/src/cache"
	"strings"
)

type Core interface {
	GetCache() cache.Cache
	GetDriver(key string) drivers.Driver
	GetDriverManager() drivers.DriverManager
	GetRouter() router.Router
	GetValue(key string, def interface{}) interface{}
	GetValueOrNil(key string) interface{}
}

type CoreInstance struct {
	cache.Cache
	drivers.DriverManager
	router.Router
}

func (core *CoreInstance) GetDriver(key string) drivers.Driver {
	return core.DriverManager.GetDriver(key)
}

func (core *CoreInstance) GetDriverManager() drivers.DriverManager {
	return core.DriverManager
}

func (core *CoreInstance) GetRouter() router.Router {
	return core.Router
}

func (core *CoreInstance) GetCache() cache.Cache {
	return core.Cache
}

func (core *CoreInstance) GetValue(key string, def interface{}) interface{} {
	var route, significantKey, isMatch = core.GetRouter().Test(key)

	if isMatch {
		var keys = strings.Split(significantKey, ".")
		var sValue, sIsPresent = core.GetCache().GetKey(keys[0])

		if ! sIsPresent {
			var dValue, dIsPresent = route.GetDriver().GetKey(keys[0])

			if ! dIsPresent {
				return def
			}

			sValue, sIsPresent = core.Cache.SetKeyFromRaw(keys[0], dValue).GetKey(keys[0])
		}

		return sValue.Get(significantKey, def)
	}

	return def
}

func (core *CoreInstance) GetValueOrNil(key string) interface{} {
	return core.GetValue(key, nil)
}

func NewCore(
	driverManager drivers.DriverManager,
	router router.Router,
	cache cache.Cache,
) CoreInstance {
	return CoreInstance{
		cache,
		driverManager,
		router,
	}
}
