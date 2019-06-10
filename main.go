package main

import (
	"html/template"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var (
	ideatable    = "ideas"
	seq          = 1
	connect, err = dbr.Open("mysql", "root:@tcp(localhost:3306)/ideabox", nil)
	session      = connect.NewSession(nil)
)

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("statics/templates/*.html")),
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "top", echo.Map{
			"title": "Top",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
