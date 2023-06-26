package controllers

import (
	"github.com/jeremyauchter/adjutor/connect"
)

type Server struct {
	database connect.Server
}

func (server *Server) Initialize() {
	server.database.Connect()

}
