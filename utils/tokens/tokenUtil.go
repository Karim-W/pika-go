package tokens

// import (
// 	"fmt"

// 	"github.com/golang-jwt/jwt"
// 	"go.uber.org/fx"
// 	"go.uber.org/zap"
// 	"techunicorn.com/UDC/Middleware/utils/http"
// )

// type TokenUtil struct {
// 	publicKey []byte
// }

// func (t TokenUtil) ParseEUJWTtoken(tokenString string) (jwt.Token, error) {
// 	// Read the private key which will be used to parse the JWT
// 	signed, _ := jwt.ParseRSAPublicKeyFromPEM(t.publicKey)
// 	parsedToken, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
// 		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
// 		}
// 		return signed, nil
// 	})

// 	return *parsedToken, err
// }

// func (t TokenUtil) ParseADJWTtoken(tokenString string) (jwt.MapClaims, error) {
// 	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
// 	if err != nil {
// 		return jwt.MapClaims{}, err
// 	}
// 	claims, _ := token.Claims.(jwt.MapClaims)

// 	return claims, nil
// }

// var authEndpoint = "https://golem-difc-cxstg.azurewebsites.net"

// func NewTokenUtil(client *http.HttpClient, logger *zap.SugaredLogger) *TokenUtil {
// 	var resource = "/api/v1/.well-known/key"
// 	var keys map[string]interface{}
// 	statusCode, ok, err := client.Get(authEndpoint+resource, nil).Result(&keys)

// 	if !ok {
// 		logger.Errorf("Error while getting public key from auth endpoint: %s", err)
// 		panic(err)
// 	} else {
// 		if statusCode >= 200 && statusCode < 300 {
// 			logger.Infof("Public key retrieved from auth endpoint: %s", keys["PublicKey"].(string))
// 			return &TokenUtil{[]byte(keys["PublicKey"].(string))}
// 		} else {
// 			panic(err)
// 		}
// 	}
// }

// var Module = fx.Option(fx.Provide(NewTokenUtil))
