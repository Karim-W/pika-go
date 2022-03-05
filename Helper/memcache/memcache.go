package memcache

import (
	"net"
	"time"

	"github.com/patrickmn/go-cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Memcache struct {
	logger *zap.SugaredLogger
	cache  *cache.Cache
}

func (m *Memcache) HandleIncomingSocketConnection(key string, value net.Conn) {
	m.logger.Infof("Incoming connection: %s", key)
	m.cache.Set(key, value, cache.NoExpiration)
}
func (m *Memcache) HandleTerminateSocketConnection(key string) {
	m.cache.Delete(key)
}
func (m *Memcache) FetchSocketConnection(key string) (net.Conn, bool) {
	ret, ok := m.cache.Get(key)
	if ok {
		result := ret.(net.Conn)
		return result, ok
	}
	return nil, false
}
func (m *Memcache) FetchSocketConnections(keys []string) []net.Conn {
	var connections []net.Conn
	for _, key := range keys {
		if conn, ok := m.FetchSocketConnection(key); ok {
			connections = append(connections, conn)
		}
	}
	return connections
}

func NewMemcache(l *zap.SugaredLogger) *Memcache {
	c := cache.New(24*time.Hour, 1*time.Minute)
	return &Memcache{l, c}
}

var FXMemCacheModule = fx.Option(fx.Provide(NewMemcache))
