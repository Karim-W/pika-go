package server

import (
	"context"
	"net/http"

	"github.com/gobwas/ws"
	connections "github.com/karim-w/go-cket/handlers/connections"
	"github.com/karim-w/go-cket/handlers/mayfair"
	jwtdecoder "github.com/karim-w/go-cket/utils/JWTDecoder"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Test struct {
	str string
}

func Server(logger *zap.SugaredLogger, handler *connections.ConnectionHandler, mayfair *mayfair.Mayfair) *Test {

	logger.Info("Server Started ")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		if userToken, err := jwtdecoder.Decode(r.Header.Get("AuthToken")); err != nil {
			logger.Error(err)
		} else {
			handler.HandleIncomingSocketConnection(userToken.UserID, conn, userToken)
			logger.Info("New Client connection	", userToken.UserID)
			go func() {
				defer logger.Info("Socket connection closed	", userToken.UserID)
				mayfair.Navigate(conn, userToken)
			}()
		}
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
