package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/devnandito/echogolang/models"
	"github.com/gookit/validate"
	"github.com/labstack/echo"
)

// SearchForm render form
func SearchForm(c echo.Context) error {
	return c.Render(http.StatusOK, "search.html", map[string]interface{}{
		"Title": "Search Form",
	})
}

// ShowFormClient render client form
func ShowFormClient(c echo.Context) error {
	return c.Render(http.StatusOK, "create.html", map[string]interface{}{
		"Title": "Create Client",
	})
}

var cls models.Client

// ResultSearch lista client 
func ResultSearch(c echo.Context) (err error) {
	document :=c.FormValue("ci")
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	data, err := validate.FromRequest(c.Request())

	if err != nil {
			panic(err)
	}

	v := data.Create()
	// v.AddRule("ci", "required")
	v.AddRule("firstname", "required")
	v.AddMessages(map[string]string{"required": "{field} requerido"})
	if v.Validate() {
		form := &models.FormClient{}
		v.BindSafeData(form)
		response, err := cls.GetClientGorm(document, firstname, lastname)
		if err != nil {
			panic(err)
		}
		return c.Render(http.StatusOK, "result.html", map[string]interface{}{
			"Title": "Result search client",
			"clients": response,
			"count": len(response),
		})
	} else {
		var s string
		data := &models.FormClient{
			Ci: document,
			FirstName: firstname,
			LastName: lastname,
		}
		for _, value := range v.Errors {
			s = fmt.Sprintf("%s", value["required"])
		}
		return c.Render(http.StatusOK, "search.html", map[string]interface{}{
			"Title": "Result search client",
			"errors": s,
			"data": data,
		})
	}
}
// func ResultSearch(c echo.Context) (err error) {
// 	document :=c.FormValue("ci")
// 	firstname := c.FormValue("firstname")
// 	lastname := c.FormValue("lastname")
// 	data, err := validate.FromRequest(c.Request())
// 	if err != nil {
// 			panic(err)
// 	}

// 	v := data.Create()
// 	v.AddRule("ci", "required")
// 	// v.AddRule("firstname", "required")
// 	v.AddMessages(map[string]string{"required": "{field} requerido"})
// 	if v.Validate() {
// 		form := &models.FormClient{}
// 		v.BindSafeData(form)
// 		response, err := cls.GetClientGorm(document, firstname, lastname)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return c.Render(http.StatusOK, "result.html", map[string]interface{}{
// 			"Title": "Result search client",
// 			"clients": response,
// 			"count": len(response),
// 		})
// 	} else {
// 		// fmt.Println(len(v.Errors.All()))
// 		var s string
// 		// var r []map[string]string
// 		for _, value := range v.Errors {
// 			// fmt.Sprintf("key:", key, "value:", value["required"])
// 			s = fmt.Sprintf("%s", value["required"])
// 			// r = append(r, v.Errors.Field(key))
// 		}
// 		// r = append(r, v.Errors.Field("firstname"))
// 		// fmt.Println(len(s))
// 		// for key, value := range r {
// 		// 	fmt.Println(key, value)
// 		// }
// 		// fmt.Println(v.Errors.FieldOne("ci"))
// 		// fmt.Println(v.Errors.FieldOne("firstname"))
// 		// ci := v.Errors.Field("ci")
// 		// firstname := v.Errors.Field("firstname")
// 		// fmt.Println(ci["required"])
// 		// fmt.Println(firstname["required"])
// 		return c.Render(http.StatusOK, "search.html", map[string]interface{}{
// 			"Title": "Result search client",
// 			"errors": s,
// 		})
// 	}
// }

// func ResultSearch(c echo.Context) (err error) {
// 	document :=c.FormValue("ci")
// 	firstname := c.FormValue("firstname")
// 	lastname := c.FormValue("lastname")
// 	data, err := validate.FromRequest(c.Request())
// 	if err != nil {
// 			panic(err)
// 	}
// 	v := data.Create()
// 	v.AddRule("ci", "required")
// 	if v.Validate() {
// 		form := &models.FormClient{}
// 		v.BindSafeData(form)
// 		response, err := cls.GetClientGorm(document, firstname, lastname)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return c.Render(http.StatusOK, "result.html", map[string]interface{}{
// 			"Title": "Result search client",
// 			"clients": response,
// 			"count": len(response),
// 		})
// 	} else {
// 		fmt.Println(v.Errors.Field("ci"))
// 		return
// 	}
// }

