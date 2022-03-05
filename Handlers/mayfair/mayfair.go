package mayfair

import (
	"encoding/json"
	"net"

	"github.com/gobwas/ws/wsutil"
	models "github.com/karim-w/go-cket/Models"
	"github.com/karim-w/go-cket/handlers/outgoing"
	"github.com/karim-w/go-cket/helper/memcache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Mayfair struct {
	logger *zap.SugaredLogger
	cache  *memcache.Memcache
}

func (m *Mayfair) Navigate(conn net.Conn, token models.UserToken) {
	keepAlive := true
	for keepAlive {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			if err.Error() == "EOF" {
				keepAlive = false
				conn.Close()
			}
		} else {
			adminCommand := models.SocketPayload{}
			json.Unmarshal(msg, &adminCommand)
			m.logger.Info("Incoming message:	", adminCommand)
			if len(adminCommand.Audience) > 0 {
				conList := m.cache.FetchSocketConnections(adminCommand.Audience)
				outgoing.RelayMessages(conList, adminCommand, op)
			}
		}
	}
}
func NewMayfair(logger *zap.SugaredLogger, cache *memcache.Memcache) *Mayfair {
	return &Mayfair{logger, cache}
}

var MayfairModule = fx.Provide(NewMayfair)
