package src

import (
	"envd/src/drivers"
	"envd/src/router"
)

type Core interface {
	GetDriver(key string) drivers.Driver
	GetDriverManager() drivers.DriverManager
	GetRouter() router.Router
	GetValue(key string, def interface{}) interface{}
	GetValueOrNil(key string) interface{}
}

type CoreInstance struct {
	drivers.DriverManager
	router.Router
}

func (core CoreInstance) GetDriver(key string) drivers.Driver {
	return core.DriverManager.GetDriver(key)
}

func (core CoreInstance) GetDriverManager() drivers.DriverManager {
	return core.DriverManager
}

func (core CoreInstance) GetRouter() router.Router {
	return core.Router
}

func (core CoreInstance) GetValue(key string, def interface{}) interface{} {
	return def
}

func (core CoreInstance) GetValueOrNil(key string) interface{} {
	return nil
}

func NewCore(driverManager drivers.DriverManager, router router.Router) Core {
	return CoreInstance{
		driverManager,
		router,
	}
}
