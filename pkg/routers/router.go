package routers

import (
	"envd/pkg/drivers"
)

type Router interface {
	GetDriver() drivers.Driver
	Test(key string) *RouterKeyDescriptorInstance
}

type RouterInstance struct {
	Driver drivers.Driver
	RootKey string
}

func (router RouterInstance) GetDriver() drivers.Driver {
	return router.Driver
}

func (router RouterInstance) Test(key string) *RouterKeyDescriptorInstance {
	return nil
}
