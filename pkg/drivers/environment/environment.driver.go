package environment

import (
	"os"

	"envd/pkg/drivers"
)

type EnvironmentDriverInstance struct {
	drivers.DriverInstance
}

func (driver *EnvironmentDriverInstance) GetKey(key string) (interface{}, bool) {
	v, isPresent := os.LookupEnv(key)

	if isPresent {
		return v, true
	}

	return nil, false
}

func (driver *EnvironmentDriverInstance) GetKeyDescriptorFromUniversal(key string) drivers.DriverKeyDescriptorInstance {
	return drivers.DriverKeyDescriptorInstance{
		[]string{},
		key,
	}
}

func (driver *EnvironmentDriverInstance) HasKey(key string) bool {
	_, isPresent := os.LookupEnv(key)

	return isPresent
}

func (driver *EnvironmentDriverInstance) IsLocallySignificant() bool {
	return true
}

func NewEnvDriver() *EnvironmentDriverInstance {
	return &EnvironmentDriverInstance{}
}