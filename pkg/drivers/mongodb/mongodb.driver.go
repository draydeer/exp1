package mongodb

import "envd/pkg/drivers"

type MongodbDriverInstance struct {

}

func NewDriver() drivers.Driver {
	return MongodbDriverInstance{}
}

func (driver MongodbDriverInstance) Get(key string, def interface{}) interface{} {
	return def
}
