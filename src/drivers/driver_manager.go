package drivers

type DriverManager interface {
	GetDriver(key string) *DriverInstance
}

type DriverManagerInstance struct {
	drivers map[string]*DriverInstance
}

func (driverManagerInstance DriverManagerInstance) GetDriver(key string) *DriverInstance {
	return driverManagerInstance.drivers[key]
}

func NewDriverManager(config map[string]interface{}) DriverManager {
	return DriverManagerInstance{make(map[string]*DriverInstance)}
}
