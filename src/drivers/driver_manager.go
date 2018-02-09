package drivers

type DriverManager interface {
	AddDriver(key string, driver Driver) DriverManager
	GetDriver(key string) Driver
}

type DriverManagerInstance struct {
	drivers map[string]Driver
}

func (driverManager *DriverManagerInstance) AddDriver(key string, driver Driver) DriverManager {
	driverManager.drivers[key] = driver

	return driverManager
}

func (driverManager *DriverManagerInstance) GetDriver(key string) Driver {
	return driverManager.drivers[key]
}

func NewDriverManager() DriverManagerInstance {
	return DriverManagerInstance{make(map[string]Driver)}
}
