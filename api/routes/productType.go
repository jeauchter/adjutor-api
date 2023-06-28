package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeProductTypeRoutes(s controllers.Server) {

	r.routers.HandleFunc("/product-types", middlewares.SetMiddlewareJSON(s.ProductTypes)).Methods("GET")
	r.routers.HandleFunc("/product-types", middlewares.SetMiddlewareJSON(s.CreateProductType)).Methods("POST")
	r.routers.HandleFunc("/product-types/{id}", middlewares.SetMiddlewareJSON(s.UpdateProductType)).Methods("PUT")
	r.routers.HandleFunc("/product-types/{id}", middlewares.SetMiddlewareJSON(s.DeleteProductType)).Methods("DELETE")
	r.routers.HandleFunc("/product-types/{id}", middlewares.SetMiddlewareJSON(s.GetProductTypeById)).Methods("GET")

}
