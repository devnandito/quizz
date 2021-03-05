package models

import (
	"fmt"
	"time"

	"github.com/devnandito/echogolang/lib"
)

// Client client access public
type Client struct {
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Ci string `json:"ci"`
	Birthday time.Time `json:"birthday"`
}

// SeekClient show all client
func SeekClient() ([]Client, error) {

	conn := lib.NewConfig()
	db := conn.DsnString()
	rows, err := db.Query("SELECT id, first_name, last_name, ci, birthday FROM clients")

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
func CreateClient(cls *Client) error {
	// var LastInsertId int
	// conn := lib.NewConfig()
	// db := conn.DsnString()

	// row := db.QueryRow("INSERT INTO clients(first_name,last_name,ci) VALUES($1,$2,$3) returning id;", cls.FirstName, cls.LastName, cls.LastName).Scan(&LastInsertId)
	fmt.Println(cls.FirstName, cls.LastName, cls.Ci)
	// if row != nil {
	// 	panic(row)
	// }

	// fmt.Println("Último id =", LastInsertId)
	return nil
}