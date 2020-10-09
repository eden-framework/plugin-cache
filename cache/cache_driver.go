package cache

import (
	"context"
	"encoding"
	"time"
)

//go:generate eden generate enum --type-name=CacheDriver
// api:enum
type CacheDriver uint8

// cache driver type
const (
	CACHE_DRIVER_UNKNOWN    CacheDriver = iota
	CACHE_DRIVER__BUILDIN               // buildin
	CACHE_DRIVER__MEMCACHED             // memcached
	CACHE_DRIVER__REDIS                 // redis
	CACHE_DRIVER__MONGO                 // mongo for object cache
)

type cacheDriver interface {
	Set(ctx context.Context, key string, value encoding.BinaryMarshaler, expire time.Duration) error
	Get(ctx context.Context, key string, value encoding.BinaryUnmarshaler) error
	Del(ctx context.Context, keys ...string) error
}
