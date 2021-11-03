package models

import (
	"strings"
	"time"

	"github.com/devnandito/quizz/lib"
	"gorm.io/gorm"
)

// Client client access public
type Client struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Ci string `json:"ci"`
	Birthday time.Time `json:"birthday"`
	Sex string `json:"sex"`
	Nationality string `json:"nationality"`
	DesType string `json:"destype"`
	Code1 string `json:"code1"`
	Code2 string `json:"code2"`
	Code3 string `json:"code3"`
	Direction string `json:"direction"`
	Phone string `json:"phone"`
	Code string
}

type FormClient struct {
	Ci string
	FirstName string
	LastName string
}

// BirthdayDateStr conver to string
func (c Client) BirthdayDateStr() string {
	return c.Birthday.Format("2006-01-02")
}
// BirthdayTime convert string to time
func (c Client) BirthdayTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:04:05"
	t, _ := time.Parse(Format, timeStr)
	return t
}

// GetClientGorm show all client
func (c Client) GetClientGorm(ci, firstname, lastname string) ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	name := strings.Title(firstname)
	last := strings.Title(lastname)
	var condition string
	var value1 string
	var value2 string
	var val []string
	if name == "" && last == "" {
		condition = "ci = ?"
		val = append(val, ci)
	} else if ci == "" && last == "" {
		condition = "first_name LIKE ? "
		val = append(val, name+"%")
	} else if ci == "" && name == "" {
		condition = "last_name LIKE ?"
		val = append(val, last+"%")
	} else if ci == "" {
		condition = "last_name LIKE ? OR first_name LIKE ?"
		val = append(val, last+"%")
		val = append(val, name+"%")
		} else if name == "" {
			condition = "last_name LIKE ? OR ci = ?"
			val = append(val, last+"%")
			val = append(val, ci)
	} else if last == "" {
		condition = "first_name LIKE ? OR ci = ?"
		val = append(val, name+"%")
		val = append(val, ci)
	}
	if len(val) == 2 {
		value1 = val[0]
		value2 = val[1]
	} else {
		value1 = val[0]
		value2 = ""
	}
	rows, err := db.Order("id asc").Model(&c).Where(condition, value1, value2).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Client
	for rows.Next() {
		var item Client
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}

// ApiGetClientGorm show all client
func (c Client) ApiGetClientGorm(id string) ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	rows, err := db.Order("id asc").Model(&c).Where("id = ?", id).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Client
	for rows.Next() {
		var item Client
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}

// ApiSearchClientGorm show all client
func (c Client) ApiSearchClientGorm(cls *Client) ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	list := strings.Fields(cls.FirstName)
	var name string
	var lastname string
	var condition string
	var val[] string

	if len(list) < 2 {
		condition = "first_name LIKE ? OR last_name LIKE ?"
		name = list[0]
		val = append(val, name+"%")
	} else if len(list) == 2 {
		condition = "first_name LIKE ?"
		name = list[0] + " " + list[1]
		val = append(val, name+"%")
	} else if len(list) == 3 {
		condition = "first_name LIKE ? OR last_name LIKE ?"
		name = list[0] + " " + list[1]
		lastname = list[2]
		val = append(val, name+"%")
		val = append(val, lastname+"%")
	} else {
		condition = "first_name LIKE ? OR last_name LIKE ?"
		name = list[0] + " " + list[1]
		lastname = list[2] + " " + list[3]
		val = append(val, name+"%")
		val = append(val, lastname+"%")
	}

	if len(val) == 1 {
		name = val[0]
		lastname = val[0]
	} else {
		name = val[0]
		lastname = val[1]
	}

	rows, err := db.Order("id asc").Model(&c).Where(condition, name, lastname).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Client
	for rows.Next() {
		var item Client
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}

//CreateClientGorm insert new client
func (c Client) CreateClientGorm(cls *Client) (Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Create(&cls)
	data := Client{
		FirstName: cls.FirstName,
		LastName: cls.LastName,
		Ci: cls.Ci,
		Birthday: cls.Birthday,
		Sex: cls.Sex,
	}
	return data, response.Error
}

// EditClientGorm edit client
func (c Client) EditClientGorm(id int64) (Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Find(&c, id)
	return c, response.Error
}

// SaveEditClientGorm saved client edit
func (c Client) SaveEditClientGorm(id int, cls *Client) (Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&c).Where("id = ?", id).Updates(Client{FirstName: cls.FirstName, LastName: cls.LastName, Ci: cls.Ci, Birthday: cls.Birthday, Sex: cls.Sex})
	return c, response.Error
}

// DeleteClientGorm delete client
func (c Client) DeleteClientGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&c, id)
	return response.Error
}

// ShowClientGorm show client
func (c Client) ShowClientGorm() ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Client{})
	rows, err := db.Order("id asc").Model(&c).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Client
	for rows.Next() {
		var item Client
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}