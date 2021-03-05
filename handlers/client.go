package handlers

import (
	"fmt"
	"net/http"

	"github.com/devnandito/echogolang/models"
	"github.com/labstack/echo"
)

// GetAllClients show all client in json
func GetAllClients(c echo.Context) error {
	cls, err := models.SeekClient()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, cls)
}

// CreateClient insert new client
func CreateClient(c echo.Context) error {
	firstname := c.QueryParam("firstname")
	lastname := c.FormValue("lastname")
	ci := c.QueryParam("ci")

	fmt.Println("Nomrbre:", firstname, "Apellido:", lastname, "Cedula", ci)
	// birthday := c.Param("birthday")
	// firstname, _ := strconv.Atoi(c.Param("firstname"))
	data := &models.Client{
		FirstName: firstname,
		LastName: lastname,
		Ci: ci,
	}

	i := models.CreateClient(data)

	if err := c.Bind(data); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, i)
}