package main

import (
	"fmt"

	connections "github.com/karim-w/go-cket/Handlers/Connections"
	"github.com/karim-w/go-cket/Helper/redishelper"
	server "github.com/karim-w/go-cket/Server"
	"github.com/karim-w/go-cket/helper/memcache"
	"github.com/karim-w/go-cket/utils/logs"
	"go.uber.org/fx"
)

func main() {
	fmt.Println("Starting Server...")
	app := fx.New(
		logs.LogsModule,
		redishelper.RedisModule,
		memcache.FXMemCacheModule,
		connections.ConnectionHandlerModule,
		server.ServerModule,
	)
	defer app.Run()

}
