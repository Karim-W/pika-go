package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/karim-w/hermes"
	"github.com/karim-w/pika-go/socket/internal/common/models"
	"go.uber.org/zap"
)

type Auth interface {
	ParseAndVaidate(token string) (*models.RedisUser, error)
}

type authImpl struct {
	logger            *zap.SugaredLogger
	publicKey         []byte
	tokenPublicKeyUrl string
	client            *hermes.HttpClient
	kid               string
}

var _ Auth = (*authImpl)(nil)

func AuthProvider(logger *zap.SugaredLogger) Auth {
	tokenUrl := os.Getenv("TOKEN_PUBLIC_KEY_URL")
	endpoint := tokenUrl + "/api/v1/.well-known/publicKey"
	cert := models.PublicCert{}
	cl := hermes.NewHttpClient()
	//ToDo: inject UUID in Header for Transaction Id
	if code, _, err := cl.Get(endpoint, nil).Result(&cert); err != nil {
		panic(err)
	} else {
		if code > 300 || code < 200 {
			err := fmt.Errorf("invalid status code")
			panic(err)
		}
	}
	return &authImpl{
		logger:            logger,
		client:            cl,
		publicKey:         []byte(cert.PublicKey),
		kid:               cert.Kid,
		tokenPublicKeyUrl: tokenUrl,
	}
}

func (a *authImpl) ParseAndVaidate(tokenS string) (*models.RedisUser, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenS, claims, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})
	if err == nil {
		u := models.RedisUser{}
		if val, ok := claims["given_name"]; ok {
			u.FirstName = val.(string)
			if lNameVal, ok := claims["family_name"]; ok {
				u.LastName = lNameVal.(string)
				if eVal, ok := claims["user_email"]; ok {
					u.Email = eVal.(string)
					if idVal, ok := claims["user_id"]; ok {
						u.Id = idVal.(string)
						if sessVal, ok := claims["session_id"]; ok {
							sess := sessVal.([]string)
							u.Sessions = sess
							if hearVal, ok := claims["hearing_id"]; ok {
								hears := hearVal.([]string)
								u.Hearings = hears
								return &u, nil
							}
						}
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("failed to parse token")
}
