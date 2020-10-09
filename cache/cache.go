package cache

import (
	"github.com/eden-framework/plugin-redis/redis"
	"github.com/profzone/envconfig"
)

type Cache struct {
	Driver CacheDriver
	// Host for driver except buildin
	Host string
	// Port for driver except buildin
	Port int
	// User for driver except buildin
	User string
	// Password for driver except buildin
	Password envconfig.Password
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
		}
		driver.Init()
		c.cacheDriver = driver
	case CACHE_DRIVER__MEMCACHED:
	case CACHE_DRIVER__MONGO:
	default:
		panic("[Cache] unsupported driver")
	}
}
