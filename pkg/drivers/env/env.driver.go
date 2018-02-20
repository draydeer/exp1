package env

import (
	"os"

	"envd/pkg/drivers"
)

type EnvDriverInstance struct {
	drivers.DriverInstance
}

func (driver *EnvDriverInstance) GetKey(key string) (interface{}, bool) {
	v, isPresent := os.LookupEnv(key)

	if isPresent {
		return v, true
	}

	return nil, false
}

func (driver *EnvDriverInstance) GetKeyDescriptorFromUniversal(key string) drivers.DriverKeyDescriptorInstance {
	return drivers.DriverKeyDescriptorInstance{
		[]string{},
		key,
	}
}

func (driver *EnvDriverInstance) HasKey(key string) bool {
	_, isPresent := os.LookupEnv(key)

	return isPresent
}

func (driver *EnvDriverInstance) IsLocallySignificant() bool {
	return true
}

func NewEnvDriver() *EnvDriverInstance {
	return &EnvDriverInstance{}
}
