package rdcache

import (
	"errors"
	"time"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

type CacheResolver interface {
	Resolve(k interface{}, v interface{}) error
}

type RedisCacheScheme interface {
	New() interface{}
	GetRedisKey(k interface{}) string
	GetDefaultExpiration() time.Duration
}

type RedisCache struct {
	codec *cache.Codec
}

var redisCache *RedisCache

func InitCache(client *redis.Client) {
	redisCache = NewCache(client, MsgpackEncoder, MsgpackDecoder)
}

func MsgpackEncoder(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func MsgpackDecoder(b []byte, v interface{}) error {
	return msgpack.Unmarshal(b, v)
}

//export
func NewCache(client *redis.Client, encoder func(v interface{}) ([]byte, error), decoder func(b []byte, v interface{}) error) *RedisCache {
	return &RedisCache{&cache.Codec{
		Redis:     client,
		Marshal:   encoder,
		Unmarshal: decoder,
	}}
}

func (rc *RedisCache) Set(key string, val interface{}, expirations ...time.Duration) error {
	item := &cache.Item{
		Key:    key,
		Object: val,
	}
	if len(expirations) > 0 {
		item.Expiration = expirations[0]
	}
	return rc.codec.Set(item)
}

func (rc *RedisCache) Get(key string, val interface{}) error {
	return rc.codec.Get(key, val)
}

func (rc *RedisCache) GetCache(msc RedisCacheScheme, k interface{}) (v interface{}, err error) {
	v = msc.New()
	err = rc.GetSchemeCache(msc, k, v)
	if err != nil {
		return nil, err
	}
	return
}

func (rc *RedisCache) GetSchemeCache(msc RedisCacheScheme, k interface{}, v interface{}) (err error) {
	redisKey := msc.GetRedisKey(k)
	err = rc.Get(redisKey, v)
	if err == nil {
		return
	}
	if resolver, ok := msc.(CacheResolver); ok {
		if err := resolver.Resolve(k, v); err == nil {
			err = rc.Set(msc.GetRedisKey(k), v, msc.GetDefaultExpiration())
			return nil
		}
	}
	return
}

func (rc *RedisCache) SetSchemeCache(msc RedisCacheScheme, k interface{}, v interface{}) (err error) {
	if v == nil {
		if resolver, ok := msc.(CacheResolver); ok {
			nv := msc.New()
			if err := resolver.Resolve(k, nv); err == nil {
				return rc.Set(msc.GetRedisKey(k), nv, msc.GetDefaultExpiration())
			}
		} else {
			return errors.New("resolve obj failed")
		}
	}

	return rc.Set(msc.GetRedisKey(k), v, msc.GetDefaultExpiration())
}

/*func (rc *RedisCache) GetString(key string) string {
	var strVal = ""
	if err := rc.codec.Get(key, &strVal); err != nil {
		return ""
	}
	return strVal
}*/
