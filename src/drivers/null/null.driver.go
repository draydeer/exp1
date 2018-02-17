package null

import "envd/src/drivers"

var noDriverKeyInfo = drivers.DriverKeyDescriptorInstance{}

type NullDriverInstance struct {
	drivers.DriverInstance
}

func (driver *NullDriverInstance) GetKey(key string) (interface{}, bool) {
	return nil, false
}

func (driver *NullDriverInstance) GetKeyDescriptorFromUniversal(key string) drivers.DriverKeyDescriptorInstance {
	return noDriverKeyInfo
}

func (driver *NullDriverInstance) HasKey(key string) bool {
	return false
}

func NewNullDriver() *NullDriverInstance {
	return &NullDriverInstance{}
}
