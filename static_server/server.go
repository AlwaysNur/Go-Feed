package staticserver

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Text string
}

var d string = `hello`

func Index(w http.ResponseWriter, r *http.Request) {
	// Data to be passed into the template
	data := PageData{
		Text: d,
	}

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("static_server/index.html") // Path to your HTML template
	if err != nil {
		log.Fatal(err)
	}

	// Call execute function and pass both tmpl and w
	Execute(tmpl, w, data)
}
func SubmitPage(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("static_server/submit.html")
	tmp.Execute(response, nil)
}
func Css(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("static_server/index.css")
	tmp.Execute(response, nil)
}
func Execute(tmpl *template.Template, w http.ResponseWriter, data PageData) {
	// Execute the template and inject the data into the HTML
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
