package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeCountryRoutes(s controllers.Server) {

	r.routers.HandleFunc("/countries", middlewares.SetMiddlewareJSON(s.Countrys)).Methods("GET")
	r.routers.HandleFunc("/countries", middlewares.SetMiddlewareJSON(s.CreateCountry)).Methods("POST")
	r.routers.HandleFunc("/countries/{id}", middlewares.SetMiddlewareJSON(s.UpdateCountry)).Methods("PUT")
	r.routers.HandleFunc("/countries/{id}", middlewares.SetMiddlewareJSON(s.DeleteCountry)).Methods("DELETE")
	r.routers.HandleFunc("/countries/{id}", middlewares.SetMiddlewareJSON(s.GetCountryById)).Methods("GET")

}
