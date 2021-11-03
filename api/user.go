package api

import (
	"net/http"

	"github.com/devnandito/quizz/helper"
	"github.com/devnandito/quizz/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
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

// ApiCreateUser endpoint to insert new user
func ApiCreateUser(c echo.Context) (err error) {
	usr := new(models.User)
	if err = c.Bind(usr); err != nil {
		return
	}

	// hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 5)
	if err != nil {

	}

	data := &models.User{
		Username: usr.Username,
		Email: usr.Email,
		Password: string(hash),
		Token: helper.JwtGenerator(usr.Username, usr.Email, "secretekey"),
		RoleID: usr.RoleID,
	}

	i, _ := usr.CreateUserGorm(data)
	return c.JSON(http.StatusCreated, i)
}

// func LoginController(c echo.Context) error {
// 	payload, _ := ioutil.ReadAll(c.Request().Body)
// 	err := json.Unmarshal(payload, &usr)
	
// 	if err != nil {
// 		return err
// 	}



// }