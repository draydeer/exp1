package storage

import "time"

type StorageRootKey interface {
	Set(val interface{}) StorageRootKey
}

type StorageRootKeyInstance struct {
	createdAt int64
	key string
	selectCount uint64
	val interface{}
	version uint64
	updateCount uint64
	updatedAt int64
}

func (storageRootKey StorageRootKeyInstance) Set(val interface{}) StorageRootKey {
	storageRootKey.val = val
	storageRootKey.updateCount += 1
	storageRootKey.updatedAt = time.Now().UnixNano()

	return storageRootKey
}

func NewStorageRootKey(key string, val interface{}) StorageRootKey {
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
