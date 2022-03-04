package middlewares

// import (
// 	"fmt"
// 	"strings"
// )

// func AuthenticationMiddleware(AuthizationHeader string) error {

// 	if AuthizationHeader == "" {
// 		return fmt.Errorf("no Authorization header")
// 	} else {
// 		splitToken := strings.Split(AuthizationHeader, " ")
// 		if len(splitToken) < 1 {
// 			return fmt.Errorf("no Authorization header")
// 		}
// 		parsedToken, err := tu.ParseEUJWTtoken(splitToken[1])
// 		if err != nil {
// 			return fmt.Errorf("invalid token")
// 		} else {
// 			// Handle logic for validating Auth
// 			if parsedToken.Claims.Valid() != nil {
// 				return fmt.Errorf("invalid token")
// 			} else {
// 				return nil
// 			}
// 		}

// 	}
// }
