package controllers

import (
	"github.com/jeremyauchter/adjutor/api/middlewares"
	"github.com/jeremyauchter/adjutor/api/routes"
	"github.com/jeremyauchter/adjutor/connect"
)

type Server struct {
	database connect.Server
	router   routes.Routers
}

func (server *Server) Initialize() {
	server.database.Connect()
	server.router.StartRouter()
	server.router.InitializeRoutes(middlewares.SetMiddlewareJSON(server.Home))
	server.router.InitializeTagRoutes(middlewares.SetMiddlewareJSON(server.Home))
}

func (server *Server) Run(addr string) {
	server.router.Run(addr)
}
