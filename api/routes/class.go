package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeClassRoutes(s controllers.Server) {

	r.routers.HandleFunc("/classes", middlewares.SetMiddlewareJSON(s.Classes)).Methods("GET")
	r.routers.HandleFunc("/classes", middlewares.SetMiddlewareJSON(s.CreateClass)).Methods("POST")
	r.routers.HandleFunc("/classes/{id}", middlewares.SetMiddlewareJSON(s.UpdateClass)).Methods("PUT")
	r.routers.HandleFunc("/classes/{id}", middlewares.SetMiddlewareJSON(s.DeleteClass)).Methods("DELETE")
	r.routers.HandleFunc("/classes/{id}", middlewares.SetMiddlewareJSON(s.GetClassById)).Methods("GET")

}
