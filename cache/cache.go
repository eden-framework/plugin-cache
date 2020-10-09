package cache

import (
	"github.com/eden-framework/plugin-redis/redis"
	"github.com/profzone/envconfig"
)

type Cache struct {
	Driver Driver
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
	if c.Driver == DRIVER_UNKNOWN {
		c.Driver = DRIVER__BUILDIN
	}
	if c.Driver == DRIVER__REDIS || c.Driver == DRIVER__MEMCACHED || c.Driver == DRIVER__MONGO {
		if c.Host == "" {
			panic("[Cache] must specify Host and Port when use REDIS or MEMCACHED or MONGO drivers")
		}
	}
}

func (c *Cache) Init() {
	switch c.Driver {
	case DRIVER__BUILDIN:
		c.cacheDriver = newMemoryCache()
	case DRIVER__REDIS:
		driver := &redis.Redis{
			Host:     c.Host,
			Port:     c.Port,
			User:     c.User,
			Password: c.Password,
		}
		driver.Init()
		c.cacheDriver = driver
	case DRIVER__MEMCACHED:
	case DRIVER__MONGO:
	default:
		panic("[Cache] unsupported driver")
	}
}
