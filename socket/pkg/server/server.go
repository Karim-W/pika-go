package server

import (
	"context"
	"net/http"

	"github.com/karim-w/pika-go/socket/internal/middlewares"
	"github.com/karim-w/pika-go/socket/pkg/connections"
	"github.com/karim-w/pika-go/socket/pkg/services/auth"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Test struct {
	str string
}

func Server(logger *zap.SugaredLogger,c connections.ConnectionManager,a auth.Auth) *Test {
	logger.Info("Server Started ")
	
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	go func() {
		con,u := middlewares.AuthenticateConnection(&w,r,a)
		if u != nil {
			c.HandleJoin()
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
