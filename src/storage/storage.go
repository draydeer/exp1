package storage

type Storage interface {
	GetValue(key string, def interface{}) interface{}
}

type StorageInstance struct {
	rootKeys map[string]StorageRootKeyInstance
}

func (storage *StorageInstance) GetValue(key string, def interface{}) interface{} {
	if val, ok := storage.rootKeys[key]; ok {
		return val
	}

	return nil
}

func (storage *StorageInstance) SetValue(key string, val interface{}) Storage {
	return storage
}

func NewStorage() StorageInstance {
	return StorageInstance{
		rootKeys: make(map[string]StorageRootKeyInstance),
	}
}
