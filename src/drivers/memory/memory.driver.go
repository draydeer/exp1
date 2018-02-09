package memory

type MemoryDriverInstance struct {
	keys map[string]interface{}
}

func (driver *MemoryDriverInstance) GetValue(key string) (interface{}, bool) {
	var v, present = driver.keys[key]

	if present {
		return v, true
	}

	return nil, false
}

func (driver *MemoryDriverInstance) HasValue(key string) bool {
	var _, present = driver.keys[key]

	return present
}

func NewMemoryDriver() MemoryDriverInstance {
	return MemoryDriverInstance{}
}
