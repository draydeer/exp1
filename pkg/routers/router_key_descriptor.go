package routers

type RouterKeyDescriptorInstance struct {
	LocalCacheRootKey string // local cache root key associated with driver and driver root key
	PathKey []string // internal path to key
	RootKey string // driver root key
}
