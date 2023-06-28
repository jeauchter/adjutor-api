package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeTagRoutes(s controllers.Server) {

	r.routers.HandleFunc("/tags", middlewares.SetMiddlewareJSON(s.Tags)).Methods("GET")
	r.routers.HandleFunc("/tags", middlewares.SetMiddlewareJSON(s.CreateTag)).Methods("POST")
	r.routers.HandleFunc("/tags/{id}", middlewares.SetMiddlewareJSON(s.UpdateTag)).Methods("PUT")
	r.routers.HandleFunc("/tags/{id}", middlewares.SetMiddlewareJSON(s.DeleteTag)).Methods("DELETE")
	r.routers.HandleFunc("/tags/{id}", middlewares.SetMiddlewareJSON(s.GetTagById)).Methods("GET")

}
