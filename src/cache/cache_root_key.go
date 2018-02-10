package cache

import (
	"time"
	"envd/src/ads"
)

type CacheRootKey interface {
	Get(key string, def interface{}) interface{}
	Update(val ads.AdsNode) CacheRootKey
}

type CacheRootKeyInstance struct {
	createdAt int64
	rootKey string
	selectCount uint64
	val ads.AdsNode
	version uint64
	updateCount uint64
	updatedAt int64
}

func (cacheRootKey *CacheRootKeyInstance) Get(key string, def interface{}) interface{} {
	cacheRootKey.selectCount += 1

	return ads.Get(cacheRootKey.val, key, def)
}

func (cacheRootKey *CacheRootKeyInstance) Update(val ads.AdsNode) CacheRootKey {
	cacheRootKey.val = val
	cacheRootKey.updateCount += 1
	cacheRootKey.updatedAt = time.Now().UnixNano()

	return cacheRootKey
}

func NewCacheRootKey(rootKey string, val ads.AdsNode) *CacheRootKeyInstance {
	var cacheRootKey = CacheRootKeyInstance{
		createdAt: time.Now().UnixNano(),
		rootKey: rootKey,
		selectCount: 0,
		version: 0,
		updateCount: 0,
		updatedAt: 0,
	}

	cacheRootKey.Update(val)

	return &cacheRootKey
}
