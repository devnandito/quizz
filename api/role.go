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