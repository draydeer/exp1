package memory

type MemoryDriverInstance struct {
	keys map[string]interface{}
}

func (driver *MemoryDriverInstance) GetKey(key string) (interface{}, bool) {
	var _, present = driver.keys[key]

	if present {
		return driver.keys, true
	}

	return nil, false
}

func (driver *MemoryDriverInstance) HasKey(key string) bool {
	var _, present = driver.keys[key]

	return present
}

func NewMemoryDriver() MemoryDriverInstance {
	return MemoryDriverInstance{}
}

func NewMemoryDriverWithKeys(keys map[string]interface{}) MemoryDriverInstance {
	return MemoryDriverInstance{keys}
}
