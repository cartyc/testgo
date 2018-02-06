package main

import (
	"net/http"
	"html/template"
	"fmt"
)

const templateFolder = "templates/"

var (
	templates = populateTemplates()
)


func renderTemplate(w http.ResponseWriter, tmpl string){
	t := templates.Lookup(tmpl + ".html")
	t.Execute(w, nil)
}

// Handle the request
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(r)
	renderTemplate(w, "index")
}

func viewhandler(w http.ResponseWriter, r *http.Request){
	requested := r.URL.Path
	fmt.Print(requested)
	renderTemplate(w, "test")
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/test/", viewhandler)
	http.ListenAndServe(":8080", nil)

}

// takes pointer to templates
func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"

	// Get only files that end in html
	template.Must(result.ParseGlob(templateFolder + "/*.html"))

	return result
}