package mayfair

import (
	"encoding/json"
	"net"

	"github.com/gobwas/ws/wsutil"
	models "github.com/karim-w/go-cket/Models"
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
			if err = json.Unmarshal(msg, &adminCommand); err != nil {
				m.logger.Info("Incoming message:	", adminCommand)
			}

		}
		err = wsutil.WriteServerMessage(conn, op, msg)
		if err != nil {
			if err.Error() == "EOF" {
				keepAlive = false
				conn.Close()
			}
		}
	}
}
func NewMayfair(logger *zap.SugaredLogger, cache *memcache.Memcache) *Mayfair {
	return &Mayfair{logger, cache}
}

var MayfairModule = fx.Provide(NewMayfair)
