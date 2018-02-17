package consul

import (
	"strings"
	"github.com/hashicorp/consul/api"

	"envd/src/drivers"
	"envd/src/lib"
)

var split = lib.SplitKeyFn(".")

type ConsulDriverInstance struct {
	drivers.DriverInstance
	Client *api.KV
}

func (driver *ConsulDriverInstance) GetKey(key string) (interface{}, bool) {
	pairs, _, err := driver.Client.List(key, nil)

	if err != nil {
		return nil, false
	}

	if len(pairs) == 1 {
		return string(pairs[0].Value), true
	}

	dict := make(map[string]interface{})

	for _, p := range pairs {
		curr := &dict
		keys := strings.Split(p.Key, "/")

		for i, key := range keys {
			if i == len(keys) - 1 {
				(*curr)[key] = string(p.Value)

				continue
			}

			if _, isPresent := (*curr)[key]; ! isPresent {
				(*curr)[key] = make(map[string]interface{})
			}

			temp := (*curr)[key].(map[string]interface{})

			curr = &temp
		}
	}

	return dict, true
}

func (driver *ConsulDriverInstance) GetKeyDescriptorFromUniversal(key string) drivers.DriverKeyDescriptorInstance {
	splitted := split(key)

	return drivers.DriverKeyDescriptorInstance{
		splitted,
		splitted[0],
	}
}

func (driver *ConsulDriverInstance) HasKey(key string) bool {
	return false
}

func NewConsulDriver() *ConsulDriverInstance {
	conf := api.DefaultConfig()

	client, err := api.NewClient(conf)

	if err != nil {
		return nil
	}

	return &ConsulDriverInstance{Client: client.KV()}
}
