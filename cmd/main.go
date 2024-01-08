package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Contact struct {
	Name  string
	Email string
}

func NewContact(name, email string) *Contact {
	return &Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts []*Contact

type Data struct {
	Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func NewData() *Data {
	return &Data{
		Contacts: Contacts{
			NewContact("John", "jd@gmail.com"),
			NewContact("Clara", "cd@gmail.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() *FormData {
	return &FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func NewPage() *Page {
	return &Page{
		Data: *NewData(),
		Form: *NewFormData(),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	page := NewPage()
	e.Renderer = NewTemplates()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if page.Data.hasEmail(email) {
			formData := NewFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			fmt.Printf("%+v\n", formData)

			return c.Render(http.StatusUnprocessableEntity, "form", formData)
		}

		contact := NewContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)

		c.Render(http.StatusOK, "form", NewFormData())
		return c.Render(http.StatusCreated, "oob-contact", contact)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
