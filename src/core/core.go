package core

import (
	"errors"

	"envd/src/drivers"
	"envd/src/driver_manager"
	"envd/src/lib"
	"envd/src/local_cache"
	"envd/src/router"
)

var errorCircular = errors.New("circular error")

type Core interface {
	GetCache() local_cache.LocalCache
	GetDriver(key string) drivers.Driver
	GetDriverManager() driver_manager.DriverManager
	GetKey(key string, def interface{}) (interface{}, error)
	GetKeyOrNil(key string) (interface{}, error)
	GetRouter() router.Router

	resolve(val interface{}) interface{}
}

type CoreInstance struct {
	local_cache.LocalCache
	driver_manager.DriverManager
	lockedKeys lib.Atom
	router.Router
}

func (core *CoreInstance) GetDriver(key string) drivers.Driver {
	return core.DriverManager.GetDriver(key)
}

func (core *CoreInstance) GetDriverManager() driver_manager.DriverManager {
	return core.DriverManager
}

func (core *CoreInstance) GetKey(key string, def interface{}) (interface{}, error) {
	route, significantKey, rootKeyPrefix, isMatch := core.GetRouter().Test(key)

	if isMatch {
		kd := route.GetDriver().GetKeyDescriptorFromUniversal(significantKey)
		rk := rootKeyPrefix + kd.RootKey
		sValue, sIsPresent := core.GetCache().GetKey(rk)

		if ! sIsPresent {
			if ! core.lockedKeys.Capture(kd.RootKey) {
				return nil, errorCircular
			}

			dValue, dIsPresent := route.GetDriver().GetKey(kd.RootKey)

			if ! dIsPresent {
				core.lockedKeys.Release(kd.RootKey)

				return def, nil
			}

			core.lockedKeys.Release(kd.RootKey)

			_, err := core.LocalCache.SetKeyFromRawWithMapper(rk, dValue, core.resolve)

			if err != nil {
				return nil, err
			}

			sValue, sIsPresent = core.LocalCache.GetKey(rk)
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

func (core *CoreInstance) GetCache() local_cache.LocalCache {
	return core.LocalCache
}

func (core *CoreInstance) resolve(val interface{}) (interface{}, error) {
	switch val.(type) {
	case string:
		if len(val.(string)) > 2 && val.(string)[0: 2] == "$$" {
			return core.GetKey(val.(string)[2:], nil)
		}
	}

	return val, nil
}

func NewCore(
	driverManager driver_manager.DriverManager,
	router router.Router,
	cache local_cache.LocalCache,
) *CoreInstance {
	return &CoreInstance{
		cache,
		driverManager,
		lib.NewAtom(),
		router,
	}
}
