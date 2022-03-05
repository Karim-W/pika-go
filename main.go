package main

import (
	"fmt"

	server "github.com/karim-w/go-cket/Server"
	connections "github.com/karim-w/go-cket/handlers/connections"
	"github.com/karim-w/go-cket/handlers/mayfair"
	"github.com/karim-w/go-cket/helper/memcache"
	"github.com/karim-w/go-cket/helper/redishelper"
	"github.com/karim-w/go-cket/utils/logs"
	"go.uber.org/fx"
)

func main() {
	fmt.Println(`:'######::::'#######::::::::::::'######::'##:::'##:'########:'########:
'##... ##::'##.... ##::::::::::'##... ##: ##::'##:: ##.....::... ##..::
 ##:::..::: ##:::: ##:::::::::: ##:::..:: ##:'##::: ##:::::::::: ##::::
 ##::'####: ##:::: ##:'#######: ##::::::: #####:::: ######:::::: ##::::
 ##::: ##:: ##:::: ##:........: ##::::::: ##. ##::: ##...::::::: ##::::
 ##::: ##:: ##:::: ##:::::::::: ##::: ##: ##:. ##:: ##:::::::::: ##::::
. ######:::. #######:::::::::::. ######:: ##::. ##: ########:::: ##::::
:......:::::.......:::::::::::::......:::..::::..::........:::::..:::::`)
	fmt.Println("Starting Server...")
	app := fx.New(
		logs.LogsModule,
		redishelper.RedisModule,
		memcache.FXMemCacheModule,
		connections.ConnectionHandlerModule,
		mayfair.MayfairModule,
		server.ServerModule,
	)
	defer app.Run()

}
