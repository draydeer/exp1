package drivers

type DriverManager interface {
	AddDriver(key string, driver Driver) *DriverManagerInstance
	GetDriver(key string) *DriverInstance
}

type DriverManagerInstance struct {
	drivers map[string]DriverInstance
}

func (driverManager DriverManagerInstance) AddDriver(key string, driver *DriverInstance) DriverManagerInstance {
	driverManager.drivers[key] = driver

	return driverManager
}

func (driverManager DriverManagerInstance) GetDriver(key string) *DriverInstance {
	return driverManager.drivers[key]
}

func NewDriverManager(config map[string]interface{}) DriverManager {
	return DriverManagerInstance{make(map[string]DriverInstance)}
}
