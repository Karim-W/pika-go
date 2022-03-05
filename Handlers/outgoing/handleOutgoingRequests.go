package outgoing

import (
	"net"

	"github.com/gobwas/ws/wsutil"
	models "github.com/karim-w/go-cket/Models"
)

func RelayMessages(conn net.Conn, token models.UserToken) {
	keepAlive := true
	for keepAlive {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			if err.Error() == "EOF" {
				keepAlive = false
				conn.Close()
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
