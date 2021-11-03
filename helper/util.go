package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func DateTime() string {
	currentTime := time.Now()
	result := currentTime.Format("2006-01-02 15:04:05")
	return result
}

func JwtGenerator(username, email, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
		"email": email,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

// func ErrorLog(rc int, detail, ext_ref string) models.Error {
// 	var error models.Error
// 	error.ResponseCode = rc
// 	error.Message = "Failed"
// 	error.Detail = detail
// 	error.ExternalReference = ext_ref
	
// 	return error
// }