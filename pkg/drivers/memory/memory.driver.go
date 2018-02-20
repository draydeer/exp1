package memory

import (
	"envd/pkg/drivers"
	"envd/pkg/lib"
)

var split = lib.SplitKeyFn(".")

type MemoryDriverInstance struct {
	drivers.DriverInstance
	Keys map[string]interface{}
}

func (driver *MemoryDriverInstance) GetKey(key string) (interface{}, bool) {
	var _, isPresent = driver.Keys[key]

	if isPresent {
		return driver.Keys, true
	}

	return nil, false
}

func (driver *MemoryDriverInstance) GetKeyDescriptorFromUniversal(key string) drivers.DriverKeyDescriptorInstance {
	splitted := split(key)

	return drivers.DriverKeyDescriptorInstance{
		splitted,
		splitted[0],
	}
}

func (driver *MemoryDriverInstance) HasKey(key string) bool {
	var _, isPresent = driver.Keys[key]

	return isPresent
}

func NewMemoryDriver() *MemoryDriverInstance {
	return &MemoryDriverInstance{}
}

func NewMemoryDriverWithKeys(keys map[string]interface{}) *MemoryDriverInstance {
	return &MemoryDriverInstance{Keys: keys}
}
