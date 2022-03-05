package outgoing

import (
	"encoding/json"
	"net"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	models "github.com/karim-w/go-cket/Models"
)

func RelayMessages(conns []net.Conn, payload models.SocketPayload, op ws.OpCode) {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(conns))
	for i, _ := range conns {
		go func() {
			defer waitGroup.Done()
			if text, err := json.Marshal(payload); err != nil {
			} else {
				err = wsutil.WriteServerMessage(conns[i], op, text)
				if err != nil {
					if err.Error() == "EOF" {
						conns[i].Close()
					}
				}
			}
		}()
	}
	waitGroup.Wait()
}
