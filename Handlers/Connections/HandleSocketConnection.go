package connections

import (
	"encoding/json"
	"net"

	models "github.com/karim-w/go-cket/Models"
	"github.com/karim-w/go-cket/helper/memcache"
	"github.com/karim-w/go-cket/helper/redishelper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ConnectionHandler struct {
	logger      *zap.SugaredLogger
	cache       *memcache.Memcache
	redisClient *redishelper.RedisManager
}

func (c *ConnectionHandler) HandleIncomingSocketConnection(key string, connection net.Conn, Token models.UserToken) {
	text, err := json.Marshal(connection)
	if err != nil {
		c.logger.Errorf("Error marshalling connection: %v", err)
	}
	c.logger.Infof("Incoming connection: %s", text)
	c.cache.HandleIncomingSocketConnection(key, connection)
	c.redisClient.LogUserOnRedis(Token)
}
func (c *ConnectionHandler) HandleIncomingSeverSocketConnection(key string, connection net.Conn, Token models.ServerToken) {
	text, err := json.Marshal(connection)
	if err != nil {
		c.logger.Errorf("Error marshalling connection: %v", err)
	}
	c.logger.Infof("Incoming Server connection: %s", text)
	c.cache.HandleIncomingSocketConnection(key, connection)
}
func (c *ConnectionHandler) HandleTerminateSocketConnection(key string) {
	c.cache.HandleTerminateSocketConnection(key)
}
func (c *ConnectionHandler) FetchSocketConnection(key string) net.Conn {
	if res, ok := c.cache.FetchSocketConnection(key); ok {
		return res
	} else {
		return nil
	}

}

func NewConnectionHandler(logger *zap.SugaredLogger, cache *memcache.Memcache, red *redishelper.RedisManager) *ConnectionHandler {
	return &ConnectionHandler{logger, cache, red}
}

var ConnectionHandlerModule = fx.Provide(NewConnectionHandler)
