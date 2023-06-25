package api

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize()
	server.Run(":8080")
}
