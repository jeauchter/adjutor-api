package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeremyauchter/adjutor/api/controllers"
	"github.com/jeremyauchter/adjutor/api/middlewares"
)

type Routers struct {
	routers *mux.Router
}

func (r *Routers) StartRouter() {
	r.routers = mux.NewRouter()
}

func (r *Routers) InitializeRoutes(s controllers.Server) {

	// Home Route
	r.routers.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// // Login Route
	// s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// //Users routes
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// //Posts routes
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}

func (r *Routers) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, r.routers))
}
