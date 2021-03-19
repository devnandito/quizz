package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/devnandito/echogolang/lib"
	"gorm.io/gorm"
)

// Client client access public
type Client struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Ci string `json:"ci" validate:"required"`
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

// func (f FormClient) CustomValidator(val string) bool {
// 	return len(val) == 4
// }

// func (f FormClient) Messages() map[string]string {
// 	return validate.MS{
// 		"required": "oh! the {field} is required",
// 		"Ci.required": "Ci is required",
// 	}
// }

// func (f FormClient) Translates() map[string]string {
// 	return validate.MS{
// 		"Ci": "Document client",
// 	}
// }

// func (c Client) Validate() {
// 	c.Errors = make(map[string]string)
// 	if strings.TrimSpace(c.Ci) == "" {
// 		c.Errors["ci"] = "Enter a ci"
// 	}
// 	return len(c.Errors) == 0
// }
// type Client struct {
// 	gorm.Model
// 	FirstName string `json:"firstname"`
// 	LastName string `json:"lastname"`
// 	Ci string `json:"ci" validate:"required"`
// 	Birthday time.Time `json:"birthday"`
// 	Sex string `json:"sex"`
// 	Nationality string `json:"nationality"`
// 	DesType string `json:"destype"`
// 	Code1 string `json:"code1"`
// 	Code2 string `json:"code2"`
// 	Code3 string `json:"code3"`
// 	Direction string `json:"direction"`
// 	Phone string `json:"phone"`
// }

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
func (c Client) ApiGetClientGorm(ci string) ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	rows, err := db.Order("id asc").Model(&c).Where("firstname LIKE ?", ci).Rows()
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

// SeekClient show all client
func SeekClient() ([]Client, error) {

	conn := lib.NewConfig()
	db := conn.DsnString()
	rows, err := db.Query("SELECT id, first_name, last_name, ci, birthday FROM clients LIMIT 20")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var cls []Client
	for rows.Next() {
		var cl Client
		err := rows.Scan(&cl.ID, &cl.FirstName, &cl.LastName, &cl.Ci, &cl.Birthday)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cls, nil
}

// CreateClient insert new client
func CreateClient(cls *Client)  (*Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnString()

	row, err := db.Prepare("INSERT INTO clients(first_name,last_name,ci,birthday) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err)
	}

	defer row.Close()

	row.Exec(cls.FirstName, cls.LastName, cls.Ci, cls.Birthday)
	fmt.Println(cls.FirstName, cls.LastName, cls.Ci, cls.Birthday)
	var i = &Client{
		FirstName: cls.FirstName,
		LastName: cls.LastName,
		Ci: cls.Ci,
		Birthday: cls.Birthday,
	}
	return i, nil
}

// UpdateClient update client
func UpdateClient(ci string, cls *Client)  (*Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnString()

	row, err := db.Prepare("UPDATE clients SET first_name = $1, last_name = $2 WHERE ci = $3")
	if err != nil {
		panic(err)
	}

	defer row.Close()

	row.Exec(cls.FirstName, cls.LastName, ci)
	fmt.Println(cls.FirstName, cls.LastName, ci)
	var i = &Client{
		FirstName: cls.FirstName,
		LastName: cls.LastName,
	}
	return i, nil
}

// DeleteClient delete client
func DeleteClient(ci string)  error {
	conn := lib.NewConfig()
	db := conn.DsnString()

	row, err := db.Prepare("DELETE FROM clients WHERE ci=$1")
	if err != nil {
		panic(err)
	}

	defer row.Close()
	row.Exec(ci)
	return nil
}

// GetClient search a client
func GetClient(ci int) ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnString()
	row, err := db.Query("SELECT id, first_name, last_name, ci, birthday FROM clients WHERE ci=$1", ci)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var cls []Client
	for row.Next() {
		var cl Client
		err := row.Scan(&cl.ID, &cl.FirstName, &cl.LastName, &cl.Ci, &cl.Birthday)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)
	}

	return cls, nil
}

// func CreateClient(cls *Client)  error {
// 	var LastInsertId int
// 	conn := lib.NewConfig()
// 	db := conn.DsnString()

// 	row := db.QueryRow("INSERT INTO clients(first_name,last_name,ci) VALUES($1,$2,$3) returning id;", cls.FirstName, cls.LastName, cls.LastName).Scan(&LastInsertId)
// 	fmt.Println(cls.FirstName, cls.LastName, cls.Ci)
// 	if row != nil {
// 		panic(row)
// 	}

// 	fmt.Println("Último id =", LastInsertId)
// 	return nil
// }