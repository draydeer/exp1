package consul

import "envd/src/drivers"

type ConsulDriverInstance struct {

}

func NewDriver() drivers.Driver {
	return ConsulDriverInstance{}
}

func (driver ConsulDriverInstance) Get(key string, def interface{}) interface{} {
	return def
}
