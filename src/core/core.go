package core

import (
	"errors"
	"strings"

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
		driverKeyDescriptor := route.GetDriver().GetKeyDescriptorFromUniversal(significantKey)
		prefixedRootKey := rootKeyPrefix + driverKeyDescriptor.RootKey
		cValue, cIsPresent := core.GetCache().GetKey(prefixedRootKey)

		// cache has no value by prefixed root key
		if ! cIsPresent {

			// if prefixed root key is already captured stop with circular error
			if ! core.lockedKeys.Capture(prefixedRootKey) {
				return nil, errorCircular
			}

			dValue, dIsPresent := route.GetDriver().GetKey(driverKeyDescriptor.RootKey)

			// driver has no value by key
			if ! dIsPresent {
				core.lockedKeys.Release(prefixedRootKey)

				return def, nil
			}

			// set cache value by prefixed root key resolving references
			_, err := core.LocalCache.SetKeyFromRawWithMapper(prefixedRootKey, dValue, core.resolve)

			core.lockedKeys.Release(prefixedRootKey)

			if err != nil {
				return nil, err
			}

			cValue, cIsPresent = core.LocalCache.GetKey(prefixedRootKey)
		}

		return cValue.GetPath(driverKeyDescriptor.PathKey, def), nil
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
		str := val.(string)

		for true {
			ixe := strings.Index(str, "}}")
			ixs := strings.Index(str, "{{")

			if ixe - ixs > 0 {
				rep, err := core.GetKey(str[ixs + 2: ixe], nil)

				if err != nil {
					return nil, err
				}

				isEntire := ixs == 0 && ixe == len(str) - 2

				switch rep.(type) {
				case int:
					if isEntire {
						return rep, nil
					}

					str = str[0: ixs] + string(rep.(int)) + str[ixe:]

					continue

				case string:
					if isEntire {
						return rep, nil
					}

					str = str[0: ixs] + rep.(string) + str[ixe:]

					continue

				default:
					return rep, nil
				}
			} else {
				break
			}
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
