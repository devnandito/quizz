package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Home show home page
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"Title": "Home Page",
		"msg": "This is a home page",
		"url": "clients/show",
	})
}

// ShowClient list all client table
// func ShowClient(c echo.Context) error {
// 	cls, err := models.SeekClient()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, cl := range cls {
// 		fmt.Println(cl.ID, cl.FirstName, cl.LastName, cl.Ci, cl.Birthday)
// 	}

//   return c.String(http.StatusOK, "Hello, Client Page!")
// }