// website.go
package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Id    int
	Name  string
	Login string
	Role  string
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8085", nil)
}
func createUserList() []User {
	return []User{
		{3, "Vladimir Ananyev", "vladimir.ananyev", "administrator"},
		{5, "Artem Kukharenko", "artem.kukharenko", "administrator"},
		{6, "Tatyana Murzova", "tatyana.murzova", "administrator"},
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, createUserList())
}
