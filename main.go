package main

import (
	"fmt"

	"github.com/joho/godotenv"
	server "github.com/karim-w/go-cket/Server"
	connections "github.com/karim-w/go-cket/handlers/connections"
	"github.com/karim-w/go-cket/handlers/mayfair"
	"github.com/karim-w/go-cket/helper/memcache"
	"github.com/karim-w/go-cket/helper/redishelper"
	jwtdecoder "github.com/karim-w/go-cket/utils/JWTDecoder"
	"github.com/karim-w/go-cket/utils/hermes"
	"github.com/karim-w/go-cket/utils/logs"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load(".env")
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
		hermes.Module,
		redishelper.RedisModule,
		memcache.FXMemCacheModule,
		connections.ConnectionHandlerModule,
		mayfair.MayfairModule,
		jwtdecoder.DecoderModule,
		server.ServerModule,
	)
	defer app.Run()

}
