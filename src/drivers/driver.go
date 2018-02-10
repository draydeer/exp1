package drivers

type Driver interface {
	GetKey(key string) (interface{}, bool)
	HasKey(key string) bool
}

type DriverInstance struct {

}
