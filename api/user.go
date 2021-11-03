package api

import (
	"net/http"

	"github.com/devnandito/quizz/models"
	"github.com/labstack/echo"
)

var usr models.User
// ApiShowUser show all client in json
func ApiShowUser(c echo.Context) error {
	usr, err := usr.ShowUserGorm()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, usr)
}