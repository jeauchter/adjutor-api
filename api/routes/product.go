package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeProductRoutes(s controllers.Server) {

	r.routers.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.Products)).Methods("GET")
	r.routers.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")
	r.routers.HandleFunc("/products/{id}", middlewares.SetMiddlewareJSON(s.UpdateProduct)).Methods("PUT")
	r.routers.HandleFunc("/products/{id}", middlewares.SetMiddlewareJSON(s.DeleteProduct)).Methods("DELETE")
	r.routers.HandleFunc("/products/{id}", middlewares.SetMiddlewareJSON(s.GetProductById)).Methods("GET")

}
