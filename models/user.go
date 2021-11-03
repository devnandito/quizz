package models

import (
	"github.com/devnandito/quizz/lib"
	"gorm.io/gorm"
)

// User access public
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
	RoleID int `json:"roleid"`
	Role Role
}

// type Error struct {
// 	ResponseCode int `json:"rc"`
// 	Message string `json:"message"`
// 	Detail string `json:"detail"`
// 	ExternalReference string `json:"ext_ref"`
// }

// ShowUserGorm show user
func (u User) ShowUserGorm() ([]User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&User{})
	rows, err := db.Order("id asc").Model(&u).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []User
	for rows.Next() {
		var item User
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}

// CreateUserGorm insert a new user
func (u User) CreateUserGorm(usr *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Create(&usr)

	data := User {
		Username: usr.Username,
		Email: usr.Email,
		Password: usr.Password,
		Token: usr.Token,
		RoleID: usr.RoleID,
	}

	return data, response.Error
}