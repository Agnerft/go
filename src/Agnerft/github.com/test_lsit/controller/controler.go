package controller

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*"))

func ListProducts(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "index", nil)

}
