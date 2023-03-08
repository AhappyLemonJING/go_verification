package web

import (
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("web/temp/index.html")

	t.Execute(w, "yanzhengma")
}
