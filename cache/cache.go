package cache

import (
	"context"
	"encoding"
	"github.com/eden-framework/plugin-redis/redis"
	"github.com/profzone/envconfig"
	"time"
)

type Cache struct {
	Driver CacheDriver
	// Prefix is the global prefix string of key
	Prefix string
	// Host for driver except buildin
	Host string
	// Port for driver except buildin
	Port int
	// User for driver except buildin
	User string
	// Password for driver except buildin
	Password envconfig.Password
	// DB fro driver redis
	DB int
	cacheDriver
}

func (c *Cache) SetDefaults() {
	if c.Driver == CACHE_DRIVER_UNKNOWN {
		c.Driver = CACHE_DRIVER__BUILDIN
	}
	if c.Driver == CACHE_DRIVER__REDIS || c.Driver == CACHE_DRIVER__MEMCACHED || c.Driver == CACHE_DRIVER__MONGO {
		if c.Host == "" {
			panic("[Cache] must specify Host and Port when use REDIS or MEMCACHED or MONGO drivers")
		}
	}
}

func (c *Cache) Init() {
	switch c.Driver {
	case CACHE_DRIVER__BUILDIN:
		c.cacheDriver = newMemoryCache()
	case CACHE_DRIVER__REDIS:
		driver := &redis.Redis{
			Host:     c.Host,
			Port:     c.Port,
			User:     c.User,
			Password: c.Password,
			DB:       c.DB,
		}
		driver.Init()
		c.cacheDriver = driver
	case CACHE_DRIVER__MEMCACHED:
	case CACHE_DRIVER__MONGO:
	default:
		panic("[Cache] unsupported driver")
	}
}

func (c *Cache) Set(ctx context.Context, key string, value encoding.BinaryMarshaler, expire time.Duration) error {
	return c.cacheDriver.Set(ctx, c.Prefix+key, value, expire)
}

func (c *Cache) Get(ctx context.Context, key string, value encoding.BinaryUnmarshaler) error {
	return c.cacheDriver.Get(ctx, c.Prefix+key, value)
}

func (c *Cache) Del(ctx context.Context, keys ...string) error {
	prefixKeys := make([]string, 0)
	for _, key := range keys {
		prefixKeys = append(prefixKeys, c.Prefix+key)
	}
	return c.Del(ctx, prefixKeys...)
}
