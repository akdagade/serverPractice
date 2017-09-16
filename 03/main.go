package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "index.gohtml", "Hello!!!! Time starts now!!")
	r.ParseForm();

	fmt.Println("Works!!!", r.Form["name"][0])

}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("template/index.gohtml"))
}

func main() {

	http.HandleFunc("/dog", a)
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err)
	}
}
