package conf

import (
	"html/template"
	"path/filepath"
)

var templatePath, _ = filepath.Abs("../../frontend/templates")

var Templates = map[string]*template.Template{
	"homepage": template.Must(template.ParseFiles(templatePath + "/index.tmpl")),
}
