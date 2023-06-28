package routes

import (
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

func (r *Routers) InitializeDepartmentRoutes(s controllers.Server) {

	r.routers.HandleFunc("/departments", middlewares.SetMiddlewareJSON(s.Departments)).Methods("GET")
	r.routers.HandleFunc("/departments", middlewares.SetMiddlewareJSON(s.CreateDepartment)).Methods("POST")
	r.routers.HandleFunc("/departments/{id}", middlewares.SetMiddlewareJSON(s.UpdateDepartment)).Methods("PUT")
	r.routers.HandleFunc("/departments/{id}", middlewares.SetMiddlewareJSON(s.DeleteDepartment)).Methods("DELETE")
	r.routers.HandleFunc("/departments/{id}", middlewares.SetMiddlewareJSON(s.GetDepartmentById)).Methods("GET")

}
