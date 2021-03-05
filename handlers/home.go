package handlers

import (
	"fmt"
	"net/http"

	"github.com/devnandito/echogolang/models"
	"github.com/labstack/echo"
)

// Home show home page
func Home(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, Home Page!")
}

// ShowClient list all client table
func ShowClient(c echo.Context) error {
	cls, err := models.SeekClient()
	if err != nil {
		panic(err)
	}

	for _, cl := range cls {
		fmt.Println(cl.ID, cl.FirstName, cl.LastName, cl.Ci, cl.Birthday)
	}

  return c.String(http.StatusOK, "Hello, Client Page!")
}