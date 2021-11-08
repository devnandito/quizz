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
	// Token string `json:"token"`
	RoleID int `json:"roleid"`
	Role Role
}

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
func (u User) CreateUserGorm(data *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Create(&data)
	user := User {
		Username: u.Username,
		Email: u.Email,
		Password: u.Password,
		RoleID: u.RoleID,
	}

	return user, response.Error
}

// UpdateUserGorm update user
func (u User) UpdateUserGorm(id int, usr *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&u).Where("id = ?", id).Updates(User{Username: usr.Username, Email: usr.Email, RoleID: usr.RoleID})
	return u, response.Error
}

func (u User) SearchUser(data *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Where("username = ?", data.Username).Find(&u)
	return u, response.Error
}

func (u User) SearchUserID(data string) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Where("id = ?", data).First(&u)
	return u, response.Error
}