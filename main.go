package main

import (
	"fmt"

	"github.com/joho/godotenv"
	server "github.com/karim-w/pikachu/Server"
	connections "github.com/karim-w/pikachu/handlers/connections"
	"github.com/karim-w/pikachu/handlers/mayfair"
	"github.com/karim-w/pikachu/helper/memcache"
	"github.com/karim-w/pikachu/helper/redishelper"
	jwtdecoder "github.com/karim-w/pikachu/utils/JWTDecoder"
	"github.com/karim-w/pikachu/utils/hermes"
	"github.com/karim-w/pikachu/utils/logs"
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
