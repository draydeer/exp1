package entire

import (
	"envd/pkg/lib"
	"envd/pkg/routers"
)

var split = lib.SplitKeyFn(".")

type EntireRouterInstance struct {
	routers.RouterInstance
	Entire string
}

func (router EntireRouterInstance) Test(key string) *routers.RouterKeyDescriptorInstance {
	if key == router.Entire {
		if router.RootKey != "" {
			splitted := split(key[len(router.RootKey):])

			return &routers.RouterKeyDescriptorInstance{
				LocalCacheRootKey: router.RootKey,
				PathKey: splitted,
				RootKey: router.RootKey,
			}
		} else {
			splitted := split(key)

			return &routers.RouterKeyDescriptorInstance{
				LocalCacheRootKey: splitted[0],
				PathKey: splitted[1:],
				RootKey: splitted[0],
			}
		}
	}

	return nil
}
