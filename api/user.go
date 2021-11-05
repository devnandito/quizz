package api

import (
	"net/http"
	"strconv"

	"github.com/devnandito/quizz/helper"
	"github.com/devnandito/quizz/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)


var msg helper.ErrorMessage

// ApiShowUser show all client in json
func ApiShowUser(c echo.Context) error {
	var u models.User
	usr, err := u.ShowUserGorm()
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

	if usr.Email == "" {
		msg.Message = "Email isn't empty!!"
		return c.JSON(http.StatusBadRequest, msg)  
	}

	// hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 5)
	if err != nil {
		return err
	}

	data := &models.User{
		Username: usr.Username,
		Email: usr.Email,
		Password: string(hash),
		RoleID: usr.RoleID,
	}

	i, _ := usr.CreateUserGorm(data)

	return c.JSON(http.StatusCreated, i)
}

// ApiUpdateUser endport to update user
func ApiUpdateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	data := &models.User{
		Username: u.Username,
		Email: u.Email,
		RoleID: u.RoleID,
	}

	tmp := c.Param("id")
	id, err := strconv.Atoi(tmp)
		if err != nil {
		panic(err)
	}
	i, _ := u.UpdateUserGorm(id, data)
	return c.JSON(http.StatusOK, i)
}

// ApiGeneratorToken endpoint to generate token
func ApiGeneratorToken(c echo.Context) (err error) {
	usr := new(models.User)
	if err = c.Bind(usr); err != nil {
		return
	}

	mySign := helper.GetSecretKey()
	claims := helper.JwtGen(usr.Username, usr.RoleID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(mySign))

	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, validToken)
}

func SignIn(c echo.Context) (err error) {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	authdetails := &helper.Authentication {
		Username: u.Username,
		Password: u.Password,
	}
	
	data := &models.User{
		Username: u.Username,
		Password: u.Password,
	}

	response, err := u.SearchUser(data)
	check := helper.CheckPasswordHash(authdetails.Password, response.Password)
	
	if !check {
		msg.Message = "Username or Password incorrect!!!"
		return c.JSON(http.StatusBadRequest, msg)
	}

	keySign := helper.GetSecretKey()
	claims := helper.JwtGen(response.Username, response.RoleID)
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := tokenString.SignedString([]byte(keySign))

	if err != nil {
		return err
	}
	
	var token helper.Token
	token.Username = response.Username
	token.Role = response.RoleID
	token.TokenString = validToken
	
	return c.JSON(http.StatusOK, token)
}