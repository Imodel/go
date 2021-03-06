package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloname(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_log"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " "))
	}

	fmt.Fprintf(w, "hello world")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloname)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("listenAndServer: ", err)
	}
}
