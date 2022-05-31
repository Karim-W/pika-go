package middlewares

import (
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/karim-w/pika-go/socket/internal/common/models"
	"github.com/karim-w/pika-go/socket/pkg/services/auth"
)


func AuthenticateConnection(w *http.ResponseWriter,r *http.Request,authService auth.Auth) (*net.Conn,*models.RedisUser){
  token := r.Header.Get("Authorization")
  if user,err := authService.ParseAndVaidate(token);err!=nil{
    return nil,nil
  }else{
    if conn, _, _, err := ws.UpgradeHTTP(r, *w); err!= nil {
      return nil,nil
    }else{
      return &conn,user
    }
    
  }         
}
