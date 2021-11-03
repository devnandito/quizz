package api

import (
	"net/http"

	"github.com/devnandito/quizz/models"
	"github.com/labstack/echo"
)

var mod models.Module
// ApiShowModule show all client in json
func ApiShowModule(c echo.Context) error {
	mod, err := mod.ShowModuleGorm()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, mod)
}