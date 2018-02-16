package src

import (
	"envd/src/drivers"
	"envd/src/router"
	"envd/src/cache"
)

type Core interface {
	GetCache() cache.LocalCache
	GetDriver(key string) drivers.Driver
	GetDriverManager() drivers.DriverManager
	GetKey(key string, def interface{}) (interface{}, error)
	GetKeyOrNil(key string) (interface{}, error)
	GetRouter() router.Router
}

type CoreInstance struct {
	cache.LocalCache
	drivers.DriverManager
	router.Router
	lockedKeys map[string]bool
}

func (core *CoreInstance) GetDriver(key string) drivers.Driver {
	return core.DriverManager.GetDriver(key)
}

func (core *CoreInstance) GetDriverManager() drivers.DriverManager {
	return core.DriverManager
}

func (core *CoreInstance) GetKey(key string, def interface{}) (interface{}, error) {
	var route, significantKey, isMatch = core.GetRouter().Test(key)

	if isMatch {
		var kd = route.GetDriver().GetKeyDescriptorFromUniversal(significantKey)
		var sValue, sIsPresent = core.GetCache().GetKey(kd.RootKey)

		if ! sIsPresent {
			var dValue, dIsPresent = route.GetDriver().GetKey(kd.RootKey)

			if ! dIsPresent {
				return def, nil
			}

			sValue, sIsPresent = core.LocalCache.SetKeyFromRaw(kd.RootKey, dValue).GetKey(kd.RootKey)
		}

		return sValue.GetPath(kd.PathKey, def), nil
	}

	return def, nil
}

func (core *CoreInstance) GetKeyOrNil(key string) (interface{}, error) {
	return core.GetKey(key, nil)
}

func (core *CoreInstance) GetRouter() router.Router {
	return core.Router
}

func (core *CoreInstance) GetCache() cache.LocalCache {
	return core.LocalCache
}

func NewCore(
	driverManager drivers.DriverManager,
	router router.Router,
	cache cache.LocalCache,
) CoreInstance {
	return CoreInstance{
		cache,
		driverManager,
		router,
		make(map[string]bool, {})
	}
}
