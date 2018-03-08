package local_cache

import (
	"time"
	"envd/pkg/ads"
	"sync/atomic"
)

type LocalCacheKey interface {
	GetDescriptor() *LocalCacheKeyInstance
	GetKey(key string, def interface{}) interface{}
	GetPath(path []string, def interface{}) interface{}
	Upd(val ads.AdsNode) LocalCacheKey
}

type LocalCacheKeyInstance struct {
	CreatedAt int64
	Key string
	SelectCount uint64
	Version uint64
	UpdateCount uint64
	UpdatedAt int64

	val ads.AdsNode
}

func (localCacheKey *LocalCacheKeyInstance) GetDescriptor() *LocalCacheKeyInstance {
	return localCacheKey
}

func (localCacheKey *LocalCacheKeyInstance) GetKey(key string, def interface{}) interface{} {
	atomic.AddUint64(&localCacheKey.SelectCount, 1)

	return ads.GetKey(localCacheKey.val, key, def)
}

func (localCacheKey *LocalCacheKeyInstance) GetPath(path []string, def interface{}) interface{} {
	atomic.AddUint64(&localCacheKey.SelectCount, 1)

	return ads.GetPath(localCacheKey.val, path, def)
}

func (localCacheKey *LocalCacheKeyInstance) Upd(val ads.AdsNode) LocalCacheKey {
	localCacheKey.val = val
	localCacheKey.UpdateCount += 1
	localCacheKey.UpdatedAt = time.Now().UnixNano()

	return localCacheKey
}

func NewCacheKey(key string, val ads.AdsNode) *LocalCacheKeyInstance {
	var localCacheKey = LocalCacheKeyInstance{
		CreatedAt: time.Now().UnixNano(),
		Key: key,
		SelectCount: 0,
		Version: 0,
		UpdateCount: 0,
		UpdatedAt: 0,
	}

	localCacheKey.Upd(val)

	return &localCacheKey
}
