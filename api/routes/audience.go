package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeAudienceRoutes(s controllers.Server) {

	r.routers.HandleFunc("/audiences", middlewares.SetMiddlewareJSON(s.Audiences)).Methods("GET")
	r.routers.HandleFunc("/audiences", middlewares.SetMiddlewareJSON(s.CreateAudience)).Methods("POST")
	r.routers.HandleFunc("/audiences/{id}", middlewares.SetMiddlewareJSON(s.UpdateAudience)).Methods("PUT")
	r.routers.HandleFunc("/audiences/{id}", middlewares.SetMiddlewareJSON(s.DeleteAudience)).Methods("DELETE")
	r.routers.HandleFunc("/audiences/{id}", middlewares.SetMiddlewareJSON(s.GetAudienceById)).Methods("GET")

}
