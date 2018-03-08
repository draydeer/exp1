package core

import (
	"errors"
	"strings"

	"envd/pkg/drivers"
	"envd/pkg/driver_manager"
	"envd/pkg/lib"
	"envd/pkg/local_cache"
	"envd/pkg/router_manager"
)

var errorCircular = errors.New("circular error")

type Core interface {
	GetCache() local_cache.LocalCache
	GetDriver(key string) drivers.Driver
	GetDriverManager() driver_manager.DriverManager
	GetKey(key string, def interface{}) (interface{}, error)
	GetKeyOrNil(key string) (interface{}, error)
	GetRouterManager() router_manager.RouterManager
	UpdKey(key string) (interface{}, error)

	updateBlockReferences(val interface{}) interface{}
}

type CoreInstance struct {
	DriverManager *driver_manager.DriverManager
	LocalCache *local_cache.LocalCache
	RouterManager *router_manager.RouterManager
	lockedKeys *lib.AtomInstance
}

func (core *CoreInstance) GetDriver(key string) drivers.Driver {
	return (*core.DriverManager).GetDriver(key)
}

func (core *CoreInstance) GetDriverManager() driver_manager.DriverManager {
	return *core.DriverManager
}

func (core *CoreInstance) GetKey(key string, def interface{}) (interface{}, error) {
	router, routerKeyDescriptor := core.GetRouterManager().Test(key)

	if router != nil {
		cValue, cIsPresent := core.GetCache().GetKey(routerKeyDescriptor.LocalCacheRootKey)

		// cache has no value by root key
		if ! cIsPresent {
			transactionId := lib.NewTransactionId()

			// if root key is already captured by same transaction then stop with circular error
			if ! core.lockedKeys.Capture(routerKeyDescriptor.LocalCacheRootKey, transactionId) {
				return nil, errorCircular
			}

			cValue, cIsPresent = core.GetCache().GetKey(routerKeyDescriptor.LocalCacheRootKey)

			if ! cIsPresent {
				dValue, dIsPresent := router.GetDriver().GetKey(routerKeyDescriptor.RootKey)

				// driver has no value by key
				if ! dIsPresent {
					core.lockedKeys.Release(routerKeyDescriptor.LocalCacheRootKey, transactionId)

					return def, nil
				}

				// set cache value by prefixed root key resolving references
				_, err := core.GetCache().SetKeyFromRawWithMapper(
					routerKeyDescriptor.LocalCacheRootKey,
					dValue,
					func (val interface{}) (interface{}, error) {
						return core.updateBlockReferences(val)
					},
				)

				core.lockedKeys.Release(routerKeyDescriptor.LocalCacheRootKey, transactionId)

				if err != nil {
					return nil, err
				}

				cValue, cIsPresent = core.GetCache().GetKey(routerKeyDescriptor.LocalCacheRootKey)
			} else {
				core.lockedKeys.Release(routerKeyDescriptor.LocalCacheRootKey, transactionId)
			}
		}

		return cValue.GetPath(routerKeyDescriptor.PathKey, def), nil
	}

	return def, nil
}

func (core *CoreInstance) GetKeyOrNil(key string) (interface{}, error) {
	return core.GetKey(key, nil)
}

func (core *CoreInstance) GetKeyDescriptor(key string) (local_cache.LocalCacheKey, error) {
	router, routerKeyDescriptor := core.GetRouterManager().Test(key)

	if router != nil {
		cValue, cIsPresent := core.GetCache().GetKey(routerKeyDescriptor.LocalCacheRootKey)

		if ! cIsPresent {
			return nil, nil
		}

		return cValue, nil
	}

	return nil, nil
}

func (core *CoreInstance) GetRouterManager() router_manager.RouterManager {
	return *core.RouterManager
}

func (core *CoreInstance) GetCache() local_cache.LocalCache {
	return *core.LocalCache
}

func (core *CoreInstance) updateBlockReferences(val interface{}) (interface{}, error) {
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

func (core *Core) UpdKey(key string) (interface{}, bool) {

}

func NewCore(
	driverManager driver_manager.DriverManager,
	router router_manager.RouterManager,
	cache local_cache.LocalCache,
) *CoreInstance {
	return &CoreInstance{
		DriverManager: &driverManager,
		LocalCache: &cache,
		RouterManager: &router,
		lockedKeys: lib.NewAtom(),
	}
}
