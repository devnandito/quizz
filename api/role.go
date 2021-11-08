package api

import (
	"net/http"
	"strconv"

	"github.com/devnandito/quizz/helper"
	"github.com/devnandito/quizz/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// ApiShowRole show all client in json
func ApiShowRole(c echo.Context) error {
	var r models.Role
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(*helper.JwtCustomClaims)
	// name := claims.Username
	// role := claims.Role

	rol, err := r.ShowRoleGorm()
	if err != nil {
		panic(err)
	}
	// fmt.Println(name, role)
	return c.JSON(http.StatusOK, rol)
}

// ReadCookie get to cookies
func ReadCookie(c echo.Context) error {
	u := new(models.User)
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return err
	}

	secretKey := helper.GetSecretKey()
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(secretKey), nil
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unautheticated")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	usr, _ := u.SearchUserID(claims.Issuer)

	return c.JSON(http.StatusOK, usr)
}


// ApiCreateRole endpoit to insert role
func ApiCreateRole(c echo.Context) (err error) {
	r := new(models.Role)
	if err = c.Bind(r); err != nil {
		return err
	}

	data := &models.Role{
		Description: r.Description,
	}

	i, _ := r.CreateRoleGorm(data)
	return c.JSON(http.StatusCreated, i)
}

// UpdateClient endpoint to update role
func ApiUpdateRole(c echo.Context) error {
	r := new(models.Role)
	if err := c.Bind(r); err != nil {
		return err
	}

	data := &models.Role{
		Description: r.Description,
	}

	tmp := c.Param("id")
	id, err := strconv.Atoi(tmp)
		if err != nil {
		panic(err)
	}
	i, _ := r.UpdateRoleGorm(id, data)
	return c.JSON(http.StatusOK, i)
}
