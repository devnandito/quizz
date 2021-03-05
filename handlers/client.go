package handlers

import (
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
func CreateClient(c echo.Context) (err error) {

	cli := new(models.Client)
	if err = c.Bind(cli); err != nil {
		return
	}

	data := &models.Client{
		FirstName: cli.FirstName,
		LastName: cli.LastName,
		Ci: cli.Ci,
	}

	i := models.CreateClient(data)
	return c.JSON(http.StatusCreated, i)

}