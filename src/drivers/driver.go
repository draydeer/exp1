package drivers

type Driver interface {
	Test() string
	Get(key string, def interface{}) interface{}
}

type DriverInstance struct {
	A int
}

func (driverOptions DriverInstance) Test() string {
	return "1"
}

func (driverOptions DriverInstance) Get(key string, def interface{}) interface{} {
	return def
}
