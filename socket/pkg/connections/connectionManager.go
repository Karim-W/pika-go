package connections
 
import (
	"net"
	"sync"

	"github.com/gobwas/ws/wsutil"
	"github.com/karim-w/pika-go/socket/internal/common/models"
	"github.com/karim-w/pika-go/socket/pkg/helpers"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ConnectionManager interface {
	HandleJoin(token string, conn *net.Conn)
	ManageUserConnection(id string, con *net.Conn)
	terminateSocketConnection(id string)
	SendMessages(ids []string, msg string)
}

type connectionManagerImpl struct {
	logger   *zap.SugaredLogger
	memCache helpers.MemCache 
	redis    helpers.RedisCache
}

var _ ConnectionManager = (*connectionManagerImpl)(nil)

func ConnectionManagerProvider(logger *zap.SugaredLogger, memCache helpers.MemCache, redis helpers.RedisCache) ConnectionManager {
	return &connectionManagerImpl{ 
		logger:   logger,
		memCache: memCache,
		redis:    redis,
	}
}
  
var ConnectionManagerModule = fx.Option(fx.Provide(ConnectionManagerProvider))          

func (c *connectionManagerImpl) HandleJoin(token string, con *net.Conn) {
	u := models.RedisUser{}
	go c.memCache.AddConnection(u.Id, *con)      
	go c.redis.HandleUserJoin(&u)
}   

func (c *connectionManagerImpl) ManageUserConnection(id string, conn *net.Conn) { 
	keepAlive := true
	for keepAlive {
		msg, _, err := wsutil.ReadClientData(*conn)
		if    err != nil {
			if err.Error() == "EOF" { 
				c.logger.Info("Closing connection for user: ", id)
				c.terminateSocketConnection(id)
				keepAlive = false
			}
		} else {
			c.logger.Info("Incoming message:	", string(msg))
		}
	}
}

func (c *connectionManagerImpl) terminateSocketConnection(id string) {
	go c.memCache.RemoveConnection(id)
	go c.redis.HandleUserExit(id)
}

func (c *connectionManagerImpl) SendMessages(ids []string, msg string) {
	messageBytes := []byte(msg)
	connections := c.memCache.FetchConnectionsAsync(ids)
	wg := sync.WaitGroup{}
	wg.Add(len(*connections))
	for i := range *connections {
		connection := *(*connections)[i]
		go func(connection net.Conn, i int) {
			defer wg.Done()
			if err := wsutil.WriteServerMessage(connection, 1, messageBytes); err != nil {
				c.logger.Info("Failed to Send Message to ", ids[i])
			}
		}(connection, i)
	}
	wg.Wait()
}
