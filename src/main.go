package main

import (
	"errors"
	"io"
	"text/template"

	"github.com/devnandito/echogolang/api"
	"github.com/devnandito/echogolang/handlers"
	"github.com/labstack/echo"
)

// TemplateRegistry initial
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Render template
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found"+ name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	// Instanciar echo
	e := echo.New()
	e.GET("/api/clients", api.ApiShowClients)
	e.POST("/api/clients", api.ApiCreateClient)
	e.PUT("/api/clients/:id", api.ApiUpdateClient)
	e.DELETE("/api/clients/:id", api.ApiDeleteClient)
	e.GET("/api/clients/:ci", api.ApiSearchClient)
	
	templates := make(map[string]*template.Template)
	templates["index.html"] = template.Must(template.ParseFiles("views/clients/index.html", "views/base.html"))
	templates["home.html"] = template.Must(template.ParseFiles("views/home/index.html", "views/base.html"))
	templates["search.html"] = template.Must(template.ParseFiles("views/clients/search.html", "views/base.html"))
	templates["result.html"] = template.Must(template.ParseFiles("views/clients/result.html", "views/base.html"))
	templates["create.html"] = template.Must(template.ParseFiles("views/clients/create.html", "views/base.html"))
	templates["edit.html"] = template.Must(template.ParseFiles("views/clients/edit.html", "views/base.html"))
	templates["msg.html"] = template.Must(template.ParseFiles("views/clients/msg.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}
	e.Static("/static", "assets")
	e.File("/favicon.png", "static/img/favicon.png")
	e.GET("/", handlers.Home)
	e.Static("/clients/static", "assets")
	e.GET("/clients/show", handlers.ShowClientsGorm)
	e.GET("/clients/search", handlers.SearchForm)
	e.POST("/clients/search", handlers.ResultSearch)
	e.GET("/clients/create", handlers.ShowFormClient)
	e.POST("/clients/create", handlers.SaveFormClient)
	e.Static("/clients/edit/:id", "assets")
	e.GET("/clients/edit/:id", handlers.EditFormClient)
	e.POST("/clients/edit/:id", handlers.UpdateClientGorm)
	e.Static("/clients/delete/:id", "assets")
	e.GET("/clients/delete/:id", handlers.DeleteClientGorm)
	e.Logger.Fatal(e.Start(":9000"))
}