package api

import (
	"net/http"

	"github.com/devnandito/quizz/models"
	"github.com/labstack/echo"
)

var rol models.Role

// ApiShowRole show all client in json
func ApiShowRole(c echo.Context) error {
	rol, err := rol.ShowRoleGorm()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, rol)
}

// ApiCreateRole endpoit to insert role
func ApiCreateRole(c echo.Context) (err error) {
	rol := new(models.Role)
	if err = c.Bind(rol); err != nil {
		return
	}

	data := &models.Role{
		Description: rol.Description,
	}

	i, _ := rol.CreateRoleGorm(data)
	return c.JSON(http.StatusCreated, i)
}