package regexp

import (
	"envd/pkg/routers"
)

type RegexpRouterInstance struct {
	routers.RouterInstance
	Regexp string
}

func (router RegexpRouterInstance) Test(key string) *routers.RouterKeyDescriptorInstance {
	return nil
}