// func ResultSearch(c echo.Context) (err error) {
// 	v := validate.Struct(cls)
// 	document :=c.FormValue("ci")
// 	firstname := c.FormValue("firstname")
// 	lastname := c.FormValue("lastname")
// 	if v.Validate() {
// 		response, err := cls.GetClientGorm(document, firstname, lastname)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return c.Render(http.StatusOK, "result.html", map[string]interface{}{
// 			"Title": "Result search client",
// 			"clients": response,
// 			"count": len(response),
// 		})
// 	} else {
// 		fmt.Println(v.Errors)
// 		return
// 	}
	// data := &models.Client{
	// 	Ci: document,
	// }
	// if data.Validate() == false {
	// 	fmt.Println(data.Errors)
	// }
	// val := &models.CustomValidator{validator: validator.New()}
	// if err = c.Validate(cls); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	// data := &DataForm{
	// 	FirstName: firstname,
	// 	LastName: lastname,
	// 	Ci: document,
	// }
	// var validate = validator.New()
	// errs := validate.Struct(cls)
	// var errors []string
	// if errs != nil {
	// 	for _, value := range errs.(validator.ValidationErrors) {
	// 		errors = append(errors, value.Tag())
	// 	}
	// 	return c.Render(http.StatusOK, "search.html", map[string]interface{}{
	// 		"Title": "Result search client",
	// 		"err": errors,
	// 		"data": data,
	// 	})
	// }
// }

// ShowClients list of client
func ShowClientsGorm(c echo.Context) error {
	response, err := cls.ShowClientGorm()
	if err != nil {
		panic(err)
	}
	// fmt.Println(response)
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Title": "List clients",
		"clients": response,
	})
}

// SaveFormClient save new client 
func SaveFormClient(c echo.Context) error {
	cli := new(models.Client)
	if err := c.Bind(cli); err != nil {
		return err
	}
	t := c.FormValue("birthday") + "T15:04:05"
	Btime := cls.BirthdayTime(t)
	data := &models.Client{
		FirstName: strings.Title(cli.FirstName),
		LastName: strings.Title(cli.LastName),
		Ci: cli.Ci,
		Birthday: Btime,
		Sex: strings.ToUpper(cli.Sex),
	}
	response, err := cls.CreateClientGorm(data)
	if err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, "msg.html", map[string]interface{}{
		"Title": "Create client",
		"msg": "Record saved",
		"client": response,
	})
}

// EditFormClient render form edit client
func EditFormClient(c echo.Context) error {
	tmp := c.Param("id")
	id, err := strconv.ParseInt(tmp, 10, 64)
	response, err := cls.EditClientGorm(id)
	if err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
		"Title": "Edit Client",
		"client": response,
	})
}

// UpdateClientGorm update client
func UpdateClientGorm(c echo.Context) error {
	cli := new(models.Client)
	if err := c.Bind(cli); err != nil {
		return err
	}
	t := c.FormValue("birthday") + "T15:04:05"
	Btime := cls.BirthdayTime(t)
	data := &models.Client{
		FirstName: strings.Title(cli.FirstName),
		LastName: strings.Title(cli.LastName),
		Ci: cli.Ci,
		Birthday: Btime,
		Sex: strings.ToUpper(cli.Sex),
	}

	tmp := c.Param("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		panic(err)
	}
	response, err := cls.SaveEditClientGorm(id, data)
	return c.Render(http.StatusOK, "msg.html", map[string]interface{}{
		"Title": "Updated record",
		"client": response,
	})
}

// DeleteClientGorm delete a client
func DeleteClientGorm(c echo.Context) error {
	tmp := c.Param("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		panic(err)
	}
	response := cls.DeleteClientGorm(id)
	return c.Render(http.StatusOK, "msg.html", map[string]interface{}{
		"Title": "Delete record",
		"client": response,
	})
}

// func EditFormClient (c echo.Context) error {
// 	tmp := c.Param("id")
// 	// pk, _ := strconv.Atoi(tmp)
// 	id, err := strconv.ParseInt(tmp, 10, 64)
// 	cls, err := models.EditClientGorm(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
// 		"Title": "Edit Client",
// 		"client": cls,
// 	})
// }