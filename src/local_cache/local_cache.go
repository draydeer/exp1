package local_cache

import "envd/src/ads"

var noCacheKey = LocalCacheKeyInstance{}

type LocalCache interface {
	GetKey(key string) (LocalCacheKey, bool)
	SetKey(key string, val ads.AdsNode) LocalCache
	SetKeyFromRaw(key string, val interface{}) LocalCache
	HasKey(key string) bool
}

type LocalCacheInstance struct {
	keys map[string]LocalCacheKey
}

func (localCache *LocalCacheInstance) GetKey(key string) (LocalCacheKey, bool) {
	var v, isPresent = localCache.keys[key]

	if isPresent {
		return v, true
	}

	return &noCacheKey, false
}

func (localCache *LocalCacheInstance) SetKey(key string, val ads.AdsNode) LocalCache {
	var v, isPresent = localCache.keys[key]

	if isPresent {
		v.Update(val)
	} else {
		localCache.keys[key] = NewCacheKey(key, val)
	}

	return localCache
}

func (localCache *LocalCacheInstance) SetKeyFromRaw(key string, val interface{}) LocalCache {
	var v, isPresent = localCache.keys[key]

	if isPresent {
		v.Update(ads.Create(val))
	} else {
		localCache.keys[key] = NewCacheKey(key, ads.Create(val))
	}

	return localCache
}

func (localCache *LocalCacheInstance) HasKey(key string) bool {
	var _, present = localCache.keys[key]

	return present
}

func NewCache() LocalCacheInstance {
	return LocalCacheInstance{
		keys: make(map[string]LocalCacheKey),
	}
}
