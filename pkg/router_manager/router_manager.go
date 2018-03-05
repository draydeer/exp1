package router_manager

import (
	"envd/pkg/drivers"
	"envd/pkg/routers"
	"envd/pkg/routers/entire"
	"envd/pkg/routers/prefix"
	"envd/pkg/routers/regexp"
	"envd/pkg/routers/trimmed_prefix"
)

type RouterManager interface {
	AddEntireMatch(driver drivers.Driver, rootKey string, entireMatch string) RouterManager
	AddPrefixMatch(driver drivers.Driver, rootKey string, prefixMatch string) RouterManager
	AddRegexpMatch(driver drivers.Driver, rootKey string, regexpMatch string) RouterManager
	AddTrimmedPrefixMatch(driver drivers.Driver, rootKey string, trimmedPrefixMatch string) RouterManager
	Test(key string) (routers.Router, *routers.RouterKeyDescriptorInstance)
}

type RouterManagerInstance struct {
	routers []routers.Router
}

func (rm *RouterManagerInstance) AddEntireMatch(
	driver drivers.Driver,
	rootKey string,
	entireMatch string,
) RouterManager {
	rm.routers = append(rm.routers, entire.EntireRouterInstance{
		RouterInstance: routers.RouterInstance{Driver: driver, RootKey: rootKey},
		Entire: entireMatch,
	})

	return rm
}

func (rm *RouterManagerInstance) AddPrefixMatch(
	driver drivers.Driver,
	rootKey string,
	prefixMatch string,
) RouterManager {
	rm.routers = append(rm.routers, prefix.PrefixRouterInstance{
		RouterInstance: routers.RouterInstance{Driver: driver, RootKey: rootKey},
		Prefix: prefixMatch,
	})

	return rm
}

func (rm *RouterManagerInstance) AddRegexpMatch(
	driver drivers.Driver,
	rootKey string,
	regexpMatch string,
) RouterManager {
	rm.routers = append(rm.routers, regexp.RegexpRouterInstance{
		RouterInstance: routers.RouterInstance{Driver: driver, RootKey: rootKey},
		Regexp: regexpMatch,
	})

	return rm
}

func (rm *RouterManagerInstance) AddTrimmedPrefixMatch(
	driver drivers.Driver,
	rootKey string,
	trimmedPrefixMatch string,
) RouterManager {
	rm.routers = append(rm.routers, trimmed_prefix.TrimmedPrefixRouterInstance{
		RouterInstance: routers.RouterInstance{Driver: driver, RootKey: rootKey},
		TrimmedPrefix: trimmedPrefixMatch,
	})

	return rm
}

func (rm *RouterManagerInstance) Test(key string) (routers.Router, *routers.RouterKeyDescriptorInstance) {
	for _, router := range rm.routers {
		keyDescriptor := router.Test(key)

		if keyDescriptor != nil {
			return router, keyDescriptor
		}
	}

	return nil, nil
}

func NewRouter() RouterManagerInstance {
	return RouterManagerInstance{make([]routers.Router, 0)}
}
