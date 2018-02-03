package core

import (
	"envd/src/drivers"
	"envd/src/router"
)

type Core interface {
	GetDriver(key string) *drivers.DriverInstance
	GetValue(key string, def interface{}) interface{}
}

type CoreInstance struct {
	drivers.DriverManager
	router.Router
}

func (coreInstance CoreInstance) GetDriver(key string) *drivers.DriverInstance {
	return coreInstance.DriverManager.GetDriver(key)
}

func (coreInstance CoreInstance) GetValue(key string, def interface{}) interface{} {
	return def
}

func New() Core {
	return CoreInstance{}
}
