package null

import "envd/pkg/drivers"

type NullDriverInstance struct {
	drivers.DriverInstance
}

func (driver *NullDriverInstance) GetKey(key string) (interface{}, bool) {
	return nil, false
}

func (driver *NullDriverInstance) HasKey(key string) bool {
	return false
}

func NewNullDriver() *NullDriverInstance {
	return &NullDriverInstance{}
}
