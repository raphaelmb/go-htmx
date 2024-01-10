package main

var id = 0

type Contact struct {
	Name  string
	Email string
	Id    int
}

func NewContact(name, email string) *Contact {
	id++
	return &Contact{
		Name:  name,
		Email: email,
		Id:    id,
	}
}

type Data struct {
	Contacts []*Contact
}

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func (d *Data) indexOf(id int) int {
	for i, contact := range d.Contacts {
		if contact.Id == id {
			return i
		}
	}
	return -1
}

func NewData() *Data {
	return &Data{
		Contacts: []*Contact{
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
