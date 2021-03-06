package main

import (
	"fmt"

	"github.com/joho/godotenv"
	server "github.com/karim-w/pika-go/Server"
	connections "github.com/karim-w/pika-go/handlers/connections"
	"github.com/karim-w/pika-go/handlers/mayfair"
	"github.com/karim-w/pika-go/helper/memcache"
	"github.com/karim-w/pika-go/helper/redishelper"
	jwtdecoder "github.com/karim-w/pika-go/utils/JWTDecoder"
	"github.com/karim-w/pika-go/utils/hermes"
	"github.com/karim-w/pika-go/utils/logs"
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
