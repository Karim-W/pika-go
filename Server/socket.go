package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
	connections "github.com/karim-w/go-cket/Handlers/Connections"
	jwtdecoder "github.com/karim-w/go-cket/utils/JWTDecoder"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Test struct {
	str string
}

func Server(logger *zap.SugaredLogger, handler *connections.ConnectionHandler) *Test {
	fmt.Println()
	fmt.Println()
	fmt.Println(`:'######::::'#######::::::::::::'######::'##:::'##:'########:'########:
'##... ##::'##.... ##::::::::::'##... ##: ##::'##:: ##.....::... ##..::
 ##:::..::: ##:::: ##:::::::::: ##:::..:: ##:'##::: ##:::::::::: ##::::
 ##::'####: ##:::: ##:'#######: ##::::::: #####:::: ######:::::: ##::::
 ##::: ##:: ##:::: ##:........: ##::::::: ##. ##::: ##...::::::: ##::::
 ##::: ##:: ##:::: ##:::::::::: ##::: ##: ##:. ##:: ##:::::::::: ##::::
. ######:::. #######:::::::::::. ######:: ##::. ##: ########:::: ##::::
:......:::::.......:::::::::::::......:::..::::..::........:::::..:::::`)
	logger.Info("Server Started ")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		fmt.Println(conn)
		if err != nil {
			// handle error
		}
		userToken, _ := jwtdecoder.Decode(r.Header.Get("Jwttoken"))
		handler.HandleIncomingSocketConnection(uuid.NewString(), conn)
		fmt.Println(userToken)

		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					// handle error
				}
			}
		}()
	}))
	return &Test{str: "test"}
}
func registerHooks(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, t *Test) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Initializing server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Terminating server")
			logger.Sync()
			return nil
		},
	})
}

var ServerModule = fx.Options(fx.Provide(Server), fx.Invoke(registerHooks))
