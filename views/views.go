package views

import (
	"net/http"
	"fmt"
	"gotemplater"
)

func index(w http.ResponseWriter, r *http.Request) {

	m := map[string]string{
		"name": "Chris",
	}

	gotemplater.RenderTemplate(w, "index.html", m)
}

func success(w http.ResponseWriter, r *http.Request){
	gotemplater.RenderTemplate(w, "success.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Errorf("%v", err)
		}
		user := r.Form.Get("username")
		if user == "chris"{
			http.Redirect(w, r, "/success", http.StatusTemporaryRedirect)
		}

	}

	gotemplater.RenderTemplate(w, "login.html", nil)
}


func StartUp(){

	gotemplater.LoadStatic("static")
	gotemplater.LoadConfiguration("templates/", "templates/layout/")
	gotemplater.LoadTemplates()

	loadUrls()

}