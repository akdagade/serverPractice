package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type hndlr int

func (h hndlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	fmt.Println(r.Form, " -- ", r.Form["fname"], " -- ", len(r.Form["fname"]))
	tmp.ExecuteTemplate(w, "index.gohtml", r.Form)

}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("template/index.gohtml"))
}

func main() {
	var h hndlr
	err := http.ListenAndServe(":9999 ", h)
	if err != nil {
		fmt.Println(err)
	}
}
