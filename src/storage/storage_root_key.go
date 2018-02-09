package storage

import (
	"time"
	"envd/src/ads"
)

type StorageRootKey interface {
	Get(key string, def interface{}) interface{}
	Update(val ads.AdsNode) StorageRootKey
}

type StorageRootKeyInstance struct {
	createdAt int64
	key string
	selectCount uint64
	val ads.AdsNode
	version uint64
	updateCount uint64
	updatedAt int64
}

func (storageRootKey StorageRootKeyInstance) Get(key string, def interface{}) interface{} {
	return ads.Get(storageRootKey.val, key, def)
}

func (storageRootKey StorageRootKeyInstance) Update(val ads.AdsNode) StorageRootKey {
	storageRootKey.val = val
	storageRootKey.updateCount += 1
	storageRootKey.updatedAt = time.Now().UnixNano()

	return storageRootKey
}

func NewStorageRootKey(key string, val ads.AdsNode) StorageRootKey {
	var storageRootKey = StorageRootKeyInstance{
		createdAt: time.Now().UnixNano(),
		key: key,
		selectCount: 0,
		version: 0,
		updateCount: 0,
		updatedAt: 0,
	}.Update(val)

	return storageRootKey
}
