package memcache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Memcache struct {
	logger *zap.SugaredLogger
	cache  *cache.Cache
}

func (m *Memcache) HandleIncomingSocketConnection(key string, value interface{}) {
	m.cache.Set(key, value, cache.NoExpiration)
}
func (m *Memcache) HandleTerminateSocketConnection(key string) {
	m.cache.Delete(key)
}
func (m *Memcache) FetchSocketConnection(key string) (interface{}, bool) {
	return m.cache.Get(key)
}

func NewMemcache(l *zap.SugaredLogger) *Memcache {
	c := cache.New(24*time.Hour, 1*time.Minute)
	return &Memcache{l, c}
}

var FXMemCacheModule = fx.Option(fx.Provide(NewMemcache))
