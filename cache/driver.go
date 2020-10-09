package cache

import (
	"context"
	"encoding"
	"time"
)

//go:generate eden generate enum --type-name=Driver
// api:enum
type Driver uint8

// cache driver type
const (
	DRIVER_UNKNOWN    Driver = iota
	DRIVER__BUILDIN          // buildin
	DRIVER__MEMCACHED        // memcached
	DRIVER__REDIS            // redis
	DRIVER__MONGO            // mongo for object cache
)

type cacheDriver interface {
	Set(ctx context.Context, key string, value encoding.BinaryMarshaler, expire time.Duration) error
	Get(ctx context.Context, key string, value encoding.BinaryUnmarshaler) error
	Del(ctx context.Context, keys ...string) error
}
