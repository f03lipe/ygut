package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/f03lipe/ygut/conf"
	"github.com/f03lipe/ygut/globals"
	"github.com/f03lipe/ygut/handlers"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

const (
	defaultPort int = 5000
)

func route() http.Handler {
	r := mux.NewRouter().Schemes("http").Subrouter()

	r.HandleFunc("/", handlers.GetIndex).Methods("GET")

	csrfKey := os.Getenv("CSRF_KEY_32")
	if csrfKey == "" {
		panic("CSRF_KEY_32 env variable not found.")
	}

	if conf.C.Env == "production" {
		return csrf.Protect([]byte(csrfKey))(r)
	} else {
		return csrf.Protect([]byte(csrfKey), csrf.Secure(false))(r)
	}
}

func main() {
	conf.Setup()

	fmt.Printf("Starting.")

	g := globals.Setup()
	defer globals.Close(g)

	http.Handle("/", route())

	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(defaultPort)
	}
	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}
