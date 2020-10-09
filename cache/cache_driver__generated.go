package cache

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidCacheDriver = errors.New("invalid CacheDriver")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("CacheDriver", map[string]string{
		"MONGO":     "mongo for object cache",
		"REDIS":     "redis",
		"MEMCACHED": "memcached",
		"BUILDIN":   "buildin",
	})
}

func ParseCacheDriverFromString(s string) (CacheDriver, error) {
	switch s {
	case "":
		return CACHE_DRIVER_UNKNOWN, nil
	case "MONGO":
		return CACHE_DRIVER__MONGO, nil
	case "REDIS":
		return CACHE_DRIVER__REDIS, nil
	case "MEMCACHED":
		return CACHE_DRIVER__MEMCACHED, nil
	case "BUILDIN":
		return CACHE_DRIVER__BUILDIN, nil
	}
	return CACHE_DRIVER_UNKNOWN, InvalidCacheDriver
}

func ParseCacheDriverFromLabelString(s string) (CacheDriver, error) {
	switch s {
	case "":
		return CACHE_DRIVER_UNKNOWN, nil
	case "mongo for object cache":
		return CACHE_DRIVER__MONGO, nil
	case "redis":
		return CACHE_DRIVER__REDIS, nil
	case "memcached":
		return CACHE_DRIVER__MEMCACHED, nil
	case "buildin":
		return CACHE_DRIVER__BUILDIN, nil
	}
	return CACHE_DRIVER_UNKNOWN, InvalidCacheDriver
}

func (CacheDriver) EnumType() string {
	return "CacheDriver"
}

func (CacheDriver) Enums() map[int][]string {
	return map[int][]string{
		int(CACHE_DRIVER__MONGO):     {"MONGO", "mongo for object cache"},
		int(CACHE_DRIVER__REDIS):     {"REDIS", "redis"},
		int(CACHE_DRIVER__MEMCACHED): {"MEMCACHED", "memcached"},
		int(CACHE_DRIVER__BUILDIN):   {"BUILDIN", "buildin"},
	}
}

func (v CacheDriver) String() string {
	switch v {
	case CACHE_DRIVER_UNKNOWN:
		return ""
	case CACHE_DRIVER__MONGO:
		return "MONGO"
	case CACHE_DRIVER__REDIS:
		return "REDIS"
	case CACHE_DRIVER__MEMCACHED:
		return "MEMCACHED"
	case CACHE_DRIVER__BUILDIN:
		return "BUILDIN"
	}
	return "UNKNOWN"
}

func (v CacheDriver) Label() string {
	switch v {
	case CACHE_DRIVER_UNKNOWN:
		return ""
	case CACHE_DRIVER__MONGO:
		return "mongo for object cache"
	case CACHE_DRIVER__REDIS:
		return "redis"
	case CACHE_DRIVER__MEMCACHED:
		return "memcached"
	case CACHE_DRIVER__BUILDIN:
		return "buildin"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CacheDriver)(nil)

func (v CacheDriver) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCacheDriver
	}
	return []byte(str), nil
}

func (v *CacheDriver) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCacheDriverFromString(string(bytes.ToUpper(data)))
	return
}
