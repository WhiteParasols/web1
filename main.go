package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user) //([]byte, error), Go product to Json
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
	fmt.Fprint(w, "hello foo!")
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello example!")
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	// http.NewServeMux()
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "hello world")
	// })

	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello Bar!")
	// })

	// http.Handle("/foo", &fooHandler{})

	// http.HandleFunc("/example", exampleHandler)

	// http.ListenAndServe(":3000", nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar!")
	})

	mux.Handle("/foo", &fooHandler{})

	mux.HandleFunc("/example", exampleHandler)

	mux.HandleFunc("/name", nameHandler)

	http.ListenAndServe(":3000", mux)
}
