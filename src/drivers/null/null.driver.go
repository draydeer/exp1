package null

type NullDriverInstance struct {
	keys map[string]interface{}
}

func (driver *NullDriverInstance) GetValue(key string) (interface{}, bool) {
	return nil, false
}

func (driver *NullDriverInstance) HasValue(key string) bool {
	return false
}

func NewNullDriver() NullDriverInstance {
	return NullDriverInstance{}
}
