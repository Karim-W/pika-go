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

func Server(logger *zap.SugaredLogger, handler *connections.ConnectionHandler, mayfair *mayfair.Mayfair, decoder *jwtdecoder.DecoderService) *Test {
	logger.Info("Server Started ")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		sender := r.Header.Get("Sender")
		if err != nil {
			// handle error
			logger.Error(err)
		}
		if sender == "" {
			if userToken, err := jwtdecoder.Decode(r.Header.Get("AuthToken")); err != nil {
				logger.Error(err)
			} else {
				handler.HandleIncomingSocketConnection(userToken.UserID, conn, userToken)
				logger.Info("New Client connection	", userToken.UserID)
				go func() {
					defer logger.Info("Client connection closed for user with ID: ", userToken.UserID)
					mayfair.ManageUserConnections(conn, userToken)
				}()
			}
		} else {
			if serverToken, err := jwtdecoder.DecodeServer(r.Header.Get("AuthToken")); err != nil {
				logger.Error(err)
			} else {
				handler.HandleIncomingSeverSocketConnection(serverToken.ClientID, conn, serverToken)
				logger.Info("New Client connection	", serverToken.ClientID)
				go func() {
					defer logger.Info("Socket connection closed	", serverToken.ClientID)
					mayfair.Navigate(conn, serverToken)
				}()
			}
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
