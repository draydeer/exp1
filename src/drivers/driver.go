package drivers

type Driver interface {
	GetKey(key string) (interface{}, bool)
	GetKeyDescriptorFromUniversal(key string) DriverKeyDescriptorInstance
	HasKey(key string) bool
}

type DriverInstance struct {

}
