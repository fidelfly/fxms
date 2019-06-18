package rdcache

import (
	"time"

	"github.com/fidelfly/fxgo/logx"
)

func Set(key string, val interface{}, expirations ...time.Duration) {
	if redisCache == nil {
		panic("redis cache is not initialized")
	}
	logx.CaptureError(redisCache.Set(key, val, expirations...))
}

func Get(key string, val interface{}) bool {
	if redisCache == nil {
		panic("redis cache is not initialized")
	}
	if err := redisCache.Get(key, val); err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func SetSchemeCache(rcs RedisCacheScheme, k interface{}, v interface{}) {
	if redisCache == nil {
		panic("redis cache is not initialized")
	}

	logx.CaptureError(redisCache.SetSchemeCache(rcs, k, v))
}

func GetSchemeCache(rcs RedisCacheScheme, k interface{}, v interface{}) bool {
	if redisCache == nil {
		panic("redis cache is not initialized")
	}
	if err := redisCache.GetSchemeCache(rcs, k, v); err != nil {
		logx.Error(err)
		return false
	}
	return true
}

func GetCache(rcs RedisCacheScheme, k interface{}) (v interface{}, err error) {
	if redisCache == nil {
		panic("redis cache is not initialized")
	}

	return redisCache.GetCache(rcs, k)
}
