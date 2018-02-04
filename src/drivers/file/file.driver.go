package file

import "envd/src/drivers"

type FileDriverInstance struct {

}

func NewDriver() drivers.Driver {
	return FileDriverInstance{}
}

func (driver FileDriverInstance) Get(key string, def interface{}) interface{} {
	return def
}
