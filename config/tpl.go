package config

import "html/template"

//TPL ...
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
