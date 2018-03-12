package drivers

type Driver interface {
	GetKey(key string) (interface{}, bool)
	HasKey(key string) bool
	IsClientSignificant() bool // flag showing that requested key should be resolved in "client" mode only
	IsConstant() bool // flag showing that data returned by driver is always constant and doesn't require resync
}

type DriverInstance struct {

}

func (driver *DriverInstance) IsClientSignificant() bool {
	return false
}

func (driver *DriverInstance) IsConstant() bool {
	return false
}
