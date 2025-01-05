package staticserver

import (
	"html/template"
	"net/http"
)

func SubmitPage(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("static_server/submit.html")
	tmp.Execute(response, nil)
}
func Css(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("static_server/index.css")
	tmp.Execute(response, nil)
}
