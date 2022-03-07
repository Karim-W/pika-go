package jwtdecoder

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	models "github.com/karim-w/go-cket/Models"
	"github.com/karim-w/go-cket/utils/hermes"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DecoderService struct {
	logger    *zap.SugaredLogger
	publicKey string
	Kid       string
}

func Decode(token string) (models.UserToken, error) {
	Segments := strings.Split(token, " ")
	tt, _, _ := new(jwt.Parser).ParseUnverified(Segments[1], jwt.MapClaims{})
	claims, _ := tt.Claims.(jwt.MapClaims)
	tok := models.UserToken{
		HearingID: claims["HearingId"].(string),
		Issuer:    claims["iss"].(string),
		SessionID: claims["SessionId"].(string),
		UserEmail: claims["UserEmail"].(string),
		UserID:    claims["UserId"].(string),
		Iss:       claims["iss"].(string),
		Type:      claims["type"].(string),
	}
	return tok, nil
}
func DecodeServer(token string) (models.ServerToken, error) {
	Segments := strings.Split(token, " ")
	tt, _, _ := new(jwt.Parser).ParseUnverified(Segments[1], jwt.MapClaims{})
	claims, _ := tt.Claims.(jwt.MapClaims)
	tok := models.ServerToken{
		ClientID: claims["client_Id"].(string),
	}
	return tok, nil
}
func VerifySignature(token string) (bool, error) {
	return true, nil
}

func NewDecoderService(logger *zap.SugaredLogger, hermes *hermes.HttpClient) *DecoderService {
	fetchPublicKeyUrl := os.Getenv("AuthServiceUrl") + "/api/v1/.well-known/key"
	certificate := models.PublicCert{}
	if statusCode, ok, err := hermes.Get(fetchPublicKeyUrl, nil).Result(&certificate); ok {
		if statusCode == 200 {
			return &DecoderService{
				logger:    logger,
				publicKey: certificate.PublicKey,
				Kid:       certificate.Kid,
			}
		}
	} else {
		logger.Error("Error fetching public key: ", err)
		return nil
	}
	return &DecoderService{}
}

var DecoderModule = fx.Provide(NewDecoderService)
