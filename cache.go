package cachedriver

type Cache interface {
	Reset() error

	Get(key string, a interface{}) error
	Set(key string, a interface{}) error
	Delete(key string) error
	Each(a interface{}, f func(key string) error) error
}

type NestedCache interface {
	Reset() error

	Get(parentKey string, key string, a interface{}) error
	Set(parentKey string, key string, a interface{}) error
	Delete(parentKey string) error
	DeleteNested(parentKey string, key string) error
	Each(a interface{}, f func(parentKey string, key string) error) error
	EachNested(parentKey string, a interface{}, f func(key string) error) error
}


