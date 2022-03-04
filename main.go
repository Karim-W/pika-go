package main

import (
	"fmt"

	"github.com/karim-w/go-cket/Helper/redishelper"
	server "github.com/karim-w/go-cket/Server"
	"github.com/karim-w/go-cket/utils/logs"
	"go.uber.org/fx"
)

func main() {
	fmt.Println("Starting Server...\n\n\n")
	app := fx.New(
		logs.LogsModule,
		redishelper.RedisModule,
		server.ServerModule,
	)
	defer app.Run()

}
