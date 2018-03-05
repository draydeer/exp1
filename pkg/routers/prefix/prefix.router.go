package prefix

import (
	"strings"

	"envd/pkg/lib"
	"envd/pkg/routers"
)

var split = lib.SplitKeyFn(".")

type PrefixRouterInstance struct {
	routers.RouterInstance
	Prefix string
}

func (router PrefixRouterInstance) Test(key string) *routers.RouterKeyDescriptorInstance {
	if strings.HasPrefix(key, router.Prefix) {
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
