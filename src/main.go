package main

import (
	"github.com/devnandito/quizz/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Instanciar echo
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"http://localhost:3000", "http://localhost:9000"},
  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
}))
	// EndPoint clients
	e.GET("/api/clients", api.ApiShowClients)
	e.POST("/api/clients", api.ApiCreateClient)
	e.POST("/api/clients/search", api.ApiFormSearchClient)
	e.PUT("/api/clients/:id", api.ApiUpdateClient)
	e.DELETE("/api/clients/:id", api.ApiDeleteClient)
	e.GET("/api/clients/:id", api.ApiSearchClient)
	// EndPoint users
	e.GET("/api/users", api.ApiShowUser)
	// EndPoint roles
	e.GET("/api/roles", api.ApiShowRole)
	// EndPoint modules
	e.GET("/api/modules", api.ApiShowModule)
	// EndPoint operations
	e.GET("/api/operations", api.ApiShowOperation)
	e.Logger.Fatal(e.Start(":9000"))
}