package cache

import (
	"context"
	"encoding"
	"errors"
	"github.com/cornelk/hashmap"
	"time"
)

var defaultMemoryCache *memoryCache

type memoryCache struct {
	*hashmap.HashMap
}

func newMemoryCache() *memoryCache {
	if defaultMemoryCache == nil {
		defaultMemoryCache = &memoryCache{
			hashmap.New(100),
		}
	}

	return defaultMemoryCache
}

func (m *memoryCache) Set(ctx context.Context, key string, value encoding.BinaryMarshaler, expire time.Duration) error {
	data, err := value.MarshalBinary()
	if err != nil {
		return err
	}
	m.HashMap.Set(key, data)
	return nil
}

func (m *memoryCache) Get(ctx context.Context, key string, value encoding.BinaryUnmarshaler) error {
	val, exist := m.HashMap.Get(key)
	if !exist {
		return errors.New("not found")
	}

	return value.UnmarshalBinary(val.([]byte))
}

func (m *memoryCache) Del(ctx context.Context, keys ...string) error {
	for _, k := range keys {
		m.HashMap.Del(k)
	}
	return nil
}
