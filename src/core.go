package src

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

func (core CoreInstance) GetDriver(key string) *drivers.DriverInstance {
	return core.DriverManager.GetDriver(key)
}

func (core CoreInstance) GetValue(key string, def interface{}) interface{} {
	return def
}

func NewCore() Core {
	return CoreInstance{}
}
