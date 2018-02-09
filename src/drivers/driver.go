package drivers

type Driver interface {
	GetValue(key string) (interface{}, bool)
	HasValue(key string) bool
}

type DriverInstance struct {

}
