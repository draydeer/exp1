package storage

type Storage interface {
	Get(key string, def interface{}) interface{}
}

type StorageInstance struct {
	rootKeys map[string]StorageRootKeyInstance
}

func (storage StorageInstance) Get(key string, def interface{}) interface{} {
	if val, ok := storage.rootKeys[key]; ok {
		return val
	}

	return nil
}

func (storage StorageInstance) Set(key string, val interface{}) Storage {
	return storage
}

func NewStorage() Storage {
	return StorageInstance{
		rootKeys: make(map[string]StorageRootKeyInstance),
	}
}
