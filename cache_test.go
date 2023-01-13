package cachedriver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryCache(t *testing.T) {
	cache := NewMemoryCache()

	err := cache.Set("foo", "bar")
	assert.NoError(t, err)

	var v string
	err = cache.Get("foo", &v)
	assert.NoError(t, err)
	assert.Equal(t, "bar", v)
}
