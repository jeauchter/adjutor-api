package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeremyauchter/adjutor/api/routes"
	"github.com/jeremyauchter/adjutor/connect"
)

type Server struct {
	database connect.Server
	Router   routes.Routes
}

func (server *Server) Initialize() {
	server.database.Connect()
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
