package api

import (
	"net/http"

	"github.com/devnandito/quizz/models"
	"github.com/labstack/echo"
)

var op models.Operation
// ApiShowOperation show all client in json
func ApiShowOperation(c echo.Context) error {
	op, err := op.ShowOperationGorm()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, op)
}