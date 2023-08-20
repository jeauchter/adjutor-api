package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeStyleRoutes(s controllers.Server) {

	r.routers.HandleFunc("/styles", middlewares.SetMiddlewareJSON(s.Styles)).Methods("GET")
	r.routers.HandleFunc("/styles", middlewares.SetMiddlewareJSON(s.CreateStyle)).Methods("POST")
	r.routers.HandleFunc("/styles/{id}", middlewares.SetMiddlewareJSON(s.UpdateStyle)).Methods("PUT")
	r.routers.HandleFunc("/styles/{id}", middlewares.SetMiddlewareJSON(s.DeleteStyle)).Methods("DELETE")
	r.routers.HandleFunc("/styles/{id}", middlewares.SetMiddlewareJSON(s.GetStyleById)).Methods("GET")

}
