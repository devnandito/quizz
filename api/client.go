package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/devnandito/echogolang/models"
	"github.com/labstack/echo"
)

var cls models.Client

// GetAllClients show all client in json
func ApiShowClients(c echo.Context) error {
	cls, err := cls.ShowClientGorm()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, cls)
}

// CreateClient insert new client
func ApiCreateClient(c echo.Context) (err error) {

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

	i, _ := cls.CreateClientGorm(data)
	return c.JSON(http.StatusCreated, i)

}

// UpdateClient update a client
func ApiUpdateClient(c echo.Context) error {
	cli := new(models.Client)
	if err := c.Bind(cli); err != nil {
		return err
	}

	data := &models.Client{
		FirstName: cli.FirstName,
		LastName: cli.LastName,
		Ci: cli.Ci,
	}

	tmp := c.Param("id")
	id, err := strconv.Atoi(tmp)
		if err != nil {
		panic(err)
	}
	i, _ := cls.SaveEditClientGorm(id, data)
	return c.JSON(http.StatusOK, i)
}

// DeleteClient delete a client
func ApiDeleteClient(c echo.Context) error {
	tmp := c.Param("id")
	id, err := strconv.Atoi(tmp)
	i := cls.DeleteClientGorm(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete id number:", id ,i)
	return c.NoContent(http.StatusNoContent)
}

// SearchClient find a client
func ApiSearchClient(c echo.Context) error {
	ci := c.Param("ci")
	cls, err := cls.ApiGetClientGorm(ci)
	fmt.Println(cls)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, cls)
}