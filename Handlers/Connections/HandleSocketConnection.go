package connections

import (
	"github.com/karim-w/go-cket/helper/memcache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ConnectionHandler struct {
	logger *zap.SugaredLogger
	cache  *memcache.Memcache
}

func (c *ConnectionHandler) HandleIncomingSocketConnection(key string, value interface{}) {
	c.cache.HandleIncomingSocketConnection(key, value)
}
func (c *ConnectionHandler) HandleTerminateSocketConnection(key string) {
	c.cache.HandleTerminateSocketConnection(key)
}
func (c *ConnectionHandler) FetchSocketConnection(key string) *interface{} {
	if res, ok := c.cache.FetchSocketConnection(key); ok {
		return &res
	} else {
		return nil
	}

}

func NewConnectionHandler(logger *zap.SugaredLogger, cache *memcache.Memcache) *ConnectionHandler {
	return &ConnectionHandler{logger, cache}
}

var ConnectionHandlerModule = fx.Provide(NewConnectionHandler)
