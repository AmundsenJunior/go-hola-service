package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (a *App) SayHello(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var name string
	if name = params["name"]; name == "" {
		name = "Roald"
	}

	hello := fmt.Sprintf("Hello from Go, %s!", name)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hello))
	logger.Printf("%s %s %d", "GET", r.URL, http.StatusOK)
}

func (a *App) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	logger.Printf("%s %s %d", "GET", r.URL, http.StatusOK)
}

func (a *App) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	logger.Printf("%s %s %d", "GET", r.URL, http.StatusNotFound)
}
