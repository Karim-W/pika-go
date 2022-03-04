package jwtdecoder

import (
	"strings"

	"github.com/golang-jwt/jwt"
	models "github.com/karim-w/go-cket/Models"
)

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
