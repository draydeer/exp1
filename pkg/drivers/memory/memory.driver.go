package memory

import "envd/pkg/drivers"

type MemoryDriverInstance struct {
	drivers.DriverInstance
	Keys map[string]interface{}
}

func (driver *MemoryDriverInstance) GetKey(key string) (interface{}, bool) {
	var _, isPresent = driver.Keys[key]

	if isPresent {
		return driver.Keys[key], true
	}

	return nil, false
}

func (driver *MemoryDriverInstance) HasKey(key string) bool {
	var _, isPresent = driver.Keys[key]

	return isPresent
}

func (driver *MemoryDriverInstance) IsConstant() bool {
	return true
}


func NewMemoryDriver() *MemoryDriverInstance {
	return &MemoryDriverInstance{}
}

func NewMemoryDriverWithKeys(keys map[string]interface{}) *MemoryDriverInstance {
	return &MemoryDriverInstance{Keys: keys}
}
