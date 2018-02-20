package drivers

type Driver interface {
	GetKey(key string) (interface{}, bool)
	GetKeyDescriptorFromUniversal(key string) DriverKeyDescriptorInstance
	HasKey(key string) bool
	IsLocallySignificant() bool
}

type DriverInstance struct {

}

func (driver *DriverInstance) IsLocallySignificant() bool {
	return false
}
