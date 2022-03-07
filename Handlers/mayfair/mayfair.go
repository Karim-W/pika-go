package mayfair

import (
	"encoding/json"
	"net"
	"os"

	"github.com/gobwas/ws/wsutil"
	models "github.com/karim-w/go-cket/Models"
	"github.com/karim-w/go-cket/handlers/outgoing"
	"github.com/karim-w/go-cket/helper/memcache"
	"github.com/karim-w/go-cket/utils/hermes"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Mayfair struct {
	logger    *zap.SugaredLogger
	cache     *memcache.Memcache
	client    *hermes.HttpClient
	brokerUrl string
}

func (m *Mayfair) ManageUserConnections(conn net.Conn, token models.UserToken) {
	update := models.PresenceUpdate{
		AddedUsers:   []string{token.UserID},
		RemovedUsers: []string{},
		SessionId:    token.SessionID,
	}
	m.notifyBroker(update)
	keepAlive := true
	for keepAlive {
		msg, _, err := wsutil.ReadClientData(conn)
		if err != nil {
			if err.Error() == "EOF" {
				m.logger.Info("Closing connection for user: ", token.UserID)
				keepAlive = false
				m.closeSocketConnection(conn, token)
			}
		} else {
			m.logger.Info("Incoming message:	", string(msg))
		}
	}
}

func (m *Mayfair) Navigate(conn net.Conn, token models.ServerToken) {
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

func (m *Mayfair) closeSocketConnection(conn net.Conn, token models.UserToken) {
	update := models.PresenceUpdate{
		AddedUsers:   []string{},
		RemovedUsers: []string{token.UserID},
		SessionId:    token.SessionID,
	}
	m.notifyBroker(update)
	conn.Close()
	m.cache.HandleTerminateSocketConnection(token.UserID)
}

func (m *Mayfair) notifyBroker(update models.PresenceUpdate) {
	m.logger.Info("Notifying broker of presence update for session: ", update.SessionId, " users to be added: ", update.AddedUsers, " users to be removed: ", update.RemovedUsers)
}

func NewMayfair(logger *zap.SugaredLogger, cache *memcache.Memcache, h *hermes.HttpClient) *Mayfair {
	brokerUrl := os.Getenv("BROKER_URL")
	return &Mayfair{logger, cache, h, brokerUrl}
}

var MayfairModule = fx.Provide(NewMayfair)
