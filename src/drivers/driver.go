package drivers

type Driver interface {
	Get(key string, def interface{}) interface{}
}

type DriverInstance struct {

}

func (driver DriverInstance) Get(key string, def interface{}) interface{} {
	return def
}
