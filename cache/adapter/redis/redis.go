
package redis

import (
	"time"
	redisCache "github.com/go-redis/cache"
	"github.com/go-redis/redis"
	cache "pnginx/cache"
	"github.com/vmihailenco/msgpack"
)

// Adapter is the memory adapter data structure.
type Adapter struct {
	store *redisCache.Codec
}

// RingOptions exports go-redis RingOptions type.
type RingOptions redis.RingOptions

// Get implements the cache Adapter interface Get method.
func (a *Adapter) Get(key uint64) ([]byte, bool) {
	var c []byte
	if err := a.store.Get(cache.KeyAsString(key), &c); err == nil {
		return c, true
	}

	return nil, false
}

// Set implements the cache Adapter interface Set method.
func (a *Adapter) Set(key uint64, response []byte, expiration time.Time) {
	a.store.Set(&redisCache.Item{
		Key:        cache.KeyAsString(key),
		Object:     response,
		Expiration: expiration.Sub(time.Now()),
	})
}

// Release implements the cache Adapter interface Release method.
func (a *Adapter) Release(key uint64) {
	a.store.Delete(cache.KeyAsString(key))
}

// NewAdapter initializes Redis adapter.
func NewAdapter(opt *RingOptions) cache.Adapter {
	ropt := redis.RingOptions(*opt)
	return &Adapter{
		&redisCache.Codec{
			Redis: redis.NewRing(&ropt),
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)

			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		},
	}
}
