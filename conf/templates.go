package conf

import (
	"html/template"
)

var templatePath = "../../frontend/templates/"

var Templates = map[string]*template.Template{
	"homepage": template.Must(template.ParseFiles(templatePath + "index.tmpl")),
}
