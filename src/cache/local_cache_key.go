package cache

import (
	"time"
	"envd/src/ads"
)

type LocalCacheKey interface {
	GetKey(key string, def interface{}) interface{}
	GetPath(path []string, def interface{}) interface{}
	Update(val ads.AdsNode) LocalCacheKey
}

type LocalCacheKeyInstance struct {
	createdAt int64
	key string
	selectCount uint64
	val ads.AdsNode
	version uint64
	updateCount uint64
	updatedAt int64
}

func (localCacheKey *LocalCacheKeyInstance) GetKey(key string, def interface{}) interface{} {
	localCacheKey.selectCount += 1

	return ads.GetKey(localCacheKey.val, key, def)
}

func (localCacheKey *LocalCacheKeyInstance) GetPath(path []string, def interface{}) interface{} {
	localCacheKey.selectCount += 1

	return ads.GetPath(localCacheKey.val, path, def)
}

func (localCacheKey *LocalCacheKeyInstance) Update(val ads.AdsNode) LocalCacheKey {
	localCacheKey.val = val
	localCacheKey.updateCount += 1
	localCacheKey.updatedAt = time.Now().UnixNano()

	return localCacheKey
}

func NewCacheKey(key string, val ads.AdsNode) *LocalCacheKeyInstance {
	var localCacheKey = LocalCacheKeyInstance{
		createdAt: time.Now().UnixNano(),
		key: key,
		selectCount: 0,
		version: 0,
		updateCount: 0,
		updatedAt: 0,
	}

	localCacheKey.Update(val)

	return &localCacheKey
}
