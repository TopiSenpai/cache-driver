package cachedriver

import (
	"errors"
	"reflect"
)

var ErrNotFound = errors.New("not found")

func decode(v interface{}, a interface{}) error {
	reflect.ValueOf(a).Elem().Set(reflect.ValueOf(v))
	return nil
}

func NewMemoryCache() Cache {
	return &memoryCache{
		cache: make(map[string]interface{}),
	}
}

type memoryCache struct {
	cache map[string]interface{}
}

func (c *memoryCache) Reset() error {
	c.cache = make(map[string]interface{})
	return nil
}

func (c *memoryCache) Get(key string, a interface{}) error {
	if v, ok := c.cache[key]; ok {
		return decode(v, a)
	}
	return ErrNotFound
}

func (c *memoryCache) Set(key string, a interface{}) error {
	c.cache[key] = a
	return nil
}

func (c *memoryCache) Delete(key string) error {
	delete(c.cache, key)
	return nil
}

func (c *memoryCache) Each(a interface{}, f func(key string) error) error {
	for key, v := range c.cache {
		if err := decode(v, a); err != nil {
			return err
		}
		if err := f(key); err != nil {
			return err
		}
	}
	return nil
}

func NewMemoryNestedCache() NestedCache {
	return &nestedMemoryCache{
		cache: make(map[string]map[string]interface{}),
	}
}

type nestedMemoryCache struct {
	cache map[string]map[string]interface{}
}

func (c *nestedMemoryCache) Reset() error {
	c.cache = make(map[string]map[string]interface{})
	return nil
}

func (c *nestedMemoryCache) Get(parentKey string, key string, a interface{}) error {
	c2, ok := c.cache[parentKey]
	if !ok {
		return ErrNotFound
	}
	if v, ok := c2[key]; ok {
		return decode(v, a)
	}
	return ErrNotFound
}

func (c *nestedMemoryCache) Set(parentKey string, key string, a interface{}) error {
	c2, ok := c.cache[parentKey]
	if !ok {
		c2 = make(map[string]interface{})
		c.cache[parentKey] = c2
	}
	c2[key] = a
	return nil
}

func (c *nestedMemoryCache) Delete(parentKey string) error {
	delete(c.cache, parentKey)
	return nil
}

func (c *nestedMemoryCache) DeleteNested(parentKey string, key string) error {
	c2, ok := c.cache[parentKey]
	if !ok {
		return nil
	}
	delete(c2, key)
	return nil
}

func (c *nestedMemoryCache) Each(a interface{}, f func(parentKey string, key string) error) error {
	for parentKey, c2 := range c.cache {
		for key, v := range c2 {
			if err := decode(v, a); err != nil {
				return err
			}
			if err := f(parentKey, key); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *nestedMemoryCache) EachNested(parentKey string, a interface{}, f func(key string) error) error {
	c2, ok := c.cache[parentKey]
	if !ok {
		return nil
	}
	for key, v := range c2 {
		if err := decode(v, a); err != nil {
			return err
		}
		if err := f(key); err != nil {
			return err
		}
	}
	return nil
}
