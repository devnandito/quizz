package main

import (
	"github.com/devnandito/quizz/api"
	"github.com/devnandito/quizz/helper"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Instanciar echo
	e := echo.New()
	// g := e.Group("/admin")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:9000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	// Be careful to use constant time comparison to prevent timing attacks
	// 	if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
	// 		subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	// g.GET("/main", api.ApiShowModule)
	// EndPoint clients
	e.GET("/api/clients", api.ApiShowClients)
	e.POST("/api/clients", api.ApiCreateClient)
	e.POST("/api/clients/search", api.ApiFormSearchClient)
	e.PUT("/api/clients/:id", api.ApiUpdateClient)
	e.DELETE("/api/clients/:id", api.ApiDeleteClient)
	e.GET("/api/clients/:id", api.ApiSearchClient)
	// EndPoint users
	e.GET("/api/users", api.ApiShowUser)
	e.POST("/api/users", api.ApiCreateUser)
	e.PUT("/api/users/:id", api.ApiUpdateUser)
	// EndPoint roles
	//e.GET("/api/roles", api.ApiShowRole)
	e.POST("/api/roles", api.ApiCreateRole)
	e.PUT("/api/roles/:id", api.ApiUpdateRole)
	// EndPoint modules
	e.GET("/api/modules", api.ApiShowModule)
	// EndPoint operations
	e.GET("/api/operations", api.ApiShowOperation)
	// Generator token
	e.POST("/api/gentoken", api.ApiGeneratorToken)
	// Login
	e.POST("api/sign-in", api.SignIn)
	// Restricted group
	r := e.Group("restricted")
	// Configure middleware with custom claims type
	var mySign = helper.GetSecretKey()
	config := middleware.JWTConfig{
		Claims: &helper.JwtCustomClaims{},
		SigningKey: []byte(mySign),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", api.ApiShowRole)

	e.Logger.Fatal(e.Start(":9000"))
}