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