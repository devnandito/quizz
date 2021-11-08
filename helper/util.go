package helper

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	Role int `json:"role"`
	jwt.StandardClaims
}

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Role int `json:"role"`
	Username string `json:"username"`
	TokenString string `json:"token"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func JwtGen(username string, role int, id uint) JwtCustomClaims {
	claims := &JwtCustomClaims{
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer: strconv.Itoa(int(id)),
		},
	}
	
	return *claims
}

func GenerateJWT(username string, role int) (jwt.MapClaims, string) {
	var mySigningKey = "secret"
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	return claims, mySigningKey
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}

func GetSecretKey() string {
	var mySigningKey = "secret"
	return mySigningKey
}

// func GenerateJWT(username string, role int) (string, error){
// 	var mySigningKey = []byte("secret")
// 	token := jwt.New(jwt.SigningMethodES256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["authorized"] = true
// 	claims["username"] = username
// 	claims["role"] = role
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

// 	tokenString, err := token.SignedString(mySigningKey)

// 	if err != nil {
// 		fmt.Printf("Something went wrong: %s", err.Error())
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// type httpError struct {
// 	Code int
// 	Key string `json:"error"`
// 	Message string `json:"message"`
// }

// func NewHTTPError(code int, key, msg string) *httpError {
// 	return &httpError{
// 		Code: code,
// 		Key: key,
// 		Message: msg,
// 	}
// }

// func (e *httpError) Error() string {
// 	return e.Key + ": " + e.Message 
// }

// func NewHTTPError(code int, key, msg string) httpError {
// 	httpErr := &httpError{
// 		Code: code,
// 		Key: key,
// 		Message: msg,
// 	}
// 	return *httpErr
// }