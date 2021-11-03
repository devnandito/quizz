package main

import (
	"github.com/devnandito/quizz/api"
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
	// EndPoint roles
	e.GET("/api/roles", api.ApiShowRole)
	e.POST("/api/roles", api.ApiCreateRole)
	// EndPoint modules
	e.GET("/api/modules", api.ApiShowModule)
	// EndPoint operations
	e.GET("/api/operations", api.ApiShowOperation)
	e.Logger.Fatal(e.Start(":9000"))
}