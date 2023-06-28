package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeItemVariantRoutes(s controllers.Server) {

	r.routers.HandleFunc("/item-variants", middlewares.SetMiddlewareJSON(s.ItemVariants)).Methods("GET")
	r.routers.HandleFunc("/item-variants", middlewares.SetMiddlewareJSON(s.CreateItemVariant)).Methods("POST")
	r.routers.HandleFunc("/item-variants/{id}", middlewares.SetMiddlewareJSON(s.UpdateItemVariant)).Methods("PUT")
	r.routers.HandleFunc("/item-variants/{id}", middlewares.SetMiddlewareJSON(s.DeleteItemVariant)).Methods("DELETE")
	r.routers.HandleFunc("/item-variants/{id}", middlewares.SetMiddlewareJSON(s.GetItemVariantById)).Methods("GET")

}
