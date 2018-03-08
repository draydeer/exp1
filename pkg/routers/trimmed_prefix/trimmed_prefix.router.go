package trimmed_prefix

import (
	"strings"

	"envd/pkg/lib"
	"envd/pkg/routers"
)

var split = lib.SplitKeyFn(".")

type TrimmedPrefixRouterInstance struct {
	routers.RouterInstance
	TrimmedPrefix string
}

/*

Example
=======

CustomDriver data block:
	test:
		key1: "val1"
		key2:
			key2key1: "key2val1"
			key2key2:
				hello:
					world: "hello, world!"

Configuration 1
---------------

TrimmerPrefix router configuration:
	Driver: CustomDriver
	RootKey: "test/key2"
	TrimmedPrefix: "custom:"

Requested key: "custom:key2key2.hello"

RouterKeyDescriptorInstance:
	LocalCacheKey: "custom:test/key2"  <--  by this key data block will be stored locally
	PathKey: ["key2key2", "hello"]  <--  by this path target value will be searched in data block
	RootKey: "test/key2"  <--  by this key data block will be requested via driver

Configuration 2
---------------

TrimmerPrefix router configuration:
	Driver: CustomDriver
	RootKey: ""
	TrimmedPrefix: "custom:"

Requested key: "custom:test.key2.key2key2.hello"

RouterKeyDescriptorInstance:
	LocalCacheKey: "custom:test"  <--  by this key data block will be stored locally
	PathKey: ["key2", "key2key2", "hello"]  <--  by this path target value will be searched in data block
	RootKey: "test"  <--  by this key data block will be requested via driver

 */
func (router TrimmedPrefixRouterInstance) Test(key string) *routers.RouterKeyDescriptorInstance {
	if strings.HasPrefix(key, router.TrimmedPrefix) {
		if router.RootKey != "" {
			splitted := split(strings.TrimPrefix(key, router.TrimmedPrefix))

			return &routers.RouterKeyDescriptorInstance{
				LocalCacheRootKey: router.TrimmedPrefix + router.RootKey,
				PathKey: splitted,
				RootKey: router.RootKey,
			}
		} else {
			splitted := split(strings.TrimPrefix(key, router.TrimmedPrefix))

			return &routers.RouterKeyDescriptorInstance{
				LocalCacheRootKey: router.TrimmedPrefix + splitted[0],
				PathKey: splitted[1:],
				RootKey: splitted[0],
			}
		}
	}

	return nil
}
