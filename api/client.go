package api

import (
	"fmt"
	"net/http"
	"strconv"

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

// GetAllClientsGorm show all client in json
func GetAllClientsGorm(c echo.Context) error {
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
		Birthday: cli.Birthday,
	}

	i, _ := models.CreateClient(data)
	return c.JSON(http.StatusCreated, i)

}

// UpdateClient update a client
func UpdateClient(c echo.Context) error {
	cli := new(models.Client)
	if err := c.Bind(cli); err != nil {
		return err
	}

	data := &models.Client{
		FirstName: cli.FirstName,
		LastName: cli.LastName,
		Ci: cli.Ci,
	}

	ci := c.Param("ci")
	i, _ := models.UpdateClient(ci, data)
	return c.JSON(http.StatusOK, i)
}

// DeleteClient delete a client
func DeleteClient(c echo.Context) error {
	ci := c.Param("ci")
	i := models.DeleteClient(ci)
	fmt.Println("Delete:", i)
	return c.NoContent(http.StatusNoContent)
}

// SearchClient find a client
func SearchClient(c echo.Context) error {
	ci, _ := strconv.Atoi(c.Param("ci"))
	cls, _ := models.GetClient(ci)
	return c.JSON(http.StatusOK, cls)
}