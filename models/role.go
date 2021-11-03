package models

import (
	"github.com/devnandito/quizz/lib"
	"gorm.io/gorm"
)

// Role access public
type Role struct {
	gorm.Model
	Description string `json:"description"`
	Operation []Operation `gorm:"many2many:role_operations;"`
}

// ShowRoleGorm show role
func (r Role) ShowRoleGorm() ([]Role, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Role{})
	rows, err := db.Order("id asc").Model(&r).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Role
	for rows.Next() {
		var item Role
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}