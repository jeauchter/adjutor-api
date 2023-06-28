package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeVendorRoutes(s controllers.Server) {

	r.routers.HandleFunc("/vendors", middlewares.SetMiddlewareJSON(s.Vendors)).Methods("GET")
	r.routers.HandleFunc("/vendors", middlewares.SetMiddlewareJSON(s.CreateVendor)).Methods("POST")
	r.routers.HandleFunc("/vendors/{id}", middlewares.SetMiddlewareJSON(s.UpdateVendor)).Methods("PUT")
	r.routers.HandleFunc("/vendors/{id}", middlewares.SetMiddlewareJSON(s.DeleteVendor)).Methods("DELETE")
	r.routers.HandleFunc("/vendors/{id}", middlewares.SetMiddlewareJSON(s.GetVendorById)).Methods("GET")

}
