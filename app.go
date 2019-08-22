package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router *mux.Router
}

// initialize router of app
func (a *App) Initialize() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.InitializeRoutes()
}

// initialize routes into the mux router
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/hello", a.SayHello).Methods("GET")
	a.Router.HandleFunc("/hello/{name:[A-Za-z]+}", a.SayHello).Methods("GET")
	a.Router.HandleFunc("/health", a.GetHealthStatus).Methods("GET")
	a.Router.NotFoundHandler = http.HandlerFunc(a.ErrorHandler)
}

func (a *App) Run(addr string) {
	logger.Fatal(http.ListenAndServe(addr, a.Router))
}
