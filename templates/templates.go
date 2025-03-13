package templates

import (
	"html/template"
)

var (
	Homepage *template.Template
	ContactPage *template.Template
	BlogViewPge *template.Template
)

func ParseHtmlTemplates() {
	Homepage = template.Must(template.ParseFiles("templates/homepage.html"))
	ContactPage = template.Must(template.ParseFiles("templates/contact.html"))
	BlogViewPge = template.Must(template.ParseFiles("templates/blogview.html"))
	
}