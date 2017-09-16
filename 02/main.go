package _3

import (
	"fmt"
	"html/template"
	"net/http"
)

type hndlr int

func a(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "index.gohtml", "This is the default page!!")
}

func b(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "index.gohtml", "My dog's name is Monya.")
}

func c(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "index.gohtml", "My name is Akshay.")
}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("template/index.gohtml"))
}

func main() {

	http.HandleFunc("/", a)
	http.HandleFunc("/dog", b)
	http.HandleFunc("/me", c)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
