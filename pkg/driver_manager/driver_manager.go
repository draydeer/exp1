package driver_manager

import "envd/pkg/drivers"

type DriverManager interface {
	AddDriver(key string, driver drivers.Driver) DriverManager
	GetDriver(key string) drivers.Driver
}

type DriverManagerInstance struct {
	drivers map[string]drivers.Driver
}

func (driverManager *DriverManagerInstance) AddDriver(key string, driver drivers.Driver) DriverManager {
	driverManager.drivers[key] = driver

	return driverManager
}

func (driverManager *DriverManagerInstance) GetDriver(key string) drivers.Driver {
	return driverManager.drivers[key]
}

func NewDriverManager() DriverManagerInstance {
	return DriverManagerInstance{make(map[string]drivers.Driver)}
}
