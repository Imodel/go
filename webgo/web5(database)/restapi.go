package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you get user id %s", uid)
}

func modifuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you modif user id %s", uid)
}

func deteleuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you delete user id %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	mux.Put("/user/:uid", modifuser)
	mux.Del("/user/:uid", deteleuser)
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
