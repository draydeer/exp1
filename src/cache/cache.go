package cache

import "envd/src/ads"

var noRootKey = CacheRootKeyInstance{}

type Cache interface {
	GetKey(key string) (CacheRootKey, bool)
	SetKey(key string, val ads.AdsNode) Cache
	SetKeyFromRaw(key string, val interface{}) Cache
	HasKey(key string) bool
}

type CacheInstance struct {
	keys map[string]CacheRootKey
}

func (cache *CacheInstance) GetKey(key string) (CacheRootKey, bool) {
	var v, isPresent = cache.keys[key]

	if isPresent {
		return v, true
	}

	return &noRootKey, false
}

func (cache *CacheInstance) SetKey(key string, val ads.AdsNode) Cache {
	var v, isPresent = cache.keys[key]

	if isPresent {
		v.Update(val)
	} else {
		cache.keys[key] = NewCacheRootKey(key, val)
	}

	return cache
}

func (cache *CacheInstance) SetKeyFromRaw(key string, val interface{}) Cache {
	var v, isPresent = cache.keys[key]

	if isPresent {
		v.Update(ads.Create(val))
	} else {
		cache.keys[key] = NewCacheRootKey(key, ads.Create(val))
	}

	return cache
}

func (cache *CacheInstance) HasKey(key string) bool {
	var _, present = cache.keys[key]

	return present
}

func NewCache() CacheInstance {
	return CacheInstance{
		keys: make(map[string]CacheRootKey),
	}
}
