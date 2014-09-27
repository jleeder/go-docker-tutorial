package main

import (
	"app/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type AppServer struct {
	gorilla *mux.Router
}

func (s *AppServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	go s.handleAccessLog(req)
	s.gorilla.ServeHTTP(rw, req)
}

func (s *AppServer) handleAccessLog(req *http.Request) {
	log.Println(req.URL)
}

// Go style "Constructor" - a Get* Method in the package
func NewAppServer() *AppServer {
	return &AppServer{gorilla: configureRouter()}
}

//--- ROUTES ----
func configureRouter() *mux.Router {
	r := mux.NewRouter()
	home := new(routes.Home)

	r.HandleFunc("/", home.Index)

	return r
}
